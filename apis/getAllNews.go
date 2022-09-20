package api

import (
	dto "api/DTO"
	"api/storage"
	"fmt"
)

func GetAllNews() []dto.News {
	n := []dto.News{}
	db := storage.OpenConnection()
	sqlStatement := `SELECT * FROM news;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return n
	}
	defer rows.Close()
	for rows.Next() {
		var news dto.News
		err := rows.Scan(&news.Id, &news.Title, &news.News_id, &news.Href)
		if err != nil {
			fmt.Println(err)
			return n
		}
		n = append(n, news)
	}
	return n
}
