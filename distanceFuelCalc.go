package distanceFuelCalc

func distanceFuelCalc(litersPerHundredKm float64, fuelInLiters float64) int {
	const percent = 5

	litersPerKm := litersPerHundredKm / 100.0
	var distance int = int(fuelInLiters / litersPerKm)
	var inaccuracy int = distance * percent / 100
	distance -= inaccuracy
	return distance
}