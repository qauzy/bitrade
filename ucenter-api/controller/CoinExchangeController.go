package controller

import (
	"bitrade/core/entity/transform"
	"bitrade/core/util"
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/shopspring/decimal"
)

func (this *CoinExchangeController) FindSupportedCoin(unit string) (result *arraylist.List[entity.AssetExchangeCoin]) {
	return this.AssetExchangeService.FindAllByFromCoin(unit)
}
func (this *CoinExchangeController) Exchange(member *transform.AuthMember, from string, to string, amount *math.BigDecimal) (result *util.MessageResult) {
	var coin *entity.AssetExchangeCoin = this.AssetExchangeService.FindOne(from, to)
	if coin == nil {
		return MessageResult(500, "不支持该币种兑换")
	}
	if amount.CompareTo(decimal.Zero) <= 0 {
		return MessageResult(500, "数量需大于0")
	}
	return this.AssetExchangeService.Exchange(member.GetId(), coin, amount)
}

type CoinExchangeController struct {
	AssetExchangeService *service.AssetExchangeService
}
