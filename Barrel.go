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

	return out
}
