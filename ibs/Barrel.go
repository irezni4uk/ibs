package ibs

/*
   %BARREL    Object represents gun barrel geometry.
   %   properties for BARREL are:
   %   DG          :Groove diameter, m
   %   DL          :Land diameter, m
   %   GLR         :Groove to land ratio
   %   Thickness   :wall thickness, m
   %   Temperature :wall temp, K

*/

import "math"

type Barrel struct {
	Caliber        float64
	BoreArea       float64
	FrictionFactor float64
	Thickness      float64
	Temperature    float64
	Length         float64
	Volume         float64
	Density        float64 //= 7860    % kg/m3
	Cp             float64 //= 460     % Heat capacity J/kg-K
	Q              float64
}

func (b *Barrel) state(s *State, ProjMass, ChargeMass float64) {
	s.Pbase = b.PressureGradient(s.Pmean, ProjMass, ChargeMass)
}

func (b *Barrel) PressureGradient(Pmean, ProjMass, ChargeMass float64) (Pbase float64) {
	return Pmean / (1 + ChargeMass/3/ProjMass)
}

func (b *Barrel) Area(path float64) float64 {
	return 2*b.BoreArea + math.Pi*b.Caliber*(b.Volume/b.BoreArea+path)
}

func (b *Barrel) Temperature_(path float64) float64 {
	return b.Temperature + (b.Q+0*0)/(b.Cp*b.Density*b.Area(path)*b.Thickness)
}

//Heat integrates heat transfered to the barrel over time
func (b *Barrel) heat(Tgas, HeatFlux, path float64) {
	h := b.FrictionFactor*HeatFlux + h0 //heat transfer coefficient
	b.Q += b.Area(path) * h * (Tgas - b.Temperature_(path)) * dt
}

func (b *Barrel) reset() {
	b.Q = 0
}

func NewBarrel(DG, DL, GLR float64) Barrel {
	out := Barrel{}

	out.Caliber = caliber(DG, DL, GLR)
	out.BoreArea = boreArea(out.Caliber)
	out.FrictionFactor = frictionFactor(out.Caliber)
	out.Thickness = .0045 * 25.4 * 1e-3 // IBHVG2 A USERS GUIDE
	out.Temperature = 273
	out.Length = 4.572
	out.Volume = 9832.24e-6
	out.Density = 7861.2
	out.Cp = 460.28

	return out
}

// func NewBarrel(DG, DL, GLR float64) Barrel {
// 	out := Barrel{}

// 	out.Caliber = caliber(DG, DL, GLR)
// 	out.BoreArea = boreArea(out.Caliber)
// 	out.FrictionFactor = frictionFactor(out.Caliber)
// 	out.Thickness = .0045 * 25.4 * 1e-3 // IBHVG2 A USERS GUIDE
// 	out.Temperature = 288
// 	out.Length = 3.4025
// 	out.Volume = 1600e-6
// 	out.Density = 7860
// 	out.Cp = 460

// 	return out
// }

//Caliber returns 'effective' caliber depending on groove and land diameters and groove to land ratio
func caliber(DG, DL, GLR float64) float64 {
	return math.Sqrt((GLR*math.Pow(DG, 2) + math.Pow(DL, 2)) / (GLR + 1))
}

func boreArea(Caliber float64) float64 {
	return math.Pi * math.Pow(Caliber, 2) / 4
}

//FrictionFactor Nordheim heat transfer friction factor
func frictionFactor(Caliber float64) float64 {
	return math.Pow(13.2+4*math.Log10(100*Caliber), -2)
}
