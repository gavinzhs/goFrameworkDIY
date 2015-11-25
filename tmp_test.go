package main

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	s := "/1/2/3"
	ss := strings.Split(s, "/")
	reg := "([^/]+)"
	r, err := regexp.Compile(reg)
	if err != nil {
		log.Fatalf("regexp.Compile err : %v", err)
	}
	for k, v := range ss {
		log.Printf("第%d个是%s!, 是否满足正则呢%t?", k, v, r.MatchString(v))
	}
}

func TestString(t *testing.T) {
	s := "123"
	s = "123" + "---" + s
	log.Println(s)
}

func TestGetParams(t *testing.T) {
	u := "/m/help/:id([0-9]+)/:oid([0-9]+)"
	s := "/m/help/123/456"
	//如何获取到id = 123 oid= 456
	//构造正则拿出两个数         拿出两个参数的顺序  两个一对一就能得出
	parts := strings.Split(u, "/")
	params := make([]string, 0)
	for k, v := range parts {
		if strings.HasPrefix(v, ":") {

			reg := "([^/]+)"

			if index := strings.Index(v, "("); index == -1 {
				parts[k] = reg
				params = append(params, v)
			} else {
				parts[k] = v[index:]
				params = append(params, v[:index])
			}
		}
	}

	mu := strings.Join(parts, "/")
	r, err := regexp.Compile(mu)
	if err != nil {
		log.Fatalf("regexp compile err : %v", err)
	}
	originalsubs := r.FindAllStringSubmatch(s, -1)
	subs := originalsubs[0][1:]
	for k, v := range params {
		log.Println(strings.Replace(v, ":", "", -1) + "=" + subs[k])
	}
}
