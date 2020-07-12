#include <stdio.h>

#define MAGIC_NUMBER 5
#define MAGIC_TEXT \
    defined

int sum(int a, int b) {
    return a + b;
}

int sum_magic(int a, int b) {
    return a + b + MAGIC_NUMBER;
}

//int main(int argc, char* argv[]) {
//    int result = sum(1, 2);
//    printf("Result = %d\n", result);
//}