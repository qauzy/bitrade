
package controller

import (
	"bitrade/core/log"
	"github.com/qauzy/util/maps/hashmap"
)

var (
	API_HOST          string                = "39.100.79.158"
	SIGNATURE_METHOD  string                = "HmacSHA256"
	SIGNATURE_VERSION string                = "2"
	ZONE_GMT          time.ZoneId           = ZoneId.Of("Z")
	DT_FORMAT         text.SimpleDateFormat = SimpleDateFormat("uuuu-MM-dd'T'HH:mm:ss")
)

func CreateSignature(method string, path string, apiKey string, timeStamp string, oMap *hashmap.Map[string, interface {
}], secretKey string) (result string, err error) {
	var sb StringBuilder = StringBuilder(1024)
	// GET
	sb.Append(method.ToUpperCase()).Append(('\n').Append(API_HOST.ToLowerCase()).Append(('\n').Append(path).Append(('\n')
	var joiner StringJoiner = StringJoiner("&")
	joiner.Add("accessKeyId=" + apiKey).Add("signatureMethod=" + SIGNATURE_METHOD).Add("signatureVersion=" + SIGNATURE_VERSION).Add("timestamp=" + Encode(timeStamp))
	//拼接 遍历map
	var entries Iterator = oMap.EntrySet().Iterator()
	for entries.HasNext() {
		var entry *hashmap.Map[string, string] = entries.Next()
		joiner.Add(entry.GetKey() + "=" + entry.GetValue())
	}
	log.Infof("sb=%v,joiner=%v", sb.ToString(), joiner.ToString())
	return Sign(sb.ToString()+joiner.ToString(), secretKey)
}
func Sign(message string, secret string) (result string) {
	exception := func() (err error) {
		var sha256_HMAC *crypto.Mac = Mac.GetInstance("HmacSHA256")
		var secretKeySpec *spec.SecretKeySpec = SecretKeySpec(secret.GetBytes(StandardCharsets.UTF_8), "HmacSHA256")
		sha256_HMAC.Init(secretKeySpec)
		var hash []byte = sha256_HMAC.DoFinal(message.GetBytes())
		return Base64.GetEncoder().EncodeToString(hash)
		return
	}()
	if exception != nil {
		return RuntimeException("Unable to sign message.", e)
	}
}
func Encode(code string) (result string) {
	exception := func() (err error) {
		return URLEncoder.Encode(code, "UTF-8").ReplaceAll("\\+", "%20")
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		return nil
	}
}

type TestApi struct {
}

