def parse_numbers(s: str) -> tuple[int, int] | None:
    # Track state while parsing
    comma_seen = False
    digits_left = digits_right = left = right = 0

    for c in s:
        if c.isdigit():
            digit = int(c)
            if comma_seen:
                digits_right += 1
                right = right * 10 + digit
            else:
                digits_left += 1
                left = left * 10 + digit
        elif c == ',':
            if comma_seen:
                return None
            comma_seen = True
        elif c == ')':
            break
        else:
            return None

        if digits_left > 3 or digits_right > 3:
            return None

    return (left, right) if digits_left > 0 and digits_right > 0 else None

def solve_puzzle(input_text: str, enable_conditionals: bool) -> int:
    total = 0
    text = input_text
    muls_enabled = True

    while True:
        # Find next command positions
        mul_pos = text.find('mul(')
        do_pos = text.find('do()')
        dont_pos = text.find("don't()")

        # Get earliest valid position
        positions = [p for p in (mul_pos, do_pos, dont_pos) if p != -1]
        if not positions:
            break

        pos = min(positions)

        # Handle command at position
        if pos == mul_pos:
            numbers = parse_numbers(text[pos + 4:])
            if numbers and (muls_enabled or not enable_conditionals):
                total += numbers[0] * numbers[1]
            text = text[pos + 4:]
        elif pos == do_pos:
            muls_enabled = True
            text = text[pos + 4:]
        else:  # dont_pos
            muls_enabled = False
            text = text[pos + 7:]

    return total

if __name__ == '__main__':
    with open('day03.txt') as f:
        data = f.read()
    print(f'Part 1: {solve_puzzle(data, False)}')
    print(f'Part 2: {solve_puzzle(data, True)}')
