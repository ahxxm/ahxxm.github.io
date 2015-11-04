---
title: 用Tarsnap备份VPS
author: ahxxm
layout: post
permalink: /146.moew/
categories:
  - 不要问我为什么写这个
---
<div>
  备份不一定安全，不备份也不一定不安全，但是花了钱肯定开心，所以有此文。
</div>

<div>
  Tarsnap的技术细节可以在<a href="http://www.daemonology.net/blog/2008-12-14-how-tarsnap-uses-aws.html" target="_blank">这里</a>看到，简而言之：加密、增量、压缩的备份服务，有一个还比较好用的备份/还原命令。</p> 
  
  <hr />
</div>

## 效果

第一次全盘备份：我的VPS磁盘用了5.5G，去除不必要备份的文件夹之后4G多点，Tarsnap认为unique data只有3G多点，压缩后是1G出头。这1G出头就是……大概0.25刀每个月。

随后半个小时左右的增量备份：新排除了/run，ttrss更新了一些，压缩后用了32M，可以忽略不计。

## 安装

<div>
  我VPS的系统是Debian 7，所以得编译安装，OpenBSD可以直接用port装之。安装dependencies：
</div>

> <div>
>   sudo apt-get install gcc libc6-dev make libssl-dev zlib1g-dev e2fslibs-dev
> </div>

<div>
  其中e2fslibs-dev可能会报错，先删e2fslibs再装就行，暂时没遇到什么问题。
</div>

<div>
  然后用默认配置编译安装tarsnap：
</div>

> <div>
>   wget <a href="https://www.tarsnap.com/download/tarsnap-autoconf-1.0.36.1.tgz">https://www.tarsnap.com/download/tarsnap-autoconf-1.0.36.1.tgz</a>
> </div>
> 
> <div>
>   tar xf tarsnap-autoconf-1.0.36.1.tgz
> </div>
> 
> <div>
>   cd tarsnap-autoconf-1.0.36.1
> </div>
> 
> <div>
>   ./configure
> </div>
> 
> <div>
>   make all
> </div>
> 
> <div>
>   make install
> </div>

## 配置

<div>
  配置分两部分，一个是<a href="https://www.tarsnap.com/gettingstarted.html" target="_blank">生成加密用的key</a>：
</div>

> <div>
>   tarsnap-keygen &#8211;keyfile /root/tarsnap.key &#8211;user me@example.com &#8211;machine machine-name
> </div>

<div>
  改一下邮箱地址和机器名称，运行后会让你输入tarsnap的密码。这时候可以先交点钱，比如10刀。然后把key备份到你觉得安全的地方，因为一旦丢失备份的数据也就解密不了了。
</div>

<div>
  另一部分是修改默认备份配置，把一些不需要备份的文件夹排除掉，创建文件/usr/local/etc/tarsnap.conf：
</div>

> <div>
>   # Tarsnap cache directory
> </div>
> 
> <div>
>   cachedir /usr/local/tarsnap-cache
> </div>
> 
> <div>
>   # Tarsnap key file
> </div>
> 
> <div>
>   keyfile /root/tarsnap.key
> </div>
> 
> <div>
>   # Don&#8217;t archive files which have the nodump flag set
> </div>
> 
> <div>
>   nodump
> </div>
> 
> <div>
>   # Print statistics when creating or deleting archives
> </div>
> 
> <div>
>   print-stats
> </div>
> 
> <div>
>   # Create a checkpoint once per GB of uploaded data.
> </div>
> 
> <div>
>   checkpoint-bytes 1G
> </div>
> 
> <div>
>   # Exclude and include folder/files
> </div>
> 
> <div>
>   exclude /proc
> </div>
> 
> <div>
>   exclude /boot
> </div>
> 
> <div>
>   exclude /dev
> </div>
> 
> <div>
>   exclude /run
> </div>
> 
> <div>
>   exclude /var/swap.img
> </div>
> 
> <div>
>   exclude /sys
> </div>
> 
> <div>
>   include /
> </div>

<div>
  这里exclude和include不能反，swap是交换分区的文件，如果你的交换分区路径不一样记得修改。
</div>

## 定时备份

<div>
  创建一个<a href="https://www.tarsnap.com/simple-usage.html" target="_blank">长成这样</a>的 /root/tarsnap-backup.sh :
</div>

> <div>
>   #!/bin/sh
> </div>
> 
> <div>
>   /usr/local/bin/tarsnap -c \
> </div>
> 
> <div>
>       -f mybackup-`date +%Y-%m-%d_%H-%M-%S` \
> </div>
> 
> <div>
>       /
> </div>

<div>
  chmod +x之，然后在crontab里加上：
</div>

> <div>
>   8 4 * * * /root/tarsnap-backup.sh
> </div>

<div>
  时间和频率当然可以改……
</div>

## 测试

<div>
  开tmux跑一下/root/tarsnap-backup.sh，再tarsnap &#8211;list-archives | sort即可。
</div>

<div>
  配置文件中写了nodump，所以本地是看不到备份的……
</div>

> <div>
>   tarsnap -tv -f mybackup-2015-10-05_13-49-25
> </div>

<div>
  这样可以看到备份中的文件列表。
</div>