package ibs

/*
   %BARREL    Object represents gun barrel geometry.
   %   properties for BARREL are:
   %   DG          :Groove diameter, m
   %   DL          :Land diameter, m
   %   GLR         :Groove to land ratio
   %   Thickness   :wall thickness, m
   %   Temperature :wall temp, K

   %   Copyright 2019 Reznichuk I.
*/

import "math"

//import "ibs"

type Barrel struct {
	DG          float64
	DL          float64
	GLR         float64
	Thickness   float64
	Temperature float64
	Length      float64
	Volume      float64
	Density     float64 //= 7860    % kg/m3
	Cp          float64 //= 460     % Heat capacity J/kg-K
	Q           float64
	Projectile  *Projectile
}

func (b *Barrel) Area() float64 {
	return 2*b.BoreArea() + math.Pi*b.Caliber()*b.Volume/b.BoreArea()
}

func (b *Barrel) Temperature_() float64 {
	return b.Temperature + (b.Q+0*0)/(b.Cp*b.Density*b.Area()*b.Thickness)
}

func (b *Barrel) Heat(Tgas, HeatFlux float64) {
	h := b.FrictionFactor()*HeatFlux + h0 //heat transfer coefficient
	b.Q += b.Area() * h * (Tgas - b.Temperature_()) * dt
}

func (b *Barrel) Reset() {
	b.Q = 0
}

func (b *Barrel) FrictionFactor() float64 { //heat transfer friction factor
	return math.Pow(13.2+4*math.Log10(100*b.Caliber()), -2)
}

func (b *Barrel) Caliber() float64 {
	return math.Sqrt((b.GLR*math.Pow(b.DG, 2) + math.Pow(b.DL, 2)) / (b.GLR + 1))
}

func (b *Barrel) BoreArea() float64 {
	return math.Pow(b.Caliber(), 2) * math.Pi / 4
}

func NewBarrel() Barrel {
	out := Barrel{}

	out.DG = 58.8e-3
	out.DL = 57e-3
	out.GLR = 0.683                     // pi*d/(n*w) - 1    <- pi*57/(16*(6.8-.3/2)) - 1
	out.Thickness = .0045 * 25.4 * 1e-3 // IBHVG2 A USERS GUIDE
	out.Temperature = 288
	out.Length = 3.4025
	out.Volume = 1600e-6
	out.Density = 7860
	out.Cp = 460
	out.Q = 0

	return out
}
