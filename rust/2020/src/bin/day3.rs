use std::collections::{HashMap, HashSet};

mod utils;

fn main() {}

struct TreesMap {
    map: HashMap<usize, HashSet<usize>>,
    max_x: usize,
    max_y: usize,
}

impl TreesMap {
    fn parse(&mut self, input: &str) {
        let mut y = 0;
        input.split("\n").for_each(|line| {
            let mut set = HashSet::new();
            let mut x: usize = 0;
            for c in line.chars() {
                if c == '#' {
                    set.insert(x);
                }
                x += 1
            }
            self.max_x = line.len();
            self.map.insert(y, set);
            y += 1;
        });
        self.max_y = y;
    }

    fn has(&self, x: usize, y: usize) -> bool {
        return self.map.contains_key(&y) &&
            self.map.get(&y).unwrap().contains(&x);
    }

    fn slope(&self, dx: usize, dy: usize) -> usize {
        let mut res = 0;

        let mut x = 0;
        let mut y = 0;
        loop {
            if self.has(x, y) {
                res += 1;
            }
            x = (x + dx) % self.max_x;
            y += dy;
            if y >= self.max_y {
                break;
            }
        }

        return res;
    }
}

fn part1(input: &str) -> usize {
    let mut trees_map = TreesMap {
        map: HashMap::new(),
        max_x: 0,
        max_y: 0,
    };
    trees_map.parse(input);
    return trees_map.slope(3, 1);
}

fn part2(input: &str) -> usize {
    let mut trees_map = TreesMap {
        map: HashMap::new(),
        max_x: 0,
        max_y: 0,
    };
    trees_map.parse(input);

    let slopes = vec!{
        trees_map.slope(1,1),
        trees_map.slope(3,1),
        trees_map.slope(5,1),
        trees_map.slope(7,1),
        trees_map.slope(1,2),
    };

    let mut res = 1;
    for i in 0..5 {
        res *= slopes.get(i).unwrap();
    }

    return res;
}

#[cfg(test)]
mod tests {
    use crate::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(3).as_str()), 7);
        assert_eq!(part1(input(3).as_str()), 145);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(sample(3).as_str()), 336);
        assert_eq!(part2(input(3).as_str()), 3424528800);
    }
}
