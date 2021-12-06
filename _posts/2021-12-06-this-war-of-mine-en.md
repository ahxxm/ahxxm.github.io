---
title: A Boring Security Incident
author: ahxxm
layout: post
permalink: /170.moew/
categories:
  - IT
---

A strange error message pops up when I run `go get -u -v`.

<!--more-->

`ERROR: ld.so: object '/usr/local/lib/libdb.so' from /etc/ld.so.preload cannot be preloaded (failed to map segment from shared object): ignored.`

I thought: this is fine, it is an Ubuntu server, it reboots randomly, adds random libraries to `LDPRELOAD` and all it needs is another regular `apt update && apt upgrade` maintenance. But  `apt upgrade` was so slow that I had some time to ~~overthink again: is VPS overselling that common?~~ check what's going on.

A process named `db` is using 100% CPU(a logic core), the executable is even under `PATH`:

```bash
➜  ~ which db
/usr/sbin/db
➜  ~ l `which db`
-rwxr-xr-x 1 root adm  7.5M Nov 22 12:54 db
```

[db](https://files.catbox.moe/w2jeng.gz) turns out to be a [CPU mining](https://github.com/xmrig/xmrig) program according to its unstripped symbols, but who is it working for?

<img class="alignnone" src="/images/mine/persistent.jpg" alt="LD_PRELOAD sounds more advanced"/>

[libdb.so](https://files.catbox.moe/itk651.gz) does not contain any credentials or wallet address, it is a `Trojan:Linux/Processhider!mclg`(thanks Windows Defender, I know I can run Linux executables on WSL). 

After happily exploring these two c++-compiled files in IDA for several minutes, I found a plaintext configuration file  `~/.config/xmrig.json`, it tells me the wallet address and mining pool website. Probably due to ~~lack of basic design concerns~~ the anonymous nature, I can browse mining statistics related to this address, or any address registered there.

<img class="alignnone" src="/images/mine/996.jpg" alt="996 is not the most lucky number"/>

Now the question left is, where does it come from?  `Caddy` is presumed to be safe, `sshd_config` disabled password authentication, and nobody would use Docker 0-day just to mine stupid crypto coins at this pathetic rate.

<img class="alignnone" src="/images/mine/worker.jpg" alt=""/>

"Accepted Shares" of 4 workers are close over time(they are all around 330M now), suggesting that mining programs are running on CPUs with similar performance and started on close dates.

Thus my conjecture is to blame the VPS provider and its virtualization platform:
- the attacker can append their rsa key to any `authorized_keys` because disks are not encrypted <!-- ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC9bg+z+g2qbTTCP1IFySRUtkvYRXqwkTjqJeztxIn8yrZYZSpkwnqg4gTdzSQKapAcOkBhZ+huaEHdLROpDq6zQ2UXJSVCUEPX9CLznK6fq5B9alM/8RAmJc3uF099l6ERSru9wL8wFuyH9nqma5sTvd72XLZxBoNl1Xqxj0oBHKGa6gbIDDwXcVeHiwg9JV0+JfDyF9wznAUJQ1LhYfOJuCJ1aorq2y7jR4hqzBls/3zdKqN719AyYN99/RbXzvTkLexSuHKhlcXquLjvLgulDy8OSoNFBe18qTStFWKHqqtXLHhVsxWsaYYAm8CNEcUMhMo8qGomS40FHQydn74N root -->
- however that ~~unnecessarily~~ requires a reboot to take effect: `journalctl --list-boots` reports twice reboot(they are not random!), `-2 dfb0b8c3880349fbb2a5218fb2dd41e6 Mon 2021-11-22 12:29:45 UTC—Mon 2021-11-22 12:29:48 UTC` and  `-1 aab144823db246d082730af353decc2c Mon 2021-11-22 12:47:51 UTC—Sat 2021-12-04 11:20:35 UTC`, the time matches with the mining executable

After finishing the conjecture I noticed a "Send Now" button, it says:

<img class="alignnone" src="/images/mine/fee.jpg" alt=""/>

I misread it as "0.004", which I thought was reasonable, all the attacker needs is a bot to "Send Now" once the balance exceeds 0.004, nice design again....... then, ok, it is 0.0004, never mind.

<!-- This VPS had gone through hard times, it was powered off for several days after [the OVH fire](https://network.status-ovhcloud.com/incidents/vlcqgm66ffnz), its network is not always stable but has been improving recently, it has good value for its price(2cores+6g+300g nvme ssd, less than $10 per month), I'm hesitant to move out and sincerely hope it survives the next incident. -->
