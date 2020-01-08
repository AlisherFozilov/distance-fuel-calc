package distanceFuelCalc

import "testing"

func Test_distanceFuelCalc(t *testing.T) {
	tests := []struct {
		name               string
		litersPerHundredKm float64
		fuelInLiters       float64
		want               int
	}{
		{"Distance more 20 km has inaccuracy", 10.0, 100.0, 950},
		{"Distance less 20 km hasn't inaccuracy", 30.0, 5.0, 16},
	}
	for _, test := range tests {
		got := distanceFuelCalc(test.litersPerHundredKm, test.fuelInLiters)
		if got != test.want {
			t.Error("For distanceFuelCalc test:", test.name, "got:", got, "want:", test.want)
		}
	}
}
