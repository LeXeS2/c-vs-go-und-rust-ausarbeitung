func readFile() error {
    _, err := os.Open("does_not_exist.txt")
    if err != nil {
       return fmt.Errorf("Fehler beim Oeffnen der Datei: %w", err)
    }
    return nil
}
