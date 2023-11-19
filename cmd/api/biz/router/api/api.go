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
				_create := _contest.Group("/create", _createMw()...)
				_create.POST("/", append(_contestcreateMw(), api.ContestCreate)...)
			}
			{
				_info0 := _contest.Group("/info", _info0Mw()...)
				_info0.GET("/:contest_id", append(_contestinfoMw(), api.ContestInfo)...)
			}
			{
				_list := _contest.Group("/list", _listMw()...)
				_list.GET("/", append(_contestlistMw(), api.ContestList)...)
			}
		}
		{
			_favorite := _fusion.Group("/favorite", _favoriteMw()...)
			{
				_contest0 := _favorite.Group("/contest", _contest0Mw()...)
				{
					_action := _contest0.Group("/action", _actionMw()...)
					_action.POST("/", append(_contestfavoriteactionMw(), api.ContestFavoriteAction)...)
				}
				{
					_list0 := _contest0.Group("/list", _list0Mw()...)
					_list0.GET("/", append(_contestfavoritelistMw(), api.ContestFavoriteList)...)
				}
			}
		}
		{
			_team0 := _fusion.Group("/team", _team0Mw()...)
			{
				_application := _team0.Group("/application", _applicationMw()...)
				{
					_submit := _application.Group("/submit", _submitMw()...)
					_submit.POST("/", append(_teamapplicationsubmitMw(), api.TeamApplicationSubmit)...)
				}
			}
			{
				_create0 := _team0.Group("/create", _create0Mw()...)
				_create0.POST("/", append(_teamcreateMw(), api.TeamCreate)...)
			}
			{
				_manage := _team0.Group("/manage", _manageMw()...)
				{
					_action0 := _manage.Group("/action", _action0Mw()...)
					_action0.POST("/", append(_teammanageactionMw(), api.TeamManageAction)...)
				}
				{
					_list1 := _manage.Group("/list", _list1Mw()...)
					_list1.GET("/", append(_teammanagelistMw(), api.TeamManageList)...)
				}
			}
		}
		{
			_user := _fusion.Group("/user", _userMw()...)
			{
				_info1 := _user.Group("/info", _info1Mw()...)
				_info1.GET("/", append(_userinfoMw(), api.UserInfo)...)
				{
					_upload := _info1.Group("/upload", _uploadMw()...)
					_upload.POST("/", append(_userinfouploadMw(), api.UserInfoUpload)...)
				}
			}
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), api.UserLogin)...)
			}
			{
				_profile := _user.Group("/profile", _profileMw()...)
				_profile.GET("/:user_id", append(_userprofileinfoMw(), api.UserProfileInfo)...)
				{
					_upload0 := _profile.Group("/upload", _upload0Mw()...)
					_upload0.POST("/", append(_userprofileuploadMw(), api.UserProfileUpload)...)
				}
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), api.UserRegister)...)
			}
		}
		{
			_utils := _fusion.Group("/utils", _utilsMw()...)
			{
				_upload1 := _utils.Group("/upload", _upload1Mw()...)
				_upload1.POST("/img", append(_imageuploadMw(), api.ImageUpload)...)
			}
		}
	}
}
