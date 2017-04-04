---
title: k2pdfopt One-liner
author: ahxxm
layout: post
permalink: /154.moew/
categories:
  - PDF
---

[K2pdfopt](http://www.willus.com/k2pdfopt/download/) saves kindle paperwhite with better formatted pdf after [conversion](http://www.willus.com/k2pdfopt/help/options.shtml), without jailbreak.

## TLDR

```
./k2pdfopt -as -dev kp2 -jpg 90 -mode fw [filename]
[press ENTER and wait]
```

Explain params:

- -as: Autostraighten skewed pages(no need for text pdf)
- -dev kp2: device is kindle paperwhite 2
- -jpg 90: compress output file
- -mode fw: fit width

### Optional Options

- -n: **disables** text reflow, good for lame pdf reader like `Preview.app`. Smaller output size because "the source PDF's native content is used along with additional PDF instructions to translate, scale, and crop the source content".




