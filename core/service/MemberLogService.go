package service

func NewMemberLogService() (ret *MemberLogService) {
	ret = new(MemberLogService)
	return
}

type MemberLogService struct {
}
