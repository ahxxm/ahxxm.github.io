---
title: 用Tarsnap备份VPS
author: ahxxm
layout: post
permalink: /146.moew/
categories:
  - 不要问我为什么写这个
---

备份不一定安全，不备份也不一定不安全，但是花了钱肯定开心，所以有此文。

  Tarsnap的技术细节可以在<a href="http://www.daemonology.net/blog/2008-12-14-how-tarsnap-uses-aws.html" target="_blank">这里</a>看到，简而言之：加密、增量、压缩的备份服务，有一个还比较好用的备份/还原命令。</p> 
  
## 效果

第一次全盘备份：我的VPS磁盘用了5.5G，去除不必要备份的文件夹之后4G多点，Tarsnap认为unique data只有3G多点，压缩后是1G出头。这1G出头就是……大概0.25刀每个月。

随后半个小时左右的增量备份：新排除了/run，ttrss更新了一些，压缩后用了32M，可以忽略不计。

## 安装

我VPS的系统是Debian 7，所以得编译安装，OpenBSD可以直接用port装之。安装dependencies：

    sudo apt-get install gcc libc6-dev make libssl-dev zlib1g-dev e2fslibs-dev

其中e2fslibs-dev可能会报错，先删e2fslibs再装就行，暂时没遇到什么问题。

然后用默认配置编译安装tarsnap：

    wget <a href="https://www.tarsnap.com/download/tarsnap-autoconf-1.0.36.1.tgz">https://www.tarsnap.com/download/tarsnap-autoconf-1.0.36.1.tgz</a>
    tar xf tarsnap-autoconf-1.0.36.1.tgz
    cd tarsnap-autoconf-1.0.36.1
    ./configure
    make all
    make install

## 配置

配置分两部分，一个是<a href="https://www.tarsnap.com/gettingstarted.html" target="_blank">生成加密用的key</a>：

    tarsnap-keygen –keyfile /root/tarsnap.key –user me@example.com –machine machine-name

改一下邮箱地址和机器名称，运行后会让你输入tarsnap的密码。这时候可以先交点钱，比如10刀。然后把key备份到你觉得安全的地方，因为一旦丢失备份的数据也就解密不了了。

另一部分是修改默认备份配置，把一些不需要备份的文件夹排除掉，创建文件/usr/local/etc/tarsnap.conf：

    # Tarsnap cache directory
    cachedir /usr/local/tarsnap-cache
    
    # Tarsnap key file
    keyfile /root/tarsnap.key

    # Don't archive files which have the nodump flag set
    nodump

    # Print statistics when creating or deleting archives
    print-stats

    # Create a checkpoint once per GB of uploaded data.
    checkpoint-bytes 1G

    # Exclude and include folder/files
    exclude /proc
    exclude /boot
    exclude /dev
    exclude /run
    exclude /var/swap.img
    exclude /sys
    include /

这里（匹配顺序要求）exclude和include不能反，swap是交换分区的文件，如果你的交换分区路径不一样记得修改。

## 定时备份

创建一个<a href="https://www.tarsnap.com/simple-usage.html" target="_blank">长成这样</a>的 /root/tarsnap-backup.sh :

    #!/bin/sh
    /usr/local/bin/tarsnap -c \
      -f mybackup-`date +%Y-%m-%d_%H-%M-%S` \
      /

chmod +x之，然后在crontab里加上：

    8 4 * * * /root/tarsnap-backup.sh

时间和频率当然可以改……

## 测试

开tmux跑一下/root/tarsnap-backup.sh，再tarsnap &#8211;list-archives | sort即可。

配置文件中写了nodump，所以本地是看不到备份的……

    tarsnap -tv -f mybackup-2015-10-05_13-49-25

这样可以看到备份中的文件列表。
