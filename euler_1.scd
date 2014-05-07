#!/usr/bin/sclang

bench {
  var sum=0;

  1000.do { |i|
    if ((i % 3 == 0) || (i % 5 == 0)) {
      sum = sum + i;
    };
  };

  postln("The answer is: " ++ sum);
  0.exit;
}