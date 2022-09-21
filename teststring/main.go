package main

import (
	"fmt"
	"strings"
)

func main() {
	// learning strings
	// compare string -1 less than, 1 more than, 0 equal
	fmt.Println(strings.Compare("A", "B"))             // A < B
	fmt.Println(strings.Compare("B", "A"))             // B > A
	fmt.Println(strings.Compare("a", "A"))             // a > A
	fmt.Println(strings.Compare("Japan", "Australia")) // J > A
	fmt.Println(strings.Compare("Australia", "Japan")) // A < J
	fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", " "))              // Space is less

	// check string that have contains value in text if found value in string, will send true but if not found value in string, will send false
	fmt.Println(strings.Contains("Australia", "Aus")) // Any part of string
	fmt.Println(strings.Contains("Australia", "Australian"))
	fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
	fmt.Println(strings.Contains("Japan", "JAP"))   // Case sensitive
	fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("  ", " ")) // space also consider as string
	fmt.Println(strings.Contains("12554", "1"))
	a := "Standard"
	b := "dard"
	if strings.Contains(a, b) {
		fmt.Println("b contains in a") // start at a then b check with a that have contains ?
	} else {
		fmt.Println("b contains not in a")
	}

	// Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Australia", "a"))
	fmt.Println(strings.ContainsAny("Australia", "r & a"))
	fmt.Println(strings.ContainsAny("JAPAN", "j"))
	fmt.Println(strings.ContainsAny("JAPAN", "J"))
	fmt.Println(strings.ContainsAny("JAPAN", "JAPAN"))
	fmt.Println(strings.ContainsAny("JAPAN", "japan"))
	fmt.Println(strings.ContainsAny("Shell-12541", "1"))

	//  Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Shell-12541", "88")) // true
	fmt.Println(strings.Contains("Shell-12541", "1 & 2")) // false

	// Command EqualFold don't care A a but a44 === A44
	fmt.Println(strings.EqualFold("Australia", "AUSTRALIA"))
	fmt.Println(strings.EqualFold("Australia", "aUSTRALIA"))
	fmt.Println(strings.EqualFold("Australia", "Australia"))
	fmt.Println(strings.EqualFold("Australia", "Aus"))
	fmt.Println(strings.EqualFold("Australia", "Australia & Japan"))
	fmt.Println(strings.EqualFold("JAPAN-1254", "japan-1254"))
	fmt.Println(strings.EqualFold(" ", " "))  // single space both side
	fmt.Println(strings.EqualFold(" ", "  ")) // double space right side

	// use strings.Spilt("","+*- ") string , string with divided (*/-+0 )
	k := "123+456+789+788+777"

	strSlice := strings.Split("a,b,c", ",")
	fmt.Println(strSlice)

	strSlice = strings.Split("Australia is a country and continent surrounded by the Indian and Pacific oceans.", " ")
	for i , v := range strSlice {
		fmt.Printf("strSlice[%d] = %s \n",i,v)
	}

	strSlice = strings.Split(k,"+")
	fmt.Println("\n",strSlice)
	fmt.Println("++++++++++++++++++")
	strSlice = strings.Split("abacadaeaf", "a")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("abacadaeaf", "A")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("123023403450456056706780789", "0")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("123023403450456056706780789", ",")
	fmt.Println("\n", strSlice)

	// strings.ToUpper
	fmt.Println(strings.ToUpper("Germany is a Western European country with a landscape of forests, rivers, mountain ranges and North Sea beaches."))
	fmt.Println(strings.ToUpper("BAVARIA HESSE BRANDENBURG SAARLAND"))
	fmt.Println(strings.ToUpper("towns and cities"))

	// strings.ToLower
	fmt.Println(strings.ToLower("Germany is a Western European country with a landscape of forests, rivers, mountain ranges and North Sea beaches."))
	fmt.Println(strings.ToLower("BAVARIA HESSE BRANDENBURG SAARLAND"))
	fmt.Println(strings.ToLower("towns and cities"))
}
