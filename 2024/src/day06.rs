use crate::{read_input, Grid, Vec2};
use std::collections::HashSet;

fn find_start(grid: &Grid<char>) -> Option<Vec2> {
    if let Some((x, y)) = grid.find(&'^') {
        return Some(Vec2::new(x, y));
    }
    None
}

fn sum_visited_fields(input: String) -> usize {
    let grid = Grid::from_string(&input);

    if let Some(mut current_pos) = find_start(&grid) {
        let mut visited = HashSet::new();
        let mut direction = Vec2::new(0, -1);
        let mut next_pos = current_pos.add(&direction);
        visited.insert(current_pos);

        while let Some(next_char) = grid.get(next_pos.x, next_pos.y) {
            // Update direction based on current character
            match next_char {
                '#' => {
                    // Next char is obstacle, rotate right 90°
                    direction = Vec2::new(-direction.y, direction.x);
                }
                _ => {
                    // Continue in same direction
                    current_pos = next_pos;
                    visited.insert(current_pos);
                }
            }

            next_pos = current_pos.add(&direction);
        }
        return visited.len();
    }
    0
}

pub fn part1() -> usize {
    sum_visited_fields(read_input(6))
}

pub fn part2() -> i64 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_part1() {
        let input = String::from(
            r#"....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."#,
        );
        assert_eq!(sum_visited_fields(input), 41 as usize);
    }
}
