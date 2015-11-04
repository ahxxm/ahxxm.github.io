---
title: 'Hackintosh: E3-1230 V3 + Gigabyte B85-HD3 + Inno3D GTX 760 (alpha version)'
author: ahxxm
layout: post
permalink: /126.moew/
categories:
  - 未分类
---
给台式机装了个黑苹果，显卡正常驱动，外接水星mw150us+驱动正常使用，能做开发环境，不出意外吊打未来三年内所有顶配Macbook Pro。

问题有：不能登录App Store，声卡无驱动，重启过慢，无法睡眠，CPU无法降频，以及一大堆问题。弄好<a href="http://bbs.pcbeta.com/viewthread-900017-1-1.html" target="_blank">DSDT</a>大概能解决大部分？<span style="color: #999999;"> </span>

暂时不继续，以后折腾好了再补上全程，以及后续开发环境配置。

安装过程流水帐记录如下——<!--more-->

  1. 下载<a href="http://bbs.pcbeta.com/viewthread-1550906-1-1.html" target="_blank">懒人版镜像</a>，写入随便什么鬼分区教程见2链接
  2. 用WoWPC.iso引导开始安装（<a href="http://bbs.pcbeta.com/viewthread-1518901-1-2.html" target="_blank">教程和下载地址</a>）
  3. U盘装Clover，接下来步骤统统是U盘引导，插拔切换系统挺炫酷的
  4. （Win下）S/L/E放入FakeSMC，（Clover引导进）单用户罗嗦模式（single user verbose mode）<a href="http://myhack.sojugarden.com/guide/" target="_blank">修复权限</a>，这里的myfix或者pcbeta上的都行，用法不太一样，-h之
  5. 重启进系统，安装<a href="http://pan.baidu.com/s/1sjqMMmD" target="_blank">无线网卡驱动</a>
  6. 重启进系统

待续……