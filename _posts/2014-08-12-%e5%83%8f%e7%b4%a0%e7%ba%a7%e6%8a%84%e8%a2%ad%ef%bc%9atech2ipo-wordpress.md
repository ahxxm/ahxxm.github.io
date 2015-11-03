---
title: '像素级抄袭：Tech2IPO -> WordPress'
author: ahxxm
layout: post
permalink: /20.moew
categories:
  - 不要问我为什么写这个
---
……（弱弱的说一句）请勿用于非法用途

抓数据、插入数据代码，主题包和其他资源文件： 链接: <a href="http://pan.baidu.com/s/1gd8BrcJ" target="_blank" rel="nofollow">http://pan.baidu.com/s/1gd8BrcJ</a> 密码: bpf2

所有文章，截至78755： <a href="http://pan.baidu.com/s/1hqh4OkO" target="_blank">http://pan.baidu.com/s/1hqh4OkO</a> 密码: oyes （这密码是个巧合）<!--more-->

**转移步骤**

  1. 运行wp-ipo-multithread.py抓数据，修改代码中抓取范围，或直接下载解压后抓增量。
  2. 配置本地环境（Windows下推荐用PHPStudy，Linux用LNMP/LAMP），搭建Wordpress，新建文章分类并获取id。
  3. 修改import.php中文章数据读取路径和文章分类id，访问进行导入。
  4. 启用主题，调整css、Ueditor和主题资源图片目录路径。

像素级抄袭核心思想是“抓静态并用动态重写”，下面是详细过程，分为三部分：抓取和导入数据，编写主题，待办事项。

最后附有主题效果测试链接，编写主题过程中主要用到的函数，和参考资料。

**抓取和导入数据**

抓数据没什么好说的，判断是不是404之后直接把HTML弄本地，Python有个threading.Thread module，多线程进行稍微节约了一些时间。

本想直接对比一篇文章发表前后数据库变化、直接写入数据库，发现WordPress里有直接插入文章的函数wp\_insert\_post()，就直接用PHP写插入代码了。

在后台**设置->常规**中设置网站标题**等信息**，**设置文章固定链接格式**为http://localhost/%post_id%。

创建新闻、观点、人物、公司、产品五个文章分类，添加分类描述（顶部导航的文章分类列表需要用到），并**获取分类id，修改import.php对应部分**。然后正则提取所需数据：标题、作者、分类、摘要、发表时间、正文和标签。导入过程中的数据处理如下：

