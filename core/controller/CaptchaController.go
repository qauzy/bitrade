package controller

func (this *CaptchaController) GetText() (result string) {
	return this.Text
}
func (this *CaptchaController) GetLength() (result int) {
	return this.Length
}
func (this *CaptchaController) GetWidth() (result int) {
	return this.Width
}
func (this *CaptchaController) GetHeight() (result int) {
	return this.Height
}
func (this *CaptchaController) SetText(Text string) {
	this.Text = this.Text
}
func (this *CaptchaController) SetLength(Length int) {
	len(this) = this.Length
}
func (this *CaptchaController) SetWidth(Width int) {
	this.Width = this.Width
}
func (this *CaptchaController) SetHeight(Height int) {
	this.Height = this.Height
}
func (this *CaptchaController) GenCaptcha(request *http.HttpServletRequest, response *http.HttpServletResponse) (err error) {
	// 在内存中创建图象
	var image *image.BufferedImage = BufferedImage(this.Width, this.Height, BufferedImage.TYPE_INT_RGB)
	// 获取图形上下文
	var g Graphics = image.GetGraphics()
	// 生成随机类
	var random *util.Random = new(util.Random)
	// 设定背景色
	g.SetColor(Color(255, 255, 255))
	g.FillRect(0, 0, this.Width, this.Height)
	// 设定字体
	g.SetFont(this.Font)
	// 取随机产生的验证码
	var sRand = ""
	//文字X坐标
	var codeX = this.Width / (this.Length + 1)
	for i := 0; i < this.Length; i += 1 {
		var rand string = String.ValueOf(this.GetRandomText())
		sRand += rand
		//将验证码输出到图象中
		g.SetColor(this.GetRandColor(100, 150))
		// 调用函数出来的颜色相同
		g.DrawString(rand, codeX*i+10, this.Height-20)
	}
	// 生成干扰线
	if this.CrossLine {
		for i := 0; i < random.NextInt(5)+5; i += 1 {
			g.SetColor(Color(random.NextInt(255)+1, random.NextInt(255)+1, random.NextInt(255)+1))
			g.DrawLine(random.NextInt(this.Width), random.NextInt(this.Height), random.NextInt(this.Width), random.NextInt(this.Height))
		}
	}
	if this.TwistImage {
		image = this.TwistImage(image)
	}
	var pageId string = request(request, "cid")
	var key = "CAPTCHA_" + pageId
	// 将验证码存入页面KEY值的SESSION里面
	var session *http.HttpSession = request.GetSession()
	session.SetAttribute(key, sRand)
	session.SetAttribute(key+"_time", System.CurrentTimeMillis())
	// 图象生效
	g.Dispose()
	var responseOutputStream *servlet.ServletOutputStream = response.GetOutputStream()
	// 输出图象到页面
	ImageIO.Write(image, "JPEG", responseOutputStream)
	// 以下关闭输入流！
	response.SetHeader("Cache-Control", "no-store")
	response.SetHeader("Pragma", "no-cache")
	response.SetDateHeader("Expires", 0)
	responseOutputStream.Flush()
	responseOutputStream.Close()
}
func (this *CaptchaController) TwistImage(buffImg *image.BufferedImage) (result *image.BufferedImage) {
	var random *util.Random = new(util.Random)
	var dMultValue = random.NextInt(10) + 5
	// 波形的幅度倍数，越大扭曲的程序越高，一般为3
	var dPhase float64 = random.NextInt(6)
	// 波形的起始相位，取值区间(0-2＊PI)
	var destBi *image.BufferedImage = BufferedImage(buffImg.GetWidth(), buffImg.GetHeight(), BufferedImage.TYPE_INT_RGB)
	var g Graphics = destBi.GetGraphics()
	g.SetColor(Color.WHITE)
	g.FillRect(0, 0, buffImg.GetWidth(), buffImg.GetHeight())
	for i := 0; i < destBi.GetWidth(); i += 1 {
		for j := 0; j < destBi.GetHeight(); j += 1 {
			var nOldX int = this.GetXPosition4Twist(dPhase, dMultValue, destBi.GetHeight(), i, j)
			var nOldY = j
			if nOldX >= 0 && nOldX < destBi.GetWidth() && nOldY >= 0 && nOldY < destBi.GetHeight() {
				destBi.SetRGB(nOldX, nOldY, buffImg.GetRGB(i, j))
			}
		}
	}
	return destBi
}
func (this *CaptchaController) GetXPosition4Twist(dPhase float64, dMultValue float64, Height int, xPosition int, yPosition int) (result int) {
	var PI = 0
	// 此值越大，扭曲程度越大
	var dx = (PI*yPosition/this.Height + dPhase).(float64)
	var dy float64 = Math.Sin(dx)
	return xPosition + (dy * dMultValue).(int)
}
func (this *CaptchaController) GetFont() (result Font) {
	return this.Font
}
func (this *CaptchaController) SetFont(Font Font) {
	this.Font = this.Font
}
func (this *CaptchaController) GetRandomText() (result byte) {
	var random *util.Random = new(util.Random)
	return this.Text.CharAt(random.NextInt(len(this.Text)))
}
func (this *CaptchaController) GetRandColor(fc int, bc int) (result Color) {
	var random *util.Random = new(util.Random)
	if fc > 255 {
		fc = 255
	}
	if bc > 255 {
		bc = 255
	}
	var r = fc + random.NextInt(bc-fc)
	var g = fc + random.NextInt(bc-fc)
	var b = fc + random.NextInt(bc-fc)
	return Color(r, g, b)
}

type CaptchaController struct {
	Text       string `gorm:"column:text" json:"text"`
	Length     int    `gorm:"column:length" json:"length"`
	Width      int    `gorm:"column:width" json:"width"`
	Height     int    `gorm:"column:height" json:"height"`
	Font       *Font  `gorm:"column:font" json:"font"`
	CrossLine  bool   `gorm:"column:cross_line" json:"crossLine"`
	TwistImage bool   `gorm:"column:twist_image" json:"twistImage"`
	BaseController
}
