package db

import (
	"errors"
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"gorm.io/gorm"
	"time"
)

// Contest corresponds to the 'contest' table in the database.
type Contest struct {
	ContestID               int32     `gorm:"primary_key;column:contest_id"`
	Title                   string    `gorm:"column:title;not null"`
	ImageURL                string    `gorm:"column:image_url"`
	Field                   string    `gorm:"column:field"`
	Format                  string    `gorm:"column:format"`
	Description             string    `gorm:"column:description;type:text"`
	Deadline                int32     `gorm:"column:deadline"`
	Fee                     string    `gorm:"column:fee"`
	TeamSizeMin             int32     `gorm:"column:team_size_min"`
	TeamSizeMax             int32     `gorm:"column:team_size_max"`
	ParticipantRequirements string    `gorm:"column:participant_requirements;type:text"`
	OfficialWebsite         string    `gorm:"column:official_website"`
	AdditionalInfo          string    `gorm:"column:additional_info;type:text"`
	CreatedTime             time.Time `gorm:"column:created_time"`
}

func (Contest) TableName() string {
	return "contest"
}

// Contact corresponds to the 'contact' table in the database.
type Contact struct {
	ContactID int32  `gorm:"primary_key;column:contact_id"`
	Name      string `gorm:"column:name"`
	Phone     string `gorm:"column:phone"`
	Email     string `gorm:"column:email"`
}

func (Contact) TableName() string {
	return "contact"
}

// ContestContactRelationship is the join table for the many-to-many relationship between contests and contacts.
type ContestContactRelationship struct {
	ContestContactID int32 `gorm:"primary_key;column:contest_contact_id"`
	ContactID        int32 `gorm:"column:contact_id"`
	ContestID        int32 `gorm:"column:contest_id"`
}

func (ContestContactRelationship) TableName() string {
	return "contest_contact_relationship"
}

type ContestBrief struct {
	ContestID   int32
	Title       string
	Description string
	CreatedTime time.Time
	Field       string
	Format      string
}

func CreateOrUpdateContestWithTx(tx *gorm.DB, c *Contest) (int32, error) {
	if c.ContestID == 0 {
		if err := tx.Create(c).Error; err != nil {
			return 0, err
		}
		return c.ContestID, nil
	}
	var existingContest Contest
	err := tx.Where("contest_id = ?", c.ContestID).First(&existingContest).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errno.ContestNotExistErr
	}
	if err != nil {
		return 0, err
	}
	if err := tx.Model(&existingContest).Updates(c).Error; err != nil {
		return 0, err
	}
	return c.ContestID, nil
}

func CreateOrUpdateContactWithTx(tx *gorm.DB, contacts []*Contact) error {
	for _, v := range contacts {
		var existingContact Contact
		err := tx.Where("name = ? AND phone = ? AND email = ?", v.Name, v.Phone, v.Email).First(&existingContact).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.Create(v).Error; err != nil {
				return err
			}
			continue
		}
		v.ContactID = existingContact.ContactID
		if err = tx.Model(&existingContact).Updates(v).Error; err != nil {
			return err
		}
	}
	return nil
}

func ContestAddContactsWithTx(tx *gorm.DB, contestID int32, contactIDs []int32) error {
	// 如果没有提供联系人ID，则删除所有关系
	if len(contactIDs) == 0 {
		return tx.Where("contest_id = ?", contestID).Delete(&ContestContactRelationship{}).Error
	}

	// 删除不再关联的联系人关系
	if err := tx.Where("contest_id = ? AND contact_id NOT IN ?", contestID, contactIDs).
		Delete(&ContestContactRelationship{}).Error; err != nil {
		return err
	}

	// 为每个联系人检查并创建关系
	for _, contactID := range contactIDs {
		var relation ContestContactRelationship
		err := tx.Where("contest_id = ? AND contact_id = ?", contestID, contactID).First(&relation).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRelation := ContestContactRelationship{
				ContestID: contestID,
				ContactID: contactID,
			}
			if err = tx.Create(&newRelation).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func QueryContestByContestId(contest_id int32) (*Contest, error) {
	var contest Contest
	if err := DB.Where("contest_id = ?", contest_id).First(&contest).Error; err != nil {
		return nil, err
	}
	return &contest, nil
}

func FindContestContacts(contest_id int32) ([]*ContestContactRelationship, error) {
	var contestContacts []*ContestContactRelationship
	if err := DB.Where("contest_id = ?", contest_id).Find(&contestContacts).Error; err != nil {
		return nil, err
	}
	return contestContacts, nil
}

func QueryContactsByContactIds(contactIds []int32) ([]*Contact, error) {
	var contacts []*Contact
	if err := DB.Where("contact_id in ?", contactIds).Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

// FetchContestList 根据关键字、领域、格式、限制和偏移量来获取赛事列表。
func FetchContestList(keyword string, fields []string, formats []string, limit int32, offset int32) ([]*ContestBrief, int32, error) {
	var contestBriefInfos []*ContestBrief

	query := DB.Model(&Contest{}).Order("created_time desc")

	// 根据字段和格式筛选
	if len(fields) > 0 && fields[0] != "" {
		query = query.Where("field IN ?", fields)
	}
	if len(formats) > 0 && formats[0] != "" {
		query = query.Where("format IN ?", formats)
	}

	// 根据关键字筛选
	if keyword != "" {
		likeKeyword := "%" + keyword + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", likeKeyword, likeKeyword)
	}

	// 选择指定的字段，确保字段名与 ContestBrief 结构体中的标签一致
	query = query.Select("contest_id, title, description, created_time, field, format")

	var total int64

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用分页
	query = query.Offset(int(offset)).Limit(int(limit))

	// 执行查询
	if err := query.Find(&contestBriefInfos).Error; err != nil {
		return nil, 0, err
	}
	return contestBriefInfos, int32(total), nil
}

func FetchContestListByContestIds(contestIds []int32) ([]*ContestBrief, error) {
	var contestBriefInfos []*ContestBrief

	// 结果将按照contestIds中的顺序排序
	if err := DB.Model(&Contest{}).
		Where("contest_id IN ?", contestIds).
		Select("contest_id, title, description, created_time, field, format").
		Find(&contestBriefInfos).Error; err != nil {
		return nil, err
	}

	return contestBriefInfos, nil
}
