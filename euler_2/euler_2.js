var sum = 0, prev = 1, fib = 2, nextPrev;

while (fib < 4000000) {
    if (fib % 2 === 0) {
        sum += fib;
    }

    nextPrev = fib;
    fib += prev;
    prev = nextPrev;
};

console.log("The answer is: " + sum);
