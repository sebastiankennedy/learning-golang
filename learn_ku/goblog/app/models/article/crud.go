package article

import (
    "goblog/pkg/logger"
    "goblog/pkg/model"
    "goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idStr string) (Article, error) {
    /*
       .First() 是 gorm.DB 提供的用以从结果集中获取第一条数据的查询方法。
       .Error 是 GORM 的错误处理机制。与常见的 Go 代码不同，因 GORM 提供的是链式 API，如果遇到任何错误，GORM 会设置 *gorm.DB 的 Error 字段，您需要像这样检查它。
       在 GORM 中，当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误。
    */
    var article Article
    id := types.StringToInt(idStr)
    if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
        return article, err
    }

    return article, nil
}

func GetAll() ([]Article, error) {
    // 将 map 类型的 Article 对象传参到 Find() 方法内，即可获取到所有文章数据
    var articles []Article
    if err := model.DB.Preload("User").Find(&articles).Error; err != nil {
        return articles, err
    }

    return articles, nil
}

func (article *Article) Create() (err error) {
    result := model.DB.Create(&article)
    if err = result.Error; err != nil {
        logger.LogError(err)
        return err
    }

    return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
    result := model.DB.Save(&article)
    if err = result.Error; err != nil {
        logger.LogError(err)
        return 0, err
    }

    return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
    result := model.DB.Delete(&article)
    if err = result.Error; err != nil {
        logger.LogError(err)
        return 0, err
    }

    return result.RowsAffected, nil
}