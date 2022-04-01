use md5::{Digest, Md5};

fn main() {
    println!("Hello, world!");
}

pub fn sign(param: &str, app_sec: &str) -> String {
    let mut hasher = Md5::new();
    // process input message
    hasher.update(format!("{}{}", param, app_sec));
    // acquire hash digest in the form of GenericArray,
    // which in this case is equivalent to [u8; 16]
    format!("{:x}", hasher.finalize())
}