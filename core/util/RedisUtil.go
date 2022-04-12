package util

import (
	"bitrade/core/config"
	"github.com/go-redis/redis"
	"time"
)

const Nil = redis.Nil

type RedisUtil struct {
	RedisTemplate *redis.Client
}

// RedisConf cache write
type Config struct {
	RedisHost     string `ini:"Host"`
	RedisAuth     string `ini:"Auth"`
	RedisPoolSize int    `ini:"PoolSize"`
}

func NewRedisWrite() *RedisUtil {
	var conf Config
	err := config.GetConfig().Section("redis").MapTo(&conf)
	if err != nil {
		panic(err)
	}

	RedisTemplate := redis.NewClient(&redis.Options{
		Addr:         conf.RedisHost,
		Password:     conf.RedisAuth,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     conf.RedisPoolSize,
		PoolTimeout:  30 * time.Second,
		MaxRetries:   2,
		IdleTimeout:  5 * time.Minute,
	})

	_, err = RedisTemplate.Ping().Result()
	if err != nil {
		panic(err)
	}
	redisD := &RedisUtil{
		RedisTemplate: RedisTemplate,
	}

	return redisD

}

func (this *RedisUtil) Expire(key string, expiration int64) (result bool) {
	exception := this.RedisTemplate.Expire(key, time.Duration(expiration)).Err()
	if exception != nil {
		return false
	}
	return true
}
func (this *RedisUtil) GetExpire(key string) (result int64) {

	expiration, err := this.RedisTemplate.TTL(key).Result()
	if err != nil {
		return -1
	}
	return int64(expiration.Seconds())
}
func (this *RedisUtil) HasKey(key string) (result bool) {
	count, err := this.RedisTemplate.Exists(key).Result()
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
func (this *RedisUtil) Delete(key ...string) {
	this.RedisTemplate.Del(key...).Err()
	return
}
func (this *RedisUtil) Get(key string) (result string) {
	object, err := this.RedisTemplate.Get(key).Result()
	if err != nil {
		return ""
	}
	return object
}
func (this *RedisUtil) GetString(key string) (result string) {
	object, err := this.RedisTemplate.Get(key).Result()
	if err != nil {
		return ""
	}
	return object
}
func (this *RedisUtil) Set(key string, value interface{}) (result bool) {
	exception := this.RedisTemplate.Set(key, value, -1).Err()
	if exception != nil {
		return false
	}

	return true
}
func (this *RedisUtil) SetExpireKV(key string, value interface{}, t int64) (result bool) {
	exception := this.RedisTemplate.Set(key, value, time.Duration(t)*time.Second)
	if exception != nil {
		return false
	}
	return true
}

func (this *RedisUtil) Incr(key string, delta int64) (result int64) {
	//if delta < 0 {
	//	return RuntimeException("递增因子必须大于0")
	//}
	return this.RedisTemplate.IncrBy(key, delta).Val()
}
func (this *RedisUtil) Decr(key string, delta int64) (result int64) {
	//if delta < 0 {
	//	return RuntimeException("递减因子必须大于0")
	//}
	return this.RedisTemplate.DecrBy(key, -delta).Val()
}
func (this *RedisUtil) Hget(key string, item string) (result interface{}) {
	return this.RedisTemplate.HGet(key, item)
}
func (this *RedisUtil) Hmget(key string) (result map[string]string) {
	result, exception := this.RedisTemplate.HGetAll(key).Result()
	if exception != nil {
		return nil
	}
	return
}
func (this *RedisUtil) Hmset(key string, oMap map[string]interface{}) (result bool) {
	exception := this.RedisTemplate.HMSet(key, oMap).Err()
	if exception != nil {
		return false
	}
	return true
}
func (this *RedisUtil) HmsetExpire(key string, oMap map[string]interface{}, time int64) (result bool) {
	exception := this.RedisTemplate.HMSet(key, oMap).Err()
	if exception != nil {
		return false
	}
	if time > 0 {
		this.Expire(key, time)
	}
	return true
}
func (this *RedisUtil) Hset(key string, item string, value interface{}) (result bool) {
	exception := this.RedisTemplate.HSet(key, item, value)
	if exception != nil {
		return false
	}
	return true
}
func (this *RedisUtil) HsetExpire(key string, item string, value interface{}, time int64) (result bool) {
	exception := this.RedisTemplate.HSet(key, item, value)
	if exception != nil {
		return false
	}
	if time > 0 {
		this.Expire(key, time)
	}
	return true
}
func (this *RedisUtil) Hdel(key string, item string) {
	this.RedisTemplate.HDel(key, item)
}
func (this *RedisUtil) HHasKey(key string, item string) (result bool) {
	result, err := this.RedisTemplate.HExists(key, item).Result()
	if err != nil {
		return false
	}
	return
}
func (this *RedisUtil) Hincr(key string, item string, by float64) (result float64) {
	result, err := this.RedisTemplate.HIncrByFloat(key, item, by).Result()
	if err != nil {
		return 0
	}
	return
}
func (this *RedisUtil) Hde(key string, item string, by float64) (result float64) {
	result, err := this.RedisTemplate.HIncrByFloat(key, item, -by).Result()
	if err != nil {
		return 0
	}
	return
}
func (this *RedisUtil) SGet(key string) (result map[string]struct{}) {
	result, exception := this.RedisTemplate.SMembersMap(key).Result()
	if exception != nil {
		return nil
	}
	return
}
func (this *RedisUtil) SHasKey(key string, value interface{}) (result bool) {
	result, exception := this.RedisTemplate.SIsMember(key, value).Result()
	if exception != nil {
		return false
	}
	return result
}
func (this *RedisUtil) SSet(key string, values interface{}) (result int64) {
	var count, exception = this.RedisTemplate.SAdd(key, values).Result()
	if exception != nil {

		return 0
	}
	return count
}
func (this *RedisUtil) SSetAndTime(key string, time int64, values interface{}) (result int64) {
	var count, exception = this.RedisTemplate.SAdd(key, values).Result()
	if exception != nil {

		return 0
	}
	if time > 0 {
		this.Expire(key, time)
	}
	return count
}
func (this *RedisUtil) SGetSetSize(key string) (result int64) {
	result, exception := this.RedisTemplate.SCard(key).Result()
	if exception != nil {
		return 0
	}
	return result
}
func (this *RedisUtil) SetRemove(key string, values interface{}) (result int64) {
	var count, exception = this.RedisTemplate.SRem(key, values).Result()
	if exception != nil {
		return 0
	}
	return count
}
func (this *RedisUtil) LGet(key string, start int64, end int64) (result []string) {
	result, exception := this.RedisTemplate.LRange(key, start, end).Result()
	if exception != nil {
		return nil
	}
	return result
}
func (this *RedisUtil) LGetListSize(key string) (result int64) {
	result, exception := this.RedisTemplate.LLen(key).Result()
	if exception != nil {

		return 0
	}
	return result
}
func (this *RedisUtil) LGetIndex(key string, index int64) (result interface{}) {
	result, exception := this.RedisTemplate.LIndex(key, index).Result()
	if exception != nil {
		return nil
	}
	return result
}

func (this *RedisUtil) LSet(key string, value []interface{}) (result bool) {
	exception := this.RedisTemplate.RPush(key, value).Err()
	if exception != nil {

		return false
	}
	return true
}
func (this *RedisUtil) LSetExpire(key string, value []interface{}, time int64) (result bool) {
	exception := this.RedisTemplate.RPush(key, value).Err()
	if exception != nil {
		return false
	}
	if time > 0 {
		this.Expire(key, time)
	}
	return true
}
func (this *RedisUtil) LUpdateIndex(key string, index int64, value interface{}) (result bool) {

	exception := this.RedisTemplate.LSet(key, index, value).Err()
	if exception != nil {

		return false
	}
	return true
}
func (this *RedisUtil) LRemove(key string, count int64, value interface{}) (result int64) {
	var remove, exception = this.RedisTemplate.LRem(key, count, value).Result()
	if exception != nil {

		return 0
	}
	return remove
}
