package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

import "daluren/models"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) Put() {
	var paiPanJson models.PaiPanJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &paiPanJson)
	//objectid := models.AddOne(ob)
	//this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	c.Data["json"] = paiPanJson.PaiPan()
	//c.ServeJSON()
	//fmt.Println(c.Ctx.Input.RequestBody)
	c.ServeJSON()
}
