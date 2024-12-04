use crate::{read_input, Grid};

const DIRECTIONS: [(i32, i32); 8] = [
    (-1, 0),
    (1, 0),
    (0, -1),
    (0, 1),
    (-1, -1),
    (1, -1),
    (-1, 1),
    (1, 1),
];

fn count_matches(grid: &Grid<char>, start_x: i32, start_y: i32, word: &String) -> i64 {
    let mut matches = 0;
    for direction in DIRECTIONS {
        let mut x = start_x + direction.0;
        let mut y = start_y + direction.1;
        let mut found_match = true;
        for word_letter in word.chars().skip(1) {
            match grid.get(x, y) {
                Some(grid_letter) => {
                    if word_letter != *grid_letter {
                        found_match = false;
                        break;
                    }
                }
                None => {
                    found_match = false;
                    break;
                }
            }
            x += direction.0;
            y += direction.1;
        }
        if found_match {
            matches += 1
        }
    }
    matches
}

fn sum_words(input: String, word: String) -> i64 {
    let grid = Grid::from_string(&input);

    let mut sum = 0;

    if let Some(first_letter) = word.chars().next() {
        for y in 0..grid.height() {
            for x in 0..grid.width() {
                if let Some(letter) = grid.get(x as i32, y as i32) {
                    if *letter == first_letter {
                        sum += count_matches(&grid, x as i32, y as i32, &word);
                    }
                }
            }
        }
    }

    sum
}

pub fn part1() -> i64 {
    sum_words(read_input(4), String::from("XMAS"))
}

fn sum_crossings(input: String) -> i64 {
    let grid = Grid::from_string(&input);

    let mut sum = 0;
    let middle = 'A';

    for y in 1..grid.height() - 1 {
        for x in 1..grid.width() - 1 {
            if let (Some(&letter), Some(&lt), Some(&rt), Some(&lb), Some(&rb)) = (
                grid.get(x as i32, y as i32),
                grid.get(x as i32 - 1, y as i32 - 1),
                grid.get(x as i32 + 1, y as i32 - 1),
                grid.get(x as i32 - 1, y as i32 + 1),
                grid.get(x as i32 + 1, y as i32 + 1),
            ) {
                if letter == middle {
                    // Check diagonal pairs
                    let diag1 = format!("{}{}", lt, rb);
                    let diag2 = format!("{}{}", rt, lb);

                    if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
                        sum += 1;
                    }
                }
            }
        }
    }

    sum
}

pub fn part2() -> i64 {
    sum_crossings(read_input(4))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_part1() {
        let input = String::from(
            r#"MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"#,
        );
        assert_eq!(sum_words(input, String::from("XMAS")), 18 as i64);
    }

    #[test]
    fn example_part2() {
        let input = String::from(
            r#"MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"#,
        );
        assert_eq!(sum_crossings(input), 9 as i64);
    }
}
