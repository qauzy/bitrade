package NESecretPair

import "github.com/gin-gonic/gin"

func NewNESecretPair(ctx *gin.Context, secretId string, secretKey string) (this *NESecretPair) {
	this = new(NESecretPair)
	this.SecretId = this.SecretId
	this.SecretKey = this.SecretKey
	return
}
func NewNESecretPair(secretId string, secretKey string) (ret *NESecretPair) {
	ret = new(NESecretPair)
	ret.SecretId = secretId
	ret.SecretKey = secretKey
	return
}

type NESecretPair struct {
	SecretId  string
	SecretKey string
}
