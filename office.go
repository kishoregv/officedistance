package main

import "math"

const planetRadius float64 = 6371

type office struct {
	Location location
}

func (o *office) getCustomersInRadius(customers []*customer, radius float64) []*customer {
	inRadiusCustomers := []*customer{}
	for i := 0; i < len(customers); i++ {
		cl := customers[i].location
		distance := o.getDistance(cl)
		if distance < radius {
			inRadiusCustomers = append(inRadiusCustomers, customers[i])
		}
	}
	return inRadiusCustomers
}

func (o *office) getDistance(customerLocation location) float64 {
	phi1, lam1 := customerLocation.toRadian()
	phi2, lam2 := o.Location.toRadian()
	dLam := math.Abs(lam1 - lam2)

	sin1, cos1 := math.Sincos(phi1)
	sin2, cos2 := math.Sincos(phi2)
	cosDLam := math.Cos(dLam)

	return planetRadius * math.Acos(sin1*sin2+cos1*cos2*cosDLam)
}

// distance calculation using the Spherical Law of Cosines.
func (o *office) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(radian(p1.lat))
	s2, c2 := math.Sincos(radian(p2.lat))
	clong := math.Cos(radian(p1.long - p2.long))
	return planetRadius * math.Acos(s1*s2+c1*c2*clong)
}
