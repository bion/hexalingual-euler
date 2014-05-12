#include <stdio.h>

int test_num(int num)
{
    for (int i = 2; i < 21; ++i) {
        if (num % i != 0)
            return 0;
    }
    return 1;
}

int main()
{
    int answer = 0;
    int test = 2520;

    while (!answer) {
        if (test_num(test))
            answer = test;
        test += 20;
    }

    printf("The answer is: %d\n", answer);

    return 0;
}
