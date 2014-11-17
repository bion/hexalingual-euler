(defn lcm [a b]
  (/ (* a b)
     (loop [x a
            y b]
       (if (zero? y)
         x
         (recur y (mod x y))))))

(time (reduce lcm (range 1 21)))
