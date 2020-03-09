---
title: "豆瓣读书迁移至Goodreads"
author: ahxxm
layout: post
permalink: /163.moew/
categories:
  - Run
---

常年寻Evernote和豆瓣读书替代品，终于都找到了，Joplin替代Evernote，迁移很方便，Goodreads替代豆瓣，花了一点功夫。

<!--more-->

Goodreads一直以来中文书都比较少，所以本次迁移分两部分，还涉及不少代码，感谢ipython notebook。

## 已有内容导入

用Chrome插件[豆伴](https://chrome.google.com/webstore/detail/ghppfgfeoafdcaebjoglabppkfmbcjdd)，把所有图书数据导出为xlsx。根据Goodreads导入格式给想读和读过加一栏`Bookshelves`，值分别为`to-read`和`read`。再加一栏`ISBN`，导出为csv。

用Python和[爬](https://gist.github.com/ahxxm/5b012985008708d96b22102258c1b90d)下来的[数据](https://files.catbox.moe/lq40pz.gz)，把ISBN填上：[https://gist.github.com/ahxxm/5b012985008708d96b22102258c1b90d#gistcomment-3205042](https://gist.github.com/ahxxm/5b012985008708d96b22102258c1b90d#gistcomment-3205042)

到[Goodreads Import](https://www.goodreads.com/review/import)上传该xls。

它会告诉你：

- [...] books imported
- [...] books updated
- [...] books **Consider adding these books manually**

最后这部分就是Goodreads没有的书，需要先创建再导入。

## 新书创建和导入

<del>我一看网页创建书时候，请求里有个`authenticity_token`就吓坏了，以为是js生成的……</del>

经[高人](https://kagami.moe/)提醒，原来CSRF_TOKEN就在网页里，用requests就可以批量创建新书了。

于是先整理出新书数据，[ISBN, Title, Author, Publisher]，导出为`books-remaining.csv`，这个文件要用到两次：第一次是[创建新书](https://gist.github.com/ahxxm/5b012985008708d96b22102258c1b90d#gistcomment-3205054)。第二次是等创建完，到[Goodreads Import](https://www.goodreads.com/review/import)上传。

## 最后必须黑一下Evernote

大概是这么个时间线：

- 2017年左右Evernote SDK不支持Python3，我写了一个，上传到了pypi自己用
- 2017年9月`devsupport@evernote.com`发邮件来说（大意）：我们注意到了pypi上这个evernote3，你介意转移给Evernote吗
- 我在15分钟内转移完毕，并回邮件：希望你们能保持更新
- 2019年3月我发邮件：两年前转给你们了，但是从那之后就没更新过
- 至今Evernote没有回音

本来对Evernote感情是“暂时没有替代品，我买Premium怕你们倒闭了”，现在“你们赶紧倒闭吧，爱谁谁。”

<img class="alignnone" src="/images/douban-goodreads/feedback.png" alt=""/>
