from collections import defaultdict
from dataclasses import dataclass
from typing import List, Set, Tuple

@dataclass(frozen=True)
class Point:
    x: int
    y: int

def read_map(filename: str) -> Tuple[List[List[str]], dict]:
    with open(filename, 'r') as f:
        lines = [line.strip() for line in f.readlines()]

    grid = [list(line) for line in lines]
    frequencies = defaultdict(list)

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x] != '.':
                frequencies[grid[y][x]].append(Point(x, y))

    return grid, frequencies

def are_collinear(p1: Point, p2: Point, p3: Point) -> bool:
    """Prüft ob drei Punkte auf einer Linie liegen"""
    # Verwendet die Determinante um Collinearität zu prüfen
    # (x2-x1)(y3-y1) = (y2-y1)(x3-x1)
    return (p2.x - p1.x) * (p3.y - p1.y) == (p2.y - p1.y) * (p3.x - p1.x)

def find_antinodes(antennas: List[Point], grid: List[List[str]]) -> Set[Point]:
    height = len(grid)
    width = len(grid[0])
    antinodes = set()

    # Wenn wir mindestens 2 Antennen haben, sind alle Antennenpositionen auch Antinoden
    if len(antennas) >= 2:
        antinodes.update(antennas)

    # Prüfe jeden möglichen Punkt im Grid
    for y in range(height):
        for x in range(width):
            test_point = Point(x, y)

            # Suche nach Antennenpaaren, mit denen dieser Punkt collinear ist
            for i in range(len(antennas)):
                for j in range(i + 1, len(antennas)):
                    if are_collinear(antennas[i], antennas[j], test_point):
                        antinodes.add(test_point)
                        break  # Ein Treffer reicht
                else:
                    continue
                break

    return antinodes

def solve_part2(filename: str) -> int:
    grid, frequencies = read_map(filename)
    all_antinodes = set()

    # Für jede Frequenz
    for frequency, antennas in frequencies.items():
        antinodes = find_antinodes(antennas, grid)
        all_antinodes.update(antinodes)

    return len(all_antinodes)

# Lösung ausführen
result = solve_part2("./day08.txt")
print(f"Part 2 Solution: {result}")
