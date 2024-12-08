use crate::read_input;
use std::cmp::Ordering;

fn parse_input(input: String) -> Option<(Vec<Vec<i64>>, Vec<Vec<i64>>)> {
    let mut rules: Vec<Vec<i64>> = Vec::new();
    let mut pages: Vec<Vec<i64>> = Vec::new();

    let mut sections = input.split("\n\n");

    if let Some(rules_section) = sections.next() {
        for rule_line in rules_section.lines() {
            let numbers: Vec<i64> = rule_line
                .trim()
                .split('|')
                .filter_map(|s| s.parse().ok())
                .collect();
            rules.push(numbers);
        }
    } else {
        return None;
    }

    if let Some(pages_section) = sections.next() {
        for pages_line in pages_section.lines() {
            let numbers: Vec<i64> = pages_line
                .trim()
                .split(',')
                .filter_map(|s| s.parse().ok())
                .collect();
            pages.push(numbers);
        }
        Some((rules, pages))
    } else {
        None
    }
}

fn sum_middles(rules: &Vec<Vec<i64>>, page_sequences: &Vec<Vec<i64>>) -> (i64, i64) {
    let mut sum_valid = 0;
    let mut sum_invalid = 0;
    for page_sequence in page_sequences {
        let mut is_valid = true;
        for rule in rules {
            let mut found_pages = Vec::new();
            for page_number in page_sequence {
                if let Some(_index) = rule.iter().position(|&x| x == *page_number) {
                    found_pages.push(*page_number);
                }
            }
            // If we found all numbers of the rule, the rule is effective.
            // If the found numbers are not in the same order as in the rule, the rule is violated.
            if found_pages.len() == rule.len() && found_pages != *rule {
                is_valid = false;
            }
        }
        if is_valid {
            if let Some(middle_number) = page_sequence.get(page_sequence.len() / 2) {
                sum_valid += middle_number;
            }
        } else {
            let mut corrected_sequence = page_sequence.clone();
            corrected_sequence.sort_by(|a, b| {
                let pair = vec![*a, *b];
                if let Some(rule) = rules.iter().find(|r| **r == pair) {
                    if rule[0] == *a {
                        Ordering::Less
                    } else {
                        Ordering::Greater
                    }
                } else {
                    Ordering::Equal
                }
            });
            if let Some(middle_number) = corrected_sequence.get(corrected_sequence.len() / 2) {
                sum_invalid += middle_number;
            }
        }
    }
    (sum_valid, sum_invalid)
}

fn sum_correct_page_sequences(input: String) -> i64 {
    match parse_input(input) {
        Some((rules, page_sequences)) => sum_middles(&rules, &page_sequences).0,
        None => 0,
    }
}

fn sum_incorrect_page_sequences(input: String) -> i64 {
    match parse_input(input) {
        Some((rules, page_sequences)) => sum_middles(&rules, &page_sequences).1,
        None => 0,
    }
}

pub fn part1() -> i64 {
    sum_correct_page_sequences(read_input(5))
}

pub fn part2() -> i64 {
    sum_incorrect_page_sequences(read_input(5))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_part1() {
        let input = String::from(
            r#"47|53
            97|13
            97|61
            97|47
            75|29
            61|13
            75|53
            29|13
            97|29
            53|29
            61|53
            97|53
            61|29
            47|13
            75|47
            97|75
            47|61
            75|61
            47|29
            75|13
            53|13

            75,47,61,53,29
            97,61,53,29,13
            75,29,13
            75,97,47,61,53
            61,13,29
            97,13,75,29,47"#,
        );
        assert_eq!(sum_correct_page_sequences(input), 143 as i64);
    }

    #[test]
    fn example_part2() {
        let input = String::from(
            r#"47|53
            97|13
            97|61
            97|47
            75|29
            61|13
            75|53
            29|13
            97|29
            53|29
            61|53
            97|53
            61|29
            47|13
            75|47
            97|75
            47|61
            75|61
            47|29
            75|13
            53|13

            75,47,61,53,29
            97,61,53,29,13
            75,29,13
            75,97,47,61,53
            61,13,29
            97,13,75,29,47"#,
        );
        assert_eq!(sum_incorrect_page_sequences(input), 123 as i64);
    }
}
