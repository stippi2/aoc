use crate::read_input;

fn parse_numbers(s: &String) -> Option<(i32, i32)> {
    let mut comma_seen = false;
    let mut closing_parenthesis_seen = false;
    let mut digits_left = 0;
    let mut digits_right = 0;
    let mut left: i32 = 0;
    let mut right: i32 = 0;
    let mut offset = 0;
    while offset < s.len() {
        let c = s.as_bytes()[offset];
        match c {
            b'0'..=b'9' => {
                let digit = c - b'0';
                if comma_seen {
                    digits_right += 1;
                    right = right * 10 + digit as i32;
                } else {
                    digits_left += 1;
                    left = left * 10 + digit as i32;
                }
            }
            b',' => {
                if comma_seen {
                    return None;
                }
                comma_seen = true;
            }
            b')' => closing_parenthesis_seen = true,
            _ => {
                return None;
            }
        }
        if digits_left > 3 || digits_right > 3 {
            return None;
        }
        if closing_parenthesis_seen {
            break;
        }
        offset += 1;
    }
    if digits_left == 0 || digits_right == 0 {
        return None;
    }
    Some((left, right))
}

fn sum_valid_muls(mut input: String, enable_conditionals: bool) -> i64 {
    let mut sum = 0;
    let mut muls_enabled = true;
    loop {
        // Find the earliest occurrence of any command
        let mul_pos = input.find("mul(");
        let do_pos = input.find("do()");
        let dont_pos = input.find("don't()");

        // Get the earliest position that exists
        let next_pos = match (mul_pos, do_pos, dont_pos) {
            (None, None, None) => break, // No more commands found
            _ => {
                // Filter out None values and get minimum
                let positions = [mul_pos, do_pos, dont_pos];
                let min_pos = positions.iter().filter_map(|&x| x).min().unwrap();
                min_pos
            }
        };

        // Now determine which command was found at this position
        if Some(next_pos) == mul_pos {
            input = input.split_off(next_pos + 4);
            if let Some((a, b)) = parse_numbers(&input) {
                if muls_enabled || !enable_conditionals {
                    sum += a as i64 * b as i64;
                }
            }
        } else if Some(next_pos) == do_pos {
            input = input.split_off(next_pos + 4);
            muls_enabled = true;
        } else if Some(next_pos) == dont_pos {
            input = input.split_off(next_pos + 7);
            muls_enabled = false;
        }
    }

    sum
}

pub fn part1() -> i64 {
    sum_valid_muls(read_input(3), false)
}

pub fn part2() -> i64 {
    sum_valid_muls(read_input(3), true)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_part1() {
        let input =
            String::from("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))");
        assert_eq!(sum_valid_muls(input, false), 161 as i64);
    }

    #[test]
    fn example_part2() {
        let input = String::from(
            "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
        );
        assert_eq!(sum_valid_muls(input, true), 48 as i64);
    }
}
