fn divide(dividend: i32, divisor: i32) -> Result<i32, String> {
    if divisor == 0 {
        Err("Division durch Null ist nicht erlaubt!".to_string())
    } else {
        Ok(dividend / divisor)
    }
}

fn run() -> Result<(), String> {
    let result = divide(10, 0)?;
    println!("Ergebnis: {}", result);
    Ok(())
}

fn main() {
    if let Err(e) = run() {
        println!("Fehler: {}", e);
    }
}
