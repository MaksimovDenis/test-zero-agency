package zeroAgency

type News struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type NewsCategories struct {
	NewsId     int   `json:"newsId" db:"newsId"`
	CategoryId []int `json:"categoryId" db:"categoryId"`
}

type NewsWithCategories struct {
	Id         int    `json:"Id" db:"id"`
	Title      string `json:"Title" db:"title"`
	Content    string `json:"Content" db:"content"`
	Categories []int  `json:"Categories" db:"categoryid"`
}
