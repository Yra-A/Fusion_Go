package db

import (
	"github.com/Yra-A/Fusion_Go/kitex_gen/team"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"time"
)

type TeamInfo struct {
	TeamID       int32     `gorm:"primary_key;column:team_id"`
	ContestID    int32     `gorm:"column:contest_id"`
	Title        string    `gorm:"column:title"`
	Goal         string    `gorm:"column:goal"`
	CurPeopleNum int32     `gorm:"column:cur_people_num"`
	CreatedTime  time.Time `gorm:"column:created_time"`
	LeaderID     int32     `gorm:"column:leader_id"`
	Description  string    `gorm:"column:description"`
}

func (TeamInfo) TableName() string {
	return "team_info"
}

type TeamApplication struct {
	ApplicationID   int32     `gorm:"primary_key;column:application_id"`
	UserID          int32     `gorm:"column:user_id"`
	TeamID          int32     `gorm:"column:team_id"`
	Reason          string    `gorm:"column:reason"`
	CreatedTime     time.Time `gorm:"column:created_time"`
	ApplicationType int32     `gorm:"column:application_type"`
}

func (TeamApplication) TableName() string {
	return "team_application"
}

type TeamUserRelationship struct {
	TeamUserID int32 `gorm:"primary_key;column:team_user_id"`
	UserID     int32 `gorm:"column:user_id"`
	TeamID     int32 `gorm:"column:team_id"`
}

func (TeamUserRelationship) TableName() string {
	return "team_user_relationship"
}

// CreateTeam 创建团队
func CreateTeam(user_id int32, contest_id int32, title string, goal string, description string) (int32, error) {
	team := &TeamInfo{
		Title:        title,
		ContestID:    contest_id,
		Goal:         goal,
		CurPeopleNum: 1,
		CreatedTime:  time.Now(),
		LeaderID:     user_id,
		Description:  description,
	}
	if err := DB.Create(team).Error; err != nil {
		return 0, err
	}
	var teamInfo TeamInfo
	// 获取最新创建的 team 的 team_id，Last 会返回按主键排序的最后一个满足条件的记录
	if err := DB.Select("team_id").Where("leader_id = ?", user_id).Last(&teamInfo).Error; err != nil {
		return 0, err
	}
	TeamAddUser(teamInfo.TeamID, user_id)
	return teamInfo.TeamID, nil
}

// ModifyTeam 修改团队信息
func ModifyTeam(team_id int32, title string, goal string, description string) error {
	team := &TeamInfo{
		TeamID: team_id,
	}
	if err := DB.Model(&team).Updates(map[string]interface{}{"title": title, "goal": goal, "description": description}).Error; err != nil {
		return err
	}
	return nil
}

func QueryTeamList(contest_id int32) ([]*team.TeamBriefInfo, error) {
	var teamList []*TeamInfo
	if err := DB.Where("contest_id = ?", contest_id).Find(&teamList).Error; err != nil {
		return nil, err
	}
	var teamBriefInfoList []*team.TeamBriefInfo
	for _, t := range teamList {
		teamBriefInfoList = append(teamBriefInfoList, &team.TeamBriefInfo{
			TeamId:       t.TeamID,
			Title:        t.Title,
			Goal:         t.Goal,
			CurPeopleNum: t.CurPeopleNum,
			CreatedTime:  t.CreatedTime.Unix(),
			ContestId:    contest_id,
			LeaderInfo: &team.MemberInfo{
				UserId: t.LeaderID,
			},
		})
	}
	return teamBriefInfoList, nil
}

