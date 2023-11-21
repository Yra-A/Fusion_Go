package redis

import (
    "math/rand"
    "strconv"
    "time"

    "github.com/go-redis/redis"
)

// getRandomTTL 设置随机的过期时间，防止缓存雪崩
func getRandomTTL() time.Duration {
    return time.Duration(60+rand.Intn(20)) * time.Minute
}

// add k & v to redis
func add(c *redis.Client, k string, v int64) {
    tx := c.TxPipeline()
    tx.SAdd(k, v)
    tx.Expire(k, getRandomTTL())
    tx.Exec()
}

// del k & v
func del(c *redis.Client, k string, v int64) {
    tx := c.TxPipeline()
    tx.SRem(k, v)
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

// exist 查询 set k 中是否有 v
func exist(c *redis.Client, k string, v int64) bool {
    if e, _ := c.SIsMember(k, v).Result(); e {
        c.Expire(k, getRandomTTL())
        return true
    }
    return false
}

// count 获取 set k 的元素数量
func count(c *redis.Client, k string) (sum int64, err error) {
    if sum, err = c.SCard(k).Result(); err == nil {
        c.Expire(k, getRandomTTL())
        return sum, err
    }
    return sum, err
}

// get 获取 set k 中的元素列表，每个元素为 i64 类型
func get(c *redis.Client, k string) (vt []int64) {
    v, _ := c.SMembers(k).Result()
    c.Expire(k, getRandomTTL())
    for _, vs := range v {
        v_i64, _ := strconv.ParseInt(vs, 10, 64)
        vt = append(vt, v_i64)
    }
    return vt
}
