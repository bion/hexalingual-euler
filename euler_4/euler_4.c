#include <stdio.h>

int checkPalindrome(int x)
{
    int reversed, remainder, temp;
    reversed = 0;
    temp = x;

    while (temp != 0) {
        remainder = temp % 10;
        reversed = reversed * 10 + remainder;
        temp /= 10;
    }

    return reversed == x;
}

int main()
{
    int num;
    int answer = 0;

    for (int i = 999; i > 99; --i) {
        for (int j = 999; j > 99; --j) {
            num = i * j;
            if (checkPalindrome(num) && answer < num) {
                answer = num;
            }
        }
    }

    printf("The answer is: %d\n", answer);

    return 0;
}
