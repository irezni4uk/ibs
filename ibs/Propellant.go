package ibs

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
	Psi             BurnFun
}

func (p *Propellant) Volume() float64 {
	return p.Mass/p.Density + p.gasMass()*(p.Covolume-1/p.Density)
}

func (p *Propellant) gasMass() float64 {
	return p.Mass * p.Psi(p.Z)
}

func (p *Propellant) FullForce() float64 {
	return p.Force * p.gasMass()
}

func (p *Propellant) HeatCapacity() float64 {
	return p.FullForce() * p.AdiabaticIndex / (p.AdiabaticIndex - 1) / p.BurnTemperature
}

func (p *Propellant) reset() {
	if p.IsPrimer {
		p.Z = 1
	} else {
		p.Z = 0
	}
}

//Burn integrates burned web over time
func (p *Propellant) burn(Pmean float64) {
	p.Z += Pmean / p.Impulse * dt
}

//NewPropellant returns 1 kg of '16/1 тр В/А' propellant
func NewPropellant() Propellant {

	out := Propellant{}

	// out.Mass = 1
	out.Mass = 8
	out.Density = 1600
	out.Force = 1.015e6
	out.Impulse = 1.04e6
	out.BurnTemperature = 2940
	out.AdiabaticIndex = 1.224
	out.Covolume = 1.009e-3
	out.Z = 0
	out.IsPrimer = false
	out.Psi = PsiFun(1, 1, 0, 0, 0)

	return out
}

//PsiFun returns function returning burned fraction of propellant depending on burned web fraction(nondimensional)
func PsiFun(zk, k1, l1, k2, l2 float64, varIn ...float64) func(z float64) (psi float64) {
	m1 := 0.
	if len(varIn) > 0 {
		m1 = varIn[0]
	}
	return func(z float64) (psi float64) {
		if z >= zk {
			return 1
		}
		tmp := z - 1
		if z > 1 {
			z = 1
		}
		psi = k1 * z * (1 + l1*z + m1*z*z)
		if z < 1 || zk == 1 {
			return psi
		}
		z = tmp
		return psi + k2*z*(1+l2*z)
	}
}
