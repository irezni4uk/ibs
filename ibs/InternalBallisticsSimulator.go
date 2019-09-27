package ibs

import "fmt"

const (
	wallTime = 20e-3
	dt       = 5e-6

	h0 = 11.35 //Free convective heat transfer coefficient, zero speed W/m^2-K
)

type InternalBallisticsSimulator struct {
	Barrel     *Barrel
	Projectile *Projectile
	Charge     *Charge
	Params     *SimParams
}

type SimParams struct {
	ProjMass   float64
	ChargeMass float64
	BoreArea   float64
	// Caliber    float64
}

func (i *InternalBallisticsSimulator) RunSym() float64 {
	i.Reset()

	// fmt.Println("Running golang interior ballistics simulation, got following inputs:")
	// fmt.Println(fmt.Sprintf("%#v", i.Barrel))
	// fmt.Println(fmt.Sprintf("%#v", i.Projectile))
	// fmt.Println(fmt.Sprintf("%#v", i.Charge))
	// fmt.Println(fmt.Sprintf("%#v", i.Volume()))
	// fmt.Println(fmt.Sprintf("%#v", i.Barrel.Area()))
	// fmt.Println(fmt.Sprintf("%#v", i.Barrel.FrictionFactor()))

	// n := 0
	t := 0.
	for t < wallTime && i.Projectile.Path < i.Barrel.Length {
		s := State{}
		s.Time = t
		i.State(&s)
		i.Step(&s)
		fmt.Printf("%+v\n", s)
		// fmt.Println(fmt.Sprintf("%#v", s))
		// fmt.Println(i.Projectile.Velocity, i.Charge.HeatCapacity(), i.Charge.GasMass(), i.Volume(), 1/2)
		// fmt.Println(fmt.Sprintf("t: %.1f ms\tPmean: %.0f MPa\tPbase: %.0f MPa\tTgas: %.0f K\tCp: %.0f J/K\trho: %.1f kg/m3\tV: %.0f m/s\tx: %.3f m", t*1e3, Pmean/1e6, Pbase/1e6, Tgas, i.Charge.HeatCapacity(), i.Charge.GasMass()/i.Volume(), i.Projectile.Velocity, i.Projectile.Path))

		t += dt
	}
	fmt.Println(fmt.Sprintf("Eproj: %.0f J, Eprop: %.0f J, Q: %.0f J",
		i.Projectile.KineticEnergy(),
		i.Charge.KineticEnergy(i.Projectile.Velocity),
		i.Barrel.Q))
	return 42
}

func (i *InternalBallisticsSimulator) Step(s *State) {
	i.Barrel.Heat(s.Tmean, s.Velocity, s.HeatCapacity, s.GasMass, s.Volume, s.Path)
	i.Charge.Burn(s.Pmean)
	i.Projectile.Accelerate(s.Pbase * i.Params.BoreArea)
}

func (i *InternalBallisticsSimulator) State(s *State) {
	s.Volume = i.Volume()
	s.EnergyLoss = i.EnergyLoss()
	i.Charge.State(s)
	i.Projectile.State(s)
	i.Barrel.State(s)

}

func (i *InternalBallisticsSimulator) EnergyLoss() float64 {
	var out float64
	out += i.Projectile.KineticEnergy()
	out += i.Charge.KineticEnergy(i.Projectile.Velocity)
	out += i.Barrel.Q
	return out
}

func (i *InternalBallisticsSimulator) Volume() float64 {
	return i.Barrel.Volume - i.Projectile.AftVol - i.Charge.Volume() + i.Projectile.Path*i.Barrel.BoreArea
}

func (i *InternalBallisticsSimulator) Reset() {
	i.Charge.Reset()
	i.Projectile.Reset()
	i.Barrel.Reset()
}

func (i *InternalBallisticsSimulator) LinkComponents() {
	i.Params = &SimParams{i.Projectile.Mass, i.Charge.Mass(), i.Barrel.BoreArea}
	i.Barrel.Sp = i.Params
}

// func (i *InternalBallisticsSimulator) dudt(u float64) float64 {
// 	i.Propellant.Z = u
// 	return i.Propellant.Pressure(i.Volume()) / 1e6
// }
