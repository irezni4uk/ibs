//Package ibs provides functionality for interior ballistics simulations
package ibs

// import "fmt"

const (
	wallTime = 20e-3
	dt       = 5e-6

	h0 = 11.35 //Free convective heat transfer coefficient, zero speed W/m^2-K
)

//InternalBallisticsSimulator manages all components of simulation
type InternalBallisticsSimulator struct {
	Barrel     *Barrel
	Projectile *Projectile
	Charge     *Charge
	Params     *SimParams
}

//SimParams contains common parameters of simulation
type SimParams struct {
	ProjMass        float64
	ChargeMass      float64
	BoreArea        float64
	ForcingPressure float64
}

//RunSym performs the simulation of interior ballistics problem
func (i *InternalBallisticsSimulator) RunSym() []State {
	var out = make([]State, int(wallTime/dt))
	i.Reset()
	// fmt.Println(i.Charge)

	n := 0
	// t := 0.
	for t := 0.; t < wallTime && i.Projectile.Path < i.Barrel.Length; n++ {
		s := State{}
		s.Time = t
		i.State(&s)
		// fmt.Printf("t: %.3f ms, zDot = %.4f\n", t*1e3, s.Pmean/i.Charge.Propellant[1].Impulse)
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
}

//Step makes one step through simulation time
func (i *InternalBallisticsSimulator) Step(s *State) {
	i.Barrel.Heat(s.Tmean, i.Charge.HeatFlux(s.Volume, s.Velocity), s.Path)
	// trig := true
	i.Charge.Burn(s.Pmean)
	if s.Pmean > i.Params.ForcingPressure || i.Projectile.Path > 0 {
		i.Projectile.Accelerate(s.Pbase * i.Params.BoreArea)
		// if trig {
		// 	fmt.Printf("FORCING t: %.0f mks\n", s.Time*1e6)
		// 	trig = false
		// }
	}
}

//State collects simulation parameters for present step
func (i *InternalBallisticsSimulator) State(s *State) {
	s.Volume = i.Volume()
	s.EnergyLoss = i.EnergyLoss()
	i.Charge.State(s)
	i.Projectile.State(s)
	i.Barrel.State(s)

}

//EnergyLoss calculates energy losses projectile translation, barrel heating, etc.
func (i *InternalBallisticsSimulator) EnergyLoss() float64 {
	var out float64
	out += i.Projectile.KineticEnergy()
	out += i.Charge.KineticEnergy(i.Projectile.Velocity)
	out += i.Barrel.Q
	return out
}

//Volume returns current system volume
func (i *InternalBallisticsSimulator) Volume() float64 {
	return i.Barrel.Volume - i.Projectile.AftVol - i.Charge.Volume() + i.Projectile.Path*i.Barrel.BoreArea
}

//Reset returns all components to inintial state
func (i *InternalBallisticsSimulator) Reset() {
	i.Charge.Reset()
	i.Projectile.Reset()
	i.Barrel.Reset()
}

//LinkComponents generates variable with common simulation parameters and provide it to all simulation components
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