func QueryTeamInfo(team_id int32) (*team.TeamInfo, error) {
	var teamInfo TeamInfo
	if err := DB.Where("team_id = ?", team_id).First(&teamInfo).Error; err != nil {
		return nil, err
	}
	var teamUserRelationship []*TeamUserRelationship
	if err := DB.Where("team_id = ?", team_id).Find(&teamUserRelationship).Error; err != nil {
		return nil, err
	}
	var memberList []*team.MemberInfo
	for _, t := range teamUserRelationship {
		memberList = append(memberList, &team.MemberInfo{
			UserId: t.UserID,
		})
	}
	return &team.TeamInfo{
		TeamBriefInfo: &team.TeamBriefInfo{
			TeamId:       teamInfo.TeamID,
			Title:        teamInfo.Title,
			Goal:         teamInfo.Goal,
			CurPeopleNum: teamInfo.CurPeopleNum,
			CreatedTime:  teamInfo.CreatedTime.Unix(),
			ContestId:    teamInfo.ContestID,
			LeaderInfo: &team.MemberInfo{
				UserId: teamInfo.LeaderID,
			},
		},
		Description: teamInfo.Description,
		Members:     memberList,
	}, nil
}

func CreateTeamApplication(user_id int32, team_id int32, reason string, created_time int64, application_type int32) error {
	if err := DB.Create(&TeamApplication{
		UserID:          user_id,
		TeamID:          team_id,
		Reason:          reason,
		CreatedTime:     time.Unix(created_time, 0),
		ApplicationType: application_type,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetTeamApplicationList(user_id int32, team_id int32) ([]*team.TeamApplication, error) {
	var teamInfo TeamInfo
	if err := DB.Select("leader_id").Where("team_id = ?", team_id).First(&teamInfo).Error; err != nil {
		return nil, err
	}
	// 只有队长才能处理申请
	if teamInfo.LeaderID != user_id {
		return nil, errno.AuthorizationFailedErr
	}
	var teamApplicationList []*TeamApplication
	if err := DB.Where("team_id = ?", team_id).Find(&teamApplicationList).Error; err != nil {
		return nil, err
	}
	var teamApplications []*team.TeamApplication
	for _, t := range teamApplicationList {
		if t.ApplicationType != 0 {
			teamApplications = append(teamApplications, &team.TeamApplication{
				ApplicationId:   t.ApplicationID,
				TeamId:          t.TeamID,
				Reason:          t.Reason,
				CreatedTime:     t.CreatedTime.Unix(),
				ApplicationType: t.ApplicationType,
				MemberInfo: &team.MemberInfo{
					UserId: t.UserID,
				},
			})
		}
	}
	return teamApplications, nil
}
func TeamAddUser(team_id int32, member_id int32) error {
	var record TeamUserRelationship
	DB.Where("team_id = ? AND user_id = ?", team_id, member_id).First(&record)
	if record.TeamUserID != 0 {
		return nil
	}
	if err := DB.Create(&TeamUserRelationship{
		UserID: member_id,
		TeamID: team_id,
	}).Error; err != nil {
		return err
	}
	// 更新 cur number
	var count int64
	if err := DB.Model(&TeamUserRelationship{}).Where("team_id = ?", team_id).Count(&count).Error; err != nil {
		return err
	}

	if err := DB.Model(&TeamInfo{}).Where("team_id = ?", team_id).Update("cur_people_num", count).Error; err != nil {
		return err
	}
	return nil
}

func TeamManageAction(user_id int32, application_id int32, action_type int32) error {
	var teamApplication TeamApplication
	if err := DB.Where("application_id = ?", application_id).First(&teamApplication).Error; err != nil {
		return err
	}
	var teamInfo TeamInfo
	if err := DB.Select("leader_id").Where("team_id = ?", teamApplication.TeamID).First(&teamInfo).Error; err != nil {
		return err
	}
	// 只有队长才能处理申请
	if teamInfo.LeaderID != user_id {
		return errno.AuthorizationFailedErr
	}

	// 接受申请
	if action_type == 1 {
		TeamAddUser(teamApplication.TeamID, teamApplication.UserID)
	}

	if err := DB.Model(&teamApplication).Update("application_type", 0).Error; err != nil {
		return err
	}
	return nil
}
