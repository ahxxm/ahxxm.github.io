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

This blog now gets global CDN and loses access log -- which I haven't spare any time to audit yet.


### Dirty log solution(suspended)

Update 2019-04-12

Came up with a dirty log solution using Google Cloud Stackdriver and Netlify Function(AWS lambda backed):

- add build env: project url as `GO_IMPORT_PATH`, 32 byte key as `KEY` to decode Google Cloud credentials json in main.go
- extend build command: `go get ./... && sed -i "s/todo-32-byte-key-here/$KEY/g" main.go && mkdir -p logger && go build -o logger/logger ./... && jekyll build`
- make client access logger lambda url

Stackdriver provides exuberant quota for free: "First 50 GiB per project/per month".

### AWS Amplify(NOT USED)

Updated 2019-07-20

CloudFront should be faster than DigitalOcean(I guess).

Transition is quite smooth:

- link github and add project, it generates build and deploy
- verify domain in Amplify console(including change CNAME to a CloudFront domain)
- wait for DNS propagation

**But** AWS requires full access to all my git repositories..
