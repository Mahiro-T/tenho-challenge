package main

import (
	"go.kyoto.codes/v3/component"
	"go.kyoto.codes/v3/rendering"
)

type IndexPageState struct {
	component.Disposable
	rendering.Template
}

func IndexPage(ctx *component.Context) component.State {
	state := &IndexPageState{}
	return state
}
