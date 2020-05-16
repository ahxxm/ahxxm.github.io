---
title: "GBA火焰纹章：控制战斗结果和升级加点"
author: ahxxm
layout: post
permalink: /165.moew/
categories:
  - Game
---

火焰纹章6-8（GBA几作）有这么几个设定：人物死了就没法再用了；招募部分人物需要特定人物存活；战斗回避、暴击，升级属性成长都由Random Number Generator(RNG)决定。

RNG不是真随机，但是模拟器里光S&L不能改变结果，还要做消耗RNG的操作：一个简单消耗RNG数列的方式是让游戏计算移动轨迹，但是这需要玩家记住规律、记录High(>50)和Low(<50)、反复S&L，很费时。

玩着游戏做这些很影响体验，所以有了这篇——精准控制结果。

<!--more-->

RNG[原理](https://www.ign.com/faqs/2004/fire-emblem-random-number-generator-faq-520430)比较简单，通过几个数字和特定算法生成数列，在战斗中消耗该数列计算结果：

- 命中：取接下来2个数字的平均值，如果低于命中率代表命中
- 暴击：取接下来的1个数字，如果低于暴击率就暴击，造成3倍伤害
- 如果敌方没死就按照同样的流程计算反击的命中和暴击，根据速度[判断](https://fireemblemwiki.org/wiki/Attack_speed)追击
- 结算后如果升级了，取接下来7个数字决定7个属性加不加

## 简单场景

最简单的场景：尾刀命中、暴击杀死，升级全属性增加。

我们需要接下来的：

|命中|暴击|升级|
|:---:|:--:|--:|
|LL|L|LLLLLLL|

L是Low的缩写，代表小于判定值，即判定通过。

用Python简单暴力跑一个：

```python
import math

acc = 75 # accuracy - dodge
rr = 20 # attr raise rate
crit = 5

def nextrng(r1, r2, r3):
    i = (r3 >> 5) ^ (r2 << 11) ^ (r1 << 1) ^ (r2 >> 15)
    j = i & 0xffff
    return j

def rngsim(base):
    r1, r2, r3 = base
    result = [r3, r2, r1]

    for i in range(4, 20):
        n = nextrng(result[i-4], result[i-3], result[i-2])
        result.append(n)

    result = [math.floor(x/655.36) for x in result]
    return result[3:]

def rngok(result):
    hit = result[0] + result[1]
    if hit > 2 * acc:
        return False

    if result[2] > crit:
        return False

    # [3, 9]
    for i in range(3, 10):
        if result[i] > rr:
            return False
    return True

for i in range(1, 100):
    for j in range(1, 100):
        for k in range(1, 100):
            base = (i, j, k)
            result = rngsim(base)
            if rngok(result):
                print(result)
                print(base)
                break
```

经计算, 1 32 512作为seed符合要求，RNG table为`[1, 3, 0, 3, 9, 6, 6, 0, 6, 0, 1, 37, 12, 77, 1, 0]`

## 复杂场景

不小心走位失误，此时角色会被一刀砍死，读档又要重新玩很久，你不希望这种事情发生，此时拯救角色需要的RNG table为：

|敌方命中|反击命中|暴击|升级|
|:---:|:--:|--:|:--:|
|HH|LL|L|LLLLLLL|

修改`rngok`，依次判断`result`里的数值，再跑一遍：

```python
def rngok(result):
    # enemy miss
    if (result[0]+result[1]) / 2 > miss:
        return False
    
    # 略
    return True
```


## 实际修改

了解原理后，还需要工具来修改内存：一个[支持lua](http://tasvideos.org/LuaScripting.html)的gba模拟器。

我嫌[现成脚本](http://tasvideos.org/forum/viewtopic.php?p=302216#302216)使用过于复杂，就改了改：去掉了输入模式（想改就改lua里`memory.writeword`后数字），按M直接改RNG table，按N产生随机数填充RNG table。

```lua
while true do
    filter = 1
    local nsim = 20
    rngs = rngsim(503)
    for i = 1, nsim do
        gui.text(228, 8*(i-1), string.format("%3d", rngs[i]/655.36))
    end
    --gui.text(0,0,"Filter Mode: ")

    c = input.get()
    -- M: rng for mine, N: aNother(enemy)
    -- write rng base to control
    local rngbase = 0x03000000
    if c.M then
        memory.writeword(rngbase, 1)
        memory.writeword(rngbase+2, 1)
        memory.writeword(rngbase+4, 32)
    elseif c.N then
        -- reset
        math.randomseed(os.clock() * 1000)
        memory.writeword(rngbase, math.random(1, 65535))
        memory.writeword(rngbase+2, math.random(1, 65535))
        memory.writeword(rngbase+4, math.random(1, 65535))
    end

    for i = 1, 10, 1 do
        --gui.text(0,8+(i*8),rn[i])
        --gui.text(20,8+(i*8),op[i])
        if hit[i] == true then
            --gui.text(40,8+(i*8),"<- Hit")
        end
    end
    if rn[1] ~= "" and inputs == 0 then
        dis = emptyarray(dis,3)
        compareRN()
    end
    n = input.get()

    -- just to show it's working
    if rn[1] ~= "" then
        if dis[1] == "" then
            gui.text(0,104,"Distance: ---")
        elseif dis[1] ~= "" and dis[2] == "" then
            gui.text(0,104,"Distance: " .. dis[1])
        elseif dis[1] ~= "" and dis[2] ~= "" and dis[3] == "" then
            gui.text(0,104,string.format("Distance: %d - %d",dis[1],dis[2]))
        else

            gui.text(0,104,string.format("Distance: %d - %d - %d",dis[1],dis[2],dis[3]))

        end
    end
    emu.frameadvance()
end
```
