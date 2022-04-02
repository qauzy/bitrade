package sdk

func NewNESecretPair(SecretId string, SecretKey string) (this *NESecretPair) {
	this = new(NESecretPair)
	this.SecretId = this.SecretId
	this.SecretKey = this.SecretKey
	return
}

type NESecretPair struct {
	SecretId  string
	SecretKey string
}
