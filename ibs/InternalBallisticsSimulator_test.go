package ibs

import (
	"testing"
)

func TestInternalBallisticsSimulator(t *testing.T) {
	Vtarget := 1528

	p := Projectile{Mass: 10, ForcingPressure: 100e6}
	b := NewBarrel(.125, .125, 1)
	c := NewCharge()
	i := InternalBallisticsSimulator{
		Barrel:     &b,
		Projectile: &p,
		Charge:     &c,
	}
	s := i.RunSim()
	Velocity := int(s[len(s)-1].Velocity)
	if Velocity != Vtarget {
		t.Errorf("Vmuzzle = %v m/s, expected %v m/s", Velocity, Vtarget)
	}

	// Progressive propellant grain test
	Vtarget = 1665
	c.Propellant[1] = Propellant{Mass: 8, Force: 1006000, Impulse: 500000, Density: 1600, AdiabaticIndex: 1.224, Covolume: 1.01e-3, BurnTemperature: 2900,
		Psi: PsiFun(1.607, .769, .101, .506, -.823)}
	s = i.RunSim()
	Velocity = int(s[len(s)-1].Velocity)
	if Velocity != Vtarget {
		t.Errorf("Vmuzzle = %v m/s, expected %v m/s", Velocity, Vtarget)
	}

	// Closed chamber test
	c.Propellant[0] = c.Propellant[1]
	c.Propellant[0].IsPrimer = true
	c.Propellant[0].Mass = .1
	p.ForcingPressure = 1e12
	Ptarget := int(c.Mass()*c.Propellant[0].Force/(b.Volume-c.Mass()*c.Propellant[0].Covolume)/1e5 + 1)
	s = i.RunSim()
	Pressure := int(s[len(s)-1].Pmean / 1e5)
	if Ptarget != Pressure {
		t.Errorf("Pmean = %v atm, expected %v atm", Pressure, Ptarget)
	}
}
