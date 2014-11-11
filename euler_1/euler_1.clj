(defn three-or-five-mult [n]
  (if (some #(= 0 %) [(mod n 5) (mod n 3)])
    n
    nil
    ))

(def sum-three-fives-multiples-below
  (fn [n]
    (loop [current 1 acc 0]
      (if (= n current)
        acc
        (recur
         (inc current)
         (+ acc (or (three-or-five-mult current) 0)))))))

(defn sum-three-fives-multiples-below2 [n]
  (let [
        threes (range 0 n 3)
        fives (range 0 n 5)
        distinct-multiples (distinct (concat threes fives))
        ]
    (reduce + distinct-multiples)
    ))

(use 'clojure.set)
(defn sum-three-fives-multiples-below3 [n]
  (let [
        threes (range 0 n 3)
        fives (range 0 n 5)
        distinct-multiples (set (union threes fives))
        ]
    (reduce + distinct-multiples)))

(println "The answer is: " (sum-three-fives-multiples-below 1000))
(println "The answer is: " (sum-three-fives-multiples-below2 1000))
(println "The answer is: " (sum-three-fives-multiples-below3 1000))