> *摘要、标题、正文和标签无需处理；  
> 作者，如没有则直接插入数据库，密码自动生成，Email格式：id/中文拼音*<a title="Twitter profile for @tech2ipo" href="http://twitter.com/tech2ipo" target="_blank"><em>@tech2ipo</em></a>*.com（数据库中邮箱地址不支持中文，用pinyin.php转换为拼音） ，获取user\_id后，更新display\_name为中文，就会和TECH2IPO一样显示了；  
> 时间用date()转换；  
> 标签险些被坑，一开始正则是「(‘/(?<=\/tag\/).\*(?=<\/a>)/’, 」，偶然发现HTML代码里其他地方还有“/tag/”，幸好真正的标签代码前面还有个引号，所以改成了「(‘/(?<=”\/tag\/).\*(?=<\/a>)/’, 」；  
> 根据TECH2IPO的文章超链接设置WordPress文章超链接，即转移先后文章链接不变，通过更新文章参数中的import\_id实现：在插入文章后获取真实id，以此修改该文章的import\_id（叫做display_id比较形象），最后设置文章分类。*

文章数量有些多，所以导入前重新定义超时时长，以免插入到一半被系统停止……

**主题**

**主页（index.php）**

把style.css和media.css两个CSS文件**另存至本地相应目录**。

创见主页代码很规范，header+content+footer，前后都可以直接抄HTML（细节改动后文会提到），重写正文部分即可。

正文部分也很有规律，分为左右两个部分。

左边，从上到下：

> *热门新闻一篇；  
> 观点，分左右两栏，左边一篇带图、摘要、作者、发表时间，右边只有图和标题，右边共4篇；  
> 第一个最新文章列表；  
> “人物”和“公司”共同组成和“观点”一样大小的“块”（Block），左右对称两栏：各5篇文章，其中只有第一篇有图和评论数，下方4篇是超链接；  
> 第二个最新文章列表；  
> 产品，2×2网格，每个网格由图和标题组成；  
> 第三个最新文章列表；  
> 翻页按钮。*

其中每个“最新文章列表”由5篇文章组成。

右边后文再说。

CSS是现成的，所以广告栏暂时忽略，其他内容从最新主页的HTML代码中复制过来，把**需要动态更新的部分用WordPress函数代替即可**。

现有函数能够获取：发表时间（不稳定，原因未知，暂用“最近修改时间”代替，其他主题也都是这样解决），作者，标题，摘要，评论数；能获取最新文章，也能获取某分类中最新文章。

只有文章配图没有现成的获取函数，TECH2IPO文章配图都在开头，所以直接到WordPress官网找个“获取文章第一张图片URL”的轮子，然后重写正则。  
另外TECH2IPO会生成195*130列表用图，似乎是为了节约流量。这里用「 background-size:contain;」**直接让大图缩略显示**（不能放在background:url(img_url)前面，放后面才会生效）。

翻页按钮，根据WordPress主题结构，网站分为首页（front-page）和文章页（list-page）——

首页：复制index.php为front.php，加入模板前缀，实际文章不只9页，所以翻页按钮用固定代码即可，在**后台设置->阅读中把主页设置为静态主页**；  
文章页：即根据「<a href="http://tech2ipo.com/all/2" target="_blank" rel="nofollow">http://tech2ipo.com/all/2</a>」制作模板，文件名为home.php，结构与Category.php相似，左边20个新闻，右边wrap-right照抄。由于翻页代码中「get\_pagenum\_link($i)」默认将当前页作为基准，生成链接形式为「$current\_page\_id/page/$next\_or\_prev\_page\_id」，故新建函数，用正则删掉其中的「$current\_page\_id/page」。  
另外还要在后台**设置->阅读->至多显示每页20篇**，不然$wp\_query->max\_num_pages函数会得出错误结果；

**单文章页面（single.php）**

同样左右两栏。

左边用函数替换内容（文章分类、作者链接、时间、正文），并在<?php the_content();?>后手动加入标签。

右边和主页差不多，多了个内推网，直接COPY/PASTE。

底部根据标签生成“相关文章”列表，最多10篇。

**头和足（header.php & footer.php）**

顶部：  
LOGO+五个分类+首页的超链接；  
投稿 & 申请报道，两个按钮合并，新建页面使用模板（tougao.php，似乎不会检测MySql命令，直接插数据库，所以要进行安全检测……）；  
登录按钮似乎不需要，直接/wp-admin/；  
当前文章大类的阴影效果，默认在“首页”上，点击分类时切换。  
搜索按钮取消了“点击切换显示/隐藏搜索框”效果，节约资源直接显示，引擎替换成了百度，在框输入关键字后回车即可在新窗口中进行搜索。

底部：  
RSS地址函数；  
“关于我们”的五个页面，新建页面写好对应的Template文件 ，填入内容；  
回到顶部的按钮。

**投稿页面（tougao.php）**

用的是百度UEditor，官网示范代码基础上**修改调用js的路径**，以及把内容框的name属性改成投稿页面参数content所需。

搜索框aubmit和textarea的CSS属性手动定制，加上!important，以免和投稿页面冲突。

TODO

正确传入作者昵称和联系方式（目前是“请在正文/附件中留下您的联系方式”）；

美化编辑器上方文字和输入框；

修改图片上传目录，php文件名考虑换成英文而不是拼音……（真的有人会在意这两个问题？）

**分类文章列表（Category.php）**

两栏，左边有个Header of list，包括分类名、分类描述和Feed图标+超链接，下面是20篇该分类文章，右边和主页一样。

只要改动一些细节：Feed URL，修改翻页导航生成函数以符合CSS，获取网页地址算出Post Query Offset什么的。

**丢失页面（404.php）**

三个背景图片另存本地，修改图片相对路径。

**作者（author.php） & 标签（tag.php）**

都和Category.php相似，在functions里给作者页新增了个get\_avator\_url。

**右栏（sidebar.php）**

右边，从上到下：

> *新闻，1个带图，5个超链接；  
> 内推网（只有主页front.php没有内推网）；  
> 关注我们，四个联系方式，四个<li>；  
> 热门文章列表。（主页有配图，其他没有）*

TODO：  
sidebar.php中热门文章列表的生成函数（看后台用PHP重写），写好后手动加入404页面「recommend-list」；  
主页中不用get sidebar()，手动删除内推网，加入热门文章配图。

**PS：这一部分中黑体字为手动进行的操作。**

**暂缓的TODO & 功能需求**

**在服务器上完成的——**

分享按钮；  
广告代码；  
主题文件中引用的js、css、Ueditor和图片相对路径修改；  
正文「<a href="http://tech2ipo.com/tag/6416" target="_blank" rel="nofollow">http://tech2ipo.com/tag/6416</a>」这种链接需处理，解决方案：根据标签名生成rul，格式：<?php site_url();?>/tag/[标签名]；  
网站静态化，只有页面改动时才更新（视服务器状况而定，独服就随意了……）；  
同步创之文章至创见：插件「<a href="http://wordpress.org/plugins/syndicate-out/" target="_blank" rel="nofollow">http://wordpress.org/plugins/syndicate-out/</a>」。（待重新发明的轮子：Cron Job抓Feed，如有定时需求则向后台插入草稿、手动定时）

**暂缓的——**

目前TECH2IPO首页的三个“最新文章列表”没有考虑文章重复问题；  
评论导入，以及从微博获取；  
评论框设计：类4chan，无需登录，填写id、邮箱（选填）和内容即可发布，（Akismet API: ec52118a754）；  
投稿抄/调用译言的“原文推荐”功能，输入稿件链接直接读取标题、作者、发布时间、摘要和内容等信息；投稿时间间隔在服务端完成，目前是本地检测Cookies，无法防御自动投稿软件；

**函数参考**

<?php the_title(); ?> 文章标题  
<?php the_permalink(); ?> 文章链接  
<?php the_author() ?> 作者名  
<?php the\_author\_posts_link(); ?> 作者超链接  
<?php the\_modified\_date(); ?> 最后修改日期  
<?php the_excerpt(); ?> 摘要  
<?php the_ID(); ?> 文章id  
<?php <catch\_that\_image(); ?> 获取文章第一张图片的URL

$curauth = (isset($\_GET[‘author\_name’])) ? get\_user\_by(‘slug’, $author\_name) : get\_userdata(intval($author));  
$curauth->display\_name; 作者信息，还有user\_url

preg\_match(‘/(?<=\/)\d+/’, $\_SERVER[‘REQUEST_URI’], $matches); 当前网页链接&当前页号

获取文章列表  
<?php  
$categories = get\_the\_category();  
$args = array( ‘posts\_per\_page’ => 20, ‘category’ => $categories[0]->cat\_ID, ‘offset’=> 0 ); #文章可以有多个分类，故有[0]，其属性有分类id，分类名cat\_name，分类描述category_description等等  
$myposts = get_posts($args);  
foreach ( $myposts as $post ) : setup_postdata( $post ); ?>  
#循环代码  
<?php endforeach; wp\_reset\_postdata();?> #循环结束

**测试**

「localhost/wp2ipo/」改成你的安装目录——

<a href="http://localhost/wp2ipo/tag/facebook" rel="nofollow">http://localhost/wp2ipo/tag/facebook</a>

<a href="http://localhost/wp2ipo/tag/%E8%B0%B7%E6%AD%8C%E5%9C%B0%E5%9B%BE" rel="nofollow">http://localhost/wp2ipo/tag/%E8%B0%B7%E6%AD%8C%E5%9C%B0%E5%9B%BE</a>

<a href="http://localhost/wp2ipo/author/nianyouyutech2ipo-com" rel="nofollow">http://localhost/wp2ipo/author/nianyouyutech2ipo-com</a>

<a href="http://localhost/wp2ipo/category/news" rel="nofollow">http://localhost/wp2ipo/category/news</a>

<a href="http://localhost/wp2ipo/" rel="nofollow">http://localhost/wp2ipo/</a>

<a href="http://localhost/wp2ipo/list/2" rel="nofollow">http://localhost/wp2ipo/list/2</a>

<a href="http://localhost/wp2ipo/business" rel="nofollow">http://localhost/wp2ipo/business</a>

**参考资料**

  1. <a href="http://www.oschina.net/code/snippet_862384_25415" target="_blank" rel="nofollow">http://www.oschina.net/code/snippet_862384_25415</a> 拼音转换
  2. <a href="http://wordpress.stackexchange.com/questions/7177/how-can-i-assign-post-a-specific-id-on-creation" target="_blank" rel="nofollow">http://wordpress.stackexchange.com/questions/7177/how-can-i-assign-post-a-specific-id-on-creation</a> 插入文章时import_id参数的作用
  3. <a href="http://wordpress.org/support/topic/retreive-first-image-from-post" target="_blank" rel="nofollow">http://wordpress.org/support/topic/retreive-first-image-from-post</a> 获取文章内容中第一张图片
  4. <a href="http://immmmm.com/wordpress-page-navigation-without-plugins.html" target="_blank" rel="nofollow">http://immmmm.com/wordpress-page-navigation-without-plugins.html</a> 翻页按钮
  5. <a href="http://wordpress.stackexchange.com/questions/110349/template-hierarchy-confused-with-index-php-front-page-php-home-php" target="_blank" rel="nofollow">http://wordpress.stackexchange.com/questions/110349/template-hierarchy-confused-with-index-php-front-page-php-home-php</a> Wordpress主题的架构（图片URL：<a href="http://codex.wordpress.org/images/1/18/Template_Hierarchy.png" target="_blank" rel="nofollow">http://codex.wordpress.org/images/1/18/Template_Hierarchy.png</a>）
  6. <a href="http://www.coolboke.com/post/wordpress-related-posts-without-plugin.html" target="_blank" rel="nofollow">http://www.coolboke.com/post/wordpress-related-posts-without-plugin.html</a> 免插件根据标签生成相关文章
  7. <a href="http://blog.csdn.net/qikexun/article/details/9137487" target="_blank" rel="nofollow">http://blog.csdn.net/qikexun/article/details/9137487</a> <a href="http://www.tuicool.com/articles/MbArqy" target="_blank" rel="nofollow">http://www.tuicool.com/articles/MbArqy</a> 免插件打造投稿页面
  8. <a href="http://fex-team.github.io/ueditor/" target="_blank" rel="nofollow">http://fex-team.github.io/ueditor/</a> UEditor官方调用代码示范
  9. <a href="http://codex.wordpress.org/Tag_Templates" target="_blank" rel="nofollow">http://codex.wordpress.org/Tag_Templates</a> “浏览该标签下的文章”模板
 10. <a href="http://codex.wordpress.org.cn/Author_Templates" target="_blank" rel="nofollow">http://codex.wordpress.org.cn/Author_Templates</a> “浏览该作者的文章”模板
 11. <a href="http://bavotasan.com/2009/how-to-list-your-most-popular-posts-in-wordpress/" target="_blank" rel="nofollow">http://bavotasan.com/2009/how-to-list-your-most-popular-posts-in-wordpress/</a> （TODO中待参考的）热门文章生成函数