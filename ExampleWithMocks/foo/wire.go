//go:build wireinject

package main

import "github.com/google/wire"

func initApp() *app {
	wire.Build(appSet)
	return nil
}

func initMockedAppFromArgs(mt timer) *app {
	wire.Build(appSetWithoutMocks)
	return nil
}

func initMockedApp() *appWithMocks {
	wire.Build(mockAppSet)
	return nil
}
