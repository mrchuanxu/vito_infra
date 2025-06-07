package alg

// 栈 是一种操作受限的数据结构
type TransStack struct{
	items []interface{}
	count int
	length int
}


func Init()*TransStack{
    return &TransStack{
		items: make([]interface{}, 4),
		count: 0,
		length: 4,
	}
}


func (Ts *TransStack)Push(val interface{})bool{
     // 满了 就直接扩容了
	Ts.items = append(Ts.items, val)
	Ts.count = len(Ts.items) - 1
	Ts.length = cap(Ts.items) // 直接使用切片的扩容机制
	return true
}

func (Ts *TransStack)Pop()interface{}{
	if len(Ts.items) == 0{
		return nil
	}
	count := len(Ts.items)-1
	pRet := Ts.items[count]
	Ts.items = Ts.items[0:count-1]
	Ts.count = len(Ts.items)-1
	return pRet
}




