int main() {
    FILE *f = fopen("nicht_existierend.txt", "r");
    if (f == NULL) {
        printf("Fehler: %s\n", strerror(errno));
    }
    return 0;
}