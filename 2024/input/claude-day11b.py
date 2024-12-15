from collections import Counter
from typing import Dict

class StoneCounter:
    def __init__(self, stones: list[str]):
        self.number_counts = Counter(stones)

    def _transform_number(self, number: str, count: int) -> Dict[str, int]:
        if number == '0':
            return {'1': count}
        elif len(number) % 2 == 0:
            mid = len(number) // 2
            left = number[:mid].lstrip('0') or '0'
            right = number[mid:].lstrip('0') or '0'
            result = {left: count}
            result[right] = result.get(right, 0) + count
            return result
        else:
            result = str(int(number) * 2024)
            return {result: count}

    def blink(self) -> None:
        new_counts = Counter()
        for number, count in self.number_counts.items():
            transformed = self._transform_number(number, count)
            new_counts.update(transformed)
        self.number_counts = new_counts

    def stone_count(self) -> int:
        return sum(self.number_counts.values())

def solve(input_file: str, blinks: int) -> int:
    with open(input_file, 'r') as f:
        stones = f.read().strip().split()

    counter = StoneCounter(stones)
    for _ in range(blinks):
        counter.blink()

    return counter.stone_count()

if __name__ == "__main__":
    part1 = solve("day11.txt", 25)
    print(f"Teil 1 (25 Blinks): {part1}")

    part2 = solve("day11.txt", 75)
    print(f"Teil 2 (75 Blinks): {part2}")
