def is_safe_report(levels):
    """
    Check if a report is safe according to the rules:
    - All numbers must be either increasing or decreasing
    - Difference between adjacent numbers must be between 1 and 3 inclusive
    """
    if len(levels) < 2:
        return False

    # Check first difference to determine if sequence should be increasing or decreasing
    first_diff = levels[1] - levels[0]
    if abs(first_diff) < 1 or abs(first_diff) > 3:
        return False

    should_increase = first_diff > 0

    # Check all adjacent pairs
    for i in range(1, len(levels)):
        diff = levels[i] - levels[i-1]

        # Check if difference is between 1 and 3
        if abs(diff) < 1 or abs(diff) > 3:
            return False

        # Check if maintaining increasing/decreasing pattern
        if should_increase and diff <= 0:
            return False
        if not should_increase and diff >= 0:
            return False

    return True

def is_safe_with_dampener(levels):
    """
    Check if a report is safe, considering the Problem Dampener which can remove one number
    """
    # First check if it's safe without removing any number
    if is_safe_report(levels):
        return True

    # Try removing each number one at a time
    for i in range(len(levels)):
        # Create new list without the current number
        dampened_levels = levels[:i] + levels[i+1:]
        if is_safe_report(dampened_levels):
            return True

    return False

def count_safe_reports(filename, use_dampener=False):
    """
    Read reports from file and count how many are safe
    use_dampener: If True, uses Problem Dampener logic
    """
    safe_count = 0

    try:
        with open(filename, 'r') as file:
            for line in file:
                # Convert line to list of integers, skip empty lines
                line = line.strip()
                if not line:
                    continue

                levels = [int(x) for x in line.split()]
                if use_dampener:
                    if is_safe_with_dampener(levels):
                        safe_count += 1
                else:
                    if is_safe_report(levels):
                        safe_count += 1

        return safe_count

    except FileNotFoundError:
        print(f"Error: Could not find file {filename}")
        return 0
    except ValueError:
        print("Error: File contains invalid number format")
        return 0

# Test cases from both parts
def run_tests():
    test_cases = [
        # Format: (levels, safe_without_dampener, safe_with_dampener)
        ([7, 6, 4, 2, 1], True, True),     # Safe without dampener
        ([1, 2, 7, 8, 9], False, False),   # Unsafe even with dampener
        ([9, 7, 6, 2, 1], False, False),   # Unsafe even with dampener
        ([1, 3, 2, 4, 5], False, True),    # Safe with dampener (remove 3)
        ([8, 6, 4, 4, 1], False, True),    # Safe with dampener (remove one 4)
        ([1, 3, 6, 7, 9], True, True),     # Safe without dampener
    ]

    print("Part 1 Tests (without dampener):")
    for i, (levels, expected, _) in enumerate(test_cases, 1):
        result = is_safe_report(levels)
        print(f"Test {i}: {levels}")
        print(f"Expected: {expected}, Got: {result}")
        print("PASS" if result == expected else "FAIL")
        print()

    print("\nPart 2 Tests (with dampener):")
    for i, (levels, _, expected) in enumerate(test_cases, 1):
        result = is_safe_with_dampener(levels)
        print(f"Test {i}: {levels}")
        print(f"Expected: {expected}, Got: {result}")
        print("PASS" if result == expected else "FAIL")
        print()

if __name__ == "__main__":
    # Run tests first
    print("Running tests...")
    run_tests()

    # Process actual input file for both parts
    print("\nProcessing input file...")
    result_part1 = count_safe_reports("day02.txt", use_dampener=False)
    result_part2 = count_safe_reports("day02.txt", use_dampener=True)
    print(f"Part 1 - Number of safe reports: {result_part1}")
    print(f"Part 2 - Number of safe reports with Problem Dampener: {result_part2}")
