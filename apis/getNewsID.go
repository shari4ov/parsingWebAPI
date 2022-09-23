package apis

import (
	dto "api/DTO"
	"api/storage"
	"fmt"
)

func GetNewsById(id string) dto.News {
	n := dto.News{}
	db := storage.OpenConnection()
	sqlStatement := `SELECT * FROM news WHERE news_id=$1;`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&n.Id, &n.Title, &n.News_id, &n.Href)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return n
}
