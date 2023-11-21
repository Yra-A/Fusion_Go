package dal

import (
    "github.com/Yra-A/Fusion_Go/cmd/favorite/dal/db"
    "github.com/Yra-A/Fusion_Go/cmd/favorite/dal/redis"
)

func Init() {
    db.Init()
    redis.Init()
}
