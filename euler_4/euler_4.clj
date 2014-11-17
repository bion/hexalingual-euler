(defn palindrome? [x]
  (= (reverse (str x)) (seq (str x))))

(def digits (range 100 1000))

(time (println "The answer is:"
               (apply max
                (for [a digits
                      b digits
                      :while (>= a b)
                      :let [c (* a b)]
                      :when (palindrome? c)]
                  c))))
