package sdk

var (
	VERIFY_API string = "http://c.dun.163yun.com/api/v2/verify"
	VERSION    string = "v2"
)

func NewNECaptchaVerifier(CaptchaId string, SecretPair NESecretPair) (this *NECaptchaVerifier) {
	this = new(NECaptchaVerifier)
	Validate.NotBlank(this.CaptchaId, "captchaId为空")
	Validate.NotNull(this.SecretPair, "secret为null")
	Validate.NotBlank(this.SecretPair.SecretId, "secretId为空")
	Validate.NotBlank(this.SecretPair.SecretKey, "secretKey为空")
	this.CaptchaId = this.CaptchaId
	this.SecretPair = this.SecretPair
	return
}
func (this *NECaptchaVerifier) Verify(validate string, user string) (result bool) {
	if StringUtils.IsEmpty(validate) || StringUtils.Equals(validate, "null") {
		return false
	}
	if user == nil {
		user = ""
	} else {
		user = user
	}
	// bugfix:如果user为null会出现签名错误的问题
	var params = hashset.New[string]()
	params.Put("captchaId", this.CaptchaId)
	params.Put("validate", validate)
	params.Put("user", user)
	// 公共参数
	params.Put("secretId", this.SecretPair.SecretId)
	params.Put("version", VERSION)
	params.Put("timestamp", String.ValueOf(System.CurrentTimeMillis()))
	params.Put("nonce", String.ValueOf(ThreadLocalRandom.Current().NextInt()))
	// 计算请求参数签名信息
	var signature string = Sign(this.SecretPair.SecretKey, params)
	params.Put("signature", signature)
	var resp string = HttpClient4Utils.SendPost(VERIFY_API, params)
	System.Out.Println("resp = " + resp)
	return this.VerifyRet(resp)
}
func Sign(secretKey string, params *hashmap.Map[string, string]) (result string) {
	var keys []string = params.KeySet().ToArray(make([]string, 0))
	Arrays.Sort(keys)
	var sb StringBuffer = StringUtils.NewStringBuilder()
	for _, key := range keys {
		sb.Append(key).Append(params.Get(key))
	}
	sb.Append(secretKey)
	exception := func() (err error) {
		return DigestUtils.Md5Hex(sb.ToString().GetBytes("UTF-8"))
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		// 一般编码都支持的。。
	}
	return nil
}
func (this *NECaptchaVerifier) VerifyRet(resp string) (result bool) {
	if StringUtils.IsEmpty(resp) {
		return false
	}
	exception := func() (err error) {
		var j *fastjson.JSONObject = JSONObject.ParseObject(resp)
		return j.GetBoolean("result")
		return
	}()
	if exception != nil {
		return false
	}
}

type NECaptchaVerifier struct {
	CaptchaId  string
	SecretPair *NESecretPair
}
