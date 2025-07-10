package alg

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 栈 是一种操作受限的数据结构
type transStack struct{
	items []interface{}
}


func Init()*transStack{
    return &transStack{
		items: make([]interface{}, 0),
	}
}


func (Ts *transStack)Push(val interface{})bool{
     // 满了 就直接扩容了
	Ts.items = append(Ts.items, val)
	return true
}

func (Ts *transStack)Pop()interface{}{
	if len(Ts.items) == 0{
		return nil
	}
	count := len(Ts.items)-1
	pRet := Ts.items[count]
	Ts.items = Ts.items[:count]
	return pRet
}

// 链式栈
type transStackNode struct{
	val interface{}
	next *transStackNode
}

type transLinkStack struct{
	top *transStackNode
	size int
}

func InitLinkStack()*transLinkStack{
	return &transLinkStack{
		top : nil,
		size :0,
	}
}

func (tls *transLinkStack) Push(val interface{})bool{
	newNode := &transStackNode{
		val: val,
		next: tls.top,
	}
	tls.top = newNode
	tls.size++

	return true
}

func (tls *transLinkStack) Pop()interface{}{
	if tls.size == 0{
		return nil
	}
	ret := tls.top.val
	cur := tls.top
	tls.top = tls.top.next
	tls.size--
	cur = nil
	_ = cur
	return ret
}

func (tls *transLinkStack) Peek()interface{}{
	return tls.top.val
}

// 队列 一种操作受限的数据结构
// 设计一个循环队列 能够接受并发消费

type transQue struct{
	que []interface{}
	head int
	tail int
	length int
}

func InitTransQue(length int)*transQue{
	return &transQue{
		length: length,
		head: 0,
		tail: 0,
		que: make([]interface{}, length),
	}
}
// tail = (tail+1)%length
// tail == head empty
// (tail+1) % length == head full
// (head+1)%length  proceed
func (tq *transQue)EnQue(val interface{})(bool,error){
	if (tq.tail+1)%tq.length == tq.head{
		return false, errors.New("sorry q is full")
	}
    tq.que[tq.tail] = val
	tq.tail = (tq.tail+1)%tq.length
	return true,nil
}

func (tq *transQue) DeQue()(interface{},error){
	if tq.tail == tq.head {
		return nil,errors.New("sorry q is empty")
	}
	ret := tq.que[tq.head]
	tq.head = (tq.head+1)%tq.length
	return ret,nil
}


// bubbleSort i,j > 0, i < len, j<len-i-1 flag
func BubbleSort(arr []int)error{
	if len(arr) <= 1{
		return nil
	}

	for i:=0;i<len(arr);i++{
		flag := false
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				flag = true
			}
		}
		if !flag{
				return nil
			}
	}
	return nil
}

// 插入排序 有序与无序
func InsertSort(arr []int)error{
	if len(arr) <= 1{
		return nil
	}

	for i:= 1;i<len(arr);i++{
		val := arr[i]
		j := i-1 // 建立有序
		for ;j>=0;j--{
			if arr[j]>val{ // > 顺序 <逆序
				arr[j+1] = arr[j]
			}else{
				break
			}
		}
		arr[j+1] = val
	}
	return nil
}

// 选择排序
func SelectSort(arr []int)error{
	if len(arr) <= 1{
		return nil
	}
	for i := 0;i<len(arr);i++{
		minVal := arr[i]
		j := i+1
		loc := i
		for ;j<len(arr);j++{
			if arr[j]<=minVal{
				minVal = arr[j]
				loc = j
			}
		}
		arr[loc],arr[i] = arr[i],minVal // 最小值排到前面
	}
	return nil
}

// 希尔排序 shell的想法是将数据进行分列排序，最后步长为1时，排序自动完成
func ShellSort(arr []int)error{
	if len(arr) <= 1{
		return nil
	}
	n := len(arr)

	stepLen := len(arr)/2 // 一半步长 然后再切半 直到切的只有1
	for stepLen > 0{
		for i:= stepLen;i<n;i++{
			j:=i
			for j>=stepLen&&arr[j]<arr[j-stepLen]{ // 顺序排序
				arr[j],arr[j-stepLen] = arr[j-stepLen],arr[j] // 交换
				j = j-stepLen // 
			}
		}
		stepLen = stepLen /2
	}
	return nil
}


func ShellReSort(arr []int)error{
	if len(arr) == 0{
		return nil
	}
	n := len(arr)

	stepLen := n/2

	for stepLen > 0{
		for i := stepLen;i<n;i++{
			j := i
			for j>=stepLen&&arr[j]>arr[j-stepLen]{
				arr[j],arr[j-stepLen] = arr[j-stepLen],arr[j]
				j = j-stepLen
			}
		}
		stepLen = stepLen/2
	}
	return nil
}

