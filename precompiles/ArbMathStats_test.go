package precompiles

import (
	"math/big"
	"testing"
)

func TestArbMathStatsStandardDeviation(t *testing.T) {
	mathStats := ArbMathStats{}

	data := []int32{1, 1, 1, 1, 1, 1}
	stdev, _ := mathStats.stdDev(nil, nil, data, 0)

	if stdev.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected stdev of 0 for identical elements, got: %s", stdev)
	}
}
