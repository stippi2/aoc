package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var example = `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 1 ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0`

var expected = []Scanner {
	{
		beacons: []Beacon{
			{position: Position{-1,-1,1}},
			{position: Position{-2,-2,2}},
			{position: Position{-3,-3,3}},
			{position: Position{-2,-3,1}},
			{position: Position{5,6,-4}},
			{position: Position{8,0,7}},
		},
	},
	{
		beacons: []Beacon{
			{position: Position{1,-1,1}},
			{position: Position{2,-2,2}},
			{position: Position{3,-3,3}},
			{position: Position{2,-1,3}},
			{position: Position{-5,4,-6}},
			{position: Position{-8,-7,0}},
		},
	},
}

func Test_parseInput(t *testing.T) {
	assert.Equal(t, expected, parseInput(example))
}

var scannerExamples = `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`

func Test_compareBeacons(t *testing.T) {
	scanners := parseInput(scannerExamples)
	for _, s := range scanners {
		s.setBeaconDistances()
	}
	for i := 0; i < len(scanners); i++ {
		referenceScanner := &scanners[i]
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}
			matchingBeacons := 0
			for _, refBeacon := range referenceScanner.beacons {
				for _, otherBeacon := range scanners[j].beacons {
					if refBeacon.distancesToNearest == otherBeacon.distancesToNearest {
						matchingBeacons++
					}
				}
			}
			fmt.Printf("common beacons scanner %v and scanner %v: %v of %v\n", i, j, matchingBeacons, len(scanners[j].beacons))
		}
	}
}

func Test_beaconsInVolume(t *testing.T) {
	scanner := Scanner{}
	scanner.appendBeacon(1, 1, 1)
	scanner.appendBeacon(-1, -1, -1)
	beacons := scanner.getBeaconsInVolume(Volume{
		min: Position{0, 0, 0},
		max: Position{2, 2, 2},
	})
	if assert.Len(t, beacons, 1) {
		assert.Equal(t, Position{1, 1, 1}, beacons[0].position)
	}
}

func Test_volumeIntersect(t *testing.T) {
	a := Volume{
		min: Position{-1, -2, -3},
		max: Position{3, 4, 5},
	}
	b := Volume{
		min: Position{0, 1, 2},
		max: Position{6, 7, 8},
	}
	i1, valid1 := a.intersect(b)
	i2, valid2 := b.intersect(a)

	assert.True(t, valid1)
	assert.True(t, valid2)

	assert.Equal(t, i1, i2)
	assert.Equal(t, i1, Volume{
		min: Position{0, 1, 2},
		max: Position{3, 4, 5},
	})
}

func Test_volumeIntersectInvalid(t *testing.T) {
	a := Volume{
		min: Position{-1, -2, -3},
		max: Position{0, 1, 2},
	}
	b := Volume{
		min: Position{3, 4, 5},
		max: Position{6, 7, 8},
	}
	_, valid := a.intersect(b)

	assert.False(t, valid)
}

func Test_combineScanners(t *testing.T) {
	scanners := parseInput(scannerExamples)
	combined := &CombinedScanners{}
	for len(scanners) > 0 {
		integratedOne := false
		for i, s := range scanners {
			if combined.integrate(&s) {
				last := len(scanners)-1
				scanners[i] = scanners[last]
				scanners = scanners[:last]
				integratedOne = true
				fmt.Printf("integrated scanner %v of %v\n", i, last + 1)
				break
			}
		}
		require.True(t, integratedOne)
	}
	assert.Equal(t, 79, len(combined.allBeacons()))
	assert.Equal(t, 3621, combined.largestManhattanDistance())
}

func Test_alignScannersBruteForce(t *testing.T) {
	scanners := parseInput(scannerExamples)
	for _, s := range scanners {
		s.setBeaconDistances()
	}
	for i := 0; i < len(scanners); i++ {
		referenceScanner := &scanners[i]
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}
			for rIndexRef, rotationRef := range referenceScanner.rotations() {
				for rIndex, rotation := range scanners[j].rotations() {
					for _, refBeacon := range rotationRef.beacons {
						for _, matchBeacon := range rotation.beacons {
							// If we found an alignment, we can transform both scanners to have the matching beacon as origin,
							// then form the intersecting volume, and all beacons within the intersection need to match
							matchOrigin := matchBeacon.position
							translatedMatch := rotation.translateBy(matchOrigin.x, matchOrigin.y, matchOrigin.z)

							refOrigin := refBeacon.position
							translatedRef := rotationRef.translateBy(refOrigin.x, refOrigin.y, refOrigin.z)

							intersection, overlap := translatedMatch.volume().intersect(translatedRef.volume())
							if !overlap {
								// Should not be possible when we translated both scanners to the same origin
								continue
							}
							matchBeaconsInIntersection := translatedMatch.getBeaconsInVolume(intersection)
							if len(matchBeaconsInIntersection) < 12 {
								continue
							}
							refBeaconsInIntersection := translatedRef.getBeaconsInVolume(intersection)

							if containsSameBeacons(matchBeaconsInIntersection, refBeaconsInIntersection) {
								offset := Position{
									x: matchOrigin.x - refOrigin.x,
									y: matchOrigin.y - refOrigin.y,
									z: matchOrigin.z - refOrigin.z,
								}
								fmt.Printf("found alignment between scanner %v and %v at rotation %v/%v and offset %v with %v beacons\n", i, j, rIndexRef, rIndex, offset, len(matchBeaconsInIntersection))
								break
							}
						}
					}
				}
			}
		}
	}
}

func Test_rotations(t *testing.T) {
	type vector []int

	inversions := []vector{
		{ 1, 1, 1},
		{ 1, 1,-1},
		{ 1,-1, 1},
		{-1, 1, 1},
		{-1,-1, 1},
		{-1, 1,-1},
		{ 1,-1,-1},
		{-1,-1,-1},
	}

	mappings := []vector{
		{ 0, 1, 2},
		{ 0, 2, 1},
		{ 1, 0, 2},
		{ 1, 2, 0},
		{ 2, 0, 1},
		{ 2, 1, 0},
	}

	rotations := map[string]bool{}
	axis := []string{"p.x", "p.y", "p.z"}

	for _, i := range inversions {
		for _, m := range mappings {
			x := axis[m[0]]
			y := axis[m[1]]
			z := axis[m[2]]
			if i[0] < 0 {
				x = "-" + x
			}
			if i[1] < 0 {
				y = "-" + y
			}
			if i[2] < 0 {
				z = "-" + z
			}
			rotations[x + "," + y + "," + z] = true
		}
	}

	for rotation := range rotations {
		fmt.Printf("{%s},\n", rotation)
	}
}
