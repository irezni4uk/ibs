package ibs

import "fmt"

// this is a comment

type Charge struct {
	Propellant []Propellant
	// Projectile *Projectile
	// Solution *Solution
}

func (c *Charge) State(s *State) {
	s.Tmean, s.Pmean = c.Thermodynamics(s.Volume, s.EnergyLoss)
	s.HeatCapacity = c.HeatCapacity()
	s.GasMass = c.GasMass()
}

func (c *Charge) HeatCapacity() (out float64) {
	var s1, s2 float64
	for _, p := range c.Propellant {
		s1 += p.Force * p.Z * p.Mass * p.AdiabaticIndex / (p.AdiabaticIndex - 1) / p.BurnTemperature
		s2 += p.Mass * p.Z
	}
	return s1 / s2
}

func (c *Charge) KineticEnergy(Vproj float64) float64 {
	return c.Mass() * Vproj * Vproj / 6
}

func (c *Charge) Reset() {
	for i, _ := range c.Propellant {
		p := &c.Propellant[i]
		if p.IsPrimer {
			p.Z = 1
		} else {
			p.Z = 0
		}
	}
}

func (c *Charge) GasMass() (out float64) {
	for _, p := range c.Propellant {
		out += p.Mass * p.Z
	}
	return out
}

func (c *Charge) Mass() (out float64) {
	for _, p := range c.Propellant {
		out += p.Mass
	}
	return out
}

func (c *Charge) Volume() (out float64) {
	for _, p := range c.Propellant {
		out += p.Volume()
	}
	return out
}

func (c *Charge) Thermodynamics(Vol, Enloss float64) (Tmean, Pmean float64) {
	var s1, s2, s3, out float64
	for _, p := range c.Propellant {
		out = p.Force * p.Z * p.Mass / (p.AdiabaticIndex - 1)
		s1 += out
		s2 += out / p.BurnTemperature
		s3 += out * (p.AdiabaticIndex - 1) / p.BurnTemperature
	}
	Tmean = (s1 - Enloss) / s2
	Pmean = Tmean / Vol * s3
	return Tmean, Pmean
}

func (c *Charge) Burn(Pmean float64) {
	var z float64
	for i, _ := range c.Propellant {
		p := &c.Propellant[i]
		if p.IsPrimer {
			continue
		}
		// z = p.Z + Pmean/1.04e6*dt
		z = p.Z + Pmean/p.Impulse*dt
		if z > 1 {
			p.Z = 1
		} else {
			p.Z = z
		}
		// fmt.Println(z)
	}
}

func NewCharge() Charge {

	out := Charge{}
	// out.Propellant = make([]Propellant, 4)
	out.Propellant = append(out.Propellant, Propellant{7e-3, 1700, 260e3, .1e6, 2427, 1.22, .0006, 1, true})
	out.Propellant = append(out.Propellant, NewPropellant())
	fmt.Println(out.Propellant)
	return out
}
