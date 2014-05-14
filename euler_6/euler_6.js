var sumOfSquares = 0, squareOfSum = 0, answer;

for (i = 1; i < 101; i++) {
    sumOfSquares += i * i;
    squareOfSum += i;
}

squareOfSum = squareOfSum * squareOfSum;

console.log("The answer is: " + (squareOfSum - sumOfSquares));
