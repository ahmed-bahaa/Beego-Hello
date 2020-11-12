package main

import (
	_ "hello2/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

