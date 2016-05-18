package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	ZeroK    Kilogram = 0
	Hundreds Kilogram = 100
)

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%gp", p) }

func KToP(k Kilogram) Pound { return Pound(k * 2.2046) }
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
