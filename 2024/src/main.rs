use advent_of_code_2024::*;
use std::time::Instant;

pub mod day01;
pub mod day02;
pub mod day03;

fn main() {
    let day = std::env::args()
        .nth(1)
        .expect("Please provide a day number")
        .parse::<u8>()
        .expect("Day must be a number");

    let start = Instant::now();
    match day {
        1 => {
            println!("Day 01");
            println!("Part 1: {}", day01::part1());
            println!("Part 2: {}", day01::part2());
        }
        2 => {
            println!("Day 02");
            println!("Part 1: {}", day02::part1());
            println!("Part 2: {}", day02::part2());
        }
        3 => {
            println!("Day 03");
            println!("Part 1: {}", day03::part1());
            println!("Part 2: {}", day03::part2());
        }
        _ => println!("Day {} not implemented yet", day),
    }
    println!("Time: {:?}", start.elapsed());
}
