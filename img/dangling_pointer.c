#include <stdio.h>
int* dangling_pointer() {
    int num = 1;
    return &num; 
}

int main() {
    int* p = dangling_pointer();
    printf("%d\n", *p); 
    return 0;
}