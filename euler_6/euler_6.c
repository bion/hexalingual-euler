#include <stdio.h>

int main() {

    long sum_of_squares = 0;
    long square_of_sum = 0;
    int i;

    for (i = 1; i < 101; ++i) {
        sum_of_squares += i * i;
        square_of_sum += i;
    }

    square_of_sum = square_of_sum * square_of_sum;

    printf("The answer is: %d\n", square_of_sum - sum_of_squares);

    return 0;
}
