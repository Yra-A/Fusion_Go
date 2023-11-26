package utils

import (
	"github.com/Yra-A/Fusion_Go/cmd/api/biz/model/api"
	"github.com/Yra-A/Fusion_Go/kitex_gen/article"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"github.com/Yra-A/Fusion_Go/kitex_gen/favorite"
	"github.com/Yra-A/Fusion_Go/kitex_gen/team"
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
			TeamSize:                ConvertTeamSizeToAPI(src.ContestCoreInfo.TeamSize),
			ParticipantRequirements: src.ContestCoreInfo.ParticipantRequirements,
			OfficialWebsite:         src.ContestCoreInfo.OfficialWebsite,
			AdditionalInfo:          src.ContestCoreInfo.AdditionalInfo,
			Contact:                 ConvertContactsToAPI(src.ContestCoreInfo.Contact),
		},
		IsFavorite: src.IsFavorite,
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

func ConvertBriefFavoriteInfoToAPI(contestList []*favorite.ContestBriefInfo) []*api.ContestBriefInfo {
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

func ConvertTeamSizeToAPI(src *contest.TeamSize) *api.TeamSize {
	if src == nil {
		return nil
	}
	return &api.TeamSize{
		Min: src.Min,
		Max: src.Max,
	}
}

func ConvertContactsToAPI(src []*contest.Contact) []*api.Contact {
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

func ConvertContactsToContest(src []*api.Contact) []*contest.Contact {
	contacts := make([]*contest.Contact, 0, len(src))
	for _, c := range src {
		if c != nil {
			contacts = append(contacts, &contest.Contact{
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

func ConvertTeamBriefInfoListToAPI(teamList []*team.TeamBriefInfo) (apiTeamList []*api.TeamBriefInfo) {
	for _, teamBriefInfo := range teamList {
		apiTeamList = append(apiTeamList, ConvertTeamBriefInfoToAPI(teamBriefInfo))
	}
	return
}

func ConvertMemberInfoListToAPI(memberList []*team.MemberInfo) (apiMemberList []*api.MemberInfo) {
	for _, memberInfo := range memberList {
		apiMemberList = append(apiMemberList, ConvertMemberInfoToAPI(memberInfo))
	}
	return
}

func ConvertMemberInfoToAPI(src *team.MemberInfo) *api.MemberInfo {
	if src == nil {
		return nil
	}
	return &api.MemberInfo{
		UserID:         src.UserId,
		Nickname:       src.Nickname,
		College:        src.College,
		AvatarURL:      src.AvatarUrl,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		Honors:         src.Honors,
	}
}

func ConverAPIToMemberInfo(src *api.MemberInfo) *team.MemberInfo {
	if src == nil {
		return nil
	}
	return &team.MemberInfo{
		UserId:         src.UserID,
		Nickname:       src.Nickname,
		College:        src.College,
		AvatarUrl:      src.AvatarURL,
		Gender:         src.Gender,
		EnrollmentYear: src.EnrollmentYear,
		Honors:         src.Honors,
	}
}

func ConvertTeamBriefInfoToAPI(src *team.TeamBriefInfo) *api.TeamBriefInfo {
	if src == nil {
		return nil
	}
	return &api.TeamBriefInfo{
		ContestID:    src.ContestId,
		TeamID:       src.TeamId,
		Title:        src.Title,
		Goal:         src.Goal,
		CurPeopleNum: src.CurPeopleNum,
		CreatedTime:  src.CreatedTime,
		LeaderInfo:   ConvertMemberInfoToAPI(src.LeaderInfo),
	}
}

func ConverTeamInfoToAPI(src *team.TeamInfo) *api.TeamInfo {
	if src == nil {
		return nil
	}
	return &api.TeamInfo{
		TeamBriefInfo: ConvertTeamBriefInfoToAPI(src.TeamBriefInfo),
		Description:   src.Description,
		Members:       ConvertMemberInfoListToAPI(src.Members),
	}
}

func ConvertApplicationToAPI(src *team.TeamApplication) *api.TeamApplication {
	if src == nil {
		return nil
	}
	return &api.TeamApplication{
		ApplicationID:   src.ApplicationId,
		TeamID:          src.TeamId,
		Reason:          src.Reason,
		CreatedTime:     src.CreatedTime,
		ApplicationType: src.ApplicationType,
		MemberInfo:      ConvertMemberInfoToAPI(src.MemberInfo),
	}
}

func ConvertApplicationListToAPI(applicationList []*team.TeamApplication) (apiApplicationList []*api.TeamApplication) {
	for _, application := range applicationList {
		apiApplicationList = append(apiApplicationList, ConvertApplicationToAPI(application))
	}
	return
}

func ConvertContestBriefToFavoriteBrief(src *contest.ContestBrief) *favorite.ContestBrief {
	if src == nil {
		return nil
	}
	return &favorite.ContestBrief{
		ContestId:   src.ContestId,
		Title:       src.Title,
		Description: src.Description,
		CreatedTime: src.CreatedTime,
		Field:       src.Field,
		Format:      src.Format,
	}
}

func ConvertArticleBriefInfoToAPI(articleList []*article.ArticleBriefInfo) []*api.ArticleBriefInfo {
	apiArticleList := make([]*api.ArticleBriefInfo, len(articleList))
	for i, articleInfo := range articleList {
		apiArticle := &api.ArticleBriefInfo{
			ArticleBriefInfo: &api.ArticleBrief{
				ArticleID:   articleInfo.ArticleBriefInfo.ArticleId,
				Title:       articleInfo.ArticleBriefInfo.Title,
				AuthorID:    articleInfo.ArticleBriefInfo.AuthorId,
				Author:      articleInfo.ArticleBriefInfo.Author,
				CreatedTime: articleInfo.ArticleBriefInfo.CreatedTime,
				Link:        articleInfo.ArticleBriefInfo.Link,
			},
		}
		apiArticleList[i] = apiArticle
	}
	return apiArticleList
}
