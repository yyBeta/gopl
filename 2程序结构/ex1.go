package tempconv

type KelvinTemperature float64
type CelsiusTemperature float64

const (
	AbsoluteZeroC Celsius = -273.15
)

func KtoC(k KelvinTemperature) CelsiusTemperature {
	return CelsiusTemperature(k - AbsoluteZeroC)
}

func CtoC(c CelsiusTemperature) KelvinTemperature {
	return KelvinTemperature(c + AbsoluteZeroC)
}
