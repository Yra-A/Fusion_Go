package utils

import (
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/model/api"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/kitex_gen/user"
)

func ConvertUserToAPI(src *user.UserInfo) *api.UserInfo {
	return &api.UserInfo{
		UserID:         src.UserId,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		MobilePhone:    src.MobilePhone,
		College:        src.College,
		Nickname:       src.Nickname,
		Realname:       src.Realname,
		HasProfile:     src.HasProfile,
		AvatarURL:      src.AvatarUrl,
	}
}

func ConvertUserProfileToAPI(src *user.UserProfileInfo) *api.UserProfileInfo {
	return &api.UserProfileInfo{
		Introduction: src.Introduction,
		QqNumber:     src.QqNumber,
		WechatNumber: src.WechatNumber,
		Honors:       src.Honors,
		UserInfo:     ConvertUserToAPI(src.UserInfo),
	}
}

func ConvertAPIToUser(src *api.UserInfo) *user.UserInfo {
	return &user.UserInfo{
		UserId:         src.UserID,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		MobilePhone:    src.MobilePhone,
		College:        src.College,
		Nickname:       src.Nickname,
		Realname:       src.Realname,
		HasProfile:     src.HasProfile,
		AvatarUrl:      src.AvatarURL,
	}
}

func ConvertAPIProfileToUser(src *api.UserProfileInfo) *user.UserProfileInfo {
	return &user.UserProfileInfo{
		Introduction: src.Introduction,
		QqNumber:     src.QqNumber,
		WechatNumber: src.WechatNumber,
		Honors:       src.Honors,
		UserInfo:     ConvertAPIToUser(src.UserInfo),
	}
}

func ConvertContestToAPI(src *contest.Contest) *api.Contest {
	return &api.Contest{
		ContestID:   src.ContestId,
		Title:       src.Title,
		Description: src.Description,
		CreatedTime: src.CreatedTime,
		Field:       src.Field,
		Format:      src.Format,
		ImageURL:    src.ImageUrl,
		ContestCoreInfo: &api.ContestCoreInfo{
			Deadline:                src.ContestCoreInfo.Deadline,
			Fee:                     src.ContestCoreInfo.Fee,
			TeamSize:                convertTeamSize(src.ContestCoreInfo.TeamSize),
			ParticipantRequirements: src.ContestCoreInfo.ParticipantRequirements,
			OfficialWebsite:         src.ContestCoreInfo.OfficialWebsite,
			AdditionalInfo:          src.ContestCoreInfo.AdditionalInfo,
			Contact:                 convertContacts(src.ContestCoreInfo.Contact),
		},
	}
}

func ConvertBriefInfoToAPI(contestList []*contest.ContestBriefInfo) []*api.ContestBriefInfo {
	apiContestList := make([]*api.ContestBriefInfo, len(contestList))
	for i, contestInfo := range contestList {
		apiContest := &api.ContestBriefInfo{
			ContestBriefInfo: &api.ContestBrief{
				ContestID:   contestInfo.ContestBriefInfo.ContestId,
				Title:       contestInfo.ContestBriefInfo.Title,
				Description: contestInfo.ContestBriefInfo.Description,
				CreatedTime: contestInfo.ContestBriefInfo.CreatedTime,
				Field:       contestInfo.ContestBriefInfo.Field,
				Format:      contestInfo.ContestBriefInfo.Format,
			},
		}
		apiContestList[i] = apiContest
	}
	return apiContestList
}

func convertTeamSize(src *contest.TeamSize) *api.TeamSize {
	if src == nil {
		return nil
	}
	return &api.TeamSize{
		Min: src.Min,
		Max: src.Max,
	}
}

func convertContacts(src []*contest.Contact) []*api.Contact {
	contacts := make([]*api.Contact, 0, len(src))
	for _, c := range src {
		if c != nil {
			contacts = append(contacts, &api.Contact{
				Name:  c.Name,
				Phone: c.Phone,
				Email: c.Email,
			})
		}
	}
	return contacts
}

type ContestBrief struct {
	ContestBriefInfo *api.ContestBriefInfo `json:"contest_brief_info"`
}
