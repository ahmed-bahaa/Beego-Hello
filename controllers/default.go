package controllers

import (
	beego "github.com/astaxie/beego/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "Behoo.me"
	c.Data["Email"] = "ahmedb.abdelhamid@gmail.com"
	// c.TplName = "index.tpl"		//if we removed this will be by defaulted look for this file in views/usercontroller/get.tpl

}
