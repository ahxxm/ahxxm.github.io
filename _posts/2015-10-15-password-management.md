---
title: 个人密码管理
author: ahxxm
layout: post
permalink: /149.moew/
categories:
  - 不是什么软文
  - 安全
  - 我也不知道为什么写这个
---
几周前支付宝提示帐号被人登录了，所幸支付密码是个强度高的老密码，没有任何损失。登录密码其实不太弱，缺点是在 XX、XX 和 XX 网站曾经用过。改了登录密码之后，支付密码也强行变成了 6 位数字，想到国内银行，我觉得支付宝其实安全等级挺高的，至少是持平银行。

随后花了大概半个小时，把剩下一些弱的、重复使用的都改好了。Lastpass 中安全评分、安全评分区间和主密码评分分别是 95%，top 1%（击败了四舍五入后大约 99% 的 Lastpass 用户）和 100%。

<!--more-->

## 短板到底有多短

<table cellspacing="0" cellpadding="0">
  <tr>
    <td valign="top">
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;"><b>有钱</b></span>
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;"><b>无钱</b></span>
    </td>
  </tr>
  
  <tr>
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;"><b>有数据</b></span>
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;">要命</span>
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;">麻烦</span>
    </td>
  </tr>
  
  <tr>
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;"><b>无数据</b></span>
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;">搞不好要命</span>
    </td>
    
    <td valign="top">
      <span style="color: #000000; font-family: 'PingFang SC'; font-size: small;">爱咋咋</span>
    </td>
  </tr>
</table>

&nbsp;

短板取决于你相信什么，你要是用 Mac 全是盗版软件（难，软件本来就少）、用 QQ 管家 360 还觉得无所谓，那就不用看下去了……

对于互联网上的密码我有这么几个假设：

  * 绝大多数网站都用明文存储，就算加密的，也能很容易解开
  * 绝大多数网站要么很容易攻破，要么他们自己把密码泄漏出去了
  * 你在绝大多数网站都是用同一个用户名/邮箱
  * 攻击者已经整合了你的其他信息，包括而不限于：姓名，手机号，身份证，常用用户名，常用密码

光明的假设也有，这是密码管理工具的基本假设：

  * Lastpass/1Password 不会明文存储密码（<del>这种数字开头的变量名跟我等 Python 用户是没什么关系了，其实</del><del>用了后发现似乎只适合 institutionalized apple boy</del>）；2-Factor Authentication 服务如 Authy 和 Google Authenticator 不会存你的 recovery key；总之，安全服务商不会有你的明文密码
  * Authy 和 Lastpass 不会同时爆炸（难说，搞不好就有什么 0-day，不过这一条在于强调他们得同时用）
  * 正规操作系统中没有内置键盘记录软件
  * 假设你能记住两个或以上的左右高强度密码

Update: 根据[这篇博文](http://myers.io/2015/10/22/1password-leaks-your-data/)，1Password似乎会明文存储，不过我认识的apple fan boy都不在意这一点……

## 该做什么

这一点很多文章提过<del>，我简要补充几句</del>：

  1. 找一个随机密码生成器，一套存储系统，在安全环境下改密码，安全环境的意思是系统干净、挂着**开了加密**的 VPN；
  2. 所有能开启两步验证的服务都开启了，特别是邮箱，这里点名批评 mail.com，广告多、强行不跳转、不支持两步验证；
  3. 改用户名，把重要网站的用户名改得自己都不认识；
  4. 支付宝密码不要用你熟悉的人的生日，可以用你暗恋的人，不容易猜；
  5. SIM 卡丢失后第一时间挂失然后改微信和支付宝密码、登出所有 session，因为这两东西直接用手机号和验证码就可以登录，后者开始做社交了，有欺诈可能性。

工具我选的是 Lastpass + Authy，原因如下：

  * Lastpass 能正常同步，移动设备不用在同一个 WiFi下手动输入密码跟电脑同步，不需要手动确认密码已经到移动设备上去了，当然也不用跟 Onedrive、Dropbox 或者 iCloud 关联，不会考验你的信仰。（关于这个句号请阅读<a href="http://languagelog.ldc.upenn.edu/nll/?p=21548" target="_blank">此文</a>）
  * Lastpass 订阅制，能自动从信用卡扣款，能自动更新…不要觉得这是理所当然的，著名的 1Password 就不行
  * 院士用 Authy

至于 Lastpass 在 Android 上能根据当前 App 自动找到并弹出相关密码供你选择这种事情我觉得还是不要乱说，毕竟 iOS QQ 目前连粘贴密码都不支持。

## 额外的话

公共 WiFi 下记得连个**开了加密**的 VPN 或者全局 Shadowsocks，Chrome 安装 HTTPS Everywhere，人离开电脑就加上密码，不要用国产越狱……

不是说相信外国人，是暂时还没发现国内外黑产交易的迹象。

差不多就这些了，我现在感觉还比较安全。
