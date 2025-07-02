use std::sync::{Arc, Mutex};
use std::thread;

fn increment(counter: Arc<Mutex<i32>>) {
    let mut num = counter.lock().unwrap();
    *num += 1;
}

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..5 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            increment(counter);
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}