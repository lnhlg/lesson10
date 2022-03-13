package lesson10

import (
	"container/heap"
	"sort"
)

type maxSlidingWindow struct {
	k      int
	nums   []int
	window sort.IntSlice
}

func NewMaxSlidingWindow(nums []int, k int) *maxSlidingWindow {

	q := new(maxSlidingWindow)
	q.k = k
	q.nums = nums
	q.window = make([]int, k)

	return q
}

func (q *maxSlidingWindow) Push(x interface{}) {

	q.window = append(q.window, x.(int))
}

func (q *maxSlidingWindow) Pop() interface{} {

	tmp := q.window
	value := tmp[q.LastIndex()]
	q.window = tmp[:q.LastIndex()]

	return value
}

func (q *maxSlidingWindow) Less(i, j int) bool {

	return q.nums[q.window[i]] >= q.nums[q.window[j]]
}

func (q *maxSlidingWindow) Len() int {

	return q.window.Len()
}

func (q *maxSlidingWindow) Swap(i, j int) {

	q.window.Swap(i, j)
}

func (q *maxSlidingWindow) LastIndex() int {

	return len(q.window) - 1
}

func (q *maxSlidingWindow) Front() int {

	return q.window[0]
}

func (q *maxSlidingWindow) Init() {

	for i := 0; i < q.k; i++ {

		q.window[i] = i
	}
	heap.Init(q)
}

func (q *maxSlidingWindow) Sliding() []int {

	count := len(q.nums)
	result := make([]int, 1, count-q.k+1)
	result[0] = q.nums[q.Front()]
	for i := q.k; i < count; i++ {

		heap.Push(q, i)
		for q.window[0] <= i-q.k {

			heap.Pop(q)
		}
		result = append(result, q.nums[q.Front()])
	}

	return result
}

//MaxSlidingWindow: 滑动窗口最大值
//给定一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
/*parameter
nums: 整型数组
return: 整型数组，记录窗口每一次滑动时窗口内的最大值
*/
func MaxSlidingWindow(nums []int, k int) []int {

	q := NewMaxSlidingWindow(nums, k)
	q.Init()

	return q.Sliding()
}
