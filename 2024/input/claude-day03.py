import re

class AdventDay3:
    def __init__(self):
        # Regex patterns for parsing
        self.mul_pattern = r'mul\((\d{1,3}),(\d{1,3})\)'
        self.do_pattern = r'do\(\)'
        self.dont_pattern = r'don\'t\(\)'

    def solve_part1(self, input_text):
        """
        Part 1: Find all valid multiplications and sum their results
        """
        total = 0
        found_multiplications = []

        # Find all valid matches
        matches = re.finditer(self.mul_pattern, input_text)
        for match in matches:
            num1 = int(match.group(1))
            num2 = int(match.group(2))
            result = num1 * num2
            total += result
            found_multiplications.append((num1, num2, result))

        print("\nPart 1:")
        print(f"Found {len(found_multiplications)} multiplications:")
        for num1, num2, result in found_multiplications:
            print(f"{num1} * {num2} = {result}")
        print(f"Total sum: {total}")

        return total

    def solve_part2(self, input_text):
        """
        Part 2: Handle do() and don't() instructions and sum enabled multiplications
        """
        pos = 0
        multiplications_enabled = True  # Initially enabled
        total = 0
        found_multiplications = []

        # Process input character by character
        while pos < len(input_text):
            # Check for do() instruction
            do_match = re.match(self.do_pattern, input_text[pos:])
            if do_match:
                multiplications_enabled = True
                pos += len(do_match.group(0))
                continue

            # Check for don't() instruction
            dont_match = re.match(self.dont_pattern, input_text[pos:])
            if dont_match:
                multiplications_enabled = False
                pos += len(dont_match.group(0))
                continue

            # Check for multiplication if enabled
            mul_match = re.match(self.mul_pattern, input_text[pos:])
            if mul_match:
                num1 = int(mul_match.group(1))
                num2 = int(mul_match.group(2))
                result = num1 * num2
                if multiplications_enabled:
                    total += result
                found_multiplications.append((num1, num2, result, multiplications_enabled))
                pos += len(mul_match.group(0))
            else:
                pos += 1

        print("\nPart 2:")
        print("Multiplications found:")
        for num1, num2, result, enabled in found_multiplications:
            status = "enabled" if enabled else "disabled"
            print(f"{num1} * {num2} = {result} ({status})")
        print(f"Total sum of enabled multiplications: {total}")

        return total

def main():
    # Test with examples first
    solver = AdventDay3()

    example1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    example2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)do()?mul(8,5))"

    print("Testing examples:")
    result1 = solver.solve_part1(example1)
    assert result1 == 161, f"Example 1 failed: Expected 161, got {result1}"

    result2 = solver.solve_part2(example2)
    assert result2 == 48, f"Example 2 failed: Expected 48, got {result2}"
    print("\nAll examples passed!")

    # Now solve the actual puzzle
    try:
        print("\nSolving puzzle input:")
        with open('day03.txt', 'r') as file:
            puzzle_input = file.read().strip()

        part1_solution = solver.solve_part1(puzzle_input)
        part2_solution = solver.solve_part2(puzzle_input)

        print("\nFinal Results:")
        print(f"Part 1: {part1_solution}")
        print(f"Part 2: {part2_solution}")

    except FileNotFoundError:
        print("Error: Could not find day03.txt in the current working directory")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
