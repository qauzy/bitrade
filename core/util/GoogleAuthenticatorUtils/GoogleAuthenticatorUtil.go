package GoogleAuthenticatorUtils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func NewGoogleAuth() *GoogleAuthenticatorUtil {
	return &GoogleAuthenticatorUtil{}
}

// 获取秘钥
func GenerateSecretKey() (result string) {
	g := NewGoogleAuth()
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.BigEndian, g.un())
	return strings.ToUpper(g.base32encode(g.hmacSha1(buf.Bytes(), nil)))
}
func GetQRBarcodeURL(user string, host string, secret string) (result string) {
	return fmt.Sprintf("otpauth://totp/%s@%s?secret=%s", user, host, secret)
}

func CheckCodes(code string, googleKey string) (result bool) {
	return NewGoogleAuth().Check_code(code, googleKey)
}

func (this *GoogleAuthenticatorUtil) Check_code(code string, secret string) (result bool) {
	g := NewGoogleAuth()
	secretUpper := strings.ToUpper(secret)
	secretKey, err := g.base32decode(secretUpper)
	if err != nil {
		return false
	}
	number := g.oneTimePassword(secretKey, g.toBytes(time.Now().Unix()/30))
	_code := fmt.Sprintf("%06d", number)
	return _code == code
}

func (this *GoogleAuthenticatorUtil) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (this *GoogleAuthenticatorUtil) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (this *GoogleAuthenticatorUtil) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (this *GoogleAuthenticatorUtil) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (this *GoogleAuthenticatorUtil) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (this *GoogleAuthenticatorUtil) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (this *GoogleAuthenticatorUtil) oneTimePassword(key []byte, data []byte) uint32 {
	hash := this.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := this.toUint32(hashParts)
	return number % 1000000
}

type GoogleAuthenticatorUtil struct {
}
