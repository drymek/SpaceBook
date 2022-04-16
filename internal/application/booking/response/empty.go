package response

import "dryka.pl/SpaceBook/internal/application/httpx"

type Empty struct {
	code int
}

func (b Empty) StatusCode() int {
	return b.code
}

func NewEmptyResponse(code int) httpx.StatusCodeHolder {
	return Empty{
		code: code,
	}
}
