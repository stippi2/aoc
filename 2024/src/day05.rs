use crate::read_input;

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

fn sum_middles(rules: &Vec<Vec<i64>>, page_sequences: &Vec<Vec<i64>>) -> i64 {
    let mut sum = 0;
    for page_sequence in page_sequences {
        let mut is_valid = true;
        for rule in rules {
            let mut found_pages = Vec::new();
            for page_number in page_sequence {
                if let Some(_index) = rule.iter().position(|&x| x == *page_number) {
                    found_pages.push(*page_number);
                }
            }
            if found_pages.len() == rule.len() {
                // If we found all numbers of the rule, the rule is effective.
                if found_pages != *rule {
                    // If the found numbers are in the same sequence as in the rule, the rule is not violated.
                    is_valid = false;
                } else {
                }
            }
        }
        if is_valid {
            println!("Non-violating sequence: {:?}", page_sequence);
            if let Some(middle_number) = page_sequence.get(page_sequence.len() / 2) {
                sum += middle_number;
            }
        }
    }
    sum
}

fn sum_correct_page_sequences(input: String) -> i64 {
    match parse_input(input) {
        Some((rules, page_sequences)) => sum_middles(&rules, &page_sequences),
        None => 0,
    }
}

pub fn part1() -> i64 {
    sum_correct_page_sequences(read_input(5))
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
}
