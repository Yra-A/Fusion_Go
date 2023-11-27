package rpc

// InitRPC 初始化 rpc 客户端
func InitRPC() {
	initUserRpc()
	initContestRpc()
	initTeamRpc()
	initFavoriteRpc()
	initArticleRpc()
}
