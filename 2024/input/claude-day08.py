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

def extend_line(p1: Point, p2: Point, factor: int) -> Point:
    """Verlängert die Linie von p1 durch p2 um den Faktor"""
    dx = p2.x - p1.x
    dy = p2.y - p1.y
    return Point(
        p1.x + dx * factor,
        p1.y + dy * factor
    )

def find_antinodes(a1: Point, a2: Point, grid: List[List[str]]) -> Set[Point]:
    antinodes = set()
    height = len(grid)
    width = len(grid[0])

    # Wir suchen von beiden Antennen aus
    for base, other in [(a1, a2), (a2, a1)]:
        # Verlängere die Linie über die andere Antenne hinaus
        # Der Punkt muss doppelt so weit von der anderen Antenne sein wie von der Basis
        p = extend_line(base, other, 2)

        # Prüfe ob der Punkt innerhalb der Grenzen liegt
        if 0 <= p.x < width and 0 <= p.y < height:
            antinodes.add(p)

    return antinodes

def solve_part1(filename: str) -> int:
    grid, frequencies = read_map(filename)
    all_antinodes = set()

    # Für jede Frequenz
    for frequency, antennas in frequencies.items():
        # Für jedes Antennenpaar dieser Frequenz
        for i in range(len(antennas)):
            for j in range(i + 1, len(antennas)):
                antinodes = find_antinodes(antennas[i], antennas[j], grid)
                all_antinodes.update(antinodes)

    return len(all_antinodes)

# Lösung ausführen
result = solve_part1("./day08.txt")
print(f"Part 1 Solution: {result}")
