#!/usr/bin/sclang

bench {
  var checkPalindrome, answer = 0, num;

  checkPalindrome = { |num|
    var stringed = asString(num);
    stringed.reverse == stringed;
  };

  (100 .. 999).do { |i|
    (100 .. 999).do { |j|
      num = i * j;
      if (checkPalindrome.(num) && (answer < num)) {
        answer = num;
      };
    };
  };

  postln("The answer is: " ++ answer);

};

0.exit;
