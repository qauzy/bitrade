package Coin

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/service"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/maps/hashmap"
)

func (this *CoinController) Legal(ctx *gin.Context) (result *MessageResult.MessageResult) {
	var legalAll arraylist.List[entity.Coin] = this.CoinService.FindLegalAll()
	return this.SuccessWithData(legalAll)
}
func (this *CoinController) FindLegalCoinPage(ctx *gin.Context, pageModel *PageModel.PageModel) (result *MessageResult.MessageResult) {
	var all **pagination.PageResult = this.CoinService.FindLegalCoinPage(pageModel)
	return this.SuccessWithData(all)
}
func (this *CoinController) FindCoins(ctx *gin.Context) (result arraylist.List[*hashmap.Map[string, string]]) {
	var coins, err = this.CoinService.FindAll()
	var result = arraylist.New[*hashmap.Map[string, string]]()
	coins.ForEach(func(coin *entity.Coin) {
		if coin.GetHasLegal() == false {
			var oMap = hashmap.New[string, string]()
			oMap.Put("name", coin.GetName())
			oMap.Put("nameCn", coin.GetNameCn())
			oMap.Put("withdrawFee", String.ValueOf(coin.GetMinTxFee()))
			oMap.Put("enableRecharge", String.ValueOf(coin.GetCanRecharge().GetOrdinal()))
			oMap.Put("minWithdrawAmount", String.ValueOf(coin.GetMinWithdrawAmount()))
			oMap.Put("enableWithdraw", String.ValueOf(coin.GetCanWithdraw().GetOrdinal()))
			oMap.Put("maxWithdrawAmount", coin.GetMaxWithdrawAmount().ToPlainString())
			oMap.Put("withdrawThreshold", coin.GetWithdrawThreshold().ToPlainString())
			result.Add(oMap)
		}
	})
	return result
}

type CoinController struct {
	CoinService *service.CoinService
	controller.BaseController
}
