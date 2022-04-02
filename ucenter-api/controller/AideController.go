package controller

import (
	"bitrade/core/constant/Platform"
	"bitrade/core/constant/SysAdvertiseLocation"
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/constant/SysHelpClassification"
	"bitrade/core/entity"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/util/lists/arraylist"
	"time"
)

func (this *AideController) KeyWords() (result *util.MessageResult) {
	var websiteInformation *entity.WebsiteInformation = this.WebsiteInformationService.FetchOne()
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(websiteInformation)
	return result
}
func (this *AideController) SysAdvertise(sysAdvertiseLocation *SysAdvertiseLocation.SysAdvertiseLocation) (result *util.MessageResult) {
	var list *arraylist.List[entity.SysAdvertise] = this.SysAdvertiseService.FindAllNormal(sysAdvertiseLocation)
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(list)
	return result
}
func (this *AideController) SysHelp(sysHelpClassification *SysHelpClassification.SysHelpClassification) (result *util.MessageResult) {
	var list *arraylist.List[entity.SysHelp]
	if sysHelpClassification == nil {
		list = this.SysHelpService.FindAllByStatusNotAndSort()
	} else {
		list = this.SysHelpService.FindBySysHelpClassification(sysHelpClassification)
	}
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(list)
	return result
}
func (this *AideController) SysHelp(id int64) (result *util.MessageResult) {
	//List<SysHelp> list = sysHelpService.findBySysHelpClassification(sysHelpClassification);
	var SysHelp = this.SysHelpService.FindById(id).(entity.SysHelp)
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(this.SysHelp)
	return result
}
func (this *AideController) SysHelp(platform *Platform.Platform) (result *util.MessageResult) {
	var revision *entity.AppRevision = this.AppRevisionService.FindRecentVersion(platform)
	if revision != nil {
		var result *util.MessageResult = MessageResult.Success()
		result.SetData(revision)
		return result
	} else {
		return MessageResult.Error(this.MsService.GetMessage("NO_UPDATE"))
	}
}
func (this *AideController) SysAllHelp(total int) (result *util.MessageResult) {
	var result = this.RedisUtil.Get(SysConstant.SYS_HELP).(*arraylist.List[fastjson.JSONObject])
	if result != nil {
		return success(result)
	} else {
		//HELP("新手指南"),
		var jsonResult = arraylist.New[fastjson.JSONObject]()
		var sysHelpPage *domain.Page = this.SysHelpService.FindByCondition(1, total, SysHelpClassification.HELP)
		var jsonSysHelp *fastjson.JSONObject = new(fastjson.JSONObject)
		jsonSysHelp.Put("content", sysHelpPage.GetContent())
		jsonSysHelp.Put("title", "新手指南")
		jsonSysHelp.Put("cate", "0")
		jsonResult.Add(jsonSysHelp)
		//FAQ("常见问题"),
		var sysFaqPage *domain.Page = this.SysHelpService.FindByCondition(1, total, SysHelpClassification.FAQ)
		var jsonSysFaq *fastjson.JSONObject = new(fastjson.JSONObject)
		jsonSysFaq.Put("content", sysFaqPage.GetContent())
		jsonSysFaq.Put("title", "常见问题")
		jsonSysFaq.Put("cate", "1")
		jsonResult.Add(jsonSysFaq)
		this.RedisUtil.Set(SysConstant.SYS_HELP, jsonResult, SysConstant.SYS_HELP_EXPIRE_TIME, time.Second)
		return success(jsonResult)
	}
	//RECHARGE("充值指南"),
	//Page<SysHelp> sysRechangePage = sysHelpService.findByCondition(1,total,SysHelpClassification.HELP);
	//TRANSACTION("交易指南"),
	//Page<SysHelp> sysTransactonPage = sysHelpService.findByCondition(1,total,SysHelpClassification.HELP);
}
func (this *AideController) SysHelpCate(pageNo int, pageSize int, cate *SysHelpClassification.SysHelpClassification) (result *util.MessageResult) {
	var result = this.RedisUtil.Get(SysConstant.SYS_HELP_CATE + cate).(fastjson.JSONObject)
	if result != nil {
		return success(result)
	} else {
		var jsonObject *fastjson.JSONObject = new(fastjson.JSONObject)
		var sysHelpPage *domain.Page = this.SysHelpService.FindByCondition(pageNo, pageSize, cate)
		jsonObject.Put("content", sysHelpPage.GetContent())
		jsonObject.Put("totalPage", sysHelpPage.GetTotalPages())
		jsonObject.Put("totalElements", sysHelpPage.GetTotalElements())
		this.RedisUtil.Set(SysConstant.SYS_HELP_CATE+cate, jsonObject, SysConstant.SYS_HELP_CATE_EXPIRE_TIME, time.Second)
		return success(jsonObject)
	}
}
func (this *AideController) SysHelpTop(cate string) (result *util.MessageResult) {
	var result = this.RedisUtil.Get(SysConstant.SYS_HELP_TOP + cate).(*arraylist.List[entity.SysHelp])
	if result != nil && !result.IsEmpty() {
		return success(result)
	} else {
		var sysHelps *arraylist.List[entity.SysHelp] = this.SysHelpService.GetgetCateTops(cate)
		this.RedisUtil.Set(SysConstant.SYS_HELP_TOP+cate, sysHelps, SysConstant.SYS_HELP_TOP_EXPIRE_TIME, time.Second)
		return success(sysHelps)
	}
}
func (this *AideController) SysHelpDetail(id int64) (result *util.MessageResult) {
	var result = this.RedisUtil.Get(SysConstant.SYS_HELP_DETAIL + id).(entity.SysHelp)
	if result != nil {
		return success(result)
	} else {
		var SysHelp = this.SysHelpService.FindById(id).(entity.SysHelp)
		this.RedisUtil.Set(SysConstant.SYS_HELP_DETAIL+id, this.SysHelp, SysConstant.SYS_HELP_DETAIL_EXPIRE_TIME, time.Second)
		return success(this.SysHelp)
	}
}

type AideController struct {
	WebsiteInformationService *service.WebsiteInformationService
	SysAdvertiseService       *service.SysAdvertiseService
	SysHelpService            *service.SysHelpService
	AppRevisionService        *service.AppRevisionService
	MsService                 *service.LocaleMessageSourceService
	RedisUtil                 *util.RedisUtil
	BaseController
}
