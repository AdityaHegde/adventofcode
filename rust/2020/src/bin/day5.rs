fn main() {}

mod utils;

fn parse_pass(pass: &str) -> i32 {
    let mut num = 0;
    let pass_bytes = pass.as_bytes();
    for i in 0..pass_bytes.len() {
        num *= 2;
        match pass_bytes[i] {
            66 => {
                num += 1
            }
            82 => {
                num += 1
            }
            _ => {}
        };
    }
    return num;
}

fn part1(input: &str) -> i32 {
    let mut res = 0;
    input.split("\n").for_each(|line| {
        let num = parse_pass(line);
        if res < num {
            res = num;
        }
    });
    return res;
}

fn part2(input: &str) -> i32 {
    let mut res = 0;
    let mut min = 10000;
    let mut c = 0;
    input.split("\n").for_each(|line| {
        let num = parse_pass(line);
        if min > num {
            min = num;
        }
        res += num;
        c += 1;
    });
    res -= (min-1) * c;
    return (c + 1) * (c + 2) / 2 - res + min - 1;
}

#[cfg(test)]
mod tests {
    use crate::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(5).as_str()), 820);
        assert_eq!(part1(input(5).as_str()), 922);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(input(5).as_str()), 747)
    }
}
