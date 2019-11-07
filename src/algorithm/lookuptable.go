package algorithm

//import "fmt"
//import "math"

type LookupTable_t struct {

	DBKey string
}

func (o* LookupTable_t)Compute(currentTemperature float32) float32 {
	val := currentTemperature
	if val < 25 {
		return 20
	} else if (val >= 25) && (val < 27) {
		return 25
	} else if (val >= 27) && (val < 29) {
		return 27
	} else if (val >= 29) && (val < 31) {
		return 29
	} else if (val >= 31) && (val < 33) {
		return 35	
	} else if (val >= 33) && (val < 35) {
		return 40
	} else if (val >= 35) && (val < 37) {
		return 45
	} else if (val >= 37) && (val < 43) {
		return 50
	} else if (val >= 43)  {
		return 60
	}
	return val
}




