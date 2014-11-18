(time (let [sum-of-first-100 (* 50 101)
            square-of-sum (* sum-of-first-100 sum-of-first-100)
            sum-of-squares (reduce + (map #(* % %) (range 1 101)))]
        (- square-of-sum sum-of-squares)))
