fn main() {}

mod utils;

fn part1(input: &str) -> i32 {
    let mut res = 0;

    let mut yes_count = 0;
    let mut yeses: [u8; 26] = [0; 26];

    for line in input.split("\n") {
        if line == "" {
            res += yes_count;
            yes_count = 0;
            yeses = [0; 26];
            continue;
        }
        for c in line.as_bytes() {
            let i  = (c-97) as usize;
            if yeses[i] == 0 {
                yeses[i] = 1;
                yes_count += 1;
            }
        }
    }

    return res + yes_count;
}

fn part2(input: &str) -> i32 {
    let mut res = 0;

    let mut yes_count = 0;
    let mut people = 0;
    let mut yeses: [u8; 26] = [0; 26];

    for line in input.split("\n") {
        if line == "" {
            res += yes_count;
            people = 0;
            yeses = [0; 26];
            continue;
        }
        yes_count = 0;
        people += 1;
        for c in line.as_bytes() {
            let i  = (c-97) as usize;
            yeses[i] += 1;
            if yeses[i] == people {
                yes_count += 1;
            }
        }
    }

    return res + yes_count;
}

#[cfg(test)]
mod tests {
    use crate::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(6).as_str()), 11);
        assert_eq!(part1(input(6).as_str()), 6630);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(sample(6).as_str()), 6);
        assert_eq!(part2(input(6).as_str()), 3437);
    }
}
