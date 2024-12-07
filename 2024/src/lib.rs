pub fn read_input(day: u8) -> String {
    std::fs::read_to_string(format!("input/day{:02}.txt", day)).expect("Could not read input file")
}

use std::fmt::Display;

#[derive(Debug, Clone)]
pub struct Grid<T> {
    data: Vec<T>,
    width: usize,
    height: usize,
}

impl<T: Clone + Default> Grid<T> {
    pub fn new(width: usize, height: usize) -> Self {
        Grid {
            data: vec![T::default(); width * height],
            width,
            height,
        }
    }

    // Basic getters
    pub fn width(&self) -> usize {
        self.width
    }
    pub fn height(&self) -> usize {
        self.height
    }
    pub fn data(&self) -> &[T] {
        &self.data
    }
}

impl<T: Clone> Grid<T> {
    pub fn get(&self, x: i32, y: i32) -> Option<&T> {
        if x >= 0 && x < self.width as i32 && y >= 0 && y < self.height as i32 {
            Some(&self.data[y as usize * self.width + x as usize])
        } else {
            None
        }
    }

    pub fn get_mut(&mut self, x: i32, y: i32) -> Option<&mut T> {
        if x >= 0 && x < self.width as i32 && y >= 0 && y < self.height as i32 {
            Some(&mut self.data[y as usize * self.width + x as usize])
        } else {
            None
        }
    }

    pub fn set(&mut self, x: usize, y: usize, value: T) -> bool {
        if x < self.width && y < self.height {
            self.data[y * self.width + x] = value;
            true
        } else {
            false
        }
    }

    // Utility method for iterating over grid coordinates
    pub fn iter_coords(&self) -> impl Iterator<Item = (i32, i32)> + '_ {
        (0..self.height as i32).flat_map(move |y| (0..self.width as i32).map(move |x| (x, y)))
    }
}

impl<T: Clone + PartialEq> Grid<T> {
    pub fn find(&self, target: &T) -> Option<(i32, i32)> {
        self.iter_coords()
            .find(|&(x, y)| self.get(x, y).map_or(false, |v| v == target))
    }

    pub fn find_all(&self, target: &T) -> Vec<(i32, i32)> {
        self.iter_coords()
            .filter(|&(x, y)| self.get(x, y).map_or(false, |v| v == target))
            .collect()
    }
}

// Special implementation for char grids (common in AoC)
impl Grid<char> {
    pub fn from_string(input: &str) -> Self {
        let height = input.lines().count();
        let width = input.lines().next().map(|line| line.len()).unwrap_or(0);
        let mut grid = Grid::new(width, height);

        for (y, line) in input.lines().enumerate() {
            for (x, ch) in line.chars().enumerate() {
                grid.set(x, y, ch);
            }
        }
        grid
    }
}

// Pretty printing for any grid with displayable elements
impl<T: Display> Display for Grid<T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for y in 0..self.height {
            for x in 0..self.width {
                write!(f, "{}", self.data[y * self.width + x])?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub struct Vec2 {
    pub x: i32,
    pub y: i32,
}

impl Vec2 {
    pub fn new(x: i32, y: i32) -> Self {
        Self { x, y }
    }

    pub fn add(&self, other: &Self) -> Self {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_char_grid() {
        let input = "ABC\nDEF\nGHI";
        let grid = Grid::from_string(input);
        assert_eq!(grid.width, 3);
        assert_eq!(grid.height, 3);
        assert_eq!(grid.get(0, 0), Some(&'A'));
        assert_eq!(grid.get(2, 2), Some(&'I'));
        assert_eq!(grid.get(3, 3), None);
    }

    #[test]
    fn test_numeric_grid() {
        let mut grid: Grid<i32> = Grid::new(2, 2);
        grid.set(0, 0, 1);
        grid.set(1, 1, 2);
        assert_eq!(grid.get(0, 0), Some(&1));
        assert_eq!(grid.get(1, 1), Some(&2));
    }

    #[test]
    fn test_find_char() {
        let input = "ABC\nDEF\nGHI";
        let grid = Grid::from_string(input);

        // Find first occurrence
        assert_eq!(grid.find(&'E'), Some((1, 1)));
        assert_eq!(grid.find(&'X'), None);
    }

    #[test]
    fn test_find_all_chars() {
        let input = "ABA\nBAB\nABA";
        let grid = Grid::from_string(input);

        // Find all occurrences of 'A'
        let a_positions = grid.find_all(&'A');
        assert_eq!(a_positions.len(), 4);
        assert!(a_positions.contains(&(0, 0)));
        assert!(a_positions.contains(&(2, 0)));
        assert!(a_positions.contains(&(0, 2)));
        assert!(a_positions.contains(&(2, 2)));
    }
}
