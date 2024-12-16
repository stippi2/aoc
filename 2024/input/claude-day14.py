from dataclasses import dataclass
from typing import List
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

    def simulate_step(self):
        for robot in self.robots:
            # Update position
            robot.x = (robot.x + robot.vx) % self.width
            robot.y = (robot.y + robot.vy) % self.height

    def simulate(self, steps: int):
        for _ in range(steps):
            self.simulate_step()

    def count_robots_in_quadrants(self) -> tuple[int, int, int, int]:
        # Initialize counters for each quadrant
        quadrants = [0, 0, 0, 0]  # [top-left, top-right, bottom-left, bottom-right]

        mid_x = self.width // 2
        mid_y = self.height // 2

        for robot in self.robots:
            # Skip robots on the middle lines
            if robot.x == mid_x or robot.y == mid_y:
                continue

            quadrant_idx = (
                (1 if robot.x > mid_x else 0) +
                (2 if robot.y > mid_y else 0)
            )
            quadrants[quadrant_idx] += 1

        return tuple(quadrants)

    def calculate_safety_factor(self) -> int:
        q1, q2, q3, q4 = self.count_robots_in_quadrants()
        return q1 * q2 * q3 * q4

def main():
    # Simulation für die echte Eingabe
    sim = RobotSimulation(101, 103)
    sim.parse_input('./day14.txt')
    sim.simulate(100)
    safety_factor = sim.calculate_safety_factor()
    print(f"Safety factor after 100 seconds: {safety_factor}")

    # Test mit Beispieldaten
    test_sim = RobotSimulation(11, 7)
    test_data = [
        Robot(0, 4, 3, -3),
        Robot(6, 3, -1, -3),
        Robot(10, 3, -1, 2),
        Robot(2, 0, 2, -1),
        Robot(0, 0, 1, 3),
        Robot(3, 0, -2, -2),
        Robot(7, 6, -1, -3),
        Robot(3, 0, -1, -2),
        Robot(9, 3, 2, 3),
        Robot(7, 3, -1, 2),
        Robot(2, 4, 2, -3),
        Robot(9, 5, -3, -3)
    ]
    for robot in test_data:
        test_sim.add_robot(robot)
    test_sim.simulate(100)
    test_safety_factor = test_sim.calculate_safety_factor()
    print(f"Test safety factor after 100 seconds: {test_safety_factor}")

if __name__ == "__main__":
    main()
