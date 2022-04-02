package RequestUtil

import (
	"bitrade/libs/common-lang/StringUtils"
	"net/http"
)

func RemoteIp(request *http.Request) (result string) {
	if StringUtils.IsNotBlank(request.Header.Get("X-Real-IP")) {
		return request.Header.Get("X-Real-IP")
	} else if StringUtils.IsNotBlank(request.Header.Get("X-Forwarded-For")) {
		return request.Header.Get("X-Forwarded-For")
	} else if StringUtils.IsNotBlank(request.Header.Get("Proxy-Client-IP")) {
		return request.Header.Get("Proxy-Client-IP")
	}
	return
}

//func GetAreaDetail(apiUrl string, apiKey string, apiValue string) (result *hashmap.Map[string, string]) {
//	var jsonObject *json.JSONObject = JSONObject.FromObject(HttpClientUtil.Get(apiUrl, apiKey, apiValue))
//	log.Info("getAreaDetail=" + jsonObject.ToString())
//	var resultMap = JSONObject.ToBean(jsonObject, *hashmap.Map[string, interface {
//	}]).(*hashmap.Map[string, string])
//	return resultMap
//}
