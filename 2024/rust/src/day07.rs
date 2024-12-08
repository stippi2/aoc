// Written by Claude 3.5 Sonnet
use crate::read_input;

#[derive(Debug)]
struct Calibration {
    result: i64,
    sequence: Vec<i64>,
}

fn concat_ints(a: i64, b: i64) -> i64 {
    let mut multiplier = 1;
    let mut temp = b;
    while temp > 0 {
        multiplier *= 10;
        temp /= 10;
    }
    a * multiplier + b
}

fn solve(result: i64, value: i64, sequence: &[i64], include_concatenation: bool) -> bool {
    if sequence.is_empty() {
        return result == value;
    }

    let next = sequence[0];
    let rest = &sequence[1..];

    solve(result, value + next, rest, include_concatenation)
        || solve(result, value * next, rest, include_concatenation)
        || (include_concatenation
            && solve(
                result,
                concat_ints(value, next),
                rest,
                include_concatenation,
            ))
}

fn parse_input(input: &str) -> Vec<Calibration> {
    input
        .lines()
        .map(|line| {
            let mut parts = line.split(": ");
            let result = parts.next().unwrap().parse().unwrap();
            let sequence = parts
                .next()
                .unwrap()
                .split_whitespace()
                .map(|n| n.parse().unwrap())
                .collect();
            Calibration { result, sequence }
        })
        .collect()
}

fn sum_valid_calibrations(calibrations: &[Calibration], include_concatenation: bool) -> i64 {
    calibrations
        .iter()
        .filter(|c| {
            !c.sequence.is_empty()
                && solve(
                    c.result,
                    c.sequence[0],
                    &c.sequence[1..],
                    include_concatenation,
                )
        })
        .map(|c| c.result)
        .sum()
}

pub fn part1() -> i64 {
    let input = read_input(7);
    let calibrations = parse_input(&input);
    sum_valid_calibrations(&calibrations, false)
}

pub fn part2() -> i64 {
    let input = read_input(7);
    let calibrations = parse_input(&input);
    sum_valid_calibrations(&calibrations, true)
}
