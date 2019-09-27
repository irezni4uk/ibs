package main

import (
	"fmt"
	"ibs"
	"time"
)

func main() {
	fmt.Println("Hello World")
	t := ibs.Projectile{}
	fmt.Println("%v", t)
	t = ibs.Projectile{1, 0, 0, 0, 0}
	s := fmt.Sprintf("%#v", t)
	fmt.Println(s)
	b := ibs.NewBarrel()
	s = fmt.Sprintf("%#v", b)
	fmt.Println(s)
	fmt.Println("Bore Caliber: ", b.Caliber())
	fmt.Println("Bore Area: ", b.BoreArea())
	p := ibs.NewPropellant()
	s = fmt.Sprintf("%#v", p)
	fmt.Println(s)
	i := ibs.InternalBallisticsSimulator{}
	fmt.Println(i)
	c := ibs.NewCharge()
	i = ibs.InternalBallisticsSimulator{&b, &t, &c, nil}
	i.LinkComponents()
	// test(i)
	fmt.Println(i.RunSym())
	// fmt.Println(fmt.Sprintf("%#v", c))
	fmt.Println(fmt.Sprintf("%#v", i))
}

func test(obj ibs.InternalBallisticsSimulator) {
	n := 1000
	start := time.Now()
	for i := 0; i < n; i++ {
		obj.RunSym()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("RunSym took %s", elapsed))
	// log.Printf("Binomial took %s", elapsed)
}
