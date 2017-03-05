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

* Do NOT overwrite system's python, unless you know what you're doing: Gentoo user or have full backup.
* [virtualenv](https://docs.python.org/3/library/venv.html) per project, do not `sudo pip install`, it's like `using namespace std;` + `import <bits/stdc++.h>`.
* `export PYTHONBYTECODEBASE=0` because `.pyc` are imported prior to your new `.py` files with bugs fixed.
* [Time complexity](https://wiki.python.org/moin/TimeComplexity) of common operations.
* `awesome-*` projects are mostly bullshit : if you know what library/framework to choose, you don't need them; if you don't, you can't distinguish good from bad ones(I'm not taking awesome-python personally) until you start using and reading others code.


Convenience:

* `pyflakes` + `pep8` + `pylint`(and more) to check your code, a [pre-commit hook](https://www.stavros.io/posts/more-pep8-git-hooks/) could be overkill, but integrating them with CI should be trivial.
* Meta programming: [jsondatabase](https://github.com/gunthercox/jsondb/blob/master/jsondb/db.py#L139) shows good use of `__get__`(though can't really appreciate this [silver bullet](https://github.com/gunthercox/jsondb/blob/master/jsondb/db.py#L52) method)
* [Pythonic code review](https://access.redhat.com/blogs/766093/posts/2802001) introduces pythonic/convenient syntax like `namedtuple`, and more importantly, what to focus when reviewing code.


Traps/gotchas:

* [This post](http://sopython.com/wiki/Common_Gotchas_In_Python) is a good summary.
* Gotchas are caused by too much expectations(why this does not work?!), try to have minimal expectations and write codes that you can understand and you're sure they will work.


Debug:

* `logging` module is your friend, read [official document](https://docs.python.org/3/library/logging.html) and some [best practices](https://fangpenlin.com/posts/2012/08/26/good-logging-practice-in-python/): proper logging level, `__name__` as logger name ...
* Simple script can be debugged using `ipdb` or `pudb`.
* [line_profiler](https://github.com/rkern/line_profiler) is by far the most intuitive profiling tool I've used.
* [This article](https://blog.ionelmc.ro/2013/06/05/python-debugging-tools/) introduced segfaults handler, monitoring and more debugging tools.



