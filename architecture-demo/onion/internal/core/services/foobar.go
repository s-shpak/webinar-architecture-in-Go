package services

import (
	"strconv"

	"onion/internal/core/model"
)

type foobarSimple struct{}

func (f *foobarSimple) Calculate(req *model.FoobarRequest) *model.FoobarResponse {
	resp := &model.FoobarResponse{
		Data: make([]string, 0, req.N),
	}
	for i := 1; i <= req.N; i++ {
		var res string
		if i%3 == 0 && i%5 == 0 {
			res = "foobar"
		} else if i%3 == 0 {
			res = "foo"
		} else if i%5 == 0 {
			res = "bar"
		} else {
			res = strconv.Itoa(i)
		}
		resp.Data = append(resp.Data, res)
	}
	return resp
}
