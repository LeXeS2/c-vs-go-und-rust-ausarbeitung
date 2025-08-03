#include <stdio.h>

int main() {
    int *p;       // nicht initialisiert
    *p = 10;      // undefiniertes Verhalten
    printf("%d\n", *p);
    return 0;
}