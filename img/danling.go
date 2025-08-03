package main
import "fmt"
func dangling_pointer() *int {
    num := 1
    return &num                                    
func main() {
    p := dangling_pointer()
    fmt.Println(*p)                            
}
