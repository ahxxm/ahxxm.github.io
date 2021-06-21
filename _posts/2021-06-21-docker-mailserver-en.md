---
title: Self-hosting Email Server in 2021
author: ahxxm
layout: post
permalink: /168.moew/
categories:
  - IT
---

A short guide to setup catch-all [docker-mail-server](https://github.com/docker-mailserver/docker-mailserver/), an all-in-one mail server.

<!--more-->

## All-in-one Server

As of writing, latest release tag is `v10.0.0`:

```bash
wget https://raw.githubusercontent.com/docker-mailserver/docker-mailserver/v10.0.0/docker-compose.yml
wget https://raw.githubusercontent.com/docker-mailserver/docker-mailserver/v10.0.0/mailserver.env
wget https://raw.githubusercontent.com/docker-mailserver/docker-mailserver/v10.0.0/setup.sh
```

For example my catch-all address is `catch@mail.ahxxm.com`, edit `docker-compose.yml`:
- `hostname`: `mail`
- `domainname`: `ahxxm.com`

hostname + domainname = `@domain.tld`.

## DNS

Add A record `hostname+domainname` to your server.

Then [Setup DKIM](https://docker-mailserver.github.io/docker-mailserver/edge/config/best-practices/dkim/):

```bash
./setup.sh config dkim
cat config/opendkim/keys/domain.tld/mail.txt
```

[DMARC](https://docker-mailserver.github.io/docker-mailserver/edge/config/best-practices/dmarc/) seems to be a global setting, not sure if it will affect other managed email services..

SPF record is [deprecated](https://docs.microsoft.com/en-us/answers/questions/386129/spf-record-deprecated.html).

## SSL

The server [currently](https://docker-mailserver.github.io/docker-mailserver/edge/config/security/ssl/#caddy) **only supports RSA certificates**, but `caddy` gets EC ones by default.

Assuming you have valid certificates, mount them into container by appending to `docker-compose.yml::volumes`:

```yaml
volumes:
  # ...
  - mail.domain.com.crt:/mail.crt:ro
  - mail.domain.com.key:/mail.key:ro
```

Then update `mailserver.env`:

```
SSL_TYPE=manual
SSL_CERT_PATH=/mail.crt
SSL_KEY_PATH=/mail.key
```

## Email Account, Catch-all

```bash
# add account
./setup.sh email add catch@mail.ahxxm.com "#t%+bscw??eft?xcz"

# catch all: https://github.com/docker-mailserver/docker-mailserver/issues/516#issuecomment-278750255
# echo "@domain.tld prefix@domain.tld" >> config/postfix-virtual.cf
echo "@mail.ahxxm.com catch@mail.ahxxm.com" >> config/postfix-virtual.cf
```

This will trigger hot reload, unlike env updates.

Now the service is ready:

```bash
docker-compose up -d --remove-orphans
```

## Optional: Backup

Email data resides in `data/`(in plain text, classified by domain and account), `tarsnap` to backup encrypted content incrementally.

Other files and directories are all configurations, `git` to trace changes.

## Optional Configurations(That I care)

- `SPOOF_PROTECTION=`: disable to send email from arbitrary address
- `ENABLE_CLAMAV=1`  `ENABLE_AMAVIS=1` `ENABLE_SPAMASSASSIN=1`: remove startup warnings
- `POSTFIX_MESSAGE_SIZE_LIMIT=1024000000`: increase from 10MB to 1GB.
- `POSTFIX_INET_PROTOCOLS=ipv4`: "Most likely you want this behind Docker."


## Optional: SMTP relay service

My VPS provider blocks outbound traffic to port 25, so a relay service is needed. According to [docs](https://github.com/docker-mailserver/docker-mailserver/blob/master/docs/content/config/advanced/mail-forwarding/relay-hosts.md), 4 env args will be needed.

Enabling "production access" on AWS SES can be annoying, but the simplicity pays back:
- verify sender address [here](https://console.aws.amazon.com/ses/home?region=us-east-1#verified-senders-email:)
- fill 4 `RELAY_` values at the end of `mailserver.env`, get them from [SMTP settings](https://console.aws.amazon.com/ses/home?region=us-east-1#smtp-settings:), port `587` works for me.

~~Sadly SES blocks unverified sender address..~~
