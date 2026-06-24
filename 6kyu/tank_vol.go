// To introduce the problem think to my neighbor who drives a tanker truck. The level
// indicator is down and he is worried because he does not know if he will be able
// to make deliveries. We put the truck on a horizontal ground and measured the
// height of the liquid in the tank.
//
// Fortunately the tank is a perfect cylinder and the vertical walls on each end
// are flat. The height of the remaining liquid is h, the diameter of the
// cylinder base is d, the total volume is vt (h, d, vt are positive or null
// integers). You can assume that h <= d.
//
// Could you calculate the remaining volume of the liquid? Your function
// tankvol(h, d, vt) returns an integer which is the truncated result (i.e
// floor) of your float calculation.
//
// Examples:
// tankvol(40,120,3500) should return 1021 (calculation gives about: 1021.26992027)
// tankvol(60,120,3500) should return 1750
// tankvol(80,120,3500) should return 2478 (calculation gives about: 2478.73007973)

package main

import "math"

// tankvol returns the truncated volume of liquid remaining in a horizontal
// cylindrical tank with diameter d, filled to height h, where vt is the total
// tank volume.
func tankvol(h, d, vt int) int {
	if h <= 0 {
		return 0
	}
	if h >= d {
		return vt
	}

	r := float64(d) / 2.0
	hh := float64(h)

	// Area of a circular segment for depth hh in circle radius r:
	// A = r^2 * acos((r-h)/r) - (r-h) * sqrt(h*(2r-h))
	theta := math.Acos((r - hh) / r)
	area := r*r*theta - (r-hh)*math.Sqrt(hh*(2*r-hh))

	// Cylinder length from total volume vt = pi * r^2 * L
	L := float64(vt) / (math.Pi * r * r)

	vol := area * L
	return int(vol)
}