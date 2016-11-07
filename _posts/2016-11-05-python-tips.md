---
title: Python Developing Notes
author: ahxxm
layout: post
permalink: /153.moew/
categories:
  - Python
---


This post will try not to repeat Code Complete and PEP-8 and talk about Python specifically.


Common sense:
- Do not overwrite system's python, unless you know what you're doing: Gentoo user or have full backup
- virtualenv per project, do not sudo pip install without virtualenv, it's like using namespace std + import <bits/stdc++.h>
- Disable PYTHONBYTECODE xxx because pyc are imported before py
- Time complexity of common operations
- awesome-* projects are mostly bullshit : if you know what library/framework to choose, you don't need them; if you don't, you can't distinguish good from bad ones(I'm not taking awesome-python personally) until you start using and reading others code.


Convenient:
- use pyflakes+pep8+pylint to check your code, a pre-commit hook https://www.stavros.io/posts/more-pep8-git-hooks/ could be overkill, long one: https://www.atlassian.com/git/tutorials/git-hooks
- meta programming: jsondatabase shows good use of __get__ https://github.com/gunthercox/jsondb/blob/master/jsondb/db.py#L139 (though can ntt really appreciate this https://github.com/gunthercox/jsondb/blob/master/jsondb/db.py#L52 silver-bullet method)

traps/gotchas:
- http://sopython.com/wiki/Common_Gotchas_In_Python a good summary.
- Gotchas are caused by too much expectations(why this does not work?!), try to have minimal expectations and write codes that you can understand and you're sure they will work.

debug:
- logging is your friend, read official document and some short introduction: use proper level, and __name__ as logger name for example https://fangpenlin.com/posts/2012/08/26/good-logging-practice-in-python/
- simple script can be debugged using ipdb/pudb
- line_profiler is by far the most convenient profiling tool I've used for scripts -- in the end, mid to large projects does not need profiling, they will be re-written in other languages.
- this article introduces https://blog.ionelmc.ro/2013/06/05/python-debugging-tools/  segfaults handler, monitoring and more.

