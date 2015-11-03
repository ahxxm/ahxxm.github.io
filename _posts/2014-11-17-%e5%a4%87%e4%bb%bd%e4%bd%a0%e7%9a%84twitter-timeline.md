---
ID: 89
post_title: 备份你的Twitter和微博时间线
author: ahxxm
post_date: 2014-11-17 19:22:54
post_excerpt: ""
layout: post
permalink: https://ahxxm.com/89.moew
published: true
---
<strong>TLDR</strong>

做了这么个事：
<ul>
	<li>微博同步至timeline</li>
	<li>保存timeline中重要信息至MySql</li>
	<li>保存完整的微博和Twitter timeline至MySql（json dump）</li>
</ul>
起因：突发奇想，想要做个Twitter Offline给自己用，但想破头也不知道应该怎么呈现信息，所以先完成前一半——保存。做完后偶然间刷一下微博，感觉有不少有意思的人，即然这样就同步过来看吧，一个客户端了事。所以有了这篇文(guan)章(shui)。

下面是流程。<!--more-->

<strong>需要准备</strong>

一个VPS，装好python和三个依赖：
<blockquote>easy_install requests

easy_install twitter

easy_install weibo</blockquote>
一个买买买，三个安装，非常惬意。

<strong>开始干</strong>

按照惯例，此时应该有代码链接——
<ul>
	<li>Twitter timeline抓取：<a href="https://raw.githubusercontent.com/ahxxm/timeline-storage/master/offline.py" target="_blank">点这里</a></li>
	<li>微博同步至Twitter：<a href="https://github.com/ahxxm/sina-timeline-to-twitter" target="_blank">点这里</a></li>
</ul>
创建一个数据库（见第一个链接里的注释），一个<a href="https://apps.twitter.com/app/new" target="_blank">Twitter应用</a>，一个<a href="http://open.weibo.com/" target="_blank">微博应用</a>，填进代码里。

先手动跑一下，如果出现了成功字样就可以设置Crontab任务——
<blockquote>crontab -e

0,20,40 * * * * /usr/bin/python /root/twioff/twi.py

0,15,30,45 * * * * /usr/bin/python /root/sinaoff/weiboff/py

:wq</blockquote>
大功告成！经测试，30M的SQL Dump经过7z压缩后只有1.5M不到，所以5刀的DO可以打十年。

恩下面可以继续思考怎么做Twitter Offline Client了……

<strong>Credits</strong>

http://docs.python-requests.org/en/latest/

https://pypi.python.org/pypi/twitter

http://lxyu.github.io/weibo/

<hr />

直(tuo)到今天，终于转移至Tweepy+Postgres——Mongo占用硬盘太大，据说Postgres效果好点。<a href="https://gist.github.com/ahxxm/9dd5ddfe7ff8cf6fd162" target="_blank">代码点这</a>，<del>效果还得等时间来告诉我……</del>

跑代码之前有2.5个事要做：

1. 安装
<blockquote>sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ $(lsb_release -cs)-pgdg main" &gt; /etc/apt/sources.list.d/pgdg.list'
sudo apt-get install wget ca-certificates
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install postgresql-9.4
pip install sqlalchemy psycopg2</blockquote>
1.5 和配置Postgres，其中asbot是数据库名，tweets是表名，随便改：
<blockquote>pg_createcluster 9.4 main --start
sudo -u postgres -i
createdb asbot
psql asbot
CREATE ROLE root superuser;
CREATE USER root;
GRANT ROOT TO root;
ALTER ROLE root WITH LOGIN;
\q

psql asbot
DROP TABLE "tweets";
CREATE TABLE "tweets"
(
tweet_id bigserial primary key,
content text
);
\q</blockquote>
2.5 修改路径：

备份代码的15行，数据库名要改；22和30行的last_tweet_id.txt路径也可以改。