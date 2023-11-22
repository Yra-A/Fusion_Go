package redis

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// getRandomTTL 设置随机的过期时间，防止缓存雪崩
func getRandomTTL() time.Duration {
	return time.Duration(60+rand.Intn(20)) * time.Minute
}

// add 将k和v以及分数添加到Redis zset
func add(c *redis.Client, k string, v int64, score float64) {
	tx := c.TxPipeline()
	tx.ZAdd(k, redis.Z{Score: score, Member: v})
	tx.Expire(k, getRandomTTL())
	tx.Exec()
}

// del 从zset中删除k和v
func del(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	tx.ZRem(k, v)
	tx.Expire(k, getRandomTTL())
	tx.Exec()
}

// check 查询 set k 是否存在
func check(c *redis.Client, k string) bool {
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

// exist 检查v是否在zset k中
func exist(c *redis.Client, k string, v int64) bool {
	// 将v从int64转换为字符串
	vStr := strconv.FormatInt(v, 10)

	// 使用ZScore来检查元素v是否在zset k中
	score, err := c.ZScore(k, vStr).Result()

	// 如果没有错误且score不是NaN，则元素存在
	if err == nil && !math.IsNaN(score) {
		// 更新元素的过期时间
		c.Expire(k, getRandomTTL())
		return true
	}

	// 如果err为redis.Nil或存在其他错误，返回false
	return false
}

// count 获取 set k 的元素数量
func count(c *redis.Client, k string) (sum int64, err error) {
	if sum, err = c.ZCard(k).Result(); err == nil {
		c.Expire(k, getRandomTTL())
		return sum, err
	}
	return sum, err
}

// get 获取zset k中的元素列表，每个元素为i64类型
func get(c *redis.Client, k string) (vt []int64) {
	v, _ := c.ZRange(k, 0, -1).Result()
	c.Expire(k, getRandomTTL())
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, v_i64)
	}
	return vt
}
