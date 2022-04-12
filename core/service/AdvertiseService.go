package service

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/PriceType"
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/maps/hashmap"
	"github.com/qauzy/math"
	"github.com/shopspring/decimal"
)

func (this *AdvertiseService) QueryWhereOrPage(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.Advertise], err error) {
	var list arraylist.List[entity.Advertise]
	var jpaQuery = queryFactory.SelectFrom(advertise)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *AdvertiseService) FindOne(id int64) (result *entity.Advertise, err error) {
	return this.AdvertiseDao.FindById(id).Get()
}
func (this *AdvertiseService) FindOne(id int64, memberId int64) (result *transform.MemberAdvertiseDetail, err error) {
	var advertise = this.AdvertiseDao.FindAdvertiseByIdAndMemberIdAndStatusNot(id, memberId, AdvertiseControlStatus.TURNOFF)
	if advertise != nil {
		return MemberAdvertiseDetail.ToMemberAdvertiseDetail(advertise)
	} else {
		return nil
	}
}
func (this *AdvertiseService) Find(id int64, memberId int64) (result *entity.Advertise, err error) {
	return this.AdvertiseDao.FindByIdAndMemberId(id, memberId)
}
func (this *AdvertiseService) SaveAdvertise(advertise *entity.Advertise) (result *entity.Advertise, err error) {
	return this.AdvertiseDao.Save(advertise)
}
func (this *AdvertiseService) TurnOffBatch(status *AdvertiseControlStatus.AdvertiseControlStatus, ids []int64) (result int, err error) {
	for _, advertiseId := range ids {
		var advertise = this.FindOne(advertiseId)
		if advertise == nil {
			return 0
		}
		if !advertise.GetStatus().Equals(AdvertiseControlStatus.PUT_ON_SHELVES) {
			return 0
		}
		if advertise.GetDealAmount() != nil && advertise.GetDealAmount().CompareTo(decimal.Zero) > 0 {
			return -100
		}
		var otcCoin = advertise.GetCoin()
		if advertise.GetAdvertiseType().Equals(AdvertiseType.SELL) {
			var memberWallet = this.OtcWalletService.FindByOtcCoinAndMemberId(advertise.GetMember().GetId(), otcCoin)
			var result = this.OtcWalletService.ThawBalance(memberWallet, advertise.GetRemainAmount())
			if result.GetCode() != 0 {
				return 0
			}
		}
		var ret = this.PutOffShelves(advertise.GetId(), advertise.GetRemainAmount())
		if !(ret > 0) {
			return 0
		}
	}
	return 1
}
func (this *AdvertiseService) ModifyAdvertise(advertise *entity.Advertise, old *entity.Advertise) (result *entity.Advertise, err error) {
	if advertise.GetPriceType() == PriceType.MUTATIVE {
		//变化的
		old.SetPriceType(PriceType.MUTATIVE)
		old.SetPremiseRate(advertise.GetPremiseRate())
	} else {
		//固定的
		old.SetPriceType(PriceType.REGULAR)
		old.SetPrice(advertise.GetPrice())
	}
	if advertise.GetAuto().IsIs() {
		old.SetAuto(BooleanEnum.IS_TRUE)
		old.SetAutoword(advertise.GetAutoword())
	} else {
		old.SetAuto(BooleanEnum.IS_FALSE)
	}
	old.SetMinLimit(advertise.GetMinLimit())
	old.SetMaxLimit(advertise.GetMaxLimit())
	old.SetTimeLimit(advertise.GetTimeLimit())
	old.SetRemark(advertise.GetRemark())
	old.SetPayMode(advertise.GetPayMode())
	old.SetNumber(advertise.GetNumber())
	old.SetRemainAmount(advertise.GetNumber())
	//变更为下架状态
	old.SetStatus(AdvertiseControlStatus.PUT_OFF_SHELVES)
	return this.AdvertiseDao.Save(old)
}
func (this *AdvertiseService) GetAllAdvertiseByMemberId(memberId int64, sort *domain.Sort) (result arraylist.List[transform.MemberAdvertise], err error) {
	var list = this.AdvertiseDao.FindAllByMemberIdAndStatusNot(memberId, AdvertiseControlStatus.TURNOFF, sort)
	return list.Stream().Map(func(x interface {
	}) {
		MemberAdvertise.ToMemberAdvertise(x)
	}).Collect(Collectors.ToList())
}
func (this *AdvertiseService) GetAllExcellentAdvertise(oType *AdvertiseType.AdvertiseType, list arraylist.List[*hashmap.Map[string, string]]) (result arraylist.List[transform.ScanAdvertise], err error) {
	var excellents = arraylist.New[transform.ScanAdvertise]()
	var sql = "SELECT\n" + "\td.*\n" + "FROM\n" + "\t(\n" + "\t\tSELECT\n" + "\t\t\tc.coin_id,\n" + "\t\t\t(\n" + "\t\t\t\tCASE\n" + "\t\t\t\tWHEN c.price_type = 0\n" + "\t\t\t\tAND c.price = b.minPrice THEN\n" + "\t\t\t\t\tc.id\n" + "\t\t\t\tWHEN c.price_type = 1\n" + "\t\t\t\tAND round(((c.premise_rate + 100) / 100 * ?),2) = b.minPrice THEN\n" + "\t\t\t\t\tc.id\n" + "\t\t\t\tEND\n" + "\t\t\t) advertise_id,\n" + "\t\t\tb.minPrice\n" + "\t\tFROM\n" + "\t\t\tadvertise c\n" + "\t\tJOIN (\n" + "\t\t\tSELECT\n" + "\t\t\t\ta.coin_id,\n" + func() {
		if oType.Equals(AdvertiseType.SELL) {
			return "\t\t\t\tmin(\n"
		} else {
			return "\t\t\t\tmax(\n"
		}
	}() + "\t\t\t\t\tCASE a.price_type\n" + "\t\t\t\t\tWHEN 0 THEN\n" + "\t\t\t\t\t\ta.price\n" + "\t\t\t\t\tELSE\n" + "\t\t\t\t\t\tround(((a.premise_rate + 100) / 100 * ?),2)\n" + "\t\t\t\t\tEND\n" + "\t\t\t\t) minPrice,\n" + "\t\t\t\ta.advertise_type,\n" + "\t\t\t\ta.`status`\n" + "\t\t\tFROM\n" + "\t\t\t\tadvertise a\n" + "\t\t\tWHERE\n" + "\t\t\t\ta. STATUS = 0\n" + "\t\t\tAND a.advertise_type = ?\n" + "\t\t\tGROUP BY\n" + "\t\t\t\ta.coin_id\n" + "\t\t) b ON c.coin_id = b.coin_id\n" + "\t\tAND c.advertise_type = b.advertise_type\n" + "\t\tAND c.`status` = b. STATUS\n" + "\t\tAND c.coin_id = ?\n" + "\t) d\n" + "WHERE\n" + "\td.advertise_id IS NOT NULL\n" + "GROUP BY\n" + "\td.coin_id"
	list.ParallelStream().ForEachOrdered(func(x *hashmap.Map[string, string]) {
		var otcCoin = this.OtcCoinDao.FindOtcCoinByUnitAndStatus(x.Get("name"), CommonStatus.NORMAL)
		if otcCoin != nil {
			exception := func() (err error) {
				var mapList = DB.Query(sql, x.Get("price"), x.Get("price"), oType.Ordinal(), otcCoin.GetId())
				if mapList.Size() > 0 {
					var advertise = this.AdvertiseDao.FindById(Long.ValueOf(mapList.Get(0).Get("advertise_id"))).Get()
					var member = advertise.GetMember()
					excellents.Add(new(ScanAdvertise).SetAdvertiseId(advertise.GetId()).SetCoinId(otcCoin.GetId()).SetCoinName(otcCoin.GetName()).SetCoinNameCn(otcCoin.GetNameCn()).SetCreateTime(advertise.GetCreateTime()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetMemberName(member.GetUsername()).SetAvatar(member.GetAvatar()).SetLevel(member.GetMemberLevel().Ordinal()).SetPayMode(advertise.GetPayMode()).SetUnit(otcCoin.GetUnit()).SetRemainAmount(advertise.GetRemainAmount()).SetTransactions(member.GetTransactions()).SetPrice(BigDecimalUtils.Round(float64.ValueOf(mapList.Get(0).Get("minPrice")), 2)))
				}
				return
			}()
			if exception != nil {
				e.PrintStackTrace()
			}
		}
	})
	return excellents
}
func (this *AdvertiseService) GetAllExcellentAdvertise(oType *AdvertiseType.AdvertiseType, list arraylist.List[*hashmap.Map[string, string]], country *entity.Country) (result arraylist.List[transform.ScanAdvertise], err error) {
	var excellents = arraylist.New[transform.ScanAdvertise]()
	var sql = "SELECT\n" + "\td.*\n" + "FROM\n" + "\t(\n" + "\t\tSELECT\n" + "\t\t\tc.coin_id,\n" + "\t\t\t(\n" + "\t\t\t\tCASE\n" + "\t\t\t\tWHEN c.price_type = 0\n" + "\t\t\t\tAND c.price = b.minPrice THEN\n" + "\t\t\t\t\tc.id\n" + "\t\t\t\tWHEN c.price_type = 1\n" + "\t\t\t\tAND round(((c.premise_rate + 100) / 100 * ?),2) = b.minPrice THEN\n" + "\t\t\t\t\tc.id\n" + "\t\t\t\tEND\n" + "\t\t\t) advertise_id,\n" + "\t\t\tb.minPrice\n" + "\t\tFROM\n" + "\t\t\tadvertise c\n" + "\t\tJOIN (\n" + "\t\t\tSELECT\n" + "\t\t\t\ta.coin_id,\n" + func() {
		if oType.Equals(AdvertiseType.SELL) {
			return "\t\t\t\tmin(\n"
		} else {
			return "\t\t\t\tmax(\n"
		}
	}() + "\t\t\t\t\tCASE a.price_type\n" + "\t\t\t\t\tWHEN 0 THEN\n" + "\t\t\t\t\t\ta.price\n" + "\t\t\t\t\tELSE\n" + "\t\t\t\t\t\tround(((a.premise_rate + 100) / 100 * ?),2)\n" + "\t\t\t\t\tEND\n" + "\t\t\t\t) minPrice,\n" + "\t\t\t\ta.advertise_type,\n" + "\t\t\t\ta.`status`\n" + "\t\t\tFROM\n" + "\t\t\t\tadvertise a\n" + "\t\t\tWHERE\n" + "\t\t\t\ta. STATUS = 0\n" + "\t\t\tAND a.advertise_type = ?\n" + "\t\t\tGROUP BY\n" + "\t\t\t\ta.coin_id\n" + "\t\t) b ON c.coin_id = b.coin_id\n" + "\t\tAND c.advertise_type = b.advertise_type\n" + "\t\tAND c.`status` = b. STATUS\n" + "\t\tAND c.country = ?\n" + "\t\tAND c.coin_id = ?\n" + "\t) d\n" + "WHERE\n" + "\td.advertise_id IS NOT NULL\n" + "GROUP BY\n" + "\td.coin_id"
	list.ParallelStream().ForEachOrdered(func(x *hashmap.Map[string, string]) {
		var otcCoin = this.OtcCoinDao.FindOtcCoinByUnitAndStatus(x.Get("name"), CommonStatus.NORMAL)
		if otcCoin != nil {
			exception := func() (err error) {
				var mapList = DB.Query(sql, x.Get("price"), x.Get("price"), oType.Ordinal(), country.GetZhName(), otcCoin.GetId())
				if mapList.Size() > 0 {
					var advertise = this.AdvertiseDao.FindById(Long.ValueOf(mapList.Get(0).Get("advertise_id"))).Get()
					var member = advertise.GetMember()
					excellents.Add(new(ScanAdvertise).SetAdvertiseId(advertise.GetId()).SetCoinId(otcCoin.GetId()).SetCoinName(otcCoin.GetName()).SetCoinNameCn(otcCoin.GetNameCn()).SetCreateTime(advertise.GetCreateTime()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetMemberName(member.GetUsername()).SetAvatar(member.GetAvatar()).SetLevel(member.GetMemberLevel().Ordinal()).SetPayMode(advertise.GetPayMode()).SetUnit(otcCoin.GetUnit()).SetRemainAmount(advertise.GetRemainAmount()).SetTransactions(member.GetTransactions()).SetPrice(BigDecimalUtils.Round(float64.ValueOf(mapList.Get(0).Get("minPrice")), 2)))
				}
				return
			}()
			if exception != nil {
				e.PrintStackTrace()
			}
		}
	})
	return excellents
}
func (this *AdvertiseService) PaginationAdvertise(pageNo int, pageSize int, otcCoin *entity.OtcCoin, advertiseType *AdvertiseType.AdvertiseType, marketPrice float64, isCertified int) (result transform.SpecialPage[transform.ScanAdvertise], err error) {
	var specialPage = new(transform.SpecialPage)
	var sql = "SELECT\n" + "\ta.*, (\n" + "\t\tCASE a.price_type\n" + "\t\tWHEN 0 THEN\n" + "\t\t\ta.price\n" + "\t\tELSE\n" + "\t\t\tround(((a.premise_rate + 100) / 100 * ?),2)\n" + "\t\tEND\n" + "\t) finalPrice,\n" + "\tb.avatar,\n" + "\tb.username,\n" + "\tb.member_level,\n" + "\tb.transactions\n" + "FROM\n" + "\tadvertise a\n" + "JOIN member b ON a.member_id = b.id\n" + func() {
		if isCertified == 1 {
			return "AND b.member_level = 2\n"
		} else {
			return " "
		}
	}() + "AND a.coin_id = ?\n" + "AND a.advertise_type = ?\n" + "AND a.`status` = 0\n" + "ORDER BY\n" + func() {
		if advertiseType.Equals(AdvertiseType.SELL) {
			return "\tfinalPrice,\n"
		} else {
			return "\tfinalPrice desc,\n"
		}
	}() + "\ta.id\n" + "LIMIT ?,\n" + " ?"
	var list = DB.Query(sql, marketPrice, otcCoin.GetId(), advertiseType.Ordinal(), (pageNo-1)*pageSize, pageSize)
	if list != nil && list.Size() > 0 {
		var sql1 = "SELECT\n" + "\tCOUNT(a.id) total\n" + "FROM\n" + "\tadvertise a\n" + "JOIN member b ON a.member_id = b.id\n" + func() {
			if isCertified == 1 {
				return "AND b.member_level = 2\n"
			} else {
				return " "
			}
		}() + "AND a.coin_id = ?\n" + "AND a.advertise_type = ?\n" + "AND a.`status` = 0"
		var list1 = DB.Query(sql1, otcCoin.GetId(), advertiseType.Ordinal())
		var oMap = list1.Get(0)
		var total = Integer.ValueOf(oMap.Get("total"))
		specialPage.SetTotalElement(total)
		specialPage.SetTotalPage(func() {
			if total%pageSize == 0 {
				return total / pageSize
			} else {
				return total/pageSize + 1
			}
		}())
		specialPage.SetContext(list.Stream().Map(func(x *hashmap.Map[string, string]) {
			new(ScanAdvertise).SetPrice(BigDecimalUtils.Round(float64.ValueOf(x.Get("finalPrice")), 2)).SetTransactions(Integer.ParseInt(x.Get("transactions"))).SetRemainAmount(BigDecimal.ValueOf(float64.ValueOf(x.Get("remain_amount")))).SetUnit(otcCoin.GetUnit()).SetPayMode(x.Get("pay_mode")).SetMemberName(x.Get("username")).SetAvatar(x.Get("avatar")).SetMinLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("min_limit")))).SetMaxLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("max_limit")))).SetCoinNameCn(otcCoin.GetNameCn()).SetLevel(Integer.ParseInt(x.Get("member_level"))).SetCoinId(otcCoin.GetId()).SetCoinName(otcCoin.GetName()).SetAdvertiseId(Long.ValueOf(x.Get("id"))).SetCreateTime(DateUtil.StrToDate(x.Get("create_time"))).SetAdvertiseType(advertiseType)
		}).Collect(Collectors.ToList()))
	} else {
		specialPage.SetTotalPage(1)
		specialPage.SetTotalElement(0)
	}
	specialPage.SetCurrentPage(pageNo)
	specialPage.SetPageNumber(pageSize)
	return specialPage
}
func (this *AdvertiseService) PaginationAdvertise(pageNo int, pageSize int, otcCoin *entity.OtcCoin, advertiseType *AdvertiseType.AdvertiseType, marketPrice float64, isCertified int, country *entity.Country, paymodes []string) (result transform.SpecialPage[transform.ScanAdvertise], err error) {
	var specialPage = new(transform.SpecialPage)
	var payConditon = StringUtils.NewStringBuilder(" ")
	if paymodes != nil {
		payConditon.Append("AND ( 1=2 ")
		for _, paymode := range paymodes {
			payConditon.Append("  or a.pay_mode like '%" + paymode + "%' ")
		}
		payConditon.Append(" ) ")
	}
	var sql = "SELECT\n" + "\ta.*, (\n" + "\t\tCASE a.price_type\n" + "\t\tWHEN 0 THEN\n" + "\t\t\ta.price\n" + "\t\tELSE\n" + "\t\t\tround(((a.premise_rate + 100) / 100 * ?),2)\n" + "\t\tEND\n" + "\t) finalPrice,\n" + "\tb.avatar,\n" + "\tb.username,\n" + "\tb.member_level,\n" + "\tb.transactions\n" + "FROM\n" + "\tadvertise a\n" + "JOIN member b ON a.member_id = b.id\n" + func() {
		if isCertified == 1 {
			return "AND b.member_level = 2\n"
		} else {
			return " "
		}
	}() + "AND a.coin_id = ?\n" + "AND a.advertise_type = ?\n" + "AND a.country = ?\n" + payConditon.ToString() + "AND a.`status` = 0\n" + "ORDER BY\n" + func() {
		if advertiseType.Equals(AdvertiseType.SELL) {
			return "\tfinalPrice,\n"
		} else {
			return "\tfinalPrice desc,\n"
		}
	}() + "\ta.id\n" + "LIMIT ?,\n" + " ?"
	var list = DB.Query(sql, marketPrice, otcCoin.GetId(), advertiseType.Ordinal(), country.GetZhName(), (pageNo-1)*pageSize, pageSize)
	if list.Size() > 0 {
		var sql1 = "SELECT\n" + "\tCOUNT(a.id) total\n" + "FROM\n" + "\tadvertise a\n" + "JOIN member b ON a.member_id = b.id\n" + func() {
			if isCertified == 1 {
				return "AND b.member_level = 2\n"
			} else {
				return " "
			}
		}() + "AND a.coin_id = ?\n" + payConditon.ToString() + "AND a.advertise_type = ?\n" + "AND a.country = ?\n" + "AND a.`status` = 0"
		var list1 = DB.Query(sql1, otcCoin.GetId(), advertiseType.Ordinal(), country.GetZhName())
		var oMap = list1.Get(0)
		var total = Integer.ValueOf(oMap.Get("total"))
		specialPage.SetTotalElement(total)
		specialPage.SetTotalPage(func() {
			if total%pageSize == 0 {
				return total / pageSize
			} else {
				return total/pageSize + 1
			}
		}())
		specialPage.SetContext(list.Stream().Map(func(x *hashmap.Map[string, string]) {
			new(ScanAdvertise).SetPrice(BigDecimalUtils.Round(float64.ValueOf(x.Get("finalPrice")), 2)).SetTransactions(Integer.ParseInt(x.Get("transactions"))).SetRemainAmount(BigDecimal.ValueOf(float64.ValueOf(x.Get("remain_amount")))).SetUnit(otcCoin.GetUnit()).SetPayMode(x.Get("pay_mode")).SetMemberName(x.Get("username")).SetAvatar(x.Get("avatar")).SetMinLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("min_limit")))).SetMaxLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("max_limit")))).SetCoinNameCn(otcCoin.GetNameCn()).SetLevel(Integer.ParseInt(x.Get("member_level"))).SetCoinId(otcCoin.GetId()).SetCoinName(otcCoin.GetName()).SetAdvertiseId(Long.ValueOf(x.Get("id"))).SetCreateTime(DateUtil.StrToDate(x.Get("create_time"))).SetAdvertiseType(advertiseType)
		}).Collect(Collectors.ToList()))
	} else {
		specialPage.SetTotalPage(1)
		specialPage.SetTotalElement(0)
	}
	specialPage.SetCurrentPage(pageNo)
	specialPage.SetPageNumber(pageSize)
	return specialPage
}
func (this *AdvertiseService) PaginationQuery(pageNo int, pageSize int, country string, payMode string, advertiseType *AdvertiseType.AdvertiseType, coin *entity.OtcCoin, marketPrice *math.BigDecimal) (result domain.Page[transform.ScanAdvertise], err error) {
	var order1 = Sort.Order(Sort.Direction.ASC, "price")
	var order2 = Sort.Order(Sort.Direction.DESC, "id")
	var sort = Sort.By(order1, order2)
	var pageRequest = PageRequest.Of(pageNo, pageSize, sort)
	var specification = func(root interface {
	}, criteriaQuery interface {
	}, criteriaBuilder interface {
	}) {
		var country1 = root.Get("country")
		var payMode1 = root.Get("payMode")
		var advertiseType1 = root.Get("advertiseType")
		var currency1 = root.Get("coin").Get("id")
		var status1 = root.Get("status")
		var predicate1 = criteriaBuilder.Like(payMode1, "%"+payMode+"%")
		var predicate2 = criteriaBuilder.Equal(country1, country)
		var predicate3 = criteriaBuilder.Equal(advertiseType1, advertiseType)
		var predicate4 = criteriaBuilder.Equal(currency1, coin.GetId())
		var predicate5 = criteriaBuilder.Equal(status1, CommonStatus.NORMAL)
		if country == nil && payMode == nil {
			return criteriaBuilder.And(predicate3, predicate4, predicate5)
		} else if country != nil && payMode == nil {
			return criteriaBuilder.And(predicate2, predicate3, predicate4, predicate5)
		} else if country == nil && payMode != nil {
			return criteriaBuilder.And(predicate1, predicate3, predicate4, predicate5)
		} else {
			return criteriaBuilder.And(predicate1, predicate2, predicate3, predicate4, predicate5)
		}
	}
	var page = this.AdvertiseDao.FindAll(specification, pageRequest)
	//todo:得到市场价
	//BigDecimal markerprice = BigDecimal.TEN;
	var page1 = page.Map(func(advertise entity.Advertise) {
		var member = advertise.GetMember()
		return new(ScanAdvertise).SetAdvertiseId(advertise.GetId()).SetCoinId(advertise.GetCoin().GetId()).SetCoinName(advertise.GetCoin().GetName()).SetCoinNameCn(advertise.GetCoin().GetNameCn()).SetCreateTime(advertise.GetCreateTime()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetMemberName(member.GetUsername()).SetPayMode(advertise.GetPayMode()).SetUnit(advertise.GetCoin().GetUnit()).SetRemainAmount(advertise.GetRemainAmount()).SetTransactions(member.GetTransactions()).SetPrice(func() {
			if advertise.GetPriceType().Equals(PriceType.REGULAR) {
				return advertise.GetPrice()
			} else {
				return marketPrice.Multiply(advertise.GetPremiseRate().Divide(BigDecimal(100)).Add(BigDecimal.ONE))
			}
		}())
	})
	return page1
}
func (this *AdvertiseService) GetMemberAdvertise(member *entity.Member, oMap *hashmap.Map[string, math.BigDecimal]) (result *transform.MemberAdvertiseInfo, err error) {
	var buy = this.AdvertiseDao.FindAllByMemberIdAndStatusAndAdvertiseType(member.GetId(), AdvertiseControlStatus.PUT_ON_SHELVES, AdvertiseType.BUY)
	var sell = this.AdvertiseDao.FindAllByMemberIdAndStatusAndAdvertiseType(member.GetId(), AdvertiseControlStatus.PUT_ON_SHELVES, AdvertiseType.SELL)
	return new(MemberAdvertiseInfo).SetCreateTime(member.GetRegistrationTime()).SetEmailVerified(func() {
		if StringUtils.IsEmpty(member.GetEmail()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetPhoneVerified(func() {
		if StringUtils.IsEmpty(member.GetMobilePhone()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetRealVerified(func() {
		if StringUtils.IsEmpty(member.GetRealName()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetTransactions(member.GetTransactions()).SetUsername(member.GetUsername()).SetAvatar(member.GetAvatar()).SetBuy(buy.Stream().Map(func(advertise interface {
	}) {
		var markerPrice = oMap.Get(advertise.GetCoin().GetUnit())
		var member1 = advertise.GetMember()
		return new(ScanAdvertise).SetAdvertiseId(advertise.GetId()).SetAdvertiseType(advertise.GetAdvertiseType()).SetCoinId(advertise.GetCoin().GetId()).SetCoinName(advertise.GetCoin().GetName()).SetCoinNameCn(advertise.GetCoin().GetNameCn()).SetCreateTime(advertise.GetCreateTime()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetMemberName(member1.GetUsername()).SetPayMode(advertise.GetPayMode()).SetUnit(advertise.GetCoin().GetUnit()).SetRemainAmount(advertise.GetRemainAmount()).SetTransactions(member1.GetTransactions()).SetLevel(member1.GetMemberLevel().GetOrdinal()).SetPrice(func() {
			if advertise.GetPriceType().Equals(PriceType.REGULAR) {
				return advertise.GetPrice()
			} else {
				return mulRound(markerPrice, rate(advertise.GetPremiseRate()), 2)
			}
		}())
	}).Collect(Collectors.ToList())).SetSell(sell.Stream().Map(func(advertise interface {
	}) {
		var markerPrice = oMap.Get(advertise.GetCoin().GetUnit())
		var member1 = advertise.GetMember()
		return new(ScanAdvertise).SetAdvertiseId(advertise.GetId()).SetAdvertiseType(advertise.GetAdvertiseType()).SetCoinId(advertise.GetCoin().GetId()).SetCoinName(advertise.GetCoin().GetName()).SetCoinNameCn(advertise.GetCoin().GetNameCn()).SetCreateTime(advertise.GetCreateTime()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetMemberName(member1.GetUsername()).SetPayMode(advertise.GetPayMode()).SetUnit(advertise.GetCoin().GetUnit()).SetRemainAmount(advertise.GetRemainAmount()).SetTransactions(member1.GetTransactions()).SetLevel(member1.GetMemberLevel().GetOrdinal()).SetPrice(func() {
			if advertise.GetPriceType().Equals(PriceType.REGULAR) {
				return advertise.GetPrice()
			} else {
				return mulRound(markerPrice, rate(advertise.GetPremiseRate()), 2)
			}
		}())
	}).Collect(Collectors.ToList()))
}
func (this *AdvertiseService) UpdateAdvertiseAmountForBuy(advertiseId int64, amount *math.BigDecimal) (result bool, err error) {
	var ret = this.AdvertiseDao.UpdateAdvertiseAmount(AdvertiseControlStatus.PUT_ON_SHELVES, advertiseId, amount)
	if ret > 0 {
		return true
	} else {
		return false
	}
}
func (this *AdvertiseService) UpdateAdvertiseAmountForCancel(advertiseId int64, amount *math.BigDecimal) (result bool, err error) {
	var ret = this.AdvertiseDao.UpdateAdvertiseDealAmount(advertiseId, amount)
	if ret > 0 {
		return true
	} else {
		return false
	}
}
func (this *AdvertiseService) UpdateAdvertiseAmountForRelease(advertiseId int64, amount *math.BigDecimal) (result bool, err error) {
	var ret = this.AdvertiseDao.UpdateAdvertiseDealAmount(advertiseId, amount)
	if ret > 0 {
		return true
	} else {
		return false
	}
}
func (this *AdvertiseService) SelectSellAutoOffShelves(coinId int64, marketPrice *math.BigDecimal, jyRate *math.BigDecimal) (result arraylist.List[*hashmap.Map[string, string]], err error) {
	var sql = "SELECT b.* FROM (SELECT\n" + "\ta.*, CAST(\n" + "\t\ta.min_limit / (\n" + "\t\t\tCASE a.price_type\n" + "\t\t\tWHEN 0 THEN\n" + "\t\t\t\ta.price\n" + "\t\t\tELSE\n" + "\t\t\t\tround(\n" + "\t\t\t\t\t(\n" + "\t\t\t\t\t\t(a.premise_rate + 100) / 100 * ?\n" + "\t\t\t\t\t),\n" + "\t\t\t\t\t2\n" + "\t\t\t\t)\n" + "\t\t\tEND\n" + "\t\t) AS DECIMAL (18, 8)\n" + "\t) minWithdrawAmount\n" + "FROM\n" + "\tadvertise a\n" + "WHERE\n" + "\ta.`status` = 0\n" + "AND a.advertise_type = 1\n" + "AND a.coin_id = ?) b WHERE b.remain_amount<ROUND(((? + 100) / 100 * b.minWithdrawAmount),8)"
	var list = DB.Query(sql, marketPrice, coinId, jyRate)
	return list
}
func (this *AdvertiseService) SelectBuyAutoOffShelves(coinId int64, marketPrice *math.BigDecimal) (result arraylist.List[*hashmap.Map[string, string]], err error) {
	var sql = "SELECT b.* FROM (SELECT\n" + "\ta.*, CAST(\n" + "\t\ta.min_limit / (\n" + "\t\t\tCASE a.price_type\n" + "\t\t\tWHEN 0 THEN\n" + "\t\t\t\ta.price\n" + "\t\t\tELSE\n" + "\t\t\t\tround(\n" + "\t\t\t\t\t(\n" + "\t\t\t\t\t\t(a.premise_rate + 100) / 100 * ?\n" + "\t\t\t\t\t),\n" + "\t\t\t\t\t2\n" + "\t\t\t\t)\n" + "\t\t\tEND\n" + "\t\t) AS DECIMAL (18, 8)\n" + "\t) minWithdrawAmount\n" + "FROM\n" + "\tadvertise a\n" + "WHERE\n" + "\ta.`status` = 0\n" + "AND a.advertise_type = 0\n" + "AND a.coin_id = ?) b WHERE b.remain_amount<b.minWithdrawAmount"
	var list = DB.Query(sql, marketPrice, coinId)
	return list
}
func (this *AdvertiseService) AutoPutOffShelves(oMap *hashmap.Map[string, string], otcCoin *entity.OtcCoin) (err error) {
	if oMap.Get("advertise_type").Equals(String.ValueOf(AdvertiseType.SELL.Ordinal())) {
		var memberWallet = this.OtcWalletService.FindByOtcCoinAndMemberId(Long.ValueOf(oMap.Get("member_id")), otcCoin)
		var result = this.OtcWalletService.ThawBalance(memberWallet, BigDecimal(oMap.Get("remain_amount")))
		if result.GetCode() != 0 {
			return InformationExpiredException("Information Expired")
		}
	}
	var is = this.PutOffShelves(Long.ValueOf(oMap.Get("id")), BigDecimal(oMap.Get("remain_amount")))
	if !(is > 0) {
		return InformationExpiredException("Information Expired")
	}
}
func (this *AdvertiseService) PutOffShelves(id int64, amount *math.BigDecimal) (result int, err error) {
	return this.AdvertiseDao.PutOffAdvertise(id, amount)
}
func (this *AdvertiseService) GetAllPutOnAdvertis(memberId int64) (result arraylist.List[entity.Advertise], err error) {
	return this.AdvertiseDao.FindAllByMemberIdAndStatus(memberId, AdvertiseControlStatus.PUT_ON_SHELVES)
}
func (this *AdvertiseService) FindAll(predicate predicate, pageable *domain.Pageable) (result domain.Page[entity.Advertise], err error) {
	return this.AdvertiseDao.FindAll(predicate, pageable)
}
func (this *AdvertiseService) GetLatestAdvertise() (result transform.Special[transform.ScanAdvertise], err error) {
	var special = new(transform.Special)
	var sql = "SELECT\n" + "\ta.*, \n" + "\tb.avatar,\n" + "\tb.username,\n" + "\tb.member_level,\n" + "\tb.transactions\n" + "FROM\n" + "\tadvertise a\n" + "JOIN member b ON a.member_id = b.id\n" + "AND a.`status` = 0\n" + "ORDER BY\n" + "\ta.id desc\n" + "LIMIT 0,\n" + " 10"
	var list = DB.Query(sql)
	if nil != list && list.Size() > 0 {
		special.SetContext(list.Stream().Map(func(x *hashmap.Map[string, string]) {
			new(ScanAdvertise).SetPremiseRate(func() {
				if Integer.ParseInt(x.Get("price_type")) == 0 {
					return nil
				} else {
					return BigDecimal.ValueOf(float64.ValueOf(x.Get("premise_rate")))
				}
			}()).SetPrice(BigDecimalUtils.Round(float64.ValueOf(x.Get("price")), 2)).SetTransactions(Integer.ParseInt(x.Get("transactions"))).SetRemainAmount(BigDecimal.ValueOf(float64.ValueOf(x.Get("remain_amount")))).SetUnit(x.Get("coin_unit")).SetPayMode(x.Get("pay_mode")).SetMemberName(x.Get("username")).SetAvatar(x.Get("avatar")).SetMinLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("min_limit")))).SetMaxLimit(BigDecimal.ValueOf(float64.ValueOf(x.Get("max_limit")))).SetLevel(Integer.ParseInt(x.Get("member_level"))).SetCoinId(Integer.ParseInt(x.Get("coin_id"))).SetAdvertiseId(Long.ValueOf(x.Get("id"))).SetCreateTime(DateUtil.StrToDate(x.Get("create_time"))).SetAdvType(Integer.ParseInt(x.Get("advertise_type")))
		}).Collect(Collectors.ToList()))
	}
	return special
}
func (this *AdvertiseService) CountByMemberAndStatus(member *entity.Member, status *AdvertiseControlStatus.AdvertiseControlStatus) (result int, err error) {
	return this.AdvertiseDao.CountAllByMemberAndStatus(member, status)
}
func NewAdvertiseService(advertiseDao *dao.AdvertiseDao, otcCoinDao *dao.OtcCoinDao, otcWalletService OtcWalletService) (ret *AdvertiseService) {
	ret = new(AdvertiseService)
	ret.AdvertiseDao = advertiseDao
	ret.OtcCoinDao = otcCoinDao
	ret.OtcWalletService = otcWalletService
	return
}

type AdvertiseService struct {
	AdvertiseDao     dao.AdvertiseDao
	OtcCoinDao       dao.OtcCoinDao
	OtcWalletService *OtcWalletService
	Base.BaseService
}
