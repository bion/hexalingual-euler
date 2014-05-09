def check_palindrome num
  stringed = num.to_s
  stringed.reverse == stringed
end

answer = 0

(100..999).each do |i|
  (100 .. 999).each do |j|
    num = i * j
    answer = num if check_palindrome(num) && answer < num
  end
end

p "The answer is #{answer}"
