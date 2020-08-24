package ibs

// import "fmt"

type Grain struct {
	P  float64 // PERF DIA
	D  float64 // OUTER DIA
	L  float64 // GRAIN LENGTH
	NP float64 // NUMBER OF PERFS
}

func (g *Grain) Convert(p *Propellant, burnRate float64) {
	web := (g.D - 3*g.P) / 4

	rho1 := 0.1772 * (g.P + web)
	rho2 := 0.0774 * (g.P + web)
	zk := (web/2 + rho1 + rho2) / (web / 2)

	beta := web / g.L

	p.Impulse = 0.5 * web / burnRate * 1e6

	p1 := (g.D + g.NP*g.P) / g.L
	Q1 := (g.D*g.D - g.NP*g.P*g.P) / (g.L * g.L)

	k1 := (Q1 + 2*p1) / Q1 * beta
	l1 := (g.NP - 1 - 2*p1) / (Q1 + 2*p1) * beta
	m1 := -((g.NP - 1) * beta * beta) / (Q1 + 2*p1)

	psi := PsiFun(zk, k1, l1, 0, 0, m1)

	if g.NP != 7 {
		p.Psi = psi
		return
	}

	// fmt.Println(fmt.Sprintf("psi(1): %f\n", psi(1)))

	k2 := 2 * (1 - psi(1)) / (zk - 1)
	l2 := -1 / (2 * (zk - 1))

	p.Psi = PsiFun(zk, k1, l1, k2, l2, m1)
	// fmt.Println(fmt.Sprintf("%f, %f, %f, %f, %f, %f\n", zk, k1, l1, k2, l2, m1))
}
