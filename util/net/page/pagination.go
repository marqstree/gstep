package pagination

type Pagination struct {
	Total int `json:"total"`
	List  any `json:"list"`
}
