from typing import Tuple, Optional

def read_input(filename: str) -> list[tuple[Tuple[int, int], Tuple[int, int], Tuple[int, int]]]:
    """Read and parse input file into list of (button_a, button_b, prize) tuples"""
    machines = []
    current_machine = []

    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            if line.startswith('Button A:'):
                # Parse "Button A: X+94, Y+34" format
                parts = line.split(':')[1].strip().split(',')
                x = int(parts[0].split('+')[1])
                y = int(parts[1].split('+')[1])
                current_machine.append((x, y))
            elif line.startswith('Button B:'):
                parts = line.split(':')[1].strip().split(',')
                x = int(parts[0].split('+')[1])
                y = int(parts[1].split('+')[1])
                current_machine.append((x, y))
            elif line.startswith('Prize:'):
                # Parse "Prize: X=8400, Y=5400" format
                parts = line.split(':')[1].strip().split(',')
                x = int(parts[0].split('=')[1])
                y = int(parts[1].split('=')[1])
                current_machine.append((x, y))
                machines.append(tuple(current_machine))
                current_machine = []

    return machines

def solve_machine(button_a: Tuple[int, int], button_b: Tuple[int, int],
                 prize: Tuple[int, int], max_presses: int = 100) -> Optional[Tuple[int, int]]:
    """
    Try to find solution for a single machine.
    Returns (a_presses, b_presses) if solution found, None otherwise.
    """
    # For each possible number of A presses
    for a in range(max_presses + 1):
        # Calculate remaining distance after A presses
        remaining_x = prize[0] - (a * button_a[0])
        remaining_y = prize[1] - (a * button_a[1])

        # Check if this can be solved with B presses
        if remaining_x % button_b[0] == 0 and remaining_y % button_b[1] == 0:
            b_x = remaining_x // button_b[0]
            b_y = remaining_y // button_b[1]

            # If both axes require same number of B presses and it's within limit
            if b_x == b_y and 0 <= b_x <= max_presses:
                return (a, b_x)

    return None

def calculate_tokens(a_presses: int, b_presses: int) -> int:
    """Calculate total tokens needed for given button presses"""
    return a_presses * 3 + b_presses * 1

def main():
    machines = read_input('./day13.txt')
    total_tokens = 0
    solvable_count = 0

    for i, (button_a, button_b, prize) in enumerate(machines, 1):
        solution = solve_machine(button_a, button_b, prize)
        if solution:
            a_presses, b_presses = solution
            tokens = calculate_tokens(a_presses, b_presses)
            total_tokens += tokens
            solvable_count += 1
            print(f"Machine {i}: Solvable with {a_presses}A + {b_presses}B = {tokens} tokens")
        else:
            print(f"Machine {i}: Not solvable")

    print(f"\nTotal solvable machines: {solvable_count}")
    print(f"Total tokens needed: {total_tokens}")

if __name__ == "__main__":
    main()
