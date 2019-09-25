package ibs

// import "fmt"

// this is a comment

type Propellant struct {
	Mass            float64
	Density         float64
	Force           float64
	BurnTemperature float64
	AdiabaticIndex  float64
	Covolume        float64
	Z               float64
	IsPrimer        bool
}

func (p *Propellant) Volume() float64 {
	return (1-p.Z)*p.Mass/p.Density + p.Z*p.Mass*p.Covolume
}

func (p *Propellant) Pressure(Vol float64) float64 {
	return p.Mass * p.Z * p.Force / Vol
}

func NewPropellant() Propellant {
	// fmt.Println(dt)

	out := Propellant{}

	out.Mass = 1
	out.Density = 1600
	out.Force = 1e6
	out.BurnTemperature = 3000
	out.AdiabaticIndex = 1.224
	out.Covolume = 1e-3
	out.Z = 0
	out.IsPrimer = false

	return out
}
