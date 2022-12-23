use regex::Regex;

mod utils;

fn main() {}

fn parse_line(line: &str) -> (usize, usize, u8, &[u8]) {
    let parser_re: Regex = Regex::new(r"([0-9]*)-([0-9]*) ([a-z]): ([a-z]*)").unwrap();

    let captures = parser_re.captures(line).unwrap();
    let min_char = captures.get(1).unwrap().as_str().parse::<usize>().unwrap();
    let max_char = captures.get(2).unwrap().as_str().parse::<usize>().unwrap();
    let pass_char = captures.get(3).unwrap().as_str().as_bytes()[0];
    let password = captures.get(4).unwrap().as_str().as_bytes();

    return (min_char, max_char, pass_char, password);
}

fn part1(input: &str) -> usize {
    input.split("\n").filter(|line| {
        let (min_char, max_char, pass_char, password) = parse_line(line);
        let mut char_count = 0;
        for pc in password {
            if *pc == pass_char {
                char_count += 1;
            }
        }
        return char_count >= min_char && char_count <= max_char
    }).count()
}

fn part2(input: &str) -> usize {
    input.split("\n").filter(|line| {
        let (min_char, max_char, pass_char, password) = parse_line(line);
        return (password[min_char-1] == pass_char && password[max_char-1] != pass_char) ||
            (password[min_char-1] != pass_char && password[max_char-1] == pass_char)
    }).count()
}

#[cfg(test)]
mod tests {
    use super::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(2).as_str()), 2);
        assert_eq!(part1(input(2).as_str()), 447);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(sample(2).as_str()), 1);
        assert_eq!(part2(input(2).as_str()), 249);
    }
}