// 归并排序 对排序进行递推公式的推到 找到终止条件 最后将梯队公式翻译成递归代码
func MergeSort(arr []int)[]int{
	if len(arr) <= 1{
		return arr
	}
	mid := len(arr)/2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left,right)
}


func merge(arrL []int,arrR []int)[]int{
	tmp := make([]int,0,len(arrL)+len(arrR))

	i,j := 0,0

	for i<len(arrL) && j<len(arrR){
		if arrL[i] <= arrR[j]{
			tmp = append(tmp, arrL[i])
			i++
		}else{
			tmp = append(tmp, arrR[j])
			j++
		}
	}

	tmp = append(tmp, arrL[i:]...)
	tmp = append(tmp, arrR[j:]...)
	return tmp
}


// quickSort

func QuickSort(arr []int,start,end int){
	if start >= end{
		return
	}
	// pivot
	p := partion(arr,start,end)
	// 继续递归区分
	QuickSort(arr,start,p-1)
	QuickSort(arr,p+1,end)
}

func partion(arr []int,start,end int)int{
	pivot := rand.Intn(end-start+1)+start // 随机法
	arr[pivot],arr[end] = arr[end],arr[pivot] // 做一次交换
	val := arr[end]
	i := start
	for j:=start;j<end;j++{
		if arr[j] <val{
			arr[i],arr[j] = arr[j],arr[i]
			i++
		} 
	}
	arr[i],arr[end] = arr[end],arr[i]

	return i
}


//  查找的终止条件是 arr[mid] = val
func Bsearch(arr []int,val int)int{
	low := 0
	high := len(arr) - 1
	for low <= high{
		mid := low + (high - low) / 2
		if arr[mid] == val{
			return mid
		}
		if arr[mid] < val{
			low = mid + 1
		}
		if arr[mid] > val{
			high = mid - 1
		}
	}
	return -1
}

// 递归

func BResearch(arr []int,low,high,val int) int{
	// 终止条件
	mid := low + ((high-low)>>1)
	if arr[mid] == val{
		return mid
	}

	if arr[mid] > val{ // 在左边
		return BResearch(arr,low,mid-1,val)
	}
	return BResearch(arr,mid+1,high,val)
}

// 查找存在重复数字的数组
func BsearchCopyStart(arr []int,val int)int{
	if len(arr) <= 1{
		return -1
	}

	low := 0
	high := len(arr) - 1

	for low<= high{
		mid := low+ ((high- low) >> 1)
		if arr[mid] < val{
			low = mid + 1
		}
		if arr[mid] > val{
			high = mid - 1
		}
		if mid == 0 || arr[mid - 1] != val{
			return mid
		}else{
			high = mid - 1
		}
	}
	return -1
}

// 查找最后一个存在重复数字的数组
func BsearchCopyLast(arr []int,val int)int{
	if len(arr) <= 1{
		return -1
	}

	low := 0
	high := len(arr) - 1

	for low<= high{
		mid := low+ ((high- low) >> 1)
		if arr[mid] < val{
			low = mid + 1
		}
		if arr[mid] > val{
			high = mid - 1
		}
		if mid == len(arr) - 1 || arr[mid + 1] != val{
			return mid
		}else{
			low = mid + 1
		}
	}
	return -1
}

// 循环有序数组 那就是位置不一样 不能按从小到大排序
// low := n  456123
func SearchNums(nums []int, target int) int {
    low := 0
    high := len(nums) - 1

   for low <= high {
		mid := low + ((high - low) >> 1)
        if nums[mid] == target {
			return mid
		}
		
		if nums[low] <= nums[mid]{
			if nums[low] <= target && target < nums[mid]{
				high = mid - 1
			}else{
				low = mid + 1
			}
		}else{
			if nums[mid + 1] <= target && target <= nums[high]{
				low = mid + 1
			}else{
				high = mid -1
			}
		}
		
	}
    return -1
}

// 二叉树 前序遍历
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 二叉查找树 动态数据集合的快速插入 删除 查找操作
// 要求 其节点的左子树小于节点值 右子树大于节点的值
// 查找
func SearchNode(val int,node *TreeNode)*TreeNode{
	if node == nil{
		return node
	}
	if node.Val == val{
		return node
	}
	if node.Val > val { // 在左子树查询
		return SearchNode(val,node.Left)
	}
    return SearchNode(val,node.Right)
	
}

