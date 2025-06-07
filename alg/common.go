package alg

import (
	"errors"
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
