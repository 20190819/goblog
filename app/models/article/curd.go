package article

import (
	"fmt"

	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/model"
	"github.com/yangliang4488/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	// err := model.DB.First(&article, id).Error
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	err := model.DB.Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	} else {
		return nil
	}
}

func (article *Article) Update() (affected int64, err error) {
	result := model.DB.Save(&article)
	err = result.Error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (affected int64, err error) {
	result := model.DB.Delete(&article)
	err = result.Error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return result.RowsAffected, nil
}
