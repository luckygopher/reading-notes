// @Description: Package tempconv performs Celsius and Fahrenheit temperature computations
// @Author: Arvin
// @date: 2021/2/19 5:02 下午
package tempconv

import "fmt"

type Kelvin float64     // 绝对温度
type Celsius float64    // 摄氏温标
type Fahrenheit float64 // 华氏温标

// T(K) = t(°C) + 273.15
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
