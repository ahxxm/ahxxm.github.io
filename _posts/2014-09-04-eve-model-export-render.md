---
title: EVE模型导出渲染指南
author: ahxxm
layout: post
permalink: /40.moew/
categories:
  - 我也不知道为什么要写这个
---
去年底的文章，当时EVE好友想做视频，需要导出模型。不过后来视频好像黄了……

效果图在底部。

本文相当于操作笔记，不讲原理，只说步骤，如有不懂的地方，参考文献有两篇：

- <a href="http://bbs.ngacn.cc/read.php?tid=6129126">http://bbs.ngacn.cc/read.php?tid=6129126</a>
- <a href="http://wiki.eveuniversity.org/Rendering_EVE_Models_101">http://wiki.eveuniversity.org/Rendering_EVE_Models_101</a>

还有个<a href="http://www.youtube.com/watch?v=Pw18CzYzJvg">视频</a>。看youtube不方便的可以点这里：<a href="http://www.tudou.com/programs/view/DJeXH0rnYcg/">http://www.tudou.com/programs/view/DJeXH0rnYcg/</a>

英文不好的去练听力。

<hr>

工具有三个：

* <a href="http://www.blender.org/download">Blender</a>
* <a href="http://dl.eve-files.com/media/corp/Xabora/TriExporter_2009.zip">TriExporter 2009 </a>
* <a href="http://www.gimp.org/downloads/">GIMP</a>及<a href="https://code.google.com/p/gimp-dds/downloads/list">dds插件</a>


废话不多说——


## <a href="http://bbs.ngacn.cc/read.php?tid=6129126">模型导出</a>

打开TriExporter 2009，File->Set EVE Floder，选择EVE目录，即eve.exe所在目录，例如C:\Program Files\EVE，
船的模型文件在res/dx9/ship/目录下，以amarr bc1为例，导出res/dx9/ship/amarr/battlecruiser/abc1下所有的文件，
该目录下包括了以abc1为船体的各种变种，在这里只关心bc1，需要的dx贴图文件为:

- abc1_t1_d.dds
- abc1_t1_n.dds
- abc1_t1_pgs.dds

导出模型，在TriExporter 2009中双击abc1_t1.gr2会选中并显示该模型，选中后File->Export Tri-Model，
保存类型选.3ds，导出.black，打开Black2Json将由TriExporter导出的.black拖到程序窗口，会在同一目录生成.json文本文件。


## 贴图文件

### _d.dds

用GIMP打开，color->components->decompose，选RGBA，点击右栏眼睛开关channels——关掉alpha，
把剩下三个channel导出（export to）为diffuse.png，然后把alpha导出为lights.png。

### _pgs.dds

green channel导出为specular.png。

color->components->compose，RGB，合成后导出整个为reflection.png。


### _n.pgs

RGBA decompose，再compose，此时red green blue分别是：green、alpha和mask value 255，视频说合成后是蓝色的图就对了，
不过文1截图证明合成后颜色奇怪就对了，最后导出为normal.png。

此处操作和视频略有出入，mask value和文1也略有出入，不过效果相同，以下是看视频时写的笔记作为参考：

- video: ngs, p
- now: pgs, n
- red green blue alpha
- ngs: sub-mask(black) normalX specular normalY
- p: null null null mask
- export specular -> specular.png
- RGB compose -> red:normalX, green:normalY, blue:mask 255 -> blue pic -> normal.png
- p -> reflections.png now
- pgs: reflection1 specular reflection2 null
- export specular.png
- compose and export reflection.png
- n:null normalX null normalY
- compose and export normal.png

## Blender

### 准备工作（没图说个）

右键选中初始立方体，del，然后点一下删掉它。

file->import，选3ds文件导出模型，以Incursus为例，因为它长得比较像 o/ 。

选中，r x 90+Enter，绕x轴旋转90度，按S调整大小。

tab进入编辑模式，Alt+J 将三角面转为四边面，然后Ctrl+V，Remove Doubles去除过于接近的顶点。

左侧T key Panel->Object Tools->Shading->Smooth 平滑着色

<img class="alignnone" src="/images/eve/1.png" alt="" width="143" height="270" />

右侧N key panel->display->shading改成GLSL

<img class="alignnone" src="/images/eve/2.png" alt="" width="159" height="267" />

左下Object mode/Edit mode旁边有个按钮，Shading从Rendered改成Texture。

<img class="alignnone" src="/images/eve/3.png" alt="" width="481" height="168" />

选中，右边有个material tab，点进去，新增Incursus，删掉原来的，然后tab进入编辑模式，a选中所有，点assign。

<img class="alignnone" src="/images/eve/4.png" alt="" width="346" height="284" />

### 材质！ （参数设置的原则就俩个字：好看，所以在能看的基础上随意调整吧。）

<img class="alignnone" src="/images/eve/5.png" alt="" width="346" height="236" />

### diffuse

点New，改名为diffuse，type改成Image/Moview，然后open diffuse.png。

Mapping->Coordinates:UV, Map:UVMap

Influence-> Diffuse-> Color:on

             -> Geometry->Normal:on

             -> Specular -> Hardness:on （可选）

<img class="alignnone" src="/images/eve/6.png" alt="" width="407" height="199" />

这里进入UV Editing，tab+a选中，左边N key Panel->Display，点Normalized，下面的Cursor Location中x=0.5 y=0.5 ，然后S,Y,-1+Enter 将UV坐标沿X轴翻转。

以下简写

#### normal

- type, open, Coordiantes -> UV
- sampling->normal map:on
- influence->color:off, geometry normal:on

#### specular

- UV
- diffuse->color off
- specular inte on
- blend multiply
- rgb to intensity:111

#### lights

- uv
- blend add
- RGB 111

#### reflect

- color:off
- ray mirror:on
- multiply
- rgb 111

####  world tab

- 设置填空背景
- horizon:black zenith:white blend sky5.lamp
- 选中camera，进入lamp tab设置灯光，黄色+energy3就是空间站效果了。

## 效果图

最后附上两个效果图，我不知道灯光应该丢哪儿 > <：

*Incursus: o/*

<img class="alignnone" src="/images/eve/incursus1.png">

<img class="alignnone" src="/images/eve/incursus2.png">
