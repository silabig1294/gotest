package main

import (
	"fmt"
	"sort"
)

func main() {


	a := []int{15, 4, 33, 52, 551, 90, 8, 16, 15, 105} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("\n", a) // order greater to least

	a = []int{-15, -4, -33, -52, -551, -90, -8, -16, -15, -105} // unsorted
	sort.Sort(sort.IntSlice(a))
	fmt.Println("\n", a) // order least to greater

	b := []string{"Bangkok", "Nakhonphatom", "Nakhonnayok", "Yala", "Sukothai", "Nontaburi", "Ayauthaya"} // unsorted
	sort.Sort(sort.StringSlice(b))
	fmt.Println("\n", b)

	b = []string{"Bangkok", "Nakhonphatom", "Nakhonnayok", "Yala", "Sukothai", "nontaburi", "ayauthaya"} // unsorted
	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	fmt.Println("\n", b)

	c := []float64{95.10, 85.10, 167.15, 49.15, 18.95} // unsorted
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println("\n", c)

	c = []float64{-95.10, -85.10, -167.15, -49.15, -18.95} // unsorted
	sort.Sort(sort.Float64Slice(c))
	fmt.Println("\n", c)

	d := []string{"sila","Boolerd","panitan","Siriporn","Lisa","lisa"} // unsorted
	sort.Sort(sort.StringSlice(d))
	fmt.Println("\n", d)

	d = []string{"Sila","Boolerd","Panitan","Siriporn","Lisa","Wanida","Jenny"}
	sort.Sort(sort.StringSlice(d))
	fmt.Println("\n", d)

	d = []string{"Sila","Boolerd","Panitan","Siriporn","Lisa","Wanida","Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(d)))
	fmt.Println("\n", d)
}
