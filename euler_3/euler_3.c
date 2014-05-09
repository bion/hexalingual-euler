#include <stdio.h>
#include <math.h>

int main()
{
    long num = 600851475143;
    int test = 2;

    while (pow(test, 2) < num) {
        if (num % test == 0)
            num = num / test;
        test += 1;
    }

    printf("The answer is %d\n", num);

    return 0;
}
