package controller

import (
	"bitrade/core/log"
	"bitrade/core/util/MessageResult"
	"bitrade/core/util/RequestUtil"
	"net/http"
	"strings"
)

func (this *BaseController) Success() (result *MessageResult.MessageResult) {
	return MessageResult.NewMessageResultV2(0, "SUCCESS")
}
func (this *BaseController) SuccessWithMsg(msg string) (result *MessageResult.MessageResult) {
	return MessageResult.NewMessageResultV2(0, msg)
}
func (this *BaseController) SuccessWithMsgData(msg string, obj interface{}) (result *MessageResult.MessageResult) {
	var mr = MessageResult.NewMessageResultV2(0, msg)
	mr.SetData(obj)
	return mr
}
func (this *BaseController) SuccessDataAndTotal(object interface{}, total int64) (result *MessageResult.MessageResult) {
	var mr = MessageResult.NewMessageResultV2(0, "SUCCESS")
	mr.SetData(object)
	mr.SetTotal(total)
	return mr
}
func (this *BaseController) SuccessWithData(obj interface{}) (result *MessageResult.MessageResult) {
	var mr = MessageResult.NewMessageResultV2(0, "SUCCESS")
	mr.SetData(obj)
	return mr
}
func (this *BaseController) Error(msg string) (result *MessageResult.MessageResult) {
	return MessageResult.NewMessageResultV2(500, msg)
}
func (this *BaseController) ErrorWithCodeMsg(code int, msg string) (result *MessageResult.MessageResult) {
	return MessageResult.NewMessageResultV2(code, msg)
}
func (this *BaseController) Request(Request *http.Request, name string) (result string) {
	return strings.TrimSpace(Request.FormValue(name))
}
func (this *BaseController) GetRemoteIp(Request *http.Request) (result string) {
	return RequestUtil.RemoteIp(Request)
}
func (this *BaseController) IsMobile(Request *http.Request) (result bool) {
	var IsMobile = false
	var mobileAgents = []string{"iphone", "android", "phone", "ipad", "mobile", "wap", "netfront", "java", "opera mobi", "opera mini", "ucweb", "windows ce", "symbian", "series", "webos", "sony", "blackberry", "dopod", "nokia", "samsung", "palmsource", "xda", "pieplus", "meizu", "midp", "cldc", "motorola", "foma", "docomo", "up.browser", "up.link", "blazer", "helio", "hosin", "huawei", "novarra", "coolpad", "webos", "techfaith", "palmsource", "alcatel", "amoi", "ktouch", "nexian", "ericsson", "philips", "sagem", "wellcom", "bunjalloo", "maui", "smartphone", "iemobile", "spice", "bird", "zte-", "longcos", "pantech", "gionee", "portalmmm", "jig browser", "hiptop", "benq", "haier", "^lct", "320x320", "240x320", "176x220", "w3c ", "acs-", "alav", "alca", "amoi", "audi", "avan", "benq", "bird", "blac", "blaz", "brew", "cell", "cldc", "cmd-", "dang", "doco", "eric", "hipt", "inno", "ipaq", "java", "jigs", "kddi", "keji", "leno", "lg-c", "lg-d", "lg-g", "lge-", "maui", "maxo", "midp", "mits", "mmef", "mobi", "mot-", "moto", "mwbp", "nec-", "newt", "noki", "oper", "palm", "pana", "pant", "phil", "play", "port", "prox", "qwap", "sage", "sams", "sany", "sch-", "sec-", "send", "seri", "sgh-", "shar", "sie-", "siem", "smal", "smar", "sony", "sph-", "symb", "t-mo", "teli", "tim-", "tsm-", "upg1", "upsi", "vk-v", "voda", "wap-", "wapa", "wapi", "wapp", "wapr", "webc", "winw", "winw", "xda", "xda-", "Googlebot-Mobile"}
	// 根据cookie选择模板，暂时未用
	var cookies = Request.Cookies()
	if cookies != nil && len(cookies) > 0 {
		for _, cookie := range cookies {
			if strings.ToUpper(cookie.Name) == "SPARK_THEME" {
				log.Info("SPAKR_THEME:" + cookie.Value)
				IsMobile = cookie.Value == "WAP"
				break
			}
		}
	}
	// 根据userAgent自动选择
	var userAgent = strings.ToLower(Request.UserAgent())
	if userAgent != "" {
		for _, mobileAgent := range mobileAgents {
			if strings.Index(userAgent, mobileAgent) >= 0 {
				log.Info("User-Agent HIT:" + mobileAgent)
				IsMobile = true
				break
			}
		}
	}
	return IsMobile
}

type BaseController struct {
}
