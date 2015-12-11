package main

import (
	"beegoDIY"
)

type Post struct {
	beegodiy.Controller
}

func (p *Post) Get() {
	beegodiy.Trace("id 是: ", p.Ct.Params["id"])
	beegodiy.Trace("[post]get method")
}

func (p *Post) Post() {
	beegodiy.Trace("id 是: ", p.Data["id"])
	beegodiy.Trace("[post]get method")
}

func (p *Post) Delete() {
	beegodiy.Trace("id 是: ", p.Data["id"])
	beegodiy.Trace("[post]get method")
}

type PostList struct {
	beegodiy.Controller
}

func (p *PostList) Get() {
	beegodiy.Trace("[postlist]get method")
}
