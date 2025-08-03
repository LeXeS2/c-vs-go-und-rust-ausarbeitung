int divide(int divident, int divisor, int* result) {
    if (divisor == 0) {
        return -1;
    }
    *result = divident / divisor;
    return 0;
}