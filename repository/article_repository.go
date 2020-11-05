package repository

import "go-blog/model"

func ArticleList() ([]*model.Article, error) {
	query := `select * from articles;`

	var articles []*model.Article
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}
	return articles, nil
}