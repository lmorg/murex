package adjacency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
nbutton: Really the value is not as important to me than they don't change, which happened during development.
*/
func TestCalculateDegreeQwert(t *testing.T) {
	avgDegreeQwert := BuildQwerty().CalculateAvgDegree()

	assert.Equal(t, float64(1.5319148936170213), avgDegreeQwert, "Avg degree for qwerty should be 1.5319148936170213")
}

func TestCalculateDegreeDvorak(t *testing.T) {
	avgDegreeQwert := BuildDvorak().CalculateAvgDegree()

	assert.Equal(t, float64(1.5319148936170213), avgDegreeQwert, "Avg degree for dvorak should be 1.53191489361702135")
}

func TestCalculateDegreeKeypad(t *testing.T) {
	avgDegreeQwert := BuildKeypad().CalculateAvgDegree()

	assert.Equal(t, float64(0.6333333333333333), avgDegreeQwert, "Avg degree for keypad should be 0.6333333333333333")
}

func TestCalculateDegreeMacKepad(t *testing.T) {
	avgDegreeQwert := BuildMacKeypad().CalculateAvgDegree()

	assert.Equal(t, float64(0.65625), avgDegreeQwert, "Avg degree for mackeyPad should be 0.65625")
}
