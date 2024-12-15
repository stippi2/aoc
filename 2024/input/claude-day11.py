from typing import List

class StoneLine:
    def __init__(self, stones: List[str]):
        self.stones = stones

    @staticmethod
    def _has_even_digits(stone: str) -> bool:
        return len(stone) % 2 == 0

    @staticmethod
    def _split_stone(stone: str) -> tuple[str, str]:
        mid = len(stone) // 2
        left = stone[:mid].lstrip('0') or '0'
        right = stone[mid:].lstrip('0') or '0'
        return left, right

    def blink(self) -> None:
        new_stones = []

        for stone in self.stones:
            # Regel 1: Wenn der Stein 0 ist, wird er zu 1
            if stone == '0':
                new_stones.append('1')

            # Regel 2: Wenn der Stein eine gerade Anzahl von Ziffern hat, wird er geteilt
            elif self._has_even_digits(stone):
                left, right = self._split_stone(stone)
                new_stones.extend([left, right])

            # Regel 3: Ansonsten wird die Zahl mit 2024 multipliziert
            else:
                result = str(int(stone) * 2024)
                new_stones.append(result)

        self.stones = new_stones

    def stone_count(self) -> int:
        return len(self.stones)

def solve_part1(input_file: str, blinks: int = 25) -> int:
    # Lese die Input-Datei
    with open(input_file, 'r') as f:
        stones = f.read().strip().split()

    # Erstelle eine neue StoneLine-Instanz
    stone_line = StoneLine(stones)

    # Führe die angegebene Anzahl von Blinks durch
    for i in range(blinks):
        stone_line.blink()

    return stone_line.stone_count()

def solve_part2(input_file: str) -> int:
    # Lese Input
    with open(input_file, 'r') as f:
        stones = f.read().strip().split()

    # Führe die ersten 5 Blinks aus, um die Basis zu bekommen
    stone_line = StoneLine(stones)
    for _ in range(5):
        stone_line.blink()
    base_count = stone_line.stone_count()

    # Berechne die verbleibenden vollen 5er-Zyklen
    remaining_full_cycles = (75 - 5) // 5  # = 14

    # Berechne das Ergebnis
    result = base_count * pow(8, remaining_full_cycles)

    # Führe die restlichen Blinks aus
    stone_line = StoneLine(stones)
    remaining_blinks = 75 % 5  # Sollte 0 sein
    for _ in range(remaining_blinks):
        stone_line.blink()

    return result

if __name__ == "__main__":
    result1 = solve_part1("day11.txt")
    print(f"Nach 25 Blinks gibt es {result1} Steine.")

    result2 = solve_part2("day11.txt")
    print(f"Nach 75 Blinks gibt es {result2} Steine.")
