package alg

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



