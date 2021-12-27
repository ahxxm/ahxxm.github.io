---
title: Advent of Code 2015 Writeup -- Clojure
author: ahxxm
layout: post
permalink: /161.moew/
categories:
  - Clojure
---

[Advent of Code](https://adventofcode.com/2019/about) is a set of *small programming puzzles*, like [Project-Euler](https://projecteuler.net) and [CodeJam](https://codingcompetitions.withgoogle.com/codejam).

I used Clojure because it is a practical language: expressive, composable, fun to write, easy to understand, fast, comes with power standard library, intuitive CSP library, usually does not support goto or early return(both I liked very much and considered useful).

<!--more-->

Preliminary modeling is assumed, the writeup will focus on my own findings: what function makes code cleaner, how to generalize a

## Day 1

Warm up puzzle.

```Clojure
;; lein projects come with `resources` directory, slurp makes REPL clean
(def i1 (slurp "resources/2015/i1.txt"))

;; part 1 does not care about input order
(let [f (frequencies i1)]
  (- (f \() (f \))))

;; part 2 does
(loop [[[i c] & r] (map-indexed vector i1)
       f 0]
  (let [d (if (= c \() 1 -1)
        -f (+ f d)]
    (if (= -1 f)
      i
      (recur r -f))))

```
`map-indexed` is equivelant to Python's `for i,v in enumerate(coll)`.

`[[i c] & r]` is called [destructuring](https://blog.brunobonacci.com/2014/11/16/clojure-complete-guide-to-destructuring/), i(ndex), c(har) and r(est)-as-collection.

Be careful with loop terminating condition, here we are ensured that f(loor) will be -1 before consuming all input, otherwise insert `(if (nil? i))` check just before `let`.

## Day 2

Destructure and reduce.

```clojure
;; 2x3x6 => [2 3 6]
;; [i j k] (map bigdec (clojure.string/split l #"x"))

(let [lines (read-lines "i2.txt")]
  (reduce + (map f lines))) ;; f(line) => value
```

## Day 3

What if we have more than 2 Robo-Santa?

`take-nth` still works.

```clojure
(take-nth 2 [1 2 3 4 5]) => (1 3 5)
(take-nth 2 (rest [1 2 3 4 5])) => (2 4)

;; Euclidean coordinates
(mapv + [0 1] [99 99]) => [99 100]
```

## Day 4

[Laziness](http://clojure-doc.org/articles/language/laziness.html):

- both `for` and `(iterate inc 1)` are lazy
- when `when` condition is true, comprehension returns first i then stops because of `first`

```clojure
(let [k  "laziness"
      is (for [i (iterate inc 1)
               :let [h (md5 (str k i))]
               :when (= (subs h 0 5) "00000")] ;; or 6 "000000"
           i)]
  (first is)) => 183118

(take 3 is) => (183118 1544928 1571333)
```

## Day 5

`partition` and `some`:

```clojure
(partition 3 1 (map str "abade")) => (("a" "b" "a") ("b" "a" "d") ("a" "d" "e"))
(some #(= (first %) (last %)) (partition 3 1 (map str "abade"))) => true
```

## Day 6

This is a simplified version of 2021-day22, the map is small enough to operate on each cell.

Combine comprehension and anonymous function:

```clojure
;; function f
(cond
  (clojure.string/includes? s "turn on") (fn [x] (inc x))
  (clojure.string/includes? s "turn off") (fn [x] (max 0M (dec x)))
  (clojure.string/includes? s "toggle") (fn [x] (+ x 2)))

;; force realization lazy-for
(doall
 (for [x (range x1 (inc x2))
       y (range y1 (inc y2))]
   (update-v! m [y x] f)))

;; m is an atom of the map(nested vector)
;; get-in and assoc-in works similarly
;; (assoc-in [[0 1 2] [3 4 5]] [0 0] 1) => [[1 1 2] [3 4 5]]
(defn update-v!
  [m l f]
  (let [v (get-in @m l)]
    (swap! m assoc-in l (f v))))
```
