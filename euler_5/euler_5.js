var answer = 0, test = 2520;

function test_num(num) {
    for (i=2; i < 21; i = i + 1) {
        if (num % i != 0) {
            return false;
        }
    }
    return true;
}

while (!answer) {
    if (test_num(test)) {
        answer = test;
    }
    test += 20;
}

console.log("The answer is: " + answer);
