package main

import (
	"log"
)

func main() {
	//	print("start")
	//
	//		http.Handle("/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/")))
	//	http.Handle("/curl/", http.StripPrefix("/curl/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/curl"))))
	//	http.Handle("/make/", http.StripPrefix("/make/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/make"))))
	//
	//	log.Fatalln(http.ListenAndServe(":9000", nil))

	//	var input string
	//	fmt.Scanln(&input)
	//	log.Println(input)

	//	s := "b"
	//	ss := &s
	//	log.Printf("%p", &s)
	//	log.Printf("%p", ss)
	//	log.Printf("%p", &ss)

	a := []int{1, 2, 3}

	log.Printf("%T", a)
	log.Printf("%p", &a)
	test(a)

	bb := [3]int{1, 2, 3}
	aa := AA{p: &bb, len: 3, cap: 3}
	log.Printf("%p", &aa)
	test1(&aa)
}

func test(a []int) {
	log.Printf("%p", a)
	log.Printf("%p", &a)
	a = nil
	log.Printf("%p", a)
	log.Printf("%p", &a)
}

func test1(aa *AA) {
	log.Printf("%p", &aa)
	aa = nil
	log.Printf("%p", &aa)
}

type AA struct {
	p   *[3]int
	len int
	cap int
}
