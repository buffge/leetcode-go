package main

import (
	"container/heap"
	"log"
	"sort"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	log.Println(maxSlidingWindow(nums, k))
	log.Println(maxSlidingWindowV2(nums, k))
	nums = []int{1}
	k = 1
	log.Println(maxSlidingWindow(nums, k))
	log.Println(maxSlidingWindowV2(nums, k))
	nums = []int{1, 3, 1, 2, 0, 5}
	//                  | 3 1
	//                     | 3 2
	//                        | 3 2
	k = 3
	log.Println(maxSlidingWindow(nums, k))
	log.Println(maxSlidingWindowV2(nums, k))
}

type PQ struct {
	sort.IntSlice
}

var ceNums []int

// 大堆
func (h *PQ) Less(i, j int) bool {
	return ceNums[h.IntSlice[i]] > ceNums[h.IntSlice[j]]
}
func (h *PQ) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *PQ) Pop() interface{} {
	arr := h.IntSlice
	v := arr[len(arr)-1]
	h.IntSlice = arr[:len(arr)-1]
	return v
}

/*
* 239. 滑动窗口最大值
思路: 优先队列法 假设k=3 初始化大堆3个元素 此时第一个结果就是堆顶
从k开始遍历数组 依次加入堆中
比如加入第四个元素后 最大值不在窗口中 此最大值无用可以直接弹出
即第1个元素已经无用了
如果最大值在窗口中就加入结果集
每轮加入元素后都要将所有的窗口前大值去除

核心点 加入堆后取最大值并不会弹出任何元素 只有搜索窗口内最大值时会弹出旧元素
一直弹到最大值在窗口内
*/
func maxSlidingWindow(nums []int, k int) []int {
	pqArr := make([]int, 0, len(nums)-k+1) // 堆底层数组
	for i := 0; i < k; i++ {               // 初始化k个元素 计算出最大值
		pqArr = append(pqArr, i) // 堆中存着数组索引
	}
	pq := &PQ{ // 大堆
		IntSlice: pqArr,
	}
	ceNums = nums                        // 添加全局变量 这样堆中就不用存结构体 只需要存个索引即可 正式环境不能这样 需要用结构体
	heap.Init(pq)                        // 初始化大堆
	res := make([]int, 1, len(nums)-k+1) // 结果集
	res[0] = nums[pq.IntSlice[0]]        // 第一个滑动窗口最大值
	for i := k; i < len(nums); i++ {     // 从k开始遍历
		heap.Push(pq, i)             // 将当前元素和索引加入大堆(虽然这里加入的是索引,但是堆中能从全局变量找到索引对应的元素)
		for pq.IntSlice[0] < i-k+1 { // 如果堆中最大值不在窗口内  一直弹出到最大值在窗口内 即索引小于窗口左边界索引
			heap.Pop(pq) // 弹出非窗口内的最大值
		}
		res = append(res, nums[pq.IntSlice[0]]) // 添加正确结果 这里的结果不能pop 因为此数后面还要参与计算
		// 如果 pq.IntSlice[0] 为窗口最左边 可以pop掉 但是上面还是要pop 省几行代码 所以这里不pop
		//if pq.IntSlice[0] == i-k {
		//	heap.Pop(pq)
		//}
	}
	return res
}

/*
单调队列 存放元素索引  元素实际值单调递减
例如k=3 队列中最多存3个元素
队列[0]为最大值
窗口右移时 判断栈[0]是否为旧左 是则删除
将新值加入队列

核心点 因为如果窗口最右值大于窗口左值 那么左值永远都不会成为窗口中的最大值
如 [1,5,2,3] 当窗口最右值(3)大于窗口左值(1,2) 1和2是永远不会再成为任何一个窗口的最大值
因为遍历是向右移动 最右值可以保证一直跟左值在一个窗口中 所以可以用单调递减队列完成
*/
func maxSlidingWindowV2(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1) // 结果集
	deque := make([]int, 1, k)           // 单调栈
	deque[0] = 0                         // 此时窗口最大值索引为0
	for i := 1; i < k; i++ {             // 避免了下面遍历中每次都要判断是否形成了窗口 因为n>=k  可以节省n-k次判断
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] { // 加入新值时删除栈中所有小于当前值的元素
			deque = deque[:len(deque)-1] // 因为如果窗口最右值大于左值 那么左值永远都不会成为窗口中的最大值
		} // 删除所有比窗口最右值小的元素
		deque = append(deque, i) // 将当前元素索引加入队列
	}
	res = append(res, nums[deque[0]]) // 第一个结果
	for i := k; i < len(nums); i++ {  // 从第k个开始遍历 如果从0遍历的话 下面res append就得判断i>=k-1
		if len(deque) > 0 && deque[0] == i-k { // 如果队列最大值不在窗口中 将此值删除
			deque = deque[1:] // 窗口最右值索引为i 那么上一轮窗口最左值索引就是i-k 如果当前队列最大值时i-k的元素就将它删除
		}
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] { // 加入新值时删除栈中所有小于当前值的元素
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)          // 将当前元素索引加入队列
		res = append(res, nums[deque[0]]) // 将当前窗口最大值加入结果集
	}
	return res
}

/*
*
两个数组 k个为1组
一个存前k个最大的值的索引
一个存后k个最大的值的索引
*/
func maxSlidingWindowV3(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1) // 结果集
	deque := make([]int, 1, k)           // 单调栈
	deque[0] = 0                         // 此时窗口最大值索引为0
	for i := 1; i < k; i++ {             // 避免了下面遍历中每次都要判断是否形成了窗口 因为n>=k  可以节省n-k次判断
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] { // 加入新值时删除栈中所有小于当前值的元素
			deque = deque[:len(deque)-1] // 因为如果窗口最右值大于左值 那么左值永远都不会成为窗口中的最大值
		} // 删除所有比窗口最右值小的元素
		deque = append(deque, i) // 将当前元素索引加入队列
	}
	res = append(res, nums[deque[0]]) // 第一个结果
	for i := k; i < len(nums); i++ {  // 从第k个开始遍历 如果从0遍历的话 下面res append就得判断i>=k-1
		if len(deque) > 0 && deque[0] == i-k { // 如果队列最大值不在窗口中 将此值删除
			deque = deque[1:] // 窗口最右值索引为i 那么上一轮窗口最左值索引就是i-k 如果当前队列最大值时i-k的元素就将它删除
		}
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] { // 加入新值时删除栈中所有小于当前值的元素
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)          // 将当前元素索引加入队列
		res = append(res, nums[deque[0]]) // 将当前窗口最大值加入结果集
	}
	return res
}
