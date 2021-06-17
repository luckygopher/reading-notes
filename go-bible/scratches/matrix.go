package main

import "fmt"

// 螺旋矩阵
func main() {
	// fmt.Println(generateMatrix1(10))
	// fmt.Println(generateMatrix2(10))
	// fmt.Println(cover(121))
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println(num)
				ch2 <- 100
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			fmt.Println(num)
			ch1 <- 300
		}
	}
}

// 方法一
func generateMatrix1(n int) [][]int {
	// 创建矩阵
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	num := 1
	// 判断奇偶性
	if n%2 != 0 {
		res[n/2][n/2] = n * n
	}
	// 圈数
	for i := 0; i < n/2; i++ {
		// 右走
		for r := i; r < n-i-1; r++ {
			res[i][r] = num
			num++
		}
		// 下走
		for b := i; b < n-i-1; b++ {
			res[b][n-i-1] = num
			num++
		}
		// 左走
		for l := n - i - 1; l > i; l-- {
			res[n-i-1][l] = num
			num++
		}
		// 上走
		for t := n - i - 1; t > i; t-- {
			res[t][i] = num
			num++
		}
	}

	return res
}

func generateMatrix2(n int) [][]int {
	// 右下左上跑，跑完就缩圈，直到num越界即停止
	// 创建矩阵,初始值均为0，由于我们是正整数，因此只要访问到不为0则表示已经访问过，应该转向
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 定以边界
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for num <= n*n {
		for r := left; r <= right; r++ {
			res[top][r] = num
			num++
		}
		top++

		for b := top; b <= bottom; b++ {
			res[b][right] = num
			num++
		}
		right--

		for l := right; l >= left; l-- {
			res[bottom][l] = num
			num++
		}
		bottom--

		for t := bottom; t >= top; t-- {
			res[t][left] = num
			num++
		}
		left++
	}
	return res
}

// 回文数
func cover(a int64) int64 {
	var c int64
	var i int64 = 1
	for a > 0 {
		b := a % 10
		c = c*10 + b*i
		a = a / 10
	}
	return c
}
