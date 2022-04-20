package main

import (
	"fmt"
	"sync/atomic"
)

//func AddT(addr *T, delta T)(new T)
//func LoadT(addr *T) (val T)
//func StoreT(addr *T, val T)
//func SwapT(addr *T, new T) (old T)
//func CompareAndSwapT(addr *T, old, new T) (swapped bool)

type Page struct {
	views uint32
}

func (page *Page) SetViews(n uint32) {
	atomic.StoreUint32(&page.views, n)
}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}

//----

type MyData struct {
	iValue   uint32
	strValue string
	arrs     []uint32
	m        map[string]string
}

var myDataValue atomic.Value

func main() {
	var page Page
	page.SetViews(2022)
	fmt.Println(page.Views())

	//----

	data := new(MyData)
	data.iValue = 2022
	data.strValue = "atomic-operate"
	data.arrs = make([]uint32, 0)
	data.arrs = append(data.arrs, 1)
	data.arrs = append(data.arrs, 2)
	data.arrs = append(data.arrs, 3)
	data.m = make(map[string]string)
	data.m["firstname"] = "chen"
	data.m["secondname"] = "allan"

	myDataValue.Store(data)

	storeVal := myDataValue.Load()
	if storeVal != nil {
		storeData := storeVal.(*MyData)
		fmt.Println(storeData.iValue)
		fmt.Println(storeData.strValue)
		fmt.Println(len(storeData.arrs))
		for _, item := range storeData.arrs {
			fmt.Println(item)
		}
		fmt.Println(len(storeData.m))
		for k, v := range storeData.m {
			fmt.Printf("k: %s, v: %s\n", k, v)
		}
	}

}
