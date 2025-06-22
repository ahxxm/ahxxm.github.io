#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Get markdown files from _posts directory only
function* getMarkdownFiles() {
  const dir = '_posts';
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  
  for (const entry of entries) {
    if (entry.isFile() && entry.name.endsWith('.md')) {
      yield path.join(dir, entry.name);
    }
  }
}

async function main() {
  const harper = await import('harper.js');
  const linter = new harper.LocalLinter({
    binary: harper.binary,
    dialect: harper.Dialect.American,
  });

  const fileSuggestions = new Map();  
  for (const file of getMarkdownFiles()) {
    const content = fs.readFileSync(file, 'utf8');
    if (!content.trim()) {
      continue;
    }
    
    const lints = await linter.lint(content);
    const suggestions = [];
    
    for (const lint of lints) {
      if (lint.suggestion_count() !== 0) {
        for (const sug of lint.suggestions()) {
          suggestions.push({
            kind: sug.kind() === 1 ? 'Remove' : 'Replace with',
            text: sug.get_replacement_text()
          });
        }
      }
    }
    
    if (suggestions.length > 0) {
      fileSuggestions.set(file, suggestions);
    }
  }
  
  for (const [file, suggestions] of fileSuggestions) {
    console.log(`${file}:`);
    for (const sug of suggestions) {
      console.log(`  - ${sug.kind}: ${sug.text}`);
    }
  }
  
  if (fileSuggestions.size > 0) {
    process.exit(1);
  }
}

main();
