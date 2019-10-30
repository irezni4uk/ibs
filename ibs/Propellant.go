package ibs

// import "fmt"

type BurnFun func(float64) float64

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
	psi             BurnFun
}

func (p *Propellant) Volume() float64 {
	// return (1-p.Z)*p.Mass/p.Density + p.Z*p.Mass*p.Covolume
	return p.Mass/p.Density + p.GasMass()*(p.Covolume-1/p.Density)
}

// func (p *Propellant) Pressure(Vol float64) float64 {
// 	return p.Mass * p.Z * p.Force / Vol
// }

func (p *Propellant) GasMass() float64 {
	// return p.Mass * p.Z
	return p.Mass * p.psi(p.Z)
}

func (p *Propellant) FullForce() float64 {
	return p.Force * p.GasMass()
}

func (p *Propellant) HeatCapacity() float64 {
	return p.FullForce() * p.AdiabaticIndex / (p.AdiabaticIndex - 1) / p.BurnTemperature
}

func (p *Propellant) Reset() {
	if p.IsPrimer {
		p.Z = 1
	} else {
		p.Z = 0
	}
}

//Burn integrates burned web over time
func (p *Propellant) Burn(Pmean float64) {
	p.Z += Pmean / p.Impulse * dt
}

//NewPropellant returns 1 kg of '16/1 тр В/А' propellant
func NewPropellant() Propellant {
	// fmt.Println(dt)

	out := Propellant{}

	out.Mass = 1
	out.Density = 1600
	out.Force = 1.015e6
	out.Impulse = 1.04e6
	out.BurnTemperature = 2940
	out.AdiabaticIndex = 1.224
	out.Covolume = 1.009e-3
	out.Z = 0
	out.IsPrimer = false
	// out.IsPrimer = true
	out.psi = PsiFun(1, 1, 0, 0, 0)

	return out
}

//PsiFun returns function returning burned fraction of propellant depending on burned web fraction
func PsiFun(zk, k1, l1, k2, l2 float64) func(z float64) (psi float64) {
	return func(z float64) (psi float64) {
		if z >= zk {
			return 1
		}
		tmp := z - 1
		if z > 1 {
			z = 1
		}
		psi = k1 * z * (1 + l1*z)
		if z < 1 || zk == 1 {
			return psi
		}
		z = tmp
		return psi + k2*z*(1+l2*z)
	}
}
