package CoinExchange

import (
	"bitrade/core/entity/transform"
	"bitrade/core/util"
	"bitrade/ucenter-api/service"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
	"github.com/shopspring/decimal"
)

func (this *CoinExchangeController) FindSupportedCoin(ctx *gin.Context, unit string) (result arraylist.List[entity.AssetExchangeCoin]) {
	return this.AssetExchangeService.FindAllByFromCoin(unit)
}
func (this *CoinExchangeController) Exchange(ctx *gin.Context, member *transform.AuthMember, from string, to string, amount *math.BigDecimal) (result *util.MessageResult) {
	var coin, err = this.AssetExchangeService.FindOne(from, to)
	if err != nil {
		return util.Error(err.Error())
	}
	if coin == nil {
		return util.NewMessageResultV2(500, "不支持该币种兑换")
	}
	if amount.CompareTo(decimal.Zero) <= 0 {
		return util.NewMessageResultV2(500, "数量需大于0")
	}
	return this.AssetExchangeService.Exchange(member.GetId(), coin, amount)
}
func NewCoinExchangeController(assetExchangeService *service.AssetExchangeService) (ret *CoinExchangeController) {
	ret = new(CoinExchangeController)
	ret.AssetExchangeService = assetExchangeService
	return
}

type CoinExchangeController struct {
	AssetExchangeService *service.AssetExchangeService
}
