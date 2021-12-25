package BigDecimalUtils

import (
	"github.com/qauzy/math"
)

var DEFAULT_DIV_SCALE int32 = 8

/**
 * 采用 BigDecimal 的字符串构造器进行初始化。
 *
 * @param v double 值
 * @return BigDecimal 对象
 */
func CreateBigDecimal(v float64) (result math.BigDecimal) {
	return CreateBigDecimal(v)
}

/**
 * 提供精确的加法运算。
 *
 * @param v1 被加数
 * @param v2 加数
 * @return 两个参数的和
 */
func AddBB(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return v1.Add(v2)
}

/**
 * 提供精确的加法运算。
 *
 * @param v1 被加数
 * @param v2 加数
 * @return 两个参数的和
 */
func AddFF(v1 float64, v2 float64) (result math.BigDecimal) {
	var b1 math.BigDecimal = CreateBigDecimal(v1)
	var b2 math.BigDecimal = CreateBigDecimal(v2)
	return b1.Add(b2)
}

/**
 * 提供精确的加法运算。
 *
 * @param v1 被加数
 * @param v2 加数
 * @return 两个参数的和
 */
func AddBF(v1 math.BigDecimal, v2 float64) (result math.BigDecimal) {
	var b2 math.BigDecimal = CreateBigDecimal(v2)
	return v1.Add(b2)
}

/**
 * 提供精确的减法运算。
 *
 * @param v1 被减数
 * @param v2 减数
 * @return 两个参数的差
 */
func SubFF(v1 float64, v2 float64) (result math.BigDecimal) {
	var b1 math.BigDecimal = CreateBigDecimal(v1)
	var b2 math.BigDecimal = CreateBigDecimal(v2)
	return b1.Sub(b2)
}

/**
 * 提供精确的减法运算。
 *
 * @param v1 被减数
 * @param v2 减数
 * @return 两个参数的差
 */
func SubBF(v1 math.BigDecimal, v2 float64) (result math.BigDecimal) {
	var b2 math.BigDecimal = CreateBigDecimal(v2)
	return v1.Sub(b2)
}

/**
 * 提供精确的减法运算。
 *
 * @param v1 被减数
 * @param v2 减数
 * @return 两个参数的差
 */
func SubBB(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return v1.Sub(v2)
}

/**
 * 提供精确的小数位四舍五入处理。
 *
 * @param v     需要四舍五入的数字
 * @param scale 小数点后保留几位
 * @return 四舍五入后的结果
 */
func RoundF(v float64, scale int32) (result math.BigDecimal) {
	if scale < 0 {
		return
	}
	var b = CreateBigDecimal(v)
	return b.Round(scale)
}

/**
 * 提供精确的小数位四舍五入处理。
 *
 * @param v     需要四舍五入的数字
 * @param scale 小数点后保留几位
 * @return 四舍五入后的结果
 */
func RoundB(v math.BigDecimal, scale int32) (result math.BigDecimal) {
	if scale < 0 {
		return
	}
	return v.RoundUp(scale)
}

/**
 * 提供精确的乘法运算。
 *
 * @param v1 被乘数
 * @param v2 乘数
 * @return 两个参数的积
 */
func MulFF(v1 float64, v2 float64) (result math.BigDecimal) {
	var b1 = CreateBigDecimal(v1)
	var b2 = CreateBigDecimal(v2)
	return b1.Mul(b2)
}

/**
 * 提供精确的乘法运算。
 *
 * @param v1 被乘数
 * @param v2 乘数
 * @return 两个参数的积
 */
func MulBF(v1 math.BigDecimal, v2 float64) (result math.BigDecimal) {
	var b2 math.BigDecimal = CreateBigDecimal(v2)
	return v1.Mul(b2)
}

/**
 * 提供精确的乘法运算。
 *
 * @param v1 被乘数
 * @param v2 乘数
 * @return 两个参数的积
 */
func MulBB(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return v1.Mul(v2)
}
func MulDown(v1 math.BigDecimal, v2 math.BigDecimal, x int32) (result math.BigDecimal) {
	return v1.Mul(v2).RoundDown(x)
}

/**
 * 提供相对精确的乘法运算，四舍五入保留八位小数。
 *
 * @param v1 被乘数
 * @param v2 乘数
 * @return 两个参数的积
 */
