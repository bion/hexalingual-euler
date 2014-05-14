#!/usr/bin/sclang

bench {
  var sumOfSquares = 0, squareOfSum = 0, answer;

  for (1, 100, { |i|
    sumOfSquares = sumOfSquares + (i * i);
    squareOfSum  = squareOfSum + i;
  });

  squareOfSum = squareOfSum * squareOfSum;

  postln("The answer is: " ++ (squareOfSum - sumOfSquares));
};

0.exit;