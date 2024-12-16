from enum import Enum
from dataclasses import dataclass
from typing import List, Set, Tuple, Dict
import heapq

class Direction(Enum):
    NORTH = (0, -1)
    EAST = (1, 0)
    SOUTH = (0, 1)
    WEST = (-1, 0)

    def turn_left(self):
        return {
            Direction.NORTH: Direction.WEST,
            Direction.WEST: Direction.SOUTH,
            Direction.SOUTH: Direction.EAST,
            Direction.EAST: Direction.NORTH
        }[self]

    def turn_right(self):
        return {
            Direction.NORTH: Direction.EAST,
            Direction.EAST: Direction.SOUTH,
            Direction.SOUTH: Direction.WEST,
            Direction.WEST: Direction.NORTH
        }[self]

@dataclass(frozen=True)
class State:
    x: int
    y: int
    direction: Direction

class PQEntry:
    def __init__(self, cost: int, state: State, visited: Set[Tuple[int, int]]):
        self.cost = cost
        self.state = state
        self.visited = visited

    def __lt__(self, other):
        return self.cost < other.cost

class Maze:
    def __init__(self, lines: List[str]):
        self.grid = [list(line.strip()) for line in lines]
        self.height = len(self.grid)
        self.width = len(self.grid[0])

        # Find start and end positions
        self.start = None
        self.end = None
        for y in range(self.height):
            for x in range(self.width):
                if self.grid[y][x] == 'S':
                    self.start = (x, y)
                elif self.grid[y][x] == 'E':
                    self.end = (x, y)

    def is_wall(self, x: int, y: int) -> bool:
        return not (0 <= x < self.width and 0 <= y < self.height) or self.grid[y][x] == '#'

def solve_maze_both_parts(maze: Maze) -> Tuple[int, int]:
    # Cost constants
    MOVE_COST = 1
    TURN_COST = 1000

    # Beste Pfade pro Zustand speichern
    best_paths: Dict[Tuple[int, int, Direction], Tuple[int, Set[Tuple[int, int]]]] = {}
    # Optimale Pfade zum Ziel
    optimal_tiles = set()
    min_end_cost = float('inf')

    # Start mit einem leeren Pfad, der nur die Startposition enthält
    start_visited = {maze.start}
    queue = [PQEntry(0, State(maze.start[0], maze.start[1], Direction.EAST), start_visited)]

    while queue:
        entry = heapq.heappop(queue)
        state = entry.state
        cost = entry.cost
        visited = entry.visited

        state_tuple = (state.x, state.y, state.direction)

        # Wenn wir diesen Zustand schon mit niedrigeren Kosten gesehen haben, überspringen
        if state_tuple in best_paths and best_paths[state_tuple][0] < cost:
            continue

        # Wenn die Kosten schon höher sind als der beste Pfad zum Ziel, überspringen
        if cost > min_end_cost:
            continue

        # Aktuellen Pfad als besten für diesen Zustand speichern oder mit bestehendem vereinigen
        if state_tuple not in best_paths or cost < best_paths[state_tuple][0]:
            best_paths[state_tuple] = (cost, visited.copy())
        elif cost == best_paths[state_tuple][0]:
            # Pfade vereinigen bei gleichen Kosten
            best_paths[state_tuple][1].update(visited)

        # Wenn wir am Ziel sind
        if (state.x, state.y) == maze.end:
            if cost < min_end_cost:
                min_end_cost = cost
                optimal_tiles = visited.copy()
            elif cost == min_end_cost:
                optimal_tiles.update(visited)
            continue

        # Mögliche Bewegungen probieren
        # Vorwärts
        dx, dy = state.direction.value
        new_x, new_y = state.x + dx, state.y + dy
        if not maze.is_wall(new_x, new_y):
            new_visited = visited.copy()
            new_visited.add((new_x, new_y))
            heapq.heappush(queue, PQEntry(
                cost + MOVE_COST,
                State(new_x, new_y, state.direction),
                new_visited
            ))

        # Links drehen
        new_direction = state.direction.turn_left()
        heapq.heappush(queue, PQEntry(
            cost + TURN_COST,
            State(state.x, state.y, new_direction),
            visited.copy()
        ))

        # Rechts drehen
        new_direction = state.direction.turn_right()
        heapq.heappush(queue, PQEntry(
            cost + TURN_COST,
            State(state.x, state.y, new_direction),
            visited.copy()
        ))

    return min_end_cost, len(optimal_tiles)

def main():
    # Read input from file
    with open('day16.txt', 'r') as f:
        lines = f.readlines()

    maze = Maze(lines)
    min_cost, optimal_count = solve_maze_both_parts(maze)
    print(f"Part 1 - Lowest possible score: {min_cost}")
    print(f"Part 2 - Number of tiles on optimal paths: {optimal_count}")

if __name__ == "__main__":
    main()