func MulRoundBB(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return MulRoundBBRes(v1, v2, DEFAULT_DIV_SCALE)
}

/**
 * 提供相对精确的乘法运算，四舍五入保留v3位小数。
 *
 * @param v1 被乘数
 * @param v2 乘数
 * @param v3 保留位数
 * @return 两个参数的积
 */
func MulRoundBBRes(v1 math.BigDecimal, v2 math.BigDecimal, v3 int32) (result math.BigDecimal) {
	return v1.Mul(v2).Round(v3)
}

/**
 * 提供（相对）精确的除法运算。 当发生除不尽的情况时，由scale参数指定精度，以后的数字四舍五入。
 *
 * @param v1    被除数
 * @param v2    除数
 * @param scale 表示表示需要精确到小数点以后几位。
 * @return 两个参数的商
 */
func DivFFScale(v1 float64, v2 float64, scale int32) (result math.BigDecimal) {
	if scale < 0 {
		return
	}
	var b1 = CreateBigDecimal(v1)
	var b2 = CreateBigDecimal(v2)
	return b1.Div(b2).RoundUp(scale)
}

/**
 * 提供（相对）精确的除法运算。 当发生除不尽的情况时，由scale参数指定精度，以后的数字四舍五入。
 *
 * @param v1    被除数
 * @param v2    除数
 * @param scale 表示表示需要精确到小数点以后几位。
 * @return 两个参数的商
 */
func DivBFScale(v1 math.BigDecimal, v2 float64, scale int32) (result math.BigDecimal) {
	if scale < 0 {
		return
	}
	var b2 = CreateBigDecimal(v2)
	return v1.Div(b2).RoundUp(scale)
}

/**
 * 提供（相对）精确的除法运算。 当发生除不尽的情况时，由scale参数指定精度，以后的数字四舍五入。
 *
 * @param v1    被除数
 * @param v2    除数
 * @param scale 表示表示需要精确到小数点以后几位。
 * @return 两个参数的商
 */
func Div(v1 math.BigDecimal, v2 math.BigDecimal, scale int32) (result math.BigDecimal) {
	if scale < 0 {
		return
	}
	return v1.Div(v2).RoundUp(scale)
}

/**
 * 提供（相对）精确的除法运算。 当发生除不尽的情况时，默认保留八位，以后的数字四舍五入。
 *
 * @param v1 被除数
 * @param v2 除数
 * @return 两个参数的商
 */
func DivFF(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return v1.Div(v2).RoundUp(DEFAULT_DIV_SCALE)
}
func DivDown(v1 math.BigDecimal, v2 math.BigDecimal) (result math.BigDecimal) {
	return v1.Div(v2).RoundDown(DEFAULT_DIV_SCALE)
}

/**
 * 得到利率
 *
 * @param v1
 * @return
 */
func GetRate(v1 math.BigDecimal) (result math.BigDecimal) {
	var hundred, _ = math.NewFromString("100")
	return DivFF(v1, hundred)
}

/**
 * 得到倍数
 *
 * @param v1
 * @return
 */
func Rate(v1 math.BigDecimal) (result math.BigDecimal) {
	return AddBB(GetRate(v1), math.NewFromInt(1))
}

/**
 * 比较大小，v1>=v2返回true,否则返回false
 *
 * @param v1
 * @param v2
 * @return
 */
func Compare(v1 math.BigDecimal, v2 math.BigDecimal) (result bool) {
	if v1.Cmp(v2) >= 0 {
		return true
	} else {
		return false
	}
}

/**
 * 比较大小，(v1+v2)>=v2返回true,否则返回false
 *
 * @param v1
 * @param v2
 * @param v3
 * @return
 */
func CompareV1V2V3(v1 math.BigDecimal, v2 math.BigDecimal, v3 math.BigDecimal) (result bool) {
	return Compare(AddBB(v1, v2), v3)
}

/**
 * 判断两值是否相等
 *
 * @param v1
 * @param v2
 * @return
 */
func IsEqual(v1 math.BigDecimal, v2 math.BigDecimal) (result bool) {
	if v1.Cmp(v2) == 0 {
		return true
	} else {
		return false
	}
}

type BigDecimalUtils struct {
}
