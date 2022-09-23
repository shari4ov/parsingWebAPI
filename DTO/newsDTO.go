package dto

type News struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Href    string `json:"href"`
	News_id string `json:"news_id"`
}
type CreateNews struct {
	Title   string `json:"title"`
	Href    string `json:"href"`
	News_id string `json:"id"`
}
