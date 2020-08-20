//Package ibs provides functionality for interior ballistics simulations
package ibs

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
}

//RunSym performs the simulation of interior ballistics problem
func (i *InternalBallisticsSimulator) RunSim() []State {

	var out = make([]State, int(wallTime/dt))

	i.reset()

	n := 0
	for t := 0.; t < wallTime && i.Projectile.path < i.Barrel.Length; n++ {
		s := State{}
		s.Time = t
		i.state(&s)
		i.step(&s)
		t += dt
		out[n] = s
	}
	return out[:n]
}

//Step makes one step over simulation time
func (i *InternalBallisticsSimulator) step(s *State) {
	i.Barrel.heat(s.Tmean, i.Charge.heatFlux(s.Volume, s.Velocity), s.Path)
	i.Charge.burn(s.Pmean)
	if s.Pmean > i.Projectile.ForcingPressure || i.Projectile.path > 0 {
		i.Projectile.accelerate(s.Pbase * i.Barrel.BoreArea)
	}
}

//State collects simulation parameters for present step
func (i *InternalBallisticsSimulator) state(s *State) {
	s.Volume = i.volume()
	s.EnergyLoss = i.energyLoss()
	i.Charge.state(s)
	i.Projectile.state(s)
	i.Barrel.state(s, i.Projectile.Mass, i.Charge.Mass())

}

//EnergyLoss calculates energy losses on: projectile translation, barrel heating, etc.
func (i *InternalBallisticsSimulator) energyLoss() float64 {
	var out float64
	out += i.Projectile.kineticEnergy()
	out += i.Charge.KineticEnergy(i.Projectile.velocity)
	out += i.Barrel.Q
	return out
}

//Volume returns current system volume
func (i *InternalBallisticsSimulator) volume() float64 {
	return i.Barrel.Volume - i.Projectile.AftVol - i.Charge.Volume() + i.Projectile.path*i.Barrel.BoreArea
}

//Reset returns all components to inintial state
func (i *InternalBallisticsSimulator) reset() {
	i.Charge.reset()
	i.Projectile.reset()
	i.Barrel.reset()
}
