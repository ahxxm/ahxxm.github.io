---
title: '统计稿件字数并自动生成带文章链接的xls文件 for TECH2IPO&#038;创之'
author: ahxxm
layout: post
permalink: /57.moew/
categories:
  - 不要问我为什么写这个
---
猫扑老习俗：标题要长才能被人看到。

生成的xls文件仅需**少量**编辑即可交给主编，如有侧漏请打开「投稿.csv」寻找自己。

（给可能根本不会存在的友媒读者：请勿滥用此文件，流量要钱的……）

步骤如下：

1. <a href="http://openapi.vdisk.me/?m=file&a=download_share_file&ss=24daG6N6zcDacuIw5Ml9yEaNn8PTinbZy8hZ0KPIQeMueHKFJlZj2Ler4KR5P--2FfOh3R--2B9thFMKOQ7iLpnR8kp1MyKYAEug" target="_blank">下载</a>所有文件，依次安装Python.msi和nltk.exe后，随便用什么编辑器打开generate.py。

Mac用户种族天赋自带Python，故只需按照<a href="http://www.nltk.org/install.html%20" target="_blank">此教程</a>安装NLTK。

> 文件名称：xls.7z  
> 文件大小：16.6 MB  
> 文件指纹：3ffc1fa581b6fa8c9a1b9b9a344a7191  
> 下载页面：<a href="http://openapi.vdisk.me/?m=file&a=download_share_file&ss=24daG6N6zcDacuIw5Ml9yEaNn8PTinbZy8hZ0KPIQeMueHKFJlZj2Ler4KR5P--2FfOh3R--2B9thFMKOQ7iLpnR8kp1MyKYAEug" target="_blank">http://openapi.vdisk.me/?m=file&a=download_share_file&ss=24daG6N6zcDacuIw5Ml9yEaNn8PTinbZy8hZ0KPIQeMueHKFJlZj2Ler4KR5P&#8211;2FfOh3R&#8211;2B9thFMKOQ7iLpnR8kp1MyKYAEug</a>  
> 微博地址：http://vdisk.weibo.com/s/zoSBbKD5SwS7p

2. 找到本月文章ID范围，修改代码「初始化」中「range」后面部分。Tip: 去后台找。

3. 确认译者信息都在dict中，不在其中又懒得改就把「投稿.xls」给他们，自行整理去。（虽然只要改一次就能永久使用）

4. 双击generate.py。

5. 待窗口消失后，当前目录下会出现一堆文件，分发「姓名.xls」文件给译者。

6. 适当给本文作者捐献部分稿酬，Mac用户请将50%捐献额给杜晨妹子。