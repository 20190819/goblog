package article

import (
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
