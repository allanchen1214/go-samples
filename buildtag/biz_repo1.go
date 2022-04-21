//go:build b1
// +build b1

package main

import "fmt"

func (repo BizRepoImpl) Store(data interface{}) {
	fmt.Println("biz repo1 impl")
}
