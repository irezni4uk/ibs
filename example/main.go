package main

import (
	"fmt"
	"ibs/ibs"
	// "os"
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"time"
)

func main() {

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

	fmt.Println("Hello World")
	t := ibs.Projectile{}
	fmt.Println("%v", t)
	t = ibs.Projectile{
		Mass:     1,
		AftVol:   0,
		AftLen:   0,
		Velocity: 0,
		Path:     0,
	}
	s := fmt.Sprintf("%#v", t)
	fmt.Println(s)
	b := ibs.NewBarrel()
	s = fmt.Sprintf("%#v", b)
	fmt.Println(s)
	fmt.Println("Bore Caliber: ", b.Caliber)
	fmt.Println("Bore Area: ", b.BoreArea)
	p := ibs.NewPropellant()
	s = fmt.Sprintf("%#v", p)
	fmt.Println(s)
	i := ibs.InternalBallisticsSimulator{}
	fmt.Println(i)
	c := ibs.NewCharge()
	i = ibs.InternalBallisticsSimulator{
		Barrel:     &b,
		Projectile: &t,
		Charge:     &c,
		Params:     &ibs.SimParams{ForcingPressure: 100e6},
	}
	i.LinkComponents()
	test(i)
	// fmt.Println(i.RunSym())
	// fmt.Println(fmt.Sprintf("%#v", c))
	fmt.Println(fmt.Sprintf("%#v", i))

	sol := i.RunSym()
	dumpSol(&sol)
	fmt.Println(sol[0])
	fmt.Println(sol[len(sol)-1])
}

func test(obj ibs.InternalBallisticsSimulator) {
	n := 100
	// n := 2
	start := time.Now()
	for i := 0; i < n; i++ {
		obj.RunSym()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("RunSym took %s", elapsed))
	// log.Printf("Binomial took %s", elapsed)
}

func dumpSol(in *[]ibs.State) {
	var bin_buf bytes.Buffer
	// x := myStruct{"1", "Hello"}
	// binary.Write(&bin_buf, binary.BigEndian, in)
	binary.Write(&bin_buf, binary.LittleEndian, in)
	err := ioutil.WriteFile("sol.bin", bin_buf.Bytes(), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
