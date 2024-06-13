package cron

import (
	"gorm.io/gorm"
	"testProject/innternal/Repository/model"
	"testProject/innternal/pkg/filesystem"
	"time"
)

type ClearArticle struct {
	DB         *gorm.DB
	Filesystem filesystem.IFilesystem
}

func (c *ClearArticle) Name() string {
	return "clear.article"
}

// 配置定时任务
func (c *ClearArticle) Spec() string {
	return "0 1 * * *"
}

func (c *ClearArticle) Enable() bool {
	return true
}

func (c *ClearArticle) clearAnnex() {
	lastId := 0
	size := 100
	for {
		items := make([]*model.ArticleAnnex, 0)
		//将查询结果填充到 items 变量中。
		err := c.DB.Model(&model.ArticleAnnex{}).Where("id > ? and status = 2 and deleted_at <= ?", lastId, time.Now().AddDate(0, 0, -30)).Order("id asc").Limit(size).Scan(&items).Error
		if err != nil {
			break
		}
		for _, item := range items {
			_ = c.Filesystem.Delete(c.Filesystem.BucketPrivateName(), item.Path)
			c.DB.Delete(&model.ArticleAnnex{}, item.Id)
		}
		if len(items) < size {
			break
		}
		lastId = items[size-1].Id
	}
}
func (c *ClearArticle) clearNote() {
	lastId := 0
	size := 100
	for {
		items := make([]*model.Article, 0)
		err := c.DB.Model(&model.Article{}).Where("id > ? and status = 2 and deleted_at <= ?", lastId, time.Now().AddDate(0, 0, -30)).Order("id asc").Limit(size).Scan(items).Error
		if err != nil {
			break
		}
		for _, item := range items {
			subItems := make([]*model.ArticleAnnex, 0)
			//查询path列
			if err := c.DB.Model(&model.ArticleAnnex{}).Select("path").Where("article_id = ?", item.Id).Scan(&subItems).Error; err != nil {
				continue
			}
			for _, subItem := range subItems {
				_ = c.Filesystem.Delete(c.Filesystem.BucketPrivateName(), subItem.Path)
				//删除表中主键=subItem.Id的记录
				c.DB.Delete(&model.ArticleAnnex{}, subItem.Id)
			}
			c.DB.Delete(&model.Article{}, item.Id)

		}
		lastId = items[size-1].Id
	}
}
