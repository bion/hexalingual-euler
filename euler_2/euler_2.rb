sum = 0
prev = 1
fib = 2

while fib < 4_000_000
  sum += fib if fib % 2 == 0

  next_prev = fib
  fib += prev
  prev = next_prev
end

p "The answer is: #{sum}"
