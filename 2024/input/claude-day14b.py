from dataclasses import dataclass
from typing import List, Tuple, Dict, Set
from collections import defaultdict
import re

@dataclass
class Robot:
    x: int
    y: int
    vx: int
    vy: int

class RobotSimulation:
    def __init__(self, width: int, height: int):
        self.width = width
        self.height = height
        self.robots: List[Robot] = []

    def add_robot(self, robot: Robot):
        self.robots.append(robot)

    def parse_input(self, filename: str):
        pattern = r'p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)'
        with open(filename, 'r') as f:
            for line in f:
                if match := re.match(pattern, line.strip()):
                    x, y, vx, vy = map(int, match.groups())
                    self.add_robot(Robot(x, y, vx, vy))

    def get_position_at_time(self, robot: Robot, t: int) -> Tuple[int, int]:
        x = (robot.x + robot.vx * t) % self.width
        y = (robot.y + robot.vy * t) % self.height
        return (x, y)

    def find_horizontal_line(self, required_robots: int = 10, max_time: int = 20000) -> int:
        for t in range(max_time):
            if t % 1000 == 0:
                print(f"Checking time {t}...")

            # Gruppiere Roboter nach y-Koordinate
            y_groups: Dict[int, Set[int]] = defaultdict(set)

            for robot in self.robots:
                x, y = self.get_position_at_time(robot, t)
                y_groups[y].add(x)

            # Für jede y-Koordinate
            for y, x_coords in y_groups.items():
                # Sortiere x-Koordinaten
                x_coords = sorted(x_coords)

                # Suche nach aufeinanderfolgenden x-Koordinaten
                consecutive = 1
                max_consecutive = 1

                for i in range(1, len(x_coords)):
                    if (x_coords[i] - x_coords[i-1]) == 1:
                        consecutive += 1
                        max_consecutive = max(max_consecutive, consecutive)
                    else:
                        # Berücksichtige auch Wrapping von der rechten zur linken Seite
                        if (x_coords[i-1] == self.width - 1 and x_coords[i] == 0):
                            consecutive += 1
                            max_consecutive = max(max_consecutive, consecutive)
                        else:
                            consecutive = 1

                if max_consecutive >= required_robots:
                    return t

        return -1

def main():
    sim = RobotSimulation(101, 103)
    sim.parse_input('./day14.txt')

    line_time = sim.find_horizontal_line()
    if line_time >= 0:
        print(f"Found horizontal line at t={line_time}")
    else:
        print("No horizontal line found")

if __name__ == "__main__":
    main()
