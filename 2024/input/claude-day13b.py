def read_input(filename: str):
    """Read and parse input file"""
    machines = []
    current_machine = []

    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            if line.startswith('Button A:'):
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
                parts = line.split(':')[1].strip().split(',')
                x = int(parts[0].split('=')[1]) + 10000000000000  # Add offset for part 2
                y = int(parts[1].split('=')[1]) + 10000000000000  # Add offset for part 2
                current_machine.append((x, y))
                machines.append(tuple(current_machine))
                current_machine = []

    return machines

def solve_machine(button_a, button_b, prize):
    """
    Solve system of linear equations:
    a_count * button_a_x + b_count * button_b_x = prize_x  (1)
    a_count * button_a_y + b_count * button_b_y = prize_y  (2)

    Multiply (1) by button_b_y and (2) by button_b_x:
    a_count * button_a_x * button_b_y + b_count * button_b_x * button_b_y = prize_x * button_b_y
    a_count * button_a_y * button_b_x + b_count * button_b_y * button_b_x = prize_y * button_b_x

    Subtract to eliminate b_count:
    a_count * (button_a_x * button_b_y - button_a_y * button_b_x) = prize_x * button_b_y - prize_y * button_b_x
    """
    ax, ay = button_a
    bx, by = button_b
    px, py = prize

    # Solve for a_count
    denominator = ax * by - ay * bx
    if denominator == 0:  # No solution if determinant is 0
        return None

    a_count = (px * by - py * bx) // denominator

    # If a_count is not an integer or is negative, no solution
    if (px * by - py * bx) % denominator != 0 or a_count < 0:
        return None

    # Calculate b_count using equation (1)
    b_count = (px - a_count * ax) // bx

    # Verify b_count is positive and the solution works
    if b_count < 0 or (px - a_count * ax) % bx != 0:
        return None

    # Verify solution works for both equations
    if a_count * ax + b_count * bx == px and a_count * ay + b_count * by == py:
        return a_count, b_count

    return None

def calculate_tokens(a_presses, b_presses):
    """Calculate total tokens needed"""
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
