import "math"

func mediaComPi(ptr *float64) {
    avg := (*ptr + math.Pi) / 2
    *ptr = avg
}
