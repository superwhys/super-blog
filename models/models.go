package models

import "time"

type ApiResponseData struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"statusCode"`
	RequestTs  int64  `json:"requestTs"`
}

func PackResponseData(statusCode int, msg string, data any) *ApiResponseData {
	return &ApiResponseData{
		Message:    msg,
		StatusCode: statusCode,
		Data:       data,
		RequestTs:  time.Now().Unix(),
	}
}

type MetaData struct {
	Layout   string   `yaml:"layout" json:"layout"`
	Title    string   `yaml:"title" json:"title"`
	SubTitle string   `yaml:"subTitle" json:"subTitle"`
	Date     string   `yaml:"date" json:"date"`
	Author   string   `yaml:"author" json:"author"`
	Tags     []string `yaml:"tags" json:"tags"`
}

type BlogItem struct {
	MetaData    *MetaData `json:"metaData"`
	FileName    string    `json:"fileName"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"subTitle"`
	PostContent string    `json:"postContent"`
	PostedTime  string    `json:"postedTime"`
	ToEndPoint  string    `json:"toEndPoint"`
}

type BlogListItems struct {
	Items []*BlogItem `json:"items"`
}
