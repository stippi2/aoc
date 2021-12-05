package main

import "strconv"

func (d *DangerMap) String() string {
	result := make([]byte, (d.width+1)*d.height-1)
	offsetDanger := 0
	offsetResult := 0
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			danger := d.danger[offsetDanger]
			if danger == 0 {
				result[offsetResult] = '.'
			} else {
				result[offsetResult] = strconv.Itoa(danger)[0]
			}
			offsetDanger++
			offsetResult++
		}
		if y < d.height-1 {
			result[offsetResult] = '\n'
			offsetResult++
		}
	}
	return string(result)
}
