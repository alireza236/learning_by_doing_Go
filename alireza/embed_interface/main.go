package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitunganKu interface {
	hitung2d
	hitung3d
	setName() string
}

type kubus struct {
	sisi float64
	nama string
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 2
}

func (k kubus) keliling() float64 {
	return k.sisi * 12
}

func (k kubus) setName() string {
	return k.nama
}

func main() {
	var bangunRuang hitunganKu = kubus{40.0, "reza"}

	fmt.Println("==== KUBUS")

	fmt.Println("Luas  :", bangunRuang.luas())
	fmt.Println("Keliling  :", bangunRuang.keliling())
	fmt.Println("Volume  :", bangunRuang.volume())
	fmt.Println("Name  :", bangunRuang.setName())

}
