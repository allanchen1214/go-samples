package main

func main() {
	repo := initBizRepo()
	var data interface{}
	repo.Store(data)

}
