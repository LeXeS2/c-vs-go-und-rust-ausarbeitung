#include <stdlib.h>
#include <stdio.h>
int main() {
    int* arr = (int*)malloc(sizeof(int) * 5);
    arr[5] = 100;
    printf("die Wert ist : %d\n", arr[5]); 
    free(arr);
    return 0;
}