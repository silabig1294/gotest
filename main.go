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

	// e := []float64{85.201, 14.74, 965.25, 125.32, 63.14} // unsorted
	// sort.Sort(sort.Float64Slice(e))
	// fmt.Println(e)	// sorted
	// fmt.Println("Length of Slice: ", sort.Float64Slice.Len(e))	// 5
	// fmt.Println("123.32 found in Slice at position: ", sort.Float64Slice(e).Search(1000.32))		//	3
	// fmt.Println("999.15 found in Slice at position: ", sort.Float64Slice(e).Search(965.25))		//	5
	// fmt.Println("12.14 found in Slice at position: ", sort.Float64Slice(e).Search(14.74))
	fmt.Println("\n######## SearchInts not works in descending order bot works in ascending order ######## ")    
	intSlice := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5}	// sorted slice in descending
	sort.Sort(sort.IntSlice(intSlice)) // add this sentences for  SerchInts will work.// if close comment , this sentenses will not work.
    x := 36
    pos := sort.SearchInts(intSlice,x)
    fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)

	intSlice2 := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5}
	fmt.Println("\n######## Search works in descending order  ########")	
	i := sort.Search(len(intSlice2), func(i int) bool { return intSlice2[i] <= x })
	fmt.Printf("Found %d at index %d in %v\n", x, i, intSlice2)
}
