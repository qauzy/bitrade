package LegalWalletScreen

import "bitrade/core/constant/LegalWalletState"

func NewLegalWalletScreen(state *LegalWalletState.LegalWalletState, coinName string) (ret *LegalWalletScreen) {
	ret = new(LegalWalletScreen)
	ret.State = state
	ret.CoinName = coinName
	return
}

type LegalWalletScreen struct {
	State    *LegalWalletState.LegalWalletState
	CoinName string
}
