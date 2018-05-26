package controllers

import (
	"fmt"
	"strings"
	"transformat/project/logadmin/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/tools/present"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) ListApp() {
	logs.Debug("enter listApp")
	appinfoList, err := models.GetAllAppInfo()
	if err != nil {
		logs.Warn("get app info failed, err:%v", err)
	}
	c.Data["applist"] = appinfoList
	c.Layout = "layout/layout.html"
	c.TplName = "log/applist.html"
}

func (c *LogController) ListLog() {
	logs.Debug("enter list log")
	appinfoList, err := models.GetAllAppInfo()
	if err != nil {
		logs.Error("get appinfo failed")
		return
	}
	var list []models.AppInfo
	for _, v := range appinfoList {
		if len(v.LogPath) == 0 {
			continue
		}
		list = append(list, v)
	}
	c.Data["list"] = list
	c.Layout = "layout/layout.html"
	c.TplName = "log/log_list.html"
}

func (p *LogController) CreateApp() {

	logs.Debug("enter list app")

	p.Layout = "layout/layout.html"
	p.TplName = "log/app.html"
}

func (p *LogController) SubmitApp() {
	appName := p.GetString("app_name")
	appType := p.GetString("app_type")
	appIP := p.GetString("app_ip")
	logPath := p.GetString("log_path")

	p.Layout = "layout/layout.html"
	p.TplName = "log/app.html"

	if len(appName) == 0 || len(appType) == 0 || len(appIP) == 0 || len(logPath) == 0 {
		p.Data["Error"] = "appName或appType或appIP参数不正确"
		p.TplName = "log/app_error.html"
		return
	}
	appInfo := &models.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.LogPath = logPath

	ips := strings.Split(appIP, ",")
	for _, v := range ips {
		ip := models.IPInfo{}
		ip.IP = v
		appInfo.IPInfo = append(appInfo.IPInfo, ip)
	}
	_, err := models.InsertAppInfo(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("insert app info failed, err:%v", err)
		p.TplName = "log/app_error.html"
		return
	}

	p.Redirect("/app/list", 302)
}
