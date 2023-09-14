package transport

type ShortURLReq struct {
	Destination string `json:"destination" validate:"required"`
	ExpireAt    string `json:"expire_at"  validate:"required"`
}
