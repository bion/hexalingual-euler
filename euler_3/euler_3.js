var num = 600851475143.0, test = 2;

while (Math.pow(test, 2) < num) {
    if (num % test === 0) {
        num = num / test;
    }
    test += 1;
}

console.log("The answer is " + num);
