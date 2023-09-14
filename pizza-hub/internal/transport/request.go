package transport

type ChefReq struct {
	Name string `json:"name"`
}

type OrderReq struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
