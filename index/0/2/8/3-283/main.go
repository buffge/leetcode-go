package main

import "log"

/*
*
283. 移动零
思路: 快慢指针 快指针遍历数组 慢指针指向所有非0
curr为慢指针 遍历数组i为快指针
将所有非0值依次与慢指针元素交换
如 0 1 0 3 12
快指针数组  |0 1 0 3 12
慢指针数组  _ _ _ _ _
第一轮后   0 |1 0 3 12
慢指针     _ _ _ _ _
第二轮后   1 0 |0 3 12
慢指针     1 _ _ _ _
第三轮后   1 3 0 |0 12
慢指针     1 3 _ _ _
第四轮后   1 3 0 0 |12
慢指针     1 3 _ _ _
第五轮后   1 3 12 0 0|
慢指针     1 3 12 _ _
*/
func moveZeroes(nums []int) {
	curr := 0                  // 设置当前插入非0值的位置
	for i, num := range nums { // 遍历数组
		if num != 0 { // 如果当前元素非0
			nums[curr], nums[i] = nums[i], nums[curr] // 交换当前元素和当前插入位置
			curr++                                    // 插入位置+1
		}
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	log.Println(nums)

}
