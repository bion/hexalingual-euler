sum = 0
1000.times.each { |i| sum += i if i % 3 == 0 || i % 5 == 0 }
p "The answer is: #{sum}"
