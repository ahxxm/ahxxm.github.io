---
title: "游戏视频剪辑笔记"
author: ahxxm
layout: post
permalink: /166.moew/
categories:
  - Game
---

以下结论仅适用2017 late macbook pro 15寸，Windows 10 Pro 2004 Build 19041.388，游戏为皇军敢死队（Shadow Tactics: Blades of the Shogun），游戏版本2.2.10.F。

<!--more-->

- [剪辑](#)
  - [视频](#)
  - [音频](#)
  - [快捷键](#)
- [导出](#)
- [屏幕录制](#)
  - [声画不同步修复（失败一大半）](#)
  - [录制中断修复（失败一半）](#)
- [结论](#)

# 剪辑

游戏玩法属于潜行类，大关卡可以分割成多个小谜题，给了充分的S&L和剪辑空间。S&L机制类似快照，读档*几乎*相当于回到存档那秒——除了音频还在照常播放。

所以剪辑用到了Premiere和Audition。

## 视频

剪辑目的有两个：一次通过关卡的成品视频（去掉解谜失败片段），尽量不影响观看体验的剪辑痕迹。

于是总结两个减小剪辑工作量的原则：

1. 打两遍，第二遍再录像，需要剪的地方就会显著变少
2. 更新显卡驱动（而非[“苹果给你什么你就乖乖用”](https://discussions.apple.com/thread/7624874)），让保存时间足够快，就不用把2-3秒的“世界暂停”再剪短了

可惜只有第一和最后一关用到了第一点，第一关是因为事后才想起可以录像玩，最后一关是因为难度太大，怕工作量爆炸……

剪辑痕迹来源是S&L差异，读档后主要是鼠标指针位置变了，另外敌人和环境*也有可能*和存档时有少许不同，我加了个`00;00;00;04`的短Cross Dissolve特效，看起来感觉可以接受。

## 音频

在切割完视频后，每段clip开头可能会有破音，右键点击选择"Edit Clip in Adobe Audition"，选中破音部分，Favorites->auto heal，保存关闭(Ctrl-s Ctrl-w)即可。

一个比较烦人的点：Premiere点完之后*有概率*让这段音频处于被占用状态，会让Audition无法保存，要切回Premiere点其他东西取消占用、切回Audition保存关闭、再切到Premiere看下一段……

另外还可以用Premiere自带的音频过渡特效，我比较喜欢Exponential Fade，译名“指数淡化”。

## 快捷键

一开始不知道Premiere快捷键，全靠鼠标在时间轴上拖，浪费了不少时间精力：

```
[1存档动画 2存档结束 3失败片段某帧 4失败片段结束前某帧 5开始读档 6读档结束]
最好是在2和6用Razor Tool
为了不总是用鼠标缩放时间轴和用鼠标选时刻，我选的3和5，第二遍才慢慢选接近26的，不光浪费时间精力，还不准……
```

解决这个用到了以下快捷键：

- `↑` `↓`：上/下切割点
- `←` `→`：Step back/forward 5帧（Edit->Preferences->Playback里修改）
- `-` `+`：缩放时间轴，Audition里修破音也用得到
- `Ctrl`-`K`：在当前时间点使用Razor Tool
- `Shift`-`Del`：Ripple delete，让后面视频自动“贴”回来
- `L`：空格控制播放和暂停，L加速当前播放（暂停后下次播放还是1倍速）

# 导出

给Premiere装了Voukoder插件——像是个ffmpeg GUI。用到的参数不多：

- preset: Slow
- profile: Main
- CRF: 19-20
- --me umh --merange 60 --subme 11： 参考[老文档](http://www.lighterra.com/papers/videoencodingh264/)，现在subme有11可以选了，原视频有2k分辨率，所以merange也提高了些

导出很是走了几步弯路，hardware encoding导出搞挂显卡驱动，黑屏数分钟后恢复。以为是浏览器同时硬件加速的的问题，关掉浏览器后好了一天，然后什么都不开什么都不干也照样黑屏。

正好读到说software encoding结果会优于前者，即同码率质量更高、同质量码率要求更低，缺点就是要个好CPU……没想到内存也很重要：VBR 2 pass 用爆16g内存报错退出，一开始以为是硬盘不够用，装系统只给windows 256g，当时可用40g，原视频20g，瞎猜中间文件太多，腾出40g后还是报错，就调成CRF模式了。

CRF内存占用可防可控，甚至还能一边打打小游戏。就是无法控制文件大小，选不对CRF值能导致上传时间显著增加……

导出太慢，用到了[防睡眠powershell脚本](https://dmitrysotnikov.wordpress.com/2009/06/29/prevent-desktop-lock-or-screensaver-with-powershell/#comment-5164)：

```powershell
param($minutes = 600)

$myshell = New-Object -com "Wscript.Shell"

for ($i = 0; $i -lt $minutes; $i++) {
  Start-Sleep -Seconds 60
  $myshell.sendkeys(".")
}
```

# 屏幕录制

bandicam 4.3.3.1498正常工作，参数为：H264 AMD VCE、关键帧间隔25、60fps（VFR动态帧率，毕竟游戏开中特效调低分辨率也就30fps出头）、品质100，能录出完整的视频。另附几个错误答案：

- dxtroy，开游戏直接崩了
- alt+win+R香菇录屏，比较依赖AMD驱动版本：bootcamp驱动版本最老，根本不能玩游戏，用两天性能还会降低到只能重启解决；AMD官网驱动可以正常录，对游戏fps影响仅比bandicacm大一点点，但是游戏本身不到10fps；游戏性能最好的[非官方驱动](https://www.bootcampdrivers.com/)下香菇录屏严重影响fps
- AMD relive: 非官方驱动说装里面的opencl就能用，结果不能，查阅AMD论坛有[帖子说](https://community.amd.com/thread/216683)装ffdshow，我看年久失修不太想下，[reddit](https://www.reddit.com/r/Amd/comments/8alrao/to_those_having_issues_with_amd_relive_not/)提供了改注册表开日志的方法，但是从日志没看出什么问题，放弃
- [obs](https://obsproject.com/)，CPU编码出来平均10fps，[obs-amd-encoder](https://github.com/obsproject/obs-amd-encoder/wiki)也一样，折腾了一会不知道该改哪里

## 声画不同步修复（失败一大半）

在录其他游戏时遇到，声音比画面慢，而且延迟越来越高，搜到三种解决方案：

- [Premiere里选](https://helpx.adobe.com/cn/premiere-pro/using/supported-file-formats.html#%E5%8F%AF%E5%8F%98%E5%B8%A7%E9%80%9F%E7%8E%87%E6%96%87%E4%BB%B6%E6%94%AF%E6%8C%81)可变帧：选项藏得很深，找到时候看已经是"Preserve Audio Sync"，后来都没找到第二次
- [整体转换成固定60fps](https://www.premiumbeat.com/blog/sync-audio-video-game-capture/)：转完还是不同步，另外转完剪好又要转码，加上youtube，就是三次了（不算录制的话）
- [把音频转成固定码率](https://zhuanlan.zhihu.com/p/72200181)：mp3没用，但是思路有效，我改了改

改后方案：

```bash
# 提取原始音频
ffmpeg -i input.mp4 -vn -acodec copy output.aac

# 还是到Audition里转换，但是选MPEG-2 .aac格式：保持采样率+固定比特率=>outputcbr.aac
# 不知道为什么ffmpeg转出来的混搭不出效果……

# 视频也转成固定帧率（简单高参数）
ffmpeg -i input.mp4 -pix_fmt yuv420p -c:v libx264 -r 30 -crf 17 -preset medium cfr.mp4

# 混搭输出最终视频
ffmpeg -i cfr.mp4 -i output.aac -c copy -map 0:v:0 -map 1:a:0  out.mp4
```

改后方案让稍短的视频不同步症状显著减轻，55分钟长视频还是明显不同步。

另外有搜索结果推荐降低录制参数，画质优先=>码率优先，改成AVI(Motion JPEG)。前者有些改善、但是画质显著降低，后者录出来文件过大，不太适用。

## 录制中断修复（失败一半）

游戏玩着突然未响应，强行结束游戏和录制进程，mp4文件就坏了，提示`moov atom not found`，在论坛上找到两个软件号称能修复。

免费的[recover mp4 to h264](https://www.videohelp.com/software/recover-mp4-to-h264)，能恢复出**静音**视频：

```bash
# https://www.videohelp.com/download/recover_mp4_192.zip
# 也可以到 https://files.catbox.moe/ybn7rc.zip 下载

# 相同bandicam录的mp4
./recover_mp4.exe reference.mp4 --analyze --ext

# 恢复出文件
./recover_mp4.exe broken.mp4 result.h264 result.aac --ext

# 合并到mp4 container里
 ./raw/magick/ffmpeg.exe -r 30.000 -i result.h264 -i result.aac -bsf:a aac_adtstoasc -c:v copy -c:a copy result.mp4
```

收费的[Video Repair Tool](https://www.videohelp.com/software/Video-Repair-Tool)有点贵，而且默认参数修不好，12分钟修成36秒，另外结尾编码还有问题……

# 结论

打游戏录屏有点为难3年前上网本级配置的电脑……
