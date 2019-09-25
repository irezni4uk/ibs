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
}

func (i *InternalBallisticsSimulator) RunSym() float64 {
	i.Reset()

	// fmt.Println("Running golang interior ballistics simulation, got following inputs:")
	// fmt.Println(fmt.Sprintf("%#v", i.Barrel))
	// fmt.Println(fmt.Sprintf("%#v", i.Projectile))
	// fmt.Println(fmt.Sprintf("%#v", i.Charge))
	// fmt.Println(fmt.Sprintf("%#v", i.Volume()))

	t := 0.
	for t < wallTime && i.Projectile.Path < i.Barrel.Length {
		Tgas, Pmean := i.Charge.Thermodynamics(i.Volume(), i.Energy())
		HeatFlux := 1 / 2 * i.Projectile.Velocity * i.Charge.HeatCapacity()
		i.Barrel.Heat(Tgas, HeatFlux)
		Force := Pmean * i.Barrel.BoreArea()
		i.Projectile.Accelerate(Force)
		// fmt.Println(fmt.Sprintf("t: %.1f ms\tP: %.1f MPa\tV: %.0f m/s\tx: %.3f m", t*1e3, Pmean/1e6, i.Projectile.Velocity, i.Projectile.Path))
		i.Charge.Burn(Pmean)
		t += dt
	}
	// fmt.Println(fmt.Sprintf("Eproj: %.0f J, Eprop: %.0f J, Q: %.0f J", i.Projectile.KineticEnergy(), i.Charge.KineticEnergy(), i.Barrel.Q))
	return 42
}

func (i *InternalBallisticsSimulator) Volume() float64 {
	return i.Barrel.Volume - i.Projectile.AftVol - i.Charge.Volume() + i.Projectile.Path*i.Barrel.BoreArea()
}

func (i *InternalBallisticsSimulator) Reset() {
	i.Charge.Reset()
	i.Projectile.Reset()
	i.Barrel.Reset()
}

func (i *InternalBallisticsSimulator) Energy() float64 {
	var out float64
	out += i.Projectile.KineticEnergy()
	out += i.Charge.KineticEnergy()
	out += i.Barrel.Q
	return out
}

func (i *InternalBallisticsSimulator) LinkComponents() {
	i.Charge.Projectile = i.Projectile
}

// func (i *InternalBallisticsSimulator) dudt(u float64) float64 {
// 	i.Propellant.Z = u
// 	return i.Propellant.Pressure(i.Volume()) / 1e6
// }
