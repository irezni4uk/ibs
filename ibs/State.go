package ibs

import "fmt"

type State struct {
	Time         float64
	Pmean        float64
	Tmean        float64
	Velocity     float64
	Path         float64
	Pbase        float64
	HeatCapacity float64
	GasMass      float64
	Volume       float64
	EnergyLoss   float64
}

func (s State) String() string {
	return fmt.Sprintf("t: %.1f ms\tPmean: %.0f MPa\tPbase: %.0f MPa\tTgas: %.0f K\tCp: %.0f J/K\trho: %.1f kg/m3\tV: %.0f m/s\tx: %.3f m\n",
		s.Time*1e3,
		s.Pmean/1e6,
		s.Pbase/1e6,
		s.Tmean,
		s.HeatCapacity,
		s.GasMass/s.Volume,
		s.Velocity,
		s.Path)
}
