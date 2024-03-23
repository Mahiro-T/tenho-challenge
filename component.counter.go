package main

import (
	"go.kyoto.codes/v3/component"
	"go.kyoto.codes/v3/rendering"
)

type CounterComponentState struct {
	component.Disposable
	rendering.Template

	Count int
}

func CounterComponent(ctx *component.Context) component.State {
	state := &CounterComponentState{}
	return state
}
