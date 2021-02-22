// @Description: Boiling prints the boiling point of water
// @Author: Arvin
// @date: 2021/2/8 9:34 上午
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
	// Output:
	// boiling point = 212ºF or 100ºC
}
