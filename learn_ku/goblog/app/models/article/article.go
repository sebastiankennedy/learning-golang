package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel

	ID    int
	Title string
	Body  string

	UserID     uint64 `gorm:"not null;index"`
	CategoryID uint64 `gorm:"not null;default:4;index"`
	User       user.User
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(article.ID), 10))
}

// 格式化创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}