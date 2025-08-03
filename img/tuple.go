func divide(dividend, divisor int) (int, error) {
  if divisor == 0 {
    return 0, errors.New("Division durch Null ist nicht erlaubt")
  }
  return dividend / divisor, nil
}

func main() {
  result, err := divide(10, 0)
  if err != nil {
    fmt.Println("Fehler:", err)
   return
  } 
  fmt.Println("Ergebnis:", result)
}
