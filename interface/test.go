package _interface

type IDuck interface {
	Swimming()
	Call()

}

type duck struct {

}

func (d duck) Swimming() {
	panic("implement me")
}

func (d duck) Call() {
	panic("implement me")
}

var _d IDuck = duck{}