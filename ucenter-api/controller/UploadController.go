package controller

import (
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"strings"
	"time"
)

func (this *UploadController) UploadOssImage(request *http.HttpServletRequest, response *http.HttpServletResponse, file *multipart.MultipartFile) (result *util.MessageResult, err error) {
	log.Info(request.GetSession().GetServletContext().GetResource("/").ToString())
	response.SetCharacterEncoding("UTF-8")
	response.SetContentType("text/html; charset=UTF-8")
	if ServletFileUpload.IsMultipartContent(request) == false {
		this.SourceService.GetMessage("FORM_FORMAT_ERROR")
	}
	if file != nil == false {
		this.SourceService.GetMessage("NOT_FIND_FILE")
	}
	var fileType string = UploadFileUtil.GetFileType(file.GetInputStream())
	log.Info("fileType=" + fileType)
	var directory string = SimpleDateFormat("yyyy/MM/dd/").Format(time.Now())
	var ossClient *oss.OSSClient = OSSClient(this.AliyunConfig.GetOssEndpoint(), this.AliyunConfig.GetAccessKeyId(), this.AliyunConfig.GetAccessKeySecret())
	exception := func() (err error) {
		var fileName string = file.GetOriginalFilename()
		var suffix string = fileName.Substring(fileName.LastIndexOf("."), len(fileName))
		if !this.AllowedFormat.Contains(suffix.Trim().ToLowerCase()) {
			return MessageResult.Error(this.SourceService.GetMessage("FORMAT_NOT_SUPPORT"))
		}
		if fileType == nil || !this.AllowedFormat.Contains(fileType.Trim().ToLowerCase()) {
			return MessageResult.Error(this.SourceService.GetMessage("FORMAT_NOT_SUPPORT"))
		}
		var key = directory + GeneratorUtil.GetUUID() + suffix
		//压缩文件
		var path = request.GetSession().GetServletContext().GetRealPath("/") + "upload/" + file.GetOriginalFilename()
		var tempFile *io.File = File(path)
		FileUtils.CopyInputStreamToFile(file.GetInputStream(), tempFile)
		log.Info("=================压缩前" + len(tempFile))
		UploadFileUtil.ZipWidthHeightImageFile(tempFile, tempFile, 425, 638, 0)
		log.Info("=================压缩后" + len(tempFile))
		ossClient.PutObject(this.AliyunConfig.GetOssBucketName(), key, file.GetInputStream())
		var uri string = this.AliyunConfig.ToUrl(key)
		var mr *util.MessageResult = MessageResult(0, this.SourceService.GetMessage("UPLOAD_SUCCESS"))
		mr.SetData(uri)
		return mr
		return
	}()
	if exception != nil {
		return MessageResult.Error(500, oe.GetErrorMessage())
	}
	if exception != nil {
		return MessageResult.Error(500, ce.GetErrorMessage())
	}
	if exception != nil {
		e.PrintStackTrace()
		return MessageResult.Error(500, this.SourceService.GetMessage("SYSTEM_ERROR"))
	}
}
func (this *UploadController) UploadVideo(request *http.HttpServletRequest, response *http.HttpServletResponse, file *multipart.MultipartFile) (result *util.MessageResult, err error) {
	log.Info(request.GetSession().GetServletContext().GetResource("/").ToString() + "==========实名认证视频上传")
	response.SetCharacterEncoding("UTF-8")
	response.SetContentType("text/html; charset=UTF-8")
	if ServletFileUpload.IsMultipartContent(request) == false {
		this.SourceService.GetMessage("FORM_FORMAT_ERROR")
	}
	if file != nil == false {
		this.SourceService.GetMessage("NOT_FIND_FILE")
	}
	log.Info("该文件的文件流转为流类型>>>>>>>>>>>" + UploadFileUtil.GetFileHeader(file.GetInputStream()))
	//文件大小判断
	var fileSize int64 = file.GetSize()
	var fileSizeMB = fileSize / (1024 * 1024)
	log.Info("==============文件大小：" + fileSize)
	if fileSizeMB > 20 {
		return MessageResult.Error("超过文件大小设置")
	}
	var directory string = SimpleDateFormat("yyyy/MM/dd/").Format(time.Now())
	var ossClient *oss.OSSClient = OSSClient(this.AliyunConfig.GetOssEndpoint(), this.AliyunConfig.GetAccessKeyId(), this.AliyunConfig.GetAccessKeySecret())
	exception := func() (err error) {
		var fileName string = file.GetOriginalFilename()
		log.Info("===========fileName:" + fileName)
		var suffix string = fileName.Substring(fileName.LastIndexOf("."), len(fileName))
		log.Info("suffix=" + suffix)
		var key = directory + GeneratorUtil.GetUUID() + suffix
		log.Info("key:" + key)
		ossClient.PutObject(this.AliyunConfig.GetOssBucketName(), key, file.GetInputStream())
		var uri string = this.AliyunConfig.ToUrl(key)
		log.Info(">>>>>>>>>>上传成功>>>>>>>>>>>>>")
		var mr *util.MessageResult = MessageResult(0, this.SourceService.GetMessage("UPLOAD_SUCCESS"))
		mr.SetData(uri)
		return mr
		return
	}()
	if exception != nil {
		return MessageResult.Error(500, oe.GetErrorMessage())
	}
	if exception != nil {
		System.Out.Println("Error Message: " + ce.GetMessage())
		return MessageResult.Error(500, ce.GetErrorMessage())
	}
	if exception != nil {
		e.PrintStackTrace()
		return MessageResult.Error(500, this.SourceService.GetMessage("SYSTEM_ERROR"))
	}
}
func (this *UploadController) UploadLocalImage(request *http.HttpServletRequest, response *http.HttpServletResponse, file *multipart.MultipartFile) (result *util.MessageResult, err error) {
	log.Info(request.GetSession().GetServletContext().GetResource("/").ToString())
	response.SetCharacterEncoding("UTF-8")
	response.SetContentType("text/html; charset=UTF-8")
	if ServletFileUpload.IsMultipartContent(request) == false {
		this.SourceService.GetMessage("FORM_FORMAT_ERROR")
	}
	if file != nil == false {
		this.SourceService.GetMessage("NOT_FIND_FILE")
	}
	//验证文件类型
	var fileName string = file.GetOriginalFilename()
	var suffix string = fileName.Substring(fileName.LastIndexOf("."), len(fileName))
	if !this.AllowedFormat.Contains(suffix.Trim().ToLowerCase()) {
		return MessageResult.Error(this.SourceService.GetMessage("FORMAT_NOT_SUPPORT"))
	}
	var result string = UploadFileUtil.UploadFile(file, fileName)
	if result != nil {
		var mr *util.MessageResult = MessageResult(0, this.SourceService.GetMessage("UPLOAD_SUCCESS"))
		mr.SetData(this.Host + result)
		return mr
	} else {
		var mr *util.MessageResult = MessageResult(0, this.SourceService.GetMessage("FAILED_TO_WRITE"))
		mr.SetData(result)
		return mr
	}
}
func (this *UploadController) Base64UpLoad(base64Data string, oss bool) (result *util.MessageResult) {
	log.Info("oss == " + oss)
	var result *util.MessageResult = new(util.MessageResult)
	var uri string
	exception := func() (err error) {
		log.Debug("上传文件的数据：" + base64Data)
		var dataPrix = ""
		var data = ""
		if base64Data == nil || "".Equals(base64Data) {
			return Exception(this.SourceService.GetMessage("NOT_FIND_FILE"))
		} else {
			var d []string = strings.Split(base64Data, "base64,")
			if d != nil && len(d) == 2 {
				dataPrix = d[0]
				data = d[1]
			} else {
				return Exception(this.SourceService.GetMessage("DATA_ILLEGAL"))
			}
		}
		log.Debug("对数据进行解析，获取文件名和流数据")
		var suffix = ""
		if "data:image/jpeg;".EqualsIgnoreCase(dataPrix) {
			//data:image/jpeg;base64,base64编码的jpeg图片数据
			suffix = ".jpg"
		} else if "data:image/x-icon;".EqualsIgnoreCase(dataPrix) {
			//data:image/x-icon;base64,base64编码的icon图片数据
			suffix = ".ico"
		} else if "data:image/gif;".EqualsIgnoreCase(dataPrix) {
			//data:image/gif;base64,base64编码的gif图片数据
			suffix = ".gif"
		} else if "data:image/png;".EqualsIgnoreCase(dataPrix) {
			//data:image/png;base64,base64编码的png图片数据
			suffix = ".png"
		} else {
			return Exception(this.SourceService.GetMessage("FORMAT_NOT_SUPPORT"))
		}
		var directory string = SimpleDateFormat("yyyy/MM/dd/").Format(time.Now())
		var key = directory + GeneratorUtil.GetUUID() + suffix
		//因为BASE64Decoder的jar问题，此处使用spring框架提供的工具包
		var bs []byte = Base64Utils.DecodeFromString(data)
		if oss == true {
			//阿里云上传
			var ossClient *oss.OSSClient = OSSClient(this.AliyunConfig.GetOssEndpoint(), this.AliyunConfig.GetAccessKeyId(), this.AliyunConfig.GetAccessKeySecret())
			exception := func() (err error) {
				//使用apache提供的工具类操作流
				var is *io.InputStream = ByteArrayInputStream(bs)
				//FileUtils.writeByteArrayToFile(new File(Global.getConfig(UPLOAD_FILE_PAHT), tempFileName), bs);
				ossClient.PutObject(this.AliyunConfig.GetOssBucketName(), key, is)
				uri = this.AliyunConfig.ToUrl(key)
				result.SetData(uri)
				result.SetMessage(this.SourceService.GetMessage("UPLOAD_SUCCESS"))
				log.Debug("上传成功,accessKey:{}", key)
				return result
				return
			}()
			if exception != nil {
				log.Info(ee.GetMessage())
				return Exception(this.SourceService.GetMessage("FAILED_TO_WRITE"))
			}
		} else {
			//本地上传
			uri = UploadFileUtil.UploadFileByBytes(bs, GeneratorUtil.GetUUID())
			result.SetData(uri)
		}
		return
	}()
	if exception != nil {
		log.Debug("上传失败," + e.GetMessage())
		result.SetCode(500)
		result.SetMessage(e.GetMessage())
	}
	return result
}

type UploadController struct {
	AllowedFormat string
	AllowedVideo  string
	AliyunConfig  *config.AliyunConfig
	SourceService *service.LocaleMessageSourceService
	Host          string
}
