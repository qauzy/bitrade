package controller

import "github.com/qauzy/util/sets/hashset"

func (this *GeetestController) StartCaptcha(request *http.HttpServletRequest) (result string) {
	var resStr = "{}"
	var userid = "bdtop"
	//自定义参数,可选择添加
	var param = hashset.New[string]()
	var ip string = getRemoteIp(request)
	param.Put("user_id", userid)
	//网站用户id
	param.Put("client_type", "web")
	//web:电脑上的浏览器；h5:手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生SDK植入APP应用的方式
	param.Put("ip_address", ip)
	//传输用户请求验证时所携带的IP
	//进行验证预处理
	var gtServerStatus int = this.GtSdk.PreProcess(param)
	//将服务器状态设置到session中
	request.GetSession().SetAttribute(this.GtSdk.GtServerStatusSessionKey, gtServerStatus)
	//将userid设置到session中
	request.GetSession().SetAttribute("userid", userid)
	resStr = this.GtSdk.GetResponseStr()
	return resStr
}

type GeetestController struct {
	GtSdk *system.GeetestLib
	BaseController
}
