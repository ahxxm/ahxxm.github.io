---
title: 用ImageMagick将PDF转换为图片
author: ahxxm
layout: post
permalink: /109.moew/
categories:
  - PDF
---
环境：OSX 10.10.2 + Homebrew

首先安装一大堆东西：

> brew install ghostscript  
> brew install imagemagick  
> brew uninstall jpeg  
> brew install jpeg  
> brew unlink jpeg && brew link &#8211;overwrite jpeg

此时输入convert可以看到一大堆参数，表示它已经安装好了。

<!--more-->

这里以江选为例，点击<a href="http://dwxc.jcet.edu.cn/xwxt/show.aspx?wzid=2c175f21-e218-467f-bef7-f7565a83a7f4&lmid=289227748" target="_blank">链接</a>，将第一卷另存为john1.pdf，建立名为john1的文件夹后，输入如下命令：

> convert \  
> -verbose \  
> -density 150 \  
> -trim \  
> -quality 100 \  
> -alpha remove \  
> -compress jpeg \  
> john1.pdf \  
> john1/john1-%02d.jpg

参数逐一解释如下：

  * density: 目标设备像素密度，横向[x纵向]，可以只写个横向的数字。密度越大转换出文件体积就越大（废话）。
  * trim: (naively)切白边，只判断像素颜色，不考虑扫描版pdf实际情况，效果堪比多看等国产PDF阅读器切白边功能。（建议删掉此行）
  * quality: 好像不用解释。
  * alpha: 对alpha channel的控制，不知道是<a href="http://www.w3.org/TR/PNG-DataRep.html" target="_blank">什么东西</a>，但若不remove转换出的图片就会是黑底。
  * compress: 好像也不用解释，jpeg效果好体积小。
  * %02d：页数，从0开始，所以封面文件名是john1-0.jpg

电脑卡一阵子后，就转换好了。

**Credits(alphabetically):**

<a href="https://twitter.com/54c3/status/568256600332328961" target="_blank">@54c3</a>

<a href="https://github.com/Homebrew/homebrew/issues/29708" target="_blank">convert in ImageMagick does not find modules on 10.9 Mavericks</a>

<a href="https://twitter.com/cry4chaos/status/568256897905614848" target="_blank">@cry4chaos</a>

<a href="http://www.imagemagick.org/script/command-line-options.php" target="_blank">ImageMagick Docs</a>

<a href="https://twitter.com/woshisanhu/status/568259264583569408" target="_blank">@woshisanhu</a>

