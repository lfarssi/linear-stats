package functions

import "math"

func Average(n []float64) float64 {
	var res float64
	for _, item := range n {
		res += float64(item)
	}
	return res/float64(len(n))
}
func Variance(arr []float64) float64 {
	var sum float64 
	avg := Average(arr)
	for i:= 0 ; i< len(arr) ; i++ {
		sum += (float64(arr[i]) - avg) * (float64(arr[i]) - avg)
	}
	return  sum / float64(len(arr))
}
func StandardDeviation(arr []float64) float64 {
	return  math.Sqrt(Variance(arr))
}

func LinearRegLine(arr []float64) (float64,float64) {
	avg := Average(arr)	
	total := float64(len(arr))
	med := (total-1)/2
	var taba_3od, taba3od_b_indice float64 
	for i:= 0; i< len(arr); i++ {
		taba_3od = (arr[i]-avg)*(float64(i)-med)
		taba3od_b_indice = (float64(i)-med)*(float64(i)-med)
	}
	cofX := taba_3od/taba3od_b_indice
	cofY := avg- cofX*med
	return cofX, cofY
}

func PearsonCoef(arr []float64) float64 {
	
	avg := Average(arr)
	total := float64(len(arr))

	med := (total - 1) / 2 

	varX := 0.0
	for i := 0; i < len(arr); i++ {
		varX += (med - float64(i)) * (med - float64(i))
	}
	varX /= total
	stdX := math.Sqrt(varX)
	stdY := StandardDeviation(arr)

	eXY := 0.0
	for i := 0; i < len(arr); i++ {
		eXY += (arr[i] - avg) * (float64(i) - med)
	}
	eXY /= total

	r := eXY / (stdX * stdY)
	return r
	}