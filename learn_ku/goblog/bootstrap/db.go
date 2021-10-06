package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	// 自动迁移
	// GORM 的自动迁移工具不支持版本，
	// 只能保持字段与传参的 Struct 一致，无法删除字段。所以我们将自动迁移的代码封装到 migration() 方法里，
	// 以后遇到需要删除字段的情况，将这些代码写到此方法即可。
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}
