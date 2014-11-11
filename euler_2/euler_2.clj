(loop [before-last-fib 1
       last-fib 2
       accumulator 0]
  (if (> last-fib 4e6)
    (print "The answer is:" accumulator)
    (let [next-fib (+ before-last-fib last-fib)
          add-last-fib? (even? last-fib)
          next-accumulator (+ accumulator (if add-last-fib? last-fib 0))]
        (recur last-fib next-fib next-accumulator))))
