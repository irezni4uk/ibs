package ibs

//Charge contains slice of propellants and is designed to manage them
type Charge struct {
	Propellant []Propellant
}

func (c *Charge) state(s *State) {
	s.Tmean, s.Pmean = c.thermodynamics(s.Volume, s.EnergyLoss)
	s.HeatCapacity = c.heatCapacity()
	s.GasMass = c.gasMass()
}

func (c *Charge) heatFlux(Vol, Vproj float64) float64 {
	return c.heatCapacity() / Vol * c.Velocity(Vproj)
}

func (c *Charge) heatCapacity() (out float64) {
	for _, p := range c.Propellant {
		out += p.HeatCapacity()
	}
	return out // / c.GasMass()
}

func (c *Charge) Velocity(Vproj float64) float64 {
	return .5 * Vproj
}

func (c *Charge) KineticEnergy(Vproj float64) float64 {
	return c.Mass() * Vproj * Vproj / 6
}

func (c *Charge) reset() {
	for i := range c.Propellant {
		c.Propellant[i].reset()
	}
}

func (c *Charge) gasDens(Vol float64) float64 {
	return c.gasMass() / Vol
}

func (c *Charge) gasMass() (out float64) {
	for _, p := range c.Propellant {
		out += p.gasMass()
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

func (c *Charge) thermodynamics(Vol, Enloss float64) (Tmean, Pmean float64) {
	var s1, s2, s3 float64
	for _, p := range c.Propellant {
		f := p.FullForce()
		s1 += f / (p.AdiabaticIndex - 1)
		s2 += f / p.BurnTemperature / (p.AdiabaticIndex - 1)
		s3 += f / p.BurnTemperature
	}
	Tmean = (s1 - Enloss) / s2
	Pmean = Tmean / Vol * s3
	return Tmean, Pmean
}

//Burn calls Burn method for Charge components
func (c *Charge) burn(Pmean float64) {
	for i := range c.Propellant {
		c.Propellant[i].burn(Pmean)
	}
}

//NewCharge returns charge containing 2 components one of which is 7 g primer
func NewCharge() Charge {
	primer := Propellant{.4712e-2, 1700, .845535E2 * 1e3, .1e6, 294, 1.4, .9755e-3, 1, true, PsiFun(1, 1, 0, 0, 0)}

	out := Charge{}
	out.Propellant = append(out.Propellant, primer)
	// out.Propellant = append(out.Propellant, Propellant{7e-3, 1700, 260e3, .1e6, 2427, 1.22, .0006, 1, true, PsiFun(1, 1, 0, 0, 0)})
	out.Propellant = append(out.Propellant, NewPropellant())
	return out
}
