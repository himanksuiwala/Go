package tempconv

type Celsius float64
type Farhanheit float64

func CToF(c Celsius) Farhanheit {
	return Farhanheit(c*9/5 + 32)
}

func FToC(f Farhanheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
