package repository

import (
	"fmt"
	zeroAgency "test-zero-agency"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type NewsWithCategoriesPostgres struct {
	db *sqlx.DB
}

func NewNewsWithCategoriesPostgres(db *sqlx.DB) *NewsWithCategoriesPostgres {
	return &NewsWithCategoriesPostgres{db: db}
}

func (n *NewsWithCategoriesPostgres) GetNewsById(newsId int) ([]zeroAgency.NewsWithCategories, error) {
	var list []zeroAgency.NewsWithCategories

	query := fmt.Sprintf("SELECT n.id, n.title, n.content, array_agg(nc.categoryId) FROM %s n INNER JOIN %s nc ON n.id=nc.newsId WHERE n.id=$1 GROUP BY n.id", newsTable, newsCategoriesTable)
	rows, err := n.db.Query(query, newsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var news zeroAgency.NewsWithCategories
		var categoryIDs []int64
		if err := rows.Scan(&news.Id, &news.Title, &news.Content, pq.Array(&categoryIDs)); err != nil {
			return nil, err
		}

		for _, id := range categoryIDs {
			news.Categories = append(news.Categories, int(id))
		}
		list = append(list, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (n *NewsWithCategoriesPostgres) GetAllNews() ([]zeroAgency.NewsWithCategories, error) {
	var list []zeroAgency.NewsWithCategories

	query := fmt.Sprintf("SELECT n.id, n.title, n.content, array_agg(nc.categoryId) FROM %s n INNER JOIN %s nc ON n.id=nc.newsId GROUP BY n.id ORDER BY n.id DESC", newsTable, newsCategoriesTable)
	rows, err := n.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var news zeroAgency.NewsWithCategories
		var categoryIDs []int64
		if err := rows.Scan(&news.Id, &news.Title, &news.Content, pq.Array(&categoryIDs)); err != nil {
			return nil, err
		}

		for _, id := range categoryIDs {
			news.Categories = append(news.Categories, int(id))
		}
		list = append(list, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (n *NewsWithCategoriesPostgres) UpdateNewsById(newsId zeroAgency.NewsWithCategories) error {

	tx := n.db.MustBegin()

	if newsId.Title != "" {
		_, err := tx.Exec("UPDATE "+newsTable+" SET title=$1 WHERE id=$2",
			newsId.Title, newsId.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if newsId.Content != "" {
		_, err := tx.Exec("UPDATE "+newsTable+" SET content=$1 WHERE id=$2",
			newsId.Content, newsId.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err := tx.Exec("DELETE FROM "+newsCategoriesTable+" WHERE newsId=$1", newsId.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(newsId.Categories) > 0 {
		for _, categoryId := range newsId.Categories {
			_, err := tx.Exec("INSERT INTO "+newsCategoriesTable+" (newsId, categoryId) VALUES ($1, $2)",
				newsId.Id, categoryId)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

func (n *NewsWithCategoriesPostgres) DeleteNewsById(newsId int) error {

	tx, err := n.db.Begin()
	if err != nil {
		return err
	}
	defer func() {

		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	_, err = tx.Exec("DELETE FROM "+newsCategoriesTable+" WHERE newsId=$1", newsId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM "+newsTable+" WHERE id=$1", newsId)
	if err != nil {
		return err
	}

	return nil
}
