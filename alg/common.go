package alg

import (
	"errors"
	"math/rand"
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
	pivot := rand.Intn(end-start+1)+start
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