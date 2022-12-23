use std::collections::{HashMap, HashSet};

mod utils;

fn main() {}

fn part1(input: &str, sum: i32) -> i32 {
    let mut nums: HashSet<i32> = HashSet::new();
    let mut res = 0;
    input.split("\n").for_each(|line| {
        let num = line.parse::<i32>().unwrap();
        let diff = sum - num;
        if nums.contains(&diff) {
            res = num * diff;
        }
        nums.insert(num);
    });
    return res;
}

fn part2(input: &str, sum: i32) -> i32 {
    let mut nums_map: HashMap<i32, (i32, i32)> = HashMap::new();
    let mut res = 0;
    let nums = input.split("\n").map(|line| line.parse::<i32>().unwrap()).collect::<Vec<i32>>();
    for i in 0..nums.len()-1 {
        for j in i+1..nums.len()-1 {
            let diff = sum - (nums[i] + nums[j]);
            nums_map.insert(diff, (nums[i], nums[j]));
        }
        if nums_map.contains_key(&nums[i]) {
            let (other_num1, other_num2) = nums_map.get(&nums[i]).unwrap();
            res = nums[i] * other_num1 * other_num2;
        }
    }
    return res
}

#[cfg(test)]
mod tests {
    use crate::utils::*;
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(part1(sample(1).as_str(), 2020), 514579);
        assert_eq!(part1(input(1).as_str(), 2020), 802011);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(sample(1).as_str(), 2020), 241861950);
        assert_eq!(part2(input(1).as_str(), 2020), 248607374);
    }
}
