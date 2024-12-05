use std::fs;

#[allow(dead_code)]
fn main() {}

pub fn sample(day: u8) -> String {
    return fs::read_to_string(format!("./data/day{}/sample.txt", day)).unwrap()
}

pub fn input(day: u8) -> String {
    return fs::read_to_string(format!("./data/day{}/input.txt", day)).unwrap()
}
