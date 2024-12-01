pub mod day01;

pub fn read_input(day: u8) -> String {
    std::fs::read_to_string(format!("input/day{:02}.txt", day)).expect("Could not read input file")
}
