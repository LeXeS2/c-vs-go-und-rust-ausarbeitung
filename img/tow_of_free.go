
package main

import ( "fmt"
    "runtime"
)
func main() {
    data := make([]int, 10)
    fmt.Println(" 10 int resievert:", data)
    // Referenz auf Slice entfernen, damit es für die GC freigegeben wird
    data = nil
    fmt.Println("10 resieverend int entfernt, GC kann Speicher freigeben")
    // Garbage Collector ausführen
    runtime.GC()
    fmt.Println("GC ausgelöst")
}

