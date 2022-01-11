# Day 19 - Beacon Scanner

[Full task description](https://adventofcode.com/2021/day/19)

Various coordinate systems need to be aligned into one contiguous 3D space.
Each coordinate system covers a portion of the same total volume and contains reference points.
The origin and orientation of each coordinate system are unknown.
If two coordinate systems intersect, the intersection contains at least 12 reference points. 

## Solution

There are 24 possible rotations for each coordinate system (Scanner).
Since the reference points (Beacons) are relative to the Scanner's origin, a rotation can be obtained by [swapping and negating](main.go#L25) the 3 coordinate components of the Beacons.
Aligning two Scanners can be achieved by assuming any two of their respective Beacons are the same Beacon.
The first Scanner does not need to be rotated, since one of the 24 rotations of the second Scanner will match the rotation of the first.
For each rotation of the second Scanner, both Scanners are translated such that the tested Beacon becomes the origin of the Scanner.
In the intersecting volume of the two Scanners, there must now be at least 12 Beacons.
In each Scanner, these Beacons are all relative to the tested Beacon, and thus, if we found the right rotation and the Beacons are indeed the same, all other Beacons in the intersection must also have the same coordinates.

### Optimization

The brute-force solution is to perform this test with all Beacons of the first Scanner against all Beacons of the second Scanner.
But we can greatly reduce the number of Beacons to test, by first computing the distances of each Beacon to its neighbors.
The distances stay the same regardless of rotation and origin.
We can sort the distance and store the 2 closest neighbors as a string for quick comparison in each Beacon.

There are a number of potential flaws with this optimization:

- For Beacons right at the edge of the intersection volume, the closest neighbors could be outside the intersection.
  However, it would be enough for a single of the 12 matching Beacons to have the same 2 closest neighbors that are within the intersection for the optimization to produce enough match candidates.
  Still, if every single one of the 12 Beacons in the intersection volume had a closest neighbor outside the intersection volume, the optimization would break. 
- If the Scanners contain the same constellation of Beacons more than once (patterns), the optimization doesn't make so much sense.

### Challenges

One of the challenges of performing the test with the intersection volume is that it needs to be done with the original Scanners.
It wouldn't work well when integrating the Beacons of aligned Scanners into one big volume, since then the total volume of the combined space has unmapped holes.
Therefore, the intersection of volumes could contain all of the Scanner which is currently being integrated.
This means it does not contain the same Beacons in the combined space and the Scanner. 
For this reason, the Scanners are [kept separate](main.go#L200) in the CombinedScanners.

### Part 2

The solution for part 2 becomes easy by treating the origin of each Scanner like an additional Beacon.
This means it is [translated along](main.go#L144) with the other Beacons of the Scanner.
