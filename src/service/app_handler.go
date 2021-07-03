package main

type Result struct {
	Value int `json:"value"`
}

type AppHandler interface {
	Sum(x, y int) Result
	Multiply(x, y int) Result
}

type AppHandlerStruct struct{}

func (a *AppHandlerStruct) Sum(x, y int) (r Result) {
	r.Value = Sum(x, y)
	return
}
func (a *AppHandlerStruct) Multiply(x, y int) (r Result) {
	r.Value = Multiply(x, y)
	return
}
