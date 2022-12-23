use std::collections::HashSet;
use regex::{Captures, Regex};

mod utils;

fn main() {}

fn part1(input: &str) -> i32 {
    let valid: HashSet<&str> = vec!("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid").into_iter().collect();
    let extract_re: Regex = Regex::new(r"([a-z]*):(.*?)(?: |$)").unwrap();

    let mut res = 0;
    let mut valid_count = 0;
    for line in input.split("\n") {
        if line == "" {
            if valid_count == 7 {
                res += 1;
            }
            valid_count = 0;
            continue;
        }
        for cap in extract_re.captures_iter(line) {
            if valid.contains(cap.get(1).unwrap().as_str()) {
                valid_count += 1;
            }
        }
    }
    if valid_count == 7 {
        res += 1;
    }

    return res;
}

fn validate(field: &str, value: &str) -> bool {
    match field {
        "byr" => {
            let val = value.parse::<i32>().unwrap();
            return val >= 1920 && val <= 2002;
        }
        "iyr" => {
            let val = value.parse::<i32>().unwrap();
            return val >= 2010 && val <= 2020;
        }
        "eyr" => {
            let val = value.parse::<i32>().unwrap();
            return val >= 2020 && val <= 2030;
        }
        "hgt" => {
            let re = Regex::new(r"([0-9]*)(in|cm)$").unwrap();
            let cap_opt: Option<Captures> = re.captures(value);
            if cap_opt.is_none() {
                return false;
            }
            let cap = cap_opt.unwrap();
            if cap.len() != 3 {
                return false;
            }
            let hgt = cap[1].parse::<i32>().unwrap();
            return if cap[2].eq("in") {
                hgt >= 59 && hgt <= 76
            } else {
                hgt >= 150 && hgt <= 193
            }
        }
        "hcl" => {
            let re = Regex::new(r"^#[0-9a-f]{6}$").unwrap();
            return re.is_match(value);
        }
        "ecl" => {
            return value == "amb" || value == "blu" || value == "brn" ||
                value == "gry" || value == "grn" || value == "hzl" || value == "oth"
        }
        "pid" => {
            let re = Regex::new(r"^[0-9]{9}$").unwrap();
            return re.is_match(value);
        }
        _ => true
    }
}

fn part2(input: &str) -> i32 {
    let extract_re: Regex = Regex::new(r"([a-z]*):(.*?)(?: |$)").unwrap();

    let mut res = 0;
    let mut valid_count = 0;
    for line in input.split("\n") {
        if line == "" {
            if valid_count == 7 {
                res += 1;
            }
            valid_count = 0;
            continue;
        }
        for cap in extract_re.captures_iter(line) {
            if validate(cap.get(1).unwrap().as_str(), cap.get(2).unwrap().as_str()) {
                valid_count += 1;
            } else {
                println!("{} {}", cap.get(1).unwrap().as_str(), cap.get(2).unwrap().as_str())
            }
        }
    }
    if valid_count == 7 {
        res += 1;
    }

    return res;
}

#[cfg(test)]
mod tests {
    use super::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(4).as_str()), 2);
        assert_eq!(part1(input(4).as_str()), 204);
    }

    #[test]
    fn part2_test() {
        println!("{}", part2(sample(4).as_str()));
        println!("{}", part2(input(4).as_str()));
    }
}
