// #!/usr/bin/sclang

bench {
  var answer = 0, test = 2520, testNum;

  testNum = { |num|
    var returnVal = true;
    (2..20).do { |i|
      if (num % i != 0) {
        returnVal = false;
      };
    };
    returnVal;
  };

  while ({answer == 0}) {
    if (testNum.(test)) {
      answer = test;
    };
    test = test + 20;
  };

  postln("The answer is: " ++ answer);

};

0.exit;
