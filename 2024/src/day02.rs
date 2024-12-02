use crate::read_input;

fn is_safe_report(mut levels: Vec<i64>) -> bool {
    if levels.len() < 2 {
        return false;
    }

    let mut last_num = levels.remove(0);

    #[derive(PartialEq)]
    enum Direction {
        Up,
        Down,
        None,
    }

    let mut direction = Direction::None;

    for num in levels {
        let diff = num - last_num;
        if diff.abs() > 3 || diff.abs() == 0 {
            return false;
        }
        if direction == Direction::None {
            if diff > 0 {
                direction = Direction::Up;
            } else if diff < 0 {
                direction = Direction::Down;
            }
        } else if diff > 0 && direction == Direction::Down {
            return false;
        } else if diff < 0 && direction == Direction::Up {
            return false;
        }
        last_num = num;
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

        if is_safe_report(levels.clone()) {
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

        // Check if the report is already save
        if is_safe_report(levels.clone()) {
            sum += 1;
            continue;
        }
        // Try to remove each level and check if the report is safe
        for i in 0..levels.len() {
            let mut levels = levels.clone();
            levels.remove(i);
            if is_safe_report(levels.clone()) {
                sum += 1;
                break;
            }
        }
    }

    sum
}
