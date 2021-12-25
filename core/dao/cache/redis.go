package cache

import (
	"bitrade/core/config"
	"github.com/go-redis/redis"
	"time"
)

// RedisUtil cache dao
const Nil = redis.Nil

type RedisUtil struct {
	redisPool *redis.Client
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

	redisPool := redis.NewClient(&redis.Options{
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

	_, err = redisPool.Ping().Result()
	if err != nil {
		panic(err)
	}
	redisD := &RedisUtil{
		redisPool: redisPool,
	}

	return redisD

}

// GetKey 根据key获取redis键值
func (redisDb *RedisUtil) Get(key string) (string, error) {
	value, err := redisDb.redisPool.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// SetNotExpireKV 设置不过期的 key
func (redisDb *RedisUtil) SetNotExpireKV(key, value string) error {
	err := redisDb.redisPool.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetExpireKV 设置过期的 key
func (redisDb *RedisUtil) SetExpireKV(key, value string, expire time.Duration) error {
	err := redisDb.redisPool.Set(key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetExpireKey 设置 key 过期
func (redisDb *RedisUtil) SetExpireKey(key string, expire time.Duration) error {
	err := redisDb.redisPool.Expire(key, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisDb *RedisUtil) SetNX(key string, value string, expire time.Duration) (bool, error) {
	return redisDb.redisPool.SetNX(key, value, expire).Result()
}

// DelKey 删除 cache 的key
func (redisDb *RedisUtil) DelKey(key string) error {
	return redisDb.redisPool.Del(key).Err()
}

// KeyExist 判断某一个key 是否存在
func (redisDb *RedisUtil) KeyExist(keys string) (bool, error) {
	count, err := redisDb.redisPool.Exists(keys).Result()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// HSet 设置 hash
func (redisDb *RedisUtil) HSet(key, field string, value interface{}) error {
	return redisDb.redisPool.HSet(key, field, value).Err()
}

// HMSet 批量存储 hash
func (redisDb *RedisUtil) HMSet(key string, fields map[string]interface{}) error {
	if len(fields) < 1 {
		return nil
	}

	err := redisDb.redisPool.HMSet(key, fields).Err()
	if err != nil {
		return err
	}

	return nil
}

// HGet 获取单个 hash
func (redisDb *RedisUtil) HGet(key, field string) (string, error) {
	return redisDb.redisPool.HGet(key, field).Result()
}

// HMGet 批量获取 hash
func (redisDb *RedisUtil) HMGet(key string, fields ...string) ([]interface{}, error) {
	res, err := redisDb.redisPool.HMGet(key, fields...).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

//// HGetAll 获取整个 hash, 公司不允许用 hgetall 这个命令
//func (redisDb *RedisUtil)HGetAll(key string) (map[string]string, error) {
//	res, err := redisDb.redisPool.HGetAll(key).Result()
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

// HDel 删除 hash key
func (redisDb *RedisUtil) HDel(key string, fields ...string) error {
	err := redisDb.redisPool.HDel(key, fields...).Err()
	if err != nil {
		return err
	}

	return nil
}

// RPush 在名称为key的list尾添加一个值为value的元素
func (redisDb *RedisUtil) RPush(key string, values ...interface{}) error {
	return redisDb.redisPool.RPush(key, values...).Err()
}

// LPush 在名称为key的list头添加一个值为value的 元素
func (redisDb *RedisUtil) LPush(key string, values ...interface{}) error {
	return redisDb.redisPool.LPush(key, values...).Err()
}

func (redisDb *RedisUtil) LPop(key string) (string, error) {
	return redisDb.redisPool.LPop(key).Result()
}

func (redisDb *RedisUtil) RPop(key string) (string, error) {
	return redisDb.redisPool.RPop(key).Result()
}

// LLen 返回名称为key的list的长度
func (redisDb *RedisUtil) LLen(key string) (int64, error) {
	return redisDb.redisPool.LLen(key).Result()
}

// LRange 返回名称为key的list中start至end之间的元素, start为0, end为-1 则是获取所有 list key
func (redisDb *RedisUtil) LRange(key string, start, end int64) ([]string, error) {
	return redisDb.redisPool.LRange(key, start, end).Result()
}

// LSet 给名称为key的list中index位置的元素赋值
func (redisDb *RedisUtil) LSet(key string, index int64, value interface{}) error {
	return redisDb.redisPool.LSet(key, index, value).Err()
}

// LRem 删除count个key的list中值为value的元素
func (redisDb *RedisUtil) LRem(key string, count int64, value interface{}) error {
	return redisDb.redisPool.LRem(key, count, value).Err()
}

// ZAdd zset 有序集合中增加一个成员
func (redisDb *RedisUtil) ZAdd(key, member string, score float64) error {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := redisDb.redisPool.ZAdd(key, z).Result()
	if err != nil {
		return err
	}

	return nil
}

// ZCount zset 有序集合中 min-max中的成员数量
func (redisDb *RedisUtil) ZCount(key, min, max string) (int64, error) {
	count, err := redisDb.redisPool.ZCount(key, min, max).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZCARD 获取 zset 中元素的数量
func (redisDb *RedisUtil) ZCARD(key string) (int64, error) {
	count, err := redisDb.redisPool.ZCard(key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZRange 通过索引区间返回有序集合成指定区间内的成员
func (redisDb *RedisUtil) ZRange(key string, start, stop int64) ([]string, error) {
	arr, err := redisDb.redisPool.ZRange(key, start, stop).Result()
	if err != nil {
		return []string{}, err
	}

	return arr, nil
}

// ZRangeByScore 通过索引区间返回有序集合成指定区间内的成员
func (redisDb *RedisUtil) ZRangeByScore(key string, min, max string) ([]string, error) {
	opt := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	arr, err := redisDb.redisPool.ZRangeByScore(key, opt).Result()
	if err != nil {
		return []string{}, err
	}

	return arr, nil
}

func (redisDb *RedisUtil) ZRem(key string, members ...string) error {
	return redisDb.redisPool.ZRem(key, members).Err()
}

func (redisDb *RedisUtil) SAdd(key string, members ...string) error {
	return redisDb.redisPool.SAdd(key, members).Err()
}

func (redisDb *RedisUtil) SPop(key string) (string, error) {
	return redisDb.redisPool.SPop(key).Result()
}

func (redisDb *RedisUtil) SMembers(key string) ([]string, error) {
	return redisDb.redisPool.SMembers(key).Result()
}

func (redisDb *RedisUtil) SIsMember(key string, members string) (bool, error) {
	return redisDb.redisPool.SIsMember(key, members).Result()
}

func (redisDb *RedisUtil) SRem(key string, members ...string) error {
	return redisDb.redisPool.SRem(key, members).Err()
}

func (redisDb *RedisUtil) SCard(key string) (int64, error) {
	count, err := redisDb.redisPool.SCard(key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (redisDb *RedisUtil) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return redisDb.redisPool.SScan(key, cursor, match, count).Result()
}

func (redisDb *RedisUtil) Subscribe(channel string) *redis.PubSub {
	return redisDb.redisPool.Subscribe(channel)
}

func (redisDb *RedisUtil) Close() {
	redisDb.redisPool.Close()
}
