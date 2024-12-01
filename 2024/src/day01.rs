use crate::read_input;
use regex::Regex;
use std::collections::HashMap;

fn parse_numbers(input: &str, re: &Regex) -> Option<(i64, i64)> {
    re.captures(input).map(|caps| {
        let first = caps[1].parse::<i64>().unwrap();
        let second = caps[2].parse::<i64>().unwrap();
        (first, second)
    })
}

pub fn part1() -> i64 {
    let input = read_input(1);

    let mut list_l = Vec::new();
    let mut list_r = Vec::new();

    let re = Regex::new(r"(\d+)\s+(\d+)").unwrap();

    for line in input.lines() {
        if let Some((left, right)) = parse_numbers(line, &re) {
            list_l.push(left);
            list_r.push(right);
        }
    }

    assert_eq!(list_l.len(), list_r.len());

    list_l.sort();
    list_r.sort();

    let mut sum = 0;

    for i in 0..list_l.len() {
        sum += (list_l[i] - list_r[i]).abs();
    }

    sum
}

pub fn part2() -> i64 {
    let input = read_input(1);

    let mut frequencies_l = HashMap::new();
    let mut frequencies_r = HashMap::new();

    let re = Regex::new(r"(\d+)\s+(\d+)").unwrap();

    for line in input.lines() {
        if let Some((left, right)) = parse_numbers(line, &re) {
            *frequencies_l.entry(left).or_insert(0) += 1;
            *frequencies_r.entry(right).or_insert(0) += 1;
        }
    }

    let mut sum = 0;

    for (left, count_l) in &frequencies_l {
        let count_r = frequencies_r.get(left).unwrap_or(&0);
        sum += left * count_l * count_r
    }

    sum
}
