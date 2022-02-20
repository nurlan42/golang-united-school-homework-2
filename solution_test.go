package square

import (
	"testing"
)

func TestCalcSquare(t *testing.T) {

	tests := []struct {
		sideLen  float64
		sidesNum sidesNum
		want     float64
	}{
		{sideLen: 4, sidesNum: SidesSquare, want: 16},
		{sideLen: 4, sidesNum: SidesTriangle, want: 6.928203230275509},
		{sideLen: 4, sidesNum: SidesCircle, want: 50.26548245743669},
	}

	for _, test := range tests {
		if got := CalcSquare(test.sideLen, test.sidesNum); test.want != got {
			t.Errorf("CalcSquare(), got: %v, wanted: %v", got, test.want)
		}
	}

}
