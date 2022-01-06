# Day 19 - Beacon Scanner

[Full task description](https://adventofcode.com/2021/day/19)

Various coordinate systems need to be aligned into one contiguous 3D space.
Each coordinate system covers a portion of the same total volume and contains reference points.
The origin and orientation of each coordinate system are unknown.
If two coordinate systems intersect, the intersection contains at least 12 reference points. 

## Solution

There are 24 possible rotations for each coordinate system (Scanner).
Since the reference points (Beacons) are relative to the Scanner's origin, a rotation can be obtained by [swapping and negating](main.go#L25) the 3 coordinates components of the Beacons.
Aligning two Scanners can be achieved by assuming any two Beacons are the same Beacons.
The first Scanner does not need to be rotated, since one of the 24 rotations of the second Scanner will match the rotation of the first.
For each rotation of the second Scanner, both Scanners are translated such that the tested beacon becomes the origin of the scanner.
In the intersecting volume of the two scanners, there must now be at least 12 beacons.
In each scanner, these beacons are all relative to the tested beacon, and thus, if we found the right rotation and the beacons are indeed the same, all other beacons in the intersection must also have the same coordinates.

### Optimization

The brute-force solution is to perform this test with all beacons of the first Scanner against all beacons of the second Scanner.
But we can greatly reduce the number of beacons to test, by first computing the distances of each beacon to its neighbors.
The distances stay the same regardless of rotation and origin.
There are a number of flaws with this optimization:

- For Beacons right at the edge of the intersection volume, the closest neighbors could be outside the intersection.
  However, it would be enough for a single of the 12 matching Beacons to have the same 2 closest neighbors that are within the intersection for the optimization to produce enough match candidates.
- If the Scanners contain the same constellation of Beacons more than once (patterns), the optimization doesn't make so mush sense.

### Challenges

One of the challenges of performing the test with the intersection volume is that it needs to be done with the original scanners.
It wouldn't work well when integrating the beacons of aligned Scanners into one big volume, since then the total volume of the combined space has unmapped holes.
Therefore, the intersection of volumes could contain all of the Scanner which is currently being integrated.
This means it does not contain the same beacons in the combined space and the Scanner. 
For this reason, the scanners are [kept separate](main.go#L200) in the CombinedScanners.
