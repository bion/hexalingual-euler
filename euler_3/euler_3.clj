(def number 600851475143)

(defn pow [base exp]
  (reduce * (repeat exp base)))

(time (loop [high number
        test 2]
   (if (< high (pow test 2))
     (println "The answer is:" high)
     (let [next-test (inc test)
           next-high (if (zero? (mod high test)) (/ high test) high)]
       (recur next-high next-test)
       ))))
