// @Description: Ftoc prints two Fahrenheit-to-Celsius conversions
// @Author: Arvin
// @date: 2021/2/8 11:00 上午
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF = %gºC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gºF = %gºC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
