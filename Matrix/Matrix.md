# 矩阵
[TOC]

## 矩阵置零
> Problem: [73. 矩阵置零](https://leetcode.cn/problems/set-matrix-zeroes/description/)



## 思路

> 对空间复杂度有要求，也就是不能开多余空间，原地算法

## 解题方法

> 使用第一行和第一列代替新开的两个数组，单独开两个bool来记录第一行和第一列有无0，因为数组的值在正负maxint，所以我用了maxint+1来记录，具体实现就是个模拟 

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
func setZeroes(matrix [][]int) {
	rowf := false
	colf := false
	if matrix[0][0] == 0 {
		rowf = true
		colf = true
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			matrix[i][0] = math.MaxInt32 + 1
			colf = true
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			matrix[0][j] = math.MaxInt32 + 1
			rowf = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = math.MaxInt32 + 1
				matrix[0][j] = math.MaxInt32 + 1
			}
		}
	}
	//fmt.Println(matrix)

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == math.MaxInt32+1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == math.MaxInt32+1 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	//fmt.Println(colf, rowf)
	if colf {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if rowf {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}
```


## 螺旋矩阵
> Problem: [54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/description/)


## 思路
> 硬核模拟

## 解题方法

> [官方题解](https://leetcode.cn/problems/spiral-matrix/solutions/275393/luo-xuan-ju-zhen-by-leetcode-solution)

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$

## Code
```Go []
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
```


## 旋转图像
> Problem: [48. 旋转图像](https://leetcode.cn/problems/rotate-image/description/)


## 思路

> 硬核模拟，就是将数组旋转九十度输出

## 解题方法

> 一个数移动之后对应的位置即可
[官方题解](https://leetcode.cn/problems/rotate-image/solutions/526980/xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m)

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
```

## 搜索二维矩阵 II
> Problem: [240. 搜索二维矩阵 II](https://leetcode.cn/problems/search-a-2d-matrix-ii/description/)

## 思路

> 遍历

## 解题方法

> 遍历，当然也可以二分，应为每行都是排好序的

## 复杂度

时间复杂度:
> $O(n^2)$

空间复杂度:
> $O(1)$



## Code
```Go []
func searchMatrix(matrix [][]int, target int) bool {
	for _, i := range matrix {
		for _, j := range i {
			if j == target {
				return true
			}
		}
	}
	return false
}
```



