#!/usr/bin/sclang

bench {
  var num = 600851475143.0, // sc can't handle big ints
  test = 2;

  while ({test**2 < num}, {

    if (num % test == 0, {
      num = (num / test);
    });

    test = test + 1;
  });

  postln("The answer is " ++ num);
};

0.exit;
