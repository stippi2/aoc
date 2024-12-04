def find_x_mas(grid):
    """
    Find all X-shaped patterns where each diagonal contains MAS (forward or backward)
    Returns the count of valid X-MAS patterns
    """
    if not grid:
        return 0

    rows = len(grid)
    cols = len(grid[0])
    count = 0

    def is_valid(x, y):
        return 0 <= x < rows and 0 <= y < cols

    def check_mas(x, y, dx, dy):
        """
        Check if MAS exists in direction (dx,dy) starting at (x,y)
        Returns True if found either MAS or SAM
        """
        # Check forward MAS
        if (is_valid(x, y) and is_valid(x + 2*dx, y + 2*dy) and
            grid[x][y] == 'M' and
            grid[x + dx][y + dy] == 'A' and
            grid[x + 2*dx][y + 2*dy] == 'S'):
            return True

        # Check backward SAM
        if (is_valid(x, y) and is_valid(x + 2*dx, y + 2*dy) and
            grid[x][y] == 'S' and
            grid[x + dx][y + dy] == 'A' and
            grid[x + 2*dx][y + 2*dy] == 'M'):
            return True

        return False

    # Check every position as potential center of X
    for i in range(1, rows-1):  # Center needs space above and below
        for j in range(1, cols-1):  # Center needs space left and right
            # For each center point, check all 4 possible X configurations:
            # 1. Upper-left to lower-right AND upper-right to lower-left
            combinations = [
                # direction pairs for the two diagonals
                [(-1,-1), (1,1), (-1,1), (1,-1)],   # normal X
            ]

            for dirs in combinations:
                # Check first diagonal
                first_diagonal = check_mas(i + dirs[0][0], j + dirs[0][1], dirs[1][0], dirs[1][1])
                # Check second diagonal
                second_diagonal = check_mas(i + dirs[2][0], j + dirs[2][1], dirs[3][0], dirs[3][1])

                if first_diagonal and second_diagonal:
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
result = find_x_mas(grid)
print(f"Found {result} X-MAS patterns")
