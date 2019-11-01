package ibs

import (
	// "fmt"
	"testing"
)

func TestInternalBallisticsSimulator(t *testing.T) {
	Vtarget := 1072

	p := Projectile{Mass: 1}
	b := NewBarrel()
	c := NewCharge()
	i := InternalBallisticsSimulator{
		Barrel:     &b,
		Projectile: &p,
		Charge:     &c,
		Params:     &SimParams{ForcingPressure: 100e6},
	}
	i.LinkComponents()
	s := i.RunSym()
	Velocity := int(s[len(s)-1].Velocity)
	// fmt.Printf("%v\t%v\t%v\n", Velocity, Velocity-1872, Velocity-1872 > 1)
	if Velocity != Vtarget {
		t.Errorf("Vmuzzle = %v m/s, expected %v m/s", Velocity, Vtarget)
	}

	Vtarget = 1507
	c.Propellant[1] = Propellant{Mass: 1, Force: 1006000, Impulse: 500000, Density: 1600, AdiabaticIndex: 1.224, Covolume: 1.01e-3, BurnTemperature: 2900, Psi: PsiFun(1.607, .769, .101, .506, -.823)}
	s = i.RunSym()
	Velocity = int(s[len(s)-1].Velocity)
	if Velocity != Vtarget {
		t.Errorf("Vmuzzle = %v m/s, expected %v m/s", Velocity, Vtarget)
	}
}
