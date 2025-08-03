fn dangling_pointer() -> Box<i32> {
    let num = Box::new(1);  
    num                     
}

fn main() {
    let p = dangling_pointer();
    println!("{}", p);  
}