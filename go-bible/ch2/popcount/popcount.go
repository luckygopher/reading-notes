// @Description:
// @Author: Arvin
// @date: 2021/2/22 5:19 下午
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 匿名函数方式初始化
//var pc [256]byte = func() (pc [256]byte) {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//	return
//}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}