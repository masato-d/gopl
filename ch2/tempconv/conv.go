package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) KelvinScale {
	return KelvinScale(-1 * AbsoluteZeroC + c)
}

func KToC(k KelvinScale) Celsius {
	return Celsius(k - KelvinScale(AbsoluteZeroC))
}
