range = (1..100)
sum_of_squares = range.map {|num| num ** 2}.inject &:+
square_of_sum = range.inject(&:+) ** 2

p "The answer is: #{square_of_sum - sum_of_squares}"
