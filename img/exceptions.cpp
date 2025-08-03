int divide(int dividend, int divisor) {
    if (divisor == 0) {
        throw std::runtime_error("Division durch Null ist nicht erlaubt!");
    }
    return dividend / divisor;
}


int main() {
    try {
        int result = divide(10, 0);
        std::cout << "Ergebnis: " << result << std::endl;
    } catch (const std::runtime_error& e) {
        std::cerr << "Fehler: " << e.what() << std::endl;
    }
    return 0;
}