func main() {
    file, err := os.Open("example.txt")
    if err != nil {
       fmt.Println("Fehler beim Öffnen der Datei:", err)
       return
    }
defer file.Close()
doStuff()
}
