int main() {
    std::ifstream file("does_not_exist.txt");
    if (!file) {
        throw std::runtime_error("Fehler beim Oeffnen der Datei: Datei nicht gefunden");
    }
    return 0;
}
