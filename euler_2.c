#include <stdio.h>

int main()
{
    int sum = 0;
    int prev = 1;
    int fib = 2;
    int next_prev;

    while (fib < 4000000) {
        if (fib % 2 == 0)
            sum += fib;

        next_prev = fib;
        fib += prev;
        prev = next_prev;
    }


    printf("The answer is: %d\n", sum);

    return 0;
}
