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

	// Create projectile
	t := ibs.Projectile{Mass: 1, ForcingPressure: 100e6}

	// Create barrel
	b := ibs.NewBarrel()
	// fmt.Println(fmt.Sprintf("%#v", b))

	// Create Propellant
	prop := ibs.Propellant{Mass: 1, Force: 1006000, Impulse: 500000, Density: 1600,
		AdiabaticIndex: 1.224, Covolume: 1.01e-3, BurnTemperature: 2900, Psi: ibs.PsiFun(1.607, .769, .101, .506, -.823)}

	// Create charge
	c := ibs.NewCharge()

	c.Propellant[1] = prop

	// Create simulation object
	i := ibs.InternalBallisticsSimulator{}

	i = ibs.InternalBallisticsSimulator{
		Barrel:     &b,
		Projectile: &t,
		Charge:     &c,
	}

	// Run simulation and save results
	sol := i.RunSim()
	dumpSol(&sol)

	fmt.Println(sol[0])
	fmt.Println(sol[len(sol)-1])

	// test(i)

	// jsn, _ := json.Marshal(i)
	// fmt.Println(string(jsn))

	// f, err := os.Create("/home/naveen/bytes")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	// n2, err := f.Write(d2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// fmt.Println(n2, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func test(obj ibs.InternalBallisticsSimulator) {
	n := 1000
	start := time.Now()
	for i := 0; i < n; i++ {
		obj.RunSim()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("RunSim took %s", elapsed))
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
