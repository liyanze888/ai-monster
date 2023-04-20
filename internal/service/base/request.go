package base

type PageBaseRequest struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}

type PageBaseResponse struct {
	TotalNum  int `json:"totalNum"`
	TotalPage int `json:"totalPage"`
}
