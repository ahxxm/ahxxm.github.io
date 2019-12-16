---
title: Advent of Code 2019 Writeup(Day 1-4) -- Clojure
author: ahxxm
layout: post
permalink: /161.moew/
categories:
  - Clojure
---

[Advent of Code](https://adventofcode.com/2019/about) is a set of *small programming puzzles*, like [Project-Euler](https://projecteuler.net) and [CodeJam](https://codingcompetitions.withgoogle.com/codejam).

<!--more-->

But questions after day 5 are just too long to read.

## Day 1

```Clojure
(def datastr "numbers splitted by spaces")
(def data (map #(Integer. %) (clojure.string/split data #"\n")))

;; a
(reduce + 0 (map #(-> % (/ 3) int (- 2)) data)))
  
;; b: data set might be too small for memoize
(defn mass->fuel
  [mass]
  (let [f1 (-> mass (/ 3) int (- 2))]
    (if (> f1 6)
      (+ (mass->fuel f1) f1)
      f1)))
      
(reduce + 0 (map mass->fuel data))
```

## Day 2

Be ware of "before running the program",

```Clojure
;; a
(def datastr "int,split,by,comma")
(def data (vec (map #(Integer. %) (clojure.string/split datastr #","))))

(defn calc
  [arg1 arg2]
  (let [end-idx (- (count data) 4)
        data    (assoc-in data [1] arg1)
        data    (assoc-in data [2] arg2)]
    (loop [i      0
           -data  data]
      (if (> i end-idx)
        (nth -data 0)
        (let [->i           (+ i 4)
              [op i1 i2 v3] (subvec -data i ->i)
              v1            (nth -data i1)
              v2            (nth -data i2)]
          (case op
            99 (recur (+ end-idx 1) -data)
            1 (recur ->i (assoc-in -data [v3] (+ v1 v2)))
            2 (recur ->i (assoc-in -data [v3] (* v1 v2)))))))))

(calc 12 2)

;; b
(loop [i 0
       j 0]
  (if (or (= (calc i j) 19690720) (> i 99))
    (+ (* i 100) j)
    (if (= j 99)
      (recur (inc i) 0)
      (recur i (inc j)))))
```

## Day 3

This one is quite longer..

```Clojure

(defn parse-move
  [action]
  ;; dir and moves both as int,
  ;; R->82, D->68, L 76, U 85
  [(-> action first int)
   (Integer. (re-find #"\d+" action))])

(defn move-one
  [point dir]
  (let [[x y] point]
    (case dir
      76 [(dec x) y]
      82 [(inc x) y]
      68 [x (dec y)]
      85 [x (inc y)])))

(defn move->path
  ;; (start x, start y), move in "R75" form, returns
  ;; [[finalx,finaly], [[point moves], ...]
  [start move moved]
  (let [[dir step] (parse-move move)
        paths      (transient [])
        [-x -y]    start]
    (loop [x -x
           y -y
           s step]
      (if (= s 0)
        [[x y] (persistent! paths) (+ moved step)]
        (let [moved-to    (move-one [x y] dir)
              [newx newy] moved-to
              used-step   (+ moved (inc (- step s)))]
          (conj! paths [moved-to used-step])
          (recur newx newy (dec s)))))))

(defn moves->points
  [movestr]
  (let [actions (clojure.string/split movestr #",")
        points  (atom [])
        start   (atom [0 0])
        moved   (atom 0)]
    (loop [[move & rem] actions]
      (if move
        (let [[end paths -moved] (move->path @start move @moved)]
          (reset! moved -moved)
          (reset! start end)
          (swap! points concat paths)
          (recur rem))
        @points))))

(defn abs [n] (max n (- n)))


;; a
(let [po1  (moves->points p1)
      po2  (moves->points p2)
      ins  (clojure.set/intersection (set (map first po1))
                                     (set (map first po2)))
      dis  (map #(+ (-> % first abs) (-> % second abs)) ins)]
  (apply min dis))

;; b
(let [po1  (moves->points p1)
      po2  (moves->points p2)
      ->p1 (into {} po1)
      ->p2 (into {} po2)
      ins  (clojure.set/intersection (set (map first po1))  ;; laziness
                                     (set (map first po2)))
      dis  (map #(+ (->p1 %) (->p2 %)) ins)]
  (apply min dis))
```

## Day 4

```Clojure

(defn valid-pass
  [num]
  (let [incr (atom true)
        adj  (atom false)
        ch   (atom (-> num first int))]
    (loop [[c & res] (rest num)]
      (if (nil? c)
        (and @incr @adj)
        (do
          (if (= (int c) @ch) (reset! adj true))
          (if (< (int c) @ch) (reset! incr false))
          (reset! ch (int c))
          (recur res))))))

(defn valid2
  [num]
  (let [freqs (vals (frequencies num))
        mfreq (filter #(> % 1) freqs)]
    (= (apply min mfreq) 2)))

(->> (range 124075 580769)
     (map str)
     (filter valid-pass)
     (filter valid2) ;; comment for a
     count)
```
