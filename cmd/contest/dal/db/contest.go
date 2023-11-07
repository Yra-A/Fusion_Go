package db

import (
	"github.com/Yra-A/Fusion_Go/cmd/user/dal/db"
	"github.com/Yra-A/Fusion_Go/kitex_gen/contest"
	"gorm.io/gorm"
	"sort"
	"strings"
)

// Contest corresponds to the 'contest' table in the database.
type Contest struct {
	gorm.Model
	ContestID               int32  `gorm:"primary_key;column:contest_id"`
	Title                   string `gorm:"column:title;not null"`
	ImageURL                string `gorm:"column:image_url"`
	Field                   string `gorm:"column:field"`
	Format                  string `gorm:"column:format"`
	Description             string `gorm:"column:description;type:text"`
	Deadline                int32  `gorm:"column:deadline"`
	Fee                     string `gorm:"column:fee"`
	TeamSizeMin             int32  `gorm:"column:team_size_min"`
	TeamSizeMax             int32  `gorm:"column:team_size_max"`
	ParticipantRequirements string `gorm:"column:participant_requirements;type:text"`
	OfficialWebsite         string `gorm:"column:official_website"`
	AdditionalInfo          string `gorm:"column:additional_info;type:text"`
	CreatedTime             int64  `gorm:"column:created_time"`
}

func (Contest) TableName() string {
	return "contests"
}

// Contact corresponds to the 'contact' table in the database.
type Contact struct {
	gorm.Model
	ContactID int32  `gorm:"primary_key;column:contact_id"`
	Name      string `gorm:"column:name"`
	Phone     string `gorm:"column:phone"`
	Email     string `gorm:"column:email"`
}

func (Contact) TableName() string {
	return "contacts"
}

// mock
var contestsData = map[int32]*Contest{
	1: {
		ContestID:   1,
		Title:       "全国青少年科技竞赛",
		Field:       "科技",
		Format:      "团队",
		Description: "面向全国的青少年科技创新团队竞赛。",
		CreatedTime: 1610000000,
	},
	2: {
		ContestID:   2,
		Title:       "城市艺术与设计大赛",
		Field:       "艺术",
		Format:      "团队",
		Description: "公开征集城市艺术设计方案，鼓励创意。",
		CreatedTime: 1620000000,
	},
	3: {
		ContestID:   3,
		Title:       "全国数学模型挑战赛",
		Field:       "科学",
		Format:      "个人",
		Description: "挑战数学模型解决实际问题的能力。",
		CreatedTime: 1630000000,
	},
	4: {
		ContestID:   4,
		Title:       "全国大学生电子设计竞赛",
		Field:       "科技",
		Format:      "个人",
		Description: "面向全国大学生的电子设计竞赛。",
		CreatedTime: 1640000000,
	},
	5: {
		ContestID:   5,
		Title:       "全国大学生数学建模竞赛",
		Field:       "科学",
		Format:      "团队",
		Description: "面向全国大学生的数学建模竞赛。",
		CreatedTime: 1650000000,
	},
	// ...其他Contest数据
}

var contactsData = map[int32]*Contact{
	1: {
		// Gorm Model fields are omitted for brevity
		ContactID: 1,
		Name:      "Alice",
		Phone:     "1234567890",
		Email:     "alice@example.com",
	},
	// 其他Contact数据...
	2: {
		ContactID: 2,
		Name:      "Bob",
		Phone:     "1234567890",
		Email:     "bob@example.com",
	},
}

var relationshipsData = map[int32]*ContestContactRelationship{
	1: {
		// Gorm Model fields are omitted for brevity
		ContestContactID: 1,
		ContactID:        1,
		ContestID:        1,
	},
	// 其他Relationship数据...
	2: {
		// Gorm Model fields are omitted for brevity
		ContestContactID: 2,
		ContactID:        2,
		ContestID:        1,
	},
	3: {
		ContestContactID: 3,
		ContactID:        1,
		ContestID:        2,
	},
}

// ContestContactRelationship is the join table for the many-to-many relationship between contests and contacts.
type ContestContactRelationship struct {
	gorm.Model
	ContestContactID int32 `gorm:"primary_key;column:contest_contact_id"`
	ContactID        int32 `gorm:"column:contact_id"`
	ContestID        int32 `gorm:"column:contest_id"`
}

func (ContestContactRelationship) TableName() string {
	return "contest_contact_relationship"
}

