def find_xmas(grid):
    """
    Find all occurrences of 'XMAS' in any direction in the grid
    Returns the count of XMAS occurrences
    """
    if not grid:
        return 0

    rows = len(grid)
    cols = len(grid[0])
    count = 0
    word = "XMAS"

    # All possible directions: right, down-right, down, down-left, left, up-left, up, up-right
    directions = [
        (0, 1),   # right
        (1, 1),   # down-right
        (1, 0),   # down
        (1, -1),  # down-left
        (0, -1),  # left
        (-1, -1), # up-left
        (-1, 0),  # up
        (-1, 1)   # up-right
    ]

    def is_valid(x, y):
        return 0 <= x < rows and 0 <= y < cols

    def check_direction(x, y, dx, dy):
        # Check if XMAS can be found starting at position (x,y) in direction (dx,dy)
        for i in range(len(word)):
            new_x = x + i * dx
            new_y = y + i * dy
            if not is_valid(new_x, new_y) or grid[new_x][new_y] != word[i]:
                return False
        return True

    # Check every position as potential start
    for i in range(rows):
        for j in range(cols):
            # Try all directions from this position
            for dx, dy in directions:
                if check_direction(i, j, dx, dy):
                    count += 1

    return count

# Example grid from the problem
example_grid = [
    "MMMSXXMASM",
    "MSAMXMSMSA",
    "AMXSXMAAMM",
    "MSAMASMSMX",
    "XMASAMXAMM",
    "XXAMMXXAMA",
    "SMSMSASXSS",
    "SAXAMASAAA",
    "MAMMMXMMMM",
    "MXMXAXMASX"
]

# Convert strings to list of characters for easier processing
grid = [list(row) for row in example_grid]
result = find_xmas(grid)
print(f"Found {result} occurrences of XMAS")  # Should print 18
