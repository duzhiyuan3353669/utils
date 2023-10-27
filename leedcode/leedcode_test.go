package leedcode

import (
	"container/list"
	"errors"
	"fmt"
	"testing"
)

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

/**用%20替换空格
 */
func TestReplaceSpace(t *testing.T) {
	//str := "你好 哈 地方非得 地方 都是"
	str := "JLJ DSFS SDFSDJ SDFSDJSDF FDSF"

	//ns := strings.ReplaceAll(str, " ", "%20")
	//t.Logf("%s", ns)
	str_arr := []byte(str)
	str_target := []byte{}
	for _, b := range str_arr {
		if b == ' ' {
			m := make([]byte, len(str_target)+3)
			c := []byte{'%', '2', '0'}
			copy(m, str_target)
			copy(m[len(str_target)-1:], c)
			str_target = m
			//copy(str_target[len(str_target)-1:], c)
		} else {
			str_target = append(str_target, b)
		}
	}

	t.Logf("%s", str_target)
}

/**
从尾部打印链表
*/
func TestPrintLinkFromTail(t *testing.T) {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	for b := l.Back(); b != nil; b = b.Prev() {
		fmt.Println(b.Value)
		fmt.Printf("%+v", b)
	}

}

/**
用两个栈实现队列
*/
type Queue struct {
	in  Stack
	out Stack
}

func (q *Queue) IsEmpty() bool {
	return q.in.IsEmpty() && q.out.IsEmpty()
}

func (q *Queue) Push(value interface{}) {
	q.in.Push(value)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	var value interface{}

	if !q.out.IsEmpty() {
		value, _ = q.out.Pop()
		return value, nil
	}

	for !q.in.IsEmpty() {
		value, _ = q.in.Pop()
		q.out.Push(value)
	}
	value, _ = q.out.Pop()
	return value, nil
}

/**
用两个栈模拟队列 先进先出【栈是先进后出】
*/
func TestListByStack(t *testing.T) {
	var myQueue Queue
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
	myQueue.Push(3)
	myQueue.Push(4)
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
}

/**
数组反转
*/
//func TestArrayRevese(t *testing.T) {
//
//	arr := [4]int{1, 2, 3, 4}
//
//	arr = *arrrevese(&arr)
//	fmt.Printf("%+v", arr)
//}
//
//func arrrevese(arr *[4]int) *[4]int {
//	l := list.New()
//
//	for _, i2 := range arr {
//		l.PushFront(i2)
//	}
//
//	i := 0
//	for v := l.Front(); v != nil; v = v.Next() {
//		arr[i] = v.Value.(int)
//
//		i++
//	}
//	return arr
//
//}
/**
获取斐波拉切数列的n项
0 1 1 2 3 5 8 13 21
*/
func TestFibonacci(t *testing.T) {
	fmt.Println(getFibonacci(1))
	fmt.Println(getFibonacci(2))
	fmt.Println(getFibonacci(3))
	fmt.Println(getFibonacci(4))

}

func getFibonacci(n int) (res int) {
	if n <= 1 {
		return n
	}

	f1, f2 := 0, 1

	for i := 2; i <= n; i++ {
		res = f1 + f2
		f1, f2 = f2, res
	}

	return res
}

func TestA(t *testing.T) {
	a := 5
	fmt.Printf("%b", a)
}
