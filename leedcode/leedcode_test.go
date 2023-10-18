package leedcode

import "testing"

/**二维数组查找
给定一个二维数组，查找是否含有对应的目标值
[1,2,3,4,5,6]
[2,3,4,5,6,7]
[3,4,5,6,7,8]
[4,5,6,7,8,9]
如：给定目标 7
解：左上、右下 为最大/最小
左下、右上可以每次剔除一行/一列
*/

func Test2DisArraySearch(t *testing.T) {
	target := 3
	target_arr := [][]int{
		{1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8},
		{4, 5, 6, 7, 8, 9},
	}

	wid := len(target_arr[0])
	hei := len(target_arr)

	for x, y := 0, wid-1; x < hei && y > 0; {
		if target_arr[x][y] == target {
			t.Logf("ok")
			t.Logf("x:%d,y:%d,value:%d", x, y, target_arr[x][y])
			return
		} else if target_arr[x][y] > target {
			t.Logf("x:%d,y:%d,value:%d", x, y, target_arr[x][y])
			y--
			continue
		} else if target_arr[x][y] < target {
			t.Logf("x:%d,y:%d,value:%d", x, y, target_arr[x][y])
			x++
			continue
		}

	}

}
