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

func (c *Charge) HeatFlux(Vol, Vproj float64) float64 {
	return c.HeatCapacity() / Vol * c.Velocity(Vproj)
}

func (c *Charge) HeatCapacity() (out float64) {
	for _, p := range c.Propellant {
		out += p.HeatCapacity()
	}
	return out
}

func (c *Charge) Velocity(Vproj float64) float64 {
	return 1 / 2 * Vproj
}

func (c *Charge) KineticEnergy(Vproj float64) float64 {
	return c.Mass() * Vproj * Vproj / 6
}

func (c *Charge) Reset() {
	for i := range c.Propellant {
		c.Propellant[i].Reset()
	}
}

func (c *Charge) GasDens(Vol float64) float64 {
	return c.GasMass() / Vol
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
	for i := range c.Propellant {
		c.Propellant[i].Burn(Pmean)
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
