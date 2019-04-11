---
title: Blog Migrated to Netlify
author: ahxxm
layout: post
permalink: /157.moew/
categories:
  - Ops
---

Linode says Tokyo1(previously [this blog's](http://ahxxm.github.io) reverse proxy frontend) will be retired.

I don't like its Tokyo2 DC, so migrated this blog to Netlify.

Steps:

- Repo build requirements: add Gemfile to repo, set build command `jekyll build`, publish dir `_site`
- Netlify Dashboard: add domain, upload wildcard certs(cert.pem, chain.pem, privkey.pem from certbot), enable all assets optimization except `Pretty URLs`
- CloudFlare: delete A record of root domain, CNAME it to Netlify-generated domain
- [Global Ping](http://ping.chinaz.com/): verify previous record was populated

This blog now gets global CDN and loses access log -- which I haven't spare any time to audit.
