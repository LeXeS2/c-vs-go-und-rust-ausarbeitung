#include <stdlib.h>

int main() {
    int* data = (int*)malloc(sizeof(int) * 10);
    free(data);
    free(data); 
    return 0;
}