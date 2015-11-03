---
title: 备份你的Twitter和微博时间线
author: ahxxm
layout: post
permalink: /89.moew
categories:
  - 不要问我为什么写这个
---
**TLDR**

做了这么个事：

  * 微博同步至timeline
  * 保存timeline中重要信息至MySql
  * 保存完整的微博和Twitter timeline至MySql（json dump）

起因：突发奇想，想要做个Twitter Offline给自己用，但想破头也不知道应该怎么呈现信息，所以先完成前一半——保存。做完后偶然间刷一下微博，感觉有不少有意思的人，即然这样就同步过来看吧，一个客户端了事。所以有了这篇文(guan)章(shui)。

下面是流程。<!--more-->

**需要准备**

一个VPS，装好python和三个依赖：

> easy_install requests
> 
> easy_install twitter
> 
> easy_install weibo

一个买买买，三个安装，非常惬意。

**开始干**

按照惯例，此时应该有代码链接——

  * Twitter timeline抓取：<a href="https://raw.githubusercontent.com/ahxxm/timeline-storage/master/offline.py" target="_blank">点这里</a>
  * 微博同步至Twitter：<a href="https://github.com/ahxxm/sina-timeline-to-twitter" target="_blank">点这里</a>

创建一个数据库（见第一个链接里的注释），一个<a href="https://apps.twitter.com/app/new" target="_blank">Twitter应用</a>，一个<a href="http://open.weibo.com/" target="_blank">微博应用</a>，填进代码里。

先手动跑一下，如果出现了成功字样就可以设置Crontab任务——

> crontab -e
> 
> 0,20,40 \* \* \* \* /usr/bin/python /root/twioff/twi.py
> 
> 0,15,30,45 \* \* \* \* /usr/bin/python /root/sinaoff/weiboff/py
> 
> :wq

大功告成！经测试，30M的SQL Dump经过7z压缩后只有1.5M不到，所以5刀的DO可以打十年。

恩下面可以继续思考怎么做Twitter Offline Client了……

**Credits**

http://docs.python-requests.org/en/latest/

https://pypi.python.org/pypi/twitter

http://lxyu.github.io/weibo/

* * *

直(tuo)到今天，终于转移至Tweepy+Postgres——Mongo占用硬盘太大，据说Postgres效果好点。<a href="https://gist.github.com/ahxxm/9dd5ddfe7ff8cf6fd162" target="_blank">代码点这</a>，<del>效果还得等时间来告诉我……</del>

跑代码之前有2.5个事要做：

1. 安装

> sudo sh -c &#8216;echo &#8220;deb http://apt.postgresql.org/pub/repos/apt/ $(lsb_release -cs)-pgdg main&#8221; > /etc/apt/sources.list.d/pgdg.list&#8217;  
> sudo apt-get install wget ca-certificates  
> wget &#8211;quiet -O &#8211; https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add &#8211;  
> sudo apt-get update  
> sudo apt-get upgrade  
> sudo apt-get install postgresql-9.4  
> pip install sqlalchemy psycopg2

1.5 和配置Postgres，其中asbot是数据库名，tweets是表名，随便改：

> pg_createcluster 9.4 main &#8211;start  
> sudo -u postgres -i  
> createdb asbot  
> psql asbot  
> CREATE ROLE root superuser;  
> CREATE USER root;  
> GRANT ROOT TO root;  
> ALTER ROLE root WITH LOGIN;  
> \q
> 
> psql asbot  
> DROP TABLE &#8220;tweets&#8221;;  
> CREATE TABLE &#8220;tweets&#8221;  
> (  
> tweet_id bigserial primary key,  
> content text  
> );  
> \q

2.5 修改路径：

备份代码的15行，数据库名要改；22和30行的last\_tweet\_id.txt路径也可以改。