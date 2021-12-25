package AliyunUtil

import (
	"bitrade/core/util/dateutils"
	"bitrade/libs/unirest"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"time"
)

func Md5EncodeToString(b []byte) string {
	hexCode := md5.Sum(b)
	return hex.EncodeToString(hexCode[:])
}

/*
 * 计算MD5+BASE64
 */
func MD5Base64(src []byte) (result string) {
	if src == nil {
		return
	}
	var encodeStr = Md5EncodeToString(src)
	result = base64.StdEncoding.EncodeToString([]byte(encodeStr))
	return result
}

/*
 * 计算 HMAC-SHA1
 */
func HMACSha1(data string, keyStr string) (result string) {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))
	//进行base64编码
	result = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return result
}
func DoPost(uri string, body []byte, accessId string, accessKey string) (result []byte, err error) {
	var method string = "POST"
	var accept string = "application/json"
	var content_type string = "application/json"
	var u *url.URL
	u, err = url.Parse(uri)
	if err != nil {
		return
	}
	var path = u.Path
	var date string = dateutils.ToGMTString(time.Now())
	// 1.对body做MD5+BASE64加密
	var bodyMd5 string = MD5Base64(body)
	var stringToSign string = method + "\n" + accept + "\n" + bodyMd5 + "\n" + content_type + "\n" + date + "\n" + path
	// 2.计算 HMAC-SHA1
	var signature string = HMACSha1(stringToSign, accessKey)
	// 3.得到 authorization header
	var authHeader string = "Dataplus " + accessId + ":" + signature

	result, err = unirest.New().Post().SetURL(uri).Header("accept", accept).Header("content-type", content_type).Header("date", date).Header("Authorization", authHeader).SetJSONBody(body).Send().AsBytes()
	if err != nil {
		return
	}
	return
}