// 插入
func SearchInsertNode(val int,node *TreeNode)*TreeNode{
	if node == nil{
		return  &TreeNode{
			 Val: val,
		}
	}
	if val > node.Val{
		if node.Right == nil{
			node.Right = &TreeNode{
				Val: val,
			}
			return node
		}
		SearchInsertNode(val,node.Right)
	}else if val < node.Val{
		if node.Left == nil{
			node.Left = &TreeNode{
				Val: val,
			}
			return node
		}
		SearchInsertNode(val,node.Left)
	}

	return node
}

// 删除
func SearchDeleteNode(val int,node *TreeNode) *TreeNode{
	if node == nil{
		return nil
	}
	rmNode := SearchNode(val,node)
	if rmNode.Left == nil && rmNode.Right == nil{ // 无子 干掉
		rmNode = nil 
	}
	if rmNode.Left != nil && rmNode.Right ==nil{
		rmNode.Val = rmNode.Left.Val
		return node
	}

	if rmNode.Right != nil && rmNode.Left == nil{
		rmNode.Val = rmNode.Right.Val
		return node
	}

	if rmNode.Right != nil && rmNode.Left != nil{
		bNode := rmNode.Right
		for bNode.Left != nil{
			bNode = bNode.Left
		}
		rmNode.Val = bNode.Val
		bNode = nil
	}
	return node
}


// 堆排序 堆 大顶堆 小顶堆 实现中位数查询 高性能定时器 合并有序的小文件的问题
func InsertHeap(arr []int,data int){
	arr = append(arr, data)
    n := len(arr)
	for n/2 > 0 && arr[n-1]>arr[(n-1)/2]{
		arr[n-1],arr[(n-1)/2] = arr[(n-1)/2],arr[n-1]
		n = (n-1)/2
	}
}

// Graph 
type Graph struct{
	points int // 顶点个数
	edges int // 边个数
	// 邻接表
	adj map[int]*redblacktree.Tree
}

func InitGraph(points int)*Graph{
	return &Graph{
		points: points,
		edges: 0,
		adj: make(map[int]*redblacktree.Tree),
	}
}

func (g *Graph) AddEdge(s,t int){
	if g.adj[s] == nil{
		g.adj[s] = redblacktree.NewWithIntComparator()
	}
	g.adj[s].Put(t,true)
	if g.adj[t] == nil{
		g.adj[t] = redblacktree.NewWithIntComparator()
	}
	g.adj[t].Put(s,true)
	g.edges++
}

// 进行图的广度优先搜索 搜索一条从s到t的路径
func (g *Graph) BFS(s,t int)[]int{
	if s == t{
		return []int{s}
	}
    // 根据三个结构进行访问的记录
	queue := make([]int,0)
	queue = append(queue,s)
	visited := make(map[int]bool)
	visited[s] = true
	prev := make([]int,g.points)
	for i:=0;i<g.points;i++{
		prev[i] = -1
	}
	
	for len(queue) > 0{
		cur := queue[0]
		queue = queue[1:]
		if g.adj[cur] == nil{
			continue
		}
		it := g.adj[cur].Iterator()
		for it.Next(){
			key := it.Key().(int)
			if !visited[key]{
				prev[key] = cur
				if key == t{
					return g.buildPath(prev,s,t)
				}
				visited[key] = true
				queue = append(queue,key)
			}
		}
	}
	return nil
}

// 修复后的buildPath函数
func (g *Graph) buildPath(prev []int, s, t int) []int {
	// 检查是否有路径
	if prev[t] == -1 && s != t {
		return nil // 没有找到路径
	}
	
	// 构建路径
	path := make([]int, 0)
	cur := t
	
	// 从终点回溯到起点
	for cur != -1 {
		path = append([]int{cur}, path...)
		cur = prev[cur]
	}
	
	// 验证路径的有效性
	if len(path) > 0 && path[0] == s && path[len(path)-1] == t {
		return path
	}
	
	return nil
}

func printPath(prev []int,s,t int){
	if prev[t] != -1 &&  t != s{
		printPath(prev,s,prev[t])
	}
	fmt.Println(t)
}

var Found = false

func (g *Graph) DFS(s,t int){
	visited := make(map[int]bool)
	visited[s] = true
	prev := make([]int,g.points)
	for i:=0;i<g.points;i++{
		prev[i] = -1
	}
	g.recurseDFS(s,t,visited,prev)
	printPath(prev,s,t)
}

func (g *Graph) recurseDFS(s,t int,visited map[int]bool,prev []int){
	if Found{
		return
	}
	visited[s] = true
	if s == t{
		Found = true
		return
	}
	if Found{
		return
	}
	it := g.adj[s].Iterator()
	for it.Next(){
		key := it.Key().(int)
		if !visited[key]{
			prev[key] = s
			g.recurseDFS(key,t,visited,prev)
		}
	}
}


