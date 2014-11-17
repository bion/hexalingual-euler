(defn palindrome? [x]
  (= (reverse (str x)) (seq (str x))))

(defn find-highest-palindrome [input-seq]
  (reduce (fn [acc product]
            (if (and (palindrome? product) (> product acc))
              product acc))
          input-seq))

(def digits (range 100 1000))

(time (println "The answer is:" (find-highest-palindrome (seq (for [a digits b digits] (* a b))))))
