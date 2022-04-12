package service

func (this *MemberApiKeyService) FindMemberApiKeyByApiKey(apiKey string) (result *entity.MemberApiKey, err error) {
	return this.ApiKeyDao.FindMemberApiKeyByApiKey(apiKey)
}
func (this *MemberApiKeyService) FindByMemberIdAndId(memberId int64, id int64) (result *entity.MemberApiKey, err error) {
	return this.ApiKeyDao.FindMemberApiKeyByMemberIdAndId(memberId, id)
}
func (this *MemberApiKeyService) Save(memberApiKey *entity.MemberApiKey) (result *entity.MemberApiKey, err error) {
	return this.ApiKeyDao.Save(memberApiKey)
}
func (this *MemberApiKeyService) Del(id int64) (err error) {
	this.ApiKeyDao.Del(id)
}
func (this *MemberApiKeyService) FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberApiKey], err error) {
	return this.ApiKeyDao.FindAllByMemberId(memberId)
}
func NewMemberApiKeyService(apiKeyDao *dao.MemberApiKeyDao) (ret *MemberApiKeyService) {
	ret = new(MemberApiKeyService)
	ret.ApiKeyDao = apiKeyDao
	return
}

type MemberApiKeyService struct {
	ApiKeyDao dao.MemberApiKeyDao
	Base.BaseService[entity.MemberApiKey]
}
