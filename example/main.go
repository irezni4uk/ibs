package main

import (
	"fmt"
	"ibs/ibs"
	// "os"
	"bytes"
	"encoding/binary"
	// "encoding/json"
	"io/ioutil"
	"time"
)

func main() {

	t := ibs.Projectile{Mass: 9.796} //, ForcingPressure: 1e12}

	b := ibs.NewBarrel(.127, .127, 1)
	// fmt.Println(fmt.Sprintf("%#v", b))

	// Dependence between burned web (nondimensional) and burned fraction of propellant grain
	psi := ibs.PsiFun(1.441, .651, .364, .6, -1.135, -.031)

	prop := ibs.Propellant{Mass: 8.7, Force: 1135990, Impulse: 1037219, Density: 1660.5,
		AdiabaticIndex: 1.23, Covolume: .9755e-3, BurnTemperature: 3142, Psi: psi}

	c := ibs.NewCharge()

	c.Propellant[1] = prop

	i := ibs.InternalBallisticsSimulator{}

	// put all objects together
	i = ibs.InternalBallisticsSimulator{
		Barrel:     &b,
		Projectile: &t,
		Charge:     &c,
	}

	// Run simulation and save results
	sol := i.RunSim()
	dumpSol(&sol)

	fmt.Println(sol[0].Pmean)
	fmt.Println(sol[0])
	fmt.Println(sol[len(sol)-700])
	fmt.Println(sol[len(sol)-1])

	test(i)
}

func test(obj ibs.InternalBallisticsSimulator) {
	n := 1000
	start := time.Now()
	for i := 0; i < n; i++ {
		obj.RunSim()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("%d calls of RunSim took %s", n, elapsed))
}

func dumpSol(in *[]ibs.State) {
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.LittleEndian, in)
	err := ioutil.WriteFile("sol.bin", bin_buf.Bytes(), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
