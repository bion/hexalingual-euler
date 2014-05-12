answer = nil
test = 2520

def test_num num
  (2...20).each do |i|
    return false if num % i != 0
  end
  true
end

while !answer
  answer = test if test_num test
  test = test + 20
end

puts "The answer is: #{answer}"
