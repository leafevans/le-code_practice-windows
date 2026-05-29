package model

import "time"

var ResumeDao _Resume

type (
	_Resume struct{}

	Resume struct {
		ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement;comment:简历ID"`
		UserID    uint      `json:"user_id" gorm:"index;not null;comment:用户ID"`
		Content   string    `json:"content" gorm:"type:longtext;not null;comment:简历内容"`
		FileName  string    `json:"file_name" gorm:"size:255;comment:原始文件名"`
		FileSize  int64     `json:"file_size" gorm:"comment:文件大小（字节）"`
		FileType  string    `json:"file_type" gorm:"size:50;comment:文件类型(pdf/doc/txt等)"`
		IsDefault int       `json:"is_default" gorm:"default:0;comment:是否为默认简历(0-否，1-是)"`
		Deleted   int       `json:"deleted" gorm:"default:0;index;comment:删除标记(0-未删除，1-已删除)"`
		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"`
	}
)

func (Resume) TableName() string {
	return "resumes"
}

// Create 创建简历记录
func (r *_Resume) Create(resume *Resume) error {
	return getDB().Create(resume).Error
}

// GetByID 根据ID获取简历
func (r *_Resume) GetByID(id uint, userID int64) (*Resume, error) {
	var resume Resume
	err := getDB().Where("id = ? AND user_id = ?", id, userID).First(&resume).Error
	return &resume, err
}

// 获取用户简历列表分页
func (r *_Resume) ListByUser(userID int64, page, pageSize int) ([]*Resume, int64, error) {
	if getDB() == nil {
		panic("数据库连接未初始化")
	}
	var resumes []*Resume
	var total int64
	query := getDB().Model(&Resume{}).Where("user_id = ?", userID).Count(&total)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Order("created_at DESC").Find(&resumes).Error; err != nil {
		return nil, 0, err
	}

	return resumes, total, nil
}

func (r *_Resume) Delete(id uint, userID int64) error {
	if getDB() == nil {
		panic("数据库连接未初始化")
	}
	return getDB().Where("id = ? AND user_id = ?", id, userID).Delete(&Resume{}).Error
}

func (r *_Resume) DeleteResume(id uint, userID int64) error {
	if getDB() == nil {
		panic("数据库连接未初始化")
	}
	return getDB().Model(&Resume{}).Where("id = ? AND user_id = ?", id, userID).Update("deleted", 1).Error
}
