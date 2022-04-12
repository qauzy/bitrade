package Aide

import (
	"bitrade/core/constant/Platform"
	"bitrade/core/constant/SysAdvertiseLocation"
	"bitrade/core/constant/SysConstant"
	"bitrade/core/constant/SysHelpClassification"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/MessageResult"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/fastjson"
	"strconv"
)

func (this *AideController) KeyWords(ctx *gin.Context) (result *MessageResult.MessageResult) {
	var websiteInformation, err = this.WebsiteInformationService.FetchOne()
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	result = MessageResult.Success()
	result.SetData(websiteInformation)
	return result
}
func (this *AideController) SysAdvertise(ctx *gin.Context, sysAdvertiseLocation *SysAdvertiseLocation.SysAdvertiseLocation) (result *MessageResult.MessageResult) {
	var list, err = this.SysAdvertiseService.FindAllNormal(sysAdvertiseLocation)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	result = MessageResult.Success()
	result.SetData(list)
	return result
}
func (this *AideController) SysHelp(ctx *gin.Context, sysHelpClassification *SysHelpClassification.SysHelpClassification) (result *MessageResult.MessageResult) {
	var list arraylist.List[entity.SysHelp]
	var err error
	if sysHelpClassification == nil {
		list, err = this.SysHelpService.FindAllByStatusNotAndSort()
		if err != nil {
			return MessageResult.Error(err.Error())
		}
	} else {
		list, err = this.SysHelpService.FindBySysHelpClassification(sysHelpClassification)
		if err != nil {
			return MessageResult.Error(err.Error())
		}
	}
	result = MessageResult.Success()
	result.SetData(list)
	return result
}
func (this *AideController) SysHelpById(ctx *gin.Context, id int64) (result *MessageResult.MessageResult) {
	//List<SysHelp> list = sysHelpService.findBySysHelpClassification(sysHelpClassification);
	var sysHelp, err = this.SysHelpService.FindById(id)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	result = MessageResult.Success()
	result.SetData(sysHelp)
	return result
}
func (this *AideController) SysHelpEx(ctx *gin.Context, platform *Platform.Platform) (result *MessageResult.MessageResult) {
	var revision, err = this.AppRevisionService.FindRecentVersion(platform)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	if revision != nil {
		var result = MessageResult.Success()
		result.SetData(revision)
		return result
	} else {
		return MessageResult.Error(this.MsService.GetMessage("NO_UPDATE"))
	}
}
func (this *AideController) SysAllHelp(ctx *gin.Context, total int) (result *MessageResult.MessageResult) {
	var resultStr = this.RedisUtil.Get(SysConstant.SYS_HELP)
	if resultStr != "" {
		var jarr arraylist.List[fastjson.JSONObject]
		err := json.Unmarshal([]byte(resultStr), &jarr)
		if err == nil {
			return this.SuccessWithData(jarr)
		}
	}
	//HELP("新手指南"),
	var jsonResult = arraylist.New[fastjson.JSONObject]()
	var sysHelpPage, err = this.SysHelpService.FindByCondition(1, total, SysHelpClassification.HELP)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	var jsonSysHelp = fastjson.NewJSONObject()
	jsonSysHelp.Put("content", sysHelpPage.GetContent())
	jsonSysHelp.Put("title", "新手指南")
	jsonSysHelp.Put("cate", "0")
	jsonResult.Add(jsonSysHelp)
	//FAQ("常见问题"),
	sysFaqPage, err := this.SysHelpService.FindByCondition(1, total, SysHelpClassification.FAQ)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	var jsonSysFaq = fastjson.NewJSONObject()
	jsonSysFaq.Put("content", sysFaqPage.GetContent())
	jsonSysFaq.Put("title", "常见问题")
	jsonSysFaq.Put("cate", "1")
	jsonResult.Add(jsonSysFaq)
	this.RedisUtil.SetExpireKV(SysConstant.SYS_HELP, jsonResult, int64(SysConstant.SYS_HELP_EXPIRE_TIME))
	return this.SuccessWithData(jsonResult)
	//RECHARGE("充值指南"),
	//Page<SysHelp> sysRechangePage = sysHelpService.findByCondition(1,total,SysHelpClassification.HELP);
	//TRANSACTION("交易指南"),
	//Page<SysHelp> sysTransactonPage = sysHelpService.findByCondition(1,total,SysHelpClassification.HELP);
}
func (this *AideController) SysHelpCate(ctx *gin.Context, pageNo int, pageSize int, cate SysHelpClassification.SysHelpClassification) (result *MessageResult.MessageResult) {
	var resultStr = this.RedisUtil.Get(SysConstant.SYS_HELP_CATE + strconv.Itoa(cate.GetOrdinal()))
	if resultStr != "" {
		var jobj fastjson.JSONObject
		err := json.Unmarshal([]byte(resultStr), &jobj)
		if err == nil {
			return this.SuccessWithData(jobj)
		}
	}
	var jsonObject = fastjson.NewJSONObject()
	var sysHelpPage, err = this.SysHelpService.FindByCondition(pageNo, pageSize, cate)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	jsonObject.Put("content", sysHelpPage.GetContent())
	jsonObject.Put("totalPage", sysHelpPage.GetTotalPages())
	jsonObject.Put("totalElements", sysHelpPage.GetTotalElements())
	this.RedisUtil.SetExpireKV(SysConstant.SYS_HELP_CATE+strconv.Itoa(cate.GetOrdinal()), jsonObject.ToJSONString(), SysConstant.SYS_HELP_CATE_EXPIRE_TIME)
	return this.SuccessWithData(jsonObject)

}
func (this *AideController) SysHelpTop(ctx *gin.Context, cate string) (result *MessageResult.MessageResult) {
	var resultStr = this.RedisUtil.Get(SysConstant.SYS_HELP_TOP + cate)
	if resultStr != "" {
		var jobj arraylist.List[entity.SysHelp]
		err := json.Unmarshal([]byte(resultStr), &jobj)
		if err == nil {
			return this.SuccessWithData(jobj)
		}
	}
	var sysHelps, err = this.SysHelpService.GetgetCateTops(cate)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	this.RedisUtil.SetExpireKV(SysConstant.SYS_HELP_TOP+cate, sysHelps, SysConstant.SYS_HELP_TOP_EXPIRE_TIME)
	return this.SuccessWithData(sysHelps)

}
func (this *AideController) SysHelpDetail(ctx *gin.Context, id int64) (result *MessageResult.MessageResult) {
	var resultStr = this.RedisUtil.Get(SysConstant.SYS_HELP_DETAIL + strconv.FormatInt(id, 10))
	if resultStr != "" {
		var jobj entity.SysHelp
		err := json.Unmarshal([]byte(resultStr), &jobj)
		if err == nil {
			return this.SuccessWithData(jobj)
		}

	}
	var sysHelp, err = this.SysHelpService.FindById(id)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	this.RedisUtil.SetExpireKV(SysConstant.SYS_HELP_DETAIL+strconv.FormatInt(id, 10), sysHelp, SysConstant.SYS_HELP_DETAIL_EXPIRE_TIME)
	return this.SuccessWithData(sysHelp)

}
func NewAideController(websiteInformationService *service.WebsiteInformationService, sysAdvertiseService *service.SysAdvertiseService, sysHelpService *service.SysHelpService, appRevisionService *service.AppRevisionService, msService *service.LocaleMessageSourceService, redisUtil *util.RedisUtil) (ret *AideController) {
	ret = new(AideController)
	ret.WebsiteInformationService = websiteInformationService
	ret.SysAdvertiseService = sysAdvertiseService
	ret.SysHelpService = sysHelpService
	ret.AppRevisionService = appRevisionService
	ret.MsService = msService
	ret.RedisUtil = redisUtil
	return
}

type AideController struct {
	WebsiteInformationService *service.WebsiteInformationService
	SysAdvertiseService       *service.SysAdvertiseService
	SysHelpService            *service.SysHelpService
	AppRevisionService        *service.AppRevisionService
	MsService                 *service.LocaleMessageSourceService
	RedisUtil                 *util.RedisUtil
	controller.BaseController
}
