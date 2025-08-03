fn main() {
    let listener = TcpListener::bind("0.0.0.0:4000").unwrap();
    for stream in listener.incoming() {
        thread::spawn(|| {
            let stream = stream.unwrap();
            let _ = panic::catch_unwind(|| {
                panic!("Opps")
            });
        });
    }
}
