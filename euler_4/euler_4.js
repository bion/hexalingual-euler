function reverseString(s) { return (s).split("").reverse().join("") };

function checkPalindrome(n) {
    var stringed = n + "";
    var reversed = reverseString(stringed);

    return stringed === reversed;
};

var answer = 0;
var num;

for (var i = 999; i > 99; --i) {
    for (var j = 999; j > 99; --j) {
        num = i * j;
        if (checkPalindrome(num) && answer < num) {
            answer = num;
        }
    }
}

console.log("The answer is: " + answer);
