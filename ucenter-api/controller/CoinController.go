package controller

import (
	"bitrade/core/entity"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
	"github.com/qauzy/util/sets/hashset"
)

func (this *CoinController) Legal() (result *util.MessageResult) {
	var legalAll *arraylist.List[entity.Coin] = this.CoinService.FindLegalAll()
	return success(legalAll)
}
func (this *CoinController) FindLegalCoinPage(pageModel *PageModel.PageModel) (result *util.MessageResult) {
	var all *domain.Page = this.CoinService.FindLegalCoinPage(pageModel)
	return success(all)
}
func (this *CoinController) FindCoins() (result *arraylist.List[*hashmap.Map[string, string]]) {
	var coins *arraylist.List[entity.Coin] = this.CoinService.FindAll()
	var result = arraylist.New[*hashmap.Map[string, string]]()
	coins.ForEach(func(coin interface {
	}) {
		if coin.GetHasLegal() == false {
			var oMap = hashset.New[string]()
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
	BaseController
}
