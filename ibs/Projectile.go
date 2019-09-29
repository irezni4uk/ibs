package ibs

// import "fmt"

// this is a comment

type Projectile struct {
	Mass     float64
	AftVol   float64
	AftLen   float64
	Velocity float64
	Path     float64
	sp       *SimParams
}

func (p *Projectile) State(s *State) {
	s.Path = p.Path
	s.Velocity = p.Velocity
}

func (p *Projectile) KineticEnergy() float64 {
	return p.Mass * p.Velocity * p.Velocity / 2
}

func (p *Projectile) Accelerate(Pbase float64) {
	if p.Path == 0 && Pbase < p.sp.ForcingPressure {
		return
	}
	accel := Pbase * p.sp.BoreArea / p.Mass
	p.Path += p.Velocity * dt
	p.Velocity += accel * dt
}

func (p *Projectile) Reset() {
	p.Velocity = 0
	p.Path = 0
}
