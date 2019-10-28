package ibs

// import "fmt"

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
	ProjMass        float64
	ChargeMass      float64
	BoreArea        float64
	ForcingPressure float64
}

func (i *InternalBallisticsSimulator) RunSym() []State {
	// func (i *InternalBallisticsSimulator) RunSym() float64 {
	var out = make([]State, int(wallTime/dt))
	i.Reset()
	// fmt.Println(i.Charge)

	n := 0
	// t := 0.
	for t := 0.; t < wallTime && i.Projectile.Path < i.Barrel.Length; n++ {
		s := State{}
		s.Time = t
		i.State(&s)
		i.Step(&s)
		// fmt.Printf("%+v\n", s)

		t += dt
		out[n] = s
	}
	// fmt.Println(fmt.Sprintf("Eproj: %.0f J, Eprop: %.0f J, Q: %.0f J",
	// 	i.Projectile.KineticEnergy(),
	// 	i.Charge.KineticEnergy(i.Projectile.Velocity),
	// 	i.Barrel.Q))
	return out[:n]
	// return 42
}

func (i *InternalBallisticsSimulator) Step(s *State) {
	i.Barrel.Heat(s.Tmean, i.Charge.HeatFlux(s.Volume, s.Velocity), s.Path)
	i.Charge.Burn(s.Pmean)
	if s.Pbase > i.Params.ForcingPressure || i.Projectile.Path > 0 {
		f := s.Pbase * i.Params.BoreArea
		i.Projectile.Accelerate(f)
	}
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
	i.Params = &SimParams{i.Projectile.Mass, i.Charge.Mass(), i.Barrel.BoreArea, i.Params.ForcingPressure}
	i.Barrel.Sp = i.Params
	// i.Projectile.sp = i.Params
}

// func (i *InternalBallisticsSimulator) String() out string {
// }

// func (i *InternalBallisticsSimulator) dudt(u float64) float64 {
// 	i.Propellant.Z = u
// 	return i.Propellant.Pressure(i.Volume()) / 1e6
// }
