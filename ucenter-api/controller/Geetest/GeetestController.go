package Geetest

import (
	"bitrade/core/controller"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/maps/hashmap"
)

func (this *GeetestController) StartCaptcha(ctx *gin.Context, request *http.HttpServletRequest) (result string) {
	var resStr = "{}"
	var userid = "bdtop"
	//自定义参数,可选择添加
	var param = hashmap.New[string, string]()
	var ip = getRemoteIp(request)
	param.Put("user_id", userid)
	//网站用户id
	param.Put("client_type", "web")
	//web:电脑上的浏览器；h5:手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生SDK植入APP应用的方式
	param.Put("ip_address", ip)
	//传输用户请求验证时所携带的IP
	//进行验证预处理
	var gtServerStatus = this.GtSdk.PreProcess(param)
	//将服务器状态设置到session中
	request.GetSession().SetAttribute(this.GtSdk.GtServerStatusSessionKey, gtServerStatus)
	//将userid设置到session中
	request.GetSession().SetAttribute("userid", userid)
	resStr = this.GtSdk.GetResponseStr()
	return resStr
}
func NewGeetestController(gtSdk *system.GeetestLib) (ret *GeetestController) {
	ret = new(GeetestController)
	ret.GtSdk = gtSdk
	return
}

type GeetestController struct {
	GtSdk *system.GeetestLib
	controller.BaseController
}
