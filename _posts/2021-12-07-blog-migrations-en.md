---
title: Blog Hosting Migration History
author: ahxxm
layout: post
permalink: /157.moew/
categories:
  - Ops
---

In reverse chronological order.

<!--more-->

## CloudFlare Pages(2021-)

It's new but quite completed, migrating from Netlify took less than 10 minutes with zero downtime:

- set up new CloudFlare Pages: link Github account, select Jekyll template and deploy
- update CNAME records

Now I don't need to worry about certificate renewal, not bad.

## Netlify(2019-2021)

Linode says Tokyo1 will be retired, I don't like its Tokyo2 DC, so I migrated this blog to Netlify.

Steps:

- Repo build requirements: add Gemfile to repo, set build command `jekyll build`, publish dir `_site`
- Netlify Dashboard: add domain, upload wildcard certs(cert.pem, chain.pem, privkey.pem from certbot), enable all assets optimization except `Pretty URLs`
- CloudFlare: delete A record of root domain, CNAME it to Netlify-generated domain
- [Global Ping](https://tools.keycdn.com/ping): verify DNS records were populated

This blog now gets global CDN and loses access log -- which I haven't spared any time to audit yet.

## Linode Tokyo1(2014-2019)

Back then SSL certificates were not free, Netlify was out but I don't know, nginx required [linking](https://github.com/ahxxm/base/blob/master/nginx/nginx.sh#L22) with `LibreSSL` to support `CHACHA20_POLY1305`, the "Continuous Deployment" is simply `docker-compose build && docker-compose up -d --remove-orphans`.

I had some fun getting an A+ in [Qualys SSL Labs](https://www.ssllabs.com/ssltest/), the efforts resulted in about 100 lines of nginx configuration files.

<img class="alignnone" src="/images/blog/blog-logs.jpg" alt="mostly web crawlers"/>

It feels good when I know there are real humans who spent their time with my writings.