func QueryContestByContestId(contest_id int32) (*Contest, error) {
	//var contest Contest
	//if err := DB.Where("contest_id = ?", contest_id).First(&contest).Error; err != nil {
	//	return nil, err
	//}
	//return &contest, nil
	contest, ok := contestsData[contest_id]
	if !ok {
		return nil, nil
	}
	return contest, nil
}

func FindContestContacts(contest_id int32) ([]*ContestContactRelationship, error) {
	//var contestContacts []*ContestContactRelationship
	//if err := DB.Where("contest_id = ?", contest_id).Find(&contestContacts).Error; err != nil {
	//	return nil, err
	//}
	//return contestContacts, nil
	var contestContacts []*ContestContactRelationship
	for _, v := range relationshipsData {
		if v.ContestID == contest_id {
			contestContacts = append(contestContacts, v)
		}
	}
	return contestContacts, nil
}

func QueryContactsByContactIds(contactIds []int32) ([]*Contact, error) {
	//var contacts []*Contact
	//if err := DB.Where("contact_id in ?", contactIds).Find(&contacts).Error; err != nil {
	//	return nil, err
	//}
	//return contacts, nil
	var contacts []*Contact
	for _, v := range contactsData {
		for _, id := range contactIds {
			if v.ContactID == id {
				contacts = append(contacts, v)
			}
		}
	}
	return contacts, nil
}

// FetchContestList 根据关键字、领域、格式、限制和偏移量来获取赛事列表。
func FetchContestList(keyword string, fields []string, formats []string, limit int32, offset int32) ([]*contest.ContestBrief, error) {
	var contestBriefInfos []*contest.ContestBrief
	query := db.DB.Model(&Contest{}).Order("created_time desc")

	// 根据字段和格式筛选
	if len(fields) > 0 {
		query = query.Where("field IN ?", fields)
	}
	if len(formats) > 0 {
		query = query.Where("format IN ?", formats)
	}

	// 根据关键字筛选
	if keyword != "" {
		likeKeyword := "%" + keyword + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", likeKeyword, likeKeyword)
	}

	// 选择指定的字段，确保字段名与 ContestBrief 结构体中的标签一致
	query = query.Select("contest_id, title, description, created_time, field, format")

	// 应用分页
	query = query.Offset(int(offset)).Limit(int(limit))

	// 执行查询
	if err := query.Find(&contestBriefInfos).Error; err != nil {
		return nil, err
	}

	return contestBriefInfos, nil
}

// FetchContestListMock is a mock function to simulate database behavior for testing purposes.
func FetchContestListMock(keyword string, fields []string, formats []string, limit int32, offset int32) ([]*contest.ContestBrief, error) {
	contestsSlice := make([]*Contest, 0, len(contestsData))
	for _, c := range contestsData {
		contestsSlice = append(contestsSlice, c)
	}

	// 按 CreatedTime 降序排序
	sort.Slice(contestsSlice, func(i, j int) bool {
		return contestsSlice[i].CreatedTime > contestsSlice[j].CreatedTime
	})

	// 过滤排序后的数据
	filteredSortedContests := []*Contest{}
	for _, c := range contestsSlice {
		if (contains(fields, c.Field) || len(fields) == 0) &&
			(contains(formats, c.Format) || len(formats) == 0) &&
			(keyword == "" || strings.Contains(strings.ToLower(c.Title), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(c.Description), strings.ToLower(keyword))) {
			filteredSortedContests = append(filteredSortedContests, c)
		}
	}

	// 应用 offset 和 limit
	start := int(offset)
	if start >= len(filteredSortedContests) {
		return nil, nil
	}
	end := start + int(limit)
	if end > len(filteredSortedContests) {
		end = len(filteredSortedContests)
	}
	paginatedContests := filteredSortedContests[start:end]

	// 创建 ContestBriefInfo 列表
	briefInfos := make([]*contest.ContestBrief, 0, len(paginatedContests))
	for _, c := range paginatedContests {
		briefInfo := &contest.ContestBrief{
			ContestId:   c.ContestID,
			Title:       c.Title,
			Description: c.Description,
			CreatedTime: c.CreatedTime,
			Field:       c.Field,
			Format:      c.Format,
		}
		briefInfos = append(briefInfos, briefInfo)
	}

	return briefInfos, nil
}

// 辅助函数，检查切片中是否包含某个字符串
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
