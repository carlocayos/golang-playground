//hello.c
#include <stdio.h>

#define CARLO_MACRO "Carlo Macro"
#define CARLO_TRUE_FALSE \
    defined(1)

void Hello(){
    printf("Hello world from hello.c\n");
    printf("CARLO_MACRO = %s\n", CARLO_MACRO);
    printf("__APPLE__ = %d\n", __APPLE__);
    printf("__SIZE_FMTu__ = %s\n", __SIZE_FMTu__);
    printf("CARLO_TRUE_FALSE = %d\n", CARLO_TRUE_FALSE);
}

int main(int argc, char* argv[]) {
    Hello();
}