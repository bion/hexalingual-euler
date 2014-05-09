num = 100
test = 2

while test**2 < num
  num = num / test if num % test == 0
  test += 1
end

p "The answer is: #{num}"
