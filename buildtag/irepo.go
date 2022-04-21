package main

type BizRepo interface {
	Store(data interface{})
}

type BizRepoImpl struct {
}

func initBizRepo() BizRepo {
	impl := BizRepoImpl{}
	return &impl
}
