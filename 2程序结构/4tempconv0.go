package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC))
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == Celsius(f))
	cc := FtoC(212.0)
	fmt.Println(cc.String())
	fmt.Printf("%v\n", cc)
	fmt.Println(cc)        // "100°C"
	fmt.Printf("%g\n", cc) // "100"; does not call String
}
