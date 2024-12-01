use advent_of_code_2024::*;
use std::time::Instant;

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
        _ => println!("Day {} not implemented yet", day),
    }
    println!("Time: {:?}", start.elapsed());
}
