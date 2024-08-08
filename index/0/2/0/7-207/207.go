package main

import (
	"fmt"
)

/*
*
拓扑序列
概念 有向无环图 再计算边数
暴力法 击败5% shit
*/
func canFinishV1(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0]) // 先完成v[1]可能才能完成v[0]课程
		if len(graph[v[0]]) == 0 {
			graph[v[0]] = nil
		}
	}
	total := numCourses
	for len(graph) > 0 && total > 0 {
		for i, arr := range graph {
			if len(arr) == 0 {
				delete(graph, i)
				for j := range graph {
					tempArr := make([]int, 0, len(graph[j]))
					for _, v := range graph[j] {
						if v != i {
							tempArr = append(tempArr, v)
						}
					}
					graph[j] = tempArr
				}
			}
		}
		total--
	}
	return len(graph) == 0
}

/*
用入度优化
*/
func canFinishV2(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 1 { // 如果只有一条边肯定是拓扑序列
		return true
	}
	inDegree := make(map[int]int, numCourses) // 入度  0->1  2->1  3->2  有2个指向1 1入度就是2 一个指向2 2的入度就是1
	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0 // 设置所有入度初始都为0
	}
	graph := make(map[int][]int, numCourses) // 邻接表图
	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1]) // 2->[3,4,5] 2指向了3,4,5 3条边
		inDegree[v[1]]++                        // 给被指向方入度+1
	}
	q := make([]int, 0, numCourses) // 所有入度为0的元素
	for i, deg := range inDegree {
		if deg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 移出所有入度为0的点
		curr := q[0]
		q = q[1:]
		for _, v := range graph[curr] { // 将入度为0的元素指向方入度也减1 如果为0 加入队列
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
		delete(graph, curr) // 从图中删除

	}
	return len(graph) == 0 //不存在环就是拓扑序列了
}

var canFinish = canFinishV2

func main() {
	prerequisites := [][]int{
		//{1, 0},
		{1, 4}, {2, 4}, {3, 1}, {3, 2},
	}
	numCourses := 5
	fmt.Println(canFinish(numCourses, prerequisites))
}
