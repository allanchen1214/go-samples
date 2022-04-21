//go:build b2
// +build b2

package main

import (
	"fmt"
)

func (repo BizRepoImpl) Store(data interface{}) {
	fmt.Println("biz repo2 impl")
}
