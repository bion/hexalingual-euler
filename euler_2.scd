#!/usr/bin/sclang

bench {
  var sum = 0, prev = 1, fib = 2, next_prev;

  while ({fib < 4000000}, {
    if (fib % 2 == 0) {
      sum = sum + fib;
    };

    next_prev = fib;
    fib = fib + prev;
    prev = next_prev;
  });

  postln("The answer is: " ++ sum);
};

0.exit;
