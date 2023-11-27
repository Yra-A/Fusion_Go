package constants

const (
	SecretKey   = "secret key"
	IdentityKey = "user_id"

	UserServiceName     = "user"
	ContestServiceName  = "contest"
	TeamServiceName     = "team"
	FavoriteServiceName = "favorite"
	ArticleServiceName  = "article"

	//CPURateLimit float64 = 80.0
	//
	//MaxVideoSize int64 = 128 * 1024 * 1024 // 可上传的单个视频大小最大为 128 MB
	//MaxFeedCount       = 30                // 视频列表最大视频个数
)

var (
	// MySQLDefaultDSN = "root:gorm@tcp(host.docker.internal:18000)/test_douyin?charset=utf8&parseTime=True&loc=Local"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:18000)/fusion_db?charset=utf8&parseTime=True&loc=Local"
	// EtcdAddress = "host.docker.internal:2379"
	EtcdAddress = "localhost:2379"

	RedisAddress  = "localhost:18003"
	RedisPassword = ""
	DBIndex       = 0 // 使用默认的数据库
)
