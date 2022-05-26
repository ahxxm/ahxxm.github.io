---
title: "逆向流水账(7): 网易云音乐分享参数"
author: ahxxm
layout: post
permalink: /173.moew/
categories:
  - Privacy
---

> 不知道从什么时候开始，在每个链接后面都有一个userid，用户名会泄露，歌单会泄露，连社交帐号都会泄露。我开始怀疑，在这个世界上，还有什么隐私可言？

<!--more-->

现在`userid`变成了`uct`，一串url encode过的base64 string，我希望是网易改善了隐私保护，非对称加密用户ID生成独一无二的trackID、用户登陆时获取这串，保护隐私同时也能收集数据。

但其实不是，否则也不会有这篇文章……

还是从移动端开始，用`apktool`解开，直接找关键词：
```bash
java -jar apktool_2.6.1.jar d -f ncm.apk
cd ncm

➜ ncm  ag '"uct"'
smali_classes8/com/netease/cloudmusic/music/biz/share/musicshare/c.smali
1583:    const-string v8, "uct"

smali_classes9/com/netease/cloudmusic/music/biz/voice/player/share/g.smali
1866:    const-string v9, "uct"
```
~~行号不一定对得上，因为读smali时用注释做了点笔记~~

在第一个结果里找到了生成逻辑：
- 获取当前用户UID
- `AES/ECB/PKCS5Padding`加密，secret key就在里面
- 如果加密结果不为空，就URLEncode之

写了一段代码验证这个逻辑：
```clojure
(let [key "hided-but-trivial-to-find" uid "also-hided" uct "try-yourself"
      ks (javax.crypto.spec.SecretKeySpec. (.getBytes key) "AES")
      enc (doto (javax.crypto.Cipher/getInstance "AES") (.init javax.crypto.Cipher/ENCRYPT_MODE ks))
      dec (doto (javax.crypto.Cipher/getInstance "AES") (.init javax.crypto.Cipher/DECRYPT_MODE ks))
      b64-encoder (.withoutPadding (java.util.Base64/getUrlEncoder)) b64-decoder (java.util.Base64/getUrlDecoder)]
  [(= uct (->> (.doFinal enc (.getBytes uid "UTF-8"))
               (.encodeToString b64-encoder)))
   (= uid (String.
           (->> (.decode b64-decoder uct) (.doFinal dec))
           "UTF-8"))]) => [true true]
```

今年初久违打开网易云音乐App，发现可能又被泄露了好些隐私——所有新出现的隐私选项，默认都被设置成**所有人可见**。

不难理解`uct`为什么这样设计也能通过评审，如果有任何评审流程的话。
