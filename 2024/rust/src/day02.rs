use crate::read_input;

fn is_safe_report(levels: &[i64]) -> bool {
    if levels.len() < 2 {
        return false;
    }

    #[derive(PartialEq)]
    enum Direction {
        Up,
        Down,
    }

    let mut direction: Option<Direction> = None;

    for window in levels.windows(2) {
        let diff = window[1] - window[0];
        if diff.abs() > 3 || diff.abs() == 0 {
            return false;
        }

        match direction {
            None => {
                direction = Some(if diff > 0 {
                    Direction::Up
                } else {
                    Direction::Down
                });
            }
            Some(Direction::Up) if diff < 0 => return false,
            Some(Direction::Down) if diff > 0 => return false,
            _ => {}
        }
    }
    true
}

pub fn part1() -> i64 {
    let input = read_input(2);

    let mut sum = 0;

    for line in input.lines() {
        let levels: Vec<i64> = line
            .split_whitespace()
            .map(|x| x.parse::<i64>().unwrap())
            .collect();

        if is_safe_report(&levels) {
            sum += 1;
        }
    }

    sum
}

pub fn part2() -> i64 {
    let input = read_input(2);

    let mut sum = 0;

    for line in input.lines() {
        let levels: Vec<i64> = line
            .split_whitespace()
            .map(|x| x.parse::<i64>().unwrap())
            .collect();

        // Check if the report is already safe
        if is_safe_report(&levels) {
            sum += 1;
            continue;
        }
        // Try to remove each level and check if the report is safe
        for i in 0..levels.len() {
            let mut levels = levels.clone();
            levels.remove(i);
            if is_safe_report(&levels) {
                sum += 1;
                break;
            }
        }
    }

    sum
}
