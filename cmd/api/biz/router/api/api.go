// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	api "github.com/Yra-A/Fusion_Go/cmd/api/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_fusion := root.Group("/fusion", _fusionMw()...)
		{
			_contest := _fusion.Group("/contest", _contestMw()...)
			_contest.POST("/create", append(_contestcreateMw(), api.ContestCreate)...)
			_contest.GET("/list", append(_contestlistMw(), api.ContestList)...)
			{
				_contest_id := _contest.Group("/:contest_id", _contest_idMw()...)
				{
					_team := _contest_id.Group("/team", _teamMw()...)
					_team.GET("/list", append(_teamlistMw(), api.TeamList)...)
					{
						_info := _team.Group("/info", _infoMw()...)
						_info.GET("/:team_id", append(_teaminfoMw(), api.TeamInfo)...)
					}
				}
			}
			{
				_info0 := _contest.Group("/info", _info0Mw()...)
				_info0.GET("/:contest_id", append(_contestinfoMw(), api.ContestInfo)...)
			}
		}
		{
			_favorite := _fusion.Group("/favorite", _favoriteMw()...)
			{
				_contest0 := _favorite.Group("/contest", _contest0Mw()...)
				_contest0.POST("/action", append(_contestfavoriteactionMw(), api.ContestFavoriteAction)...)
				_contest0.GET("/list", append(_contestfavoritelistMw(), api.ContestFavoriteList)...)
			}
		}
		{
			_team0 := _fusion.Group("/team", _team0Mw()...)
			_team0.POST("/create", append(_teamcreateMw(), api.TeamCreate)...)
			{
				_application := _team0.Group("/application", _applicationMw()...)
				_application.POST("/submit", append(_teamapplicationsubmitMw(), api.TeamApplicationSubmit)...)
			}
			{
				_manage := _team0.Group("/manage", _manageMw()...)
				_manage.POST("/action", append(_teammanageactionMw(), api.TeamManageAction)...)
				_manage.GET("/list", append(_teammanagelistMw(), api.TeamManageList)...)
			}
		}
		{
			_user := _fusion.Group("/user", _userMw()...)
			_user.POST("/login", append(_userloginMw(), api.UserLogin)...)
			_user.POST("/register", append(_userregisterMw(), api.UserRegister)...)
			_user.GET("/info", append(_userinfoMw(), api.UserInfo)...)
			_info1 := _user.Group("/info", _info1Mw()...)
			_info1.POST("/upload", append(_userinfouploadMw(), api.UserInfoUpload)...)
			{
				_profile := _user.Group("/profile", _profileMw()...)
				_profile.GET("/:user_id", append(_userprofileinfoMw(), api.UserProfileInfo)...)
				_profile.POST("/upload", append(_userprofileuploadMw(), api.UserProfileUpload)...)
			}
		}
		{
			_utils := _fusion.Group("/utils", _utilsMw()...)
			{
				_upload := _utils.Group("/upload", _uploadMw()...)
				_upload.POST("/img", append(_imageuploadMw(), api.ImageUpload)...)
			}
		}
	}
}
