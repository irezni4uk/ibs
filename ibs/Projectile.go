package ibs

type Projectile struct {
	Mass            float64
	AftVol          float64
	AftLen          float64
	ForcingPressure float64
	velocity        float64
	path            float64
}

func (p *Projectile) state(s *State) {
	s.Path = p.path
	s.Velocity = p.velocity
}

func (p *Projectile) kineticEnergy() float64 {
	return p.Mass * p.velocity * p.velocity / 2
}

func (p *Projectile) accelerate(Force float64) {
	accel := Force / p.Mass
	p.path += p.velocity * dt
	p.velocity += accel * dt
}

func (p *Projectile) reset() {
	p.velocity = 0
	p.path = 0
}
