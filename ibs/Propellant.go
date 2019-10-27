package ibs

// import "fmt"

type Propellant struct {
	Mass            float64
	Density         float64
	Force           float64
	Impulse         float64
	BurnTemperature float64
	AdiabaticIndex  float64
	Covolume        float64
	Z               float64
	IsPrimer        bool
}

func (p *Propellant) Volume() float64 {
	return (1-p.Z)*p.Mass/p.Density + p.Z*p.Mass*p.Covolume
}

// func (p *Propellant) Pressure(Vol float64) float64 {
// 	return p.Mass * p.Z * p.Force / Vol
// }

func (p *Propellant) GasMass() float64 {
	return p.Mass * p.Z
}

func (p *Propellant) HeatCapacity() float64 {
	return p.Force * p.Z * p.Mass * p.AdiabaticIndex / (p.AdiabaticIndex - 1) / p.BurnTemperature
}

func (p *Propellant) Reset() {
	if p.IsPrimer {
		p.Z = 1
	} else {
		p.Z = 0
	}
}

func (p *Propellant) Burn(Pmean float64) {
	if p.IsPrimer {
		return
	}
	z := p.Z + Pmean/p.Impulse*dt
	if z > 1 {
		p.Z = 1
	} else {
		p.Z = z
	}
}

func NewPropellant() Propellant {
	// fmt.Println(dt)

	out := Propellant{}

	out.Mass = 1
	out.Density = 1600
	out.Force = 1.015e6
	out.Impulse = 1.04e6
	out.BurnTemperature = 2940
	out.AdiabaticIndex = 1.224
	out.Covolume = 1e-3
	out.Z = 0
	out.IsPrimer = false

	return out
}
