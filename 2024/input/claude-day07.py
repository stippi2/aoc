def concat_ints(a: int, b: int) -> int:
    """Concatenate integers mathematically instead of using strings."""
    multiplier = 1
    temp = b
    while temp > 0:
        multiplier *= 10
        temp //= 10
    return a * multiplier + b

def solve(target: int, value: int, sequence: list, use_concat: bool) -> bool:
    """Recursively try to find a valid solution."""
    if not sequence:
        return target == value

    next_val = sequence[0]
    rest = sequence[1:]

    # Try addition
    if solve(target, value + next_val, rest, use_concat):
        return True

    # Try multiplication
    if solve(target, value * next_val, rest, use_concat):
        return True

    # Try concatenation if allowed
    if use_concat and solve(target, concat_ints(value, next_val), rest, use_concat):
        return True

    return False

def solve_puzzle(filename: str, part2: bool = False) -> int:
    total = 0

    with open(filename, 'r') as file:
        for line in file:
            target_str, numbers_str = line.strip().split(': ')
            target = int(target_str)
            numbers = [int(x) for x in numbers_str.split()]

            if solve(target, numbers[0], numbers[1:], part2):
                total += target

    return total

if __name__ == "__main__":
    from time import time

    start = time()
    part1 = solve_puzzle('./day07.txt')
    mid = time()
    part2 = solve_puzzle('./day07.txt', True)
    end = time()

    print(f"Part 1: {part1} (took {(mid-start)*1000:.1f}ms)")
    print(f"Part 2: {part2} (took {(end-mid)*1000:.1f}ms)")
