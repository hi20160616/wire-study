package main

import (
	"fmt"
	"time"

	"github.com/google/wire"
)

func main() {
	fmt.Printf("Real time greeting: %s [current time elided]\n", initApp().Greet()[0:15])

	fmt.Println("Approach A")
	mt := newMockTimer()
	mockedApp := initMockedAppFromArgs(mt)
	fmt.Println(mockedApp.Greet())
	mt.T = mt.T.AddDate(1999, 0, 0)
	fmt.Println(mockedApp.Greet())

	fmt.Println("Approach B")
	appWithMocks := initMockedApp()
	fmt.Println(appWithMocks.app.Greet())
	appWithMocks.mt.T = appWithMocks.mt.T.AddDate(999, 0, 0)
	fmt.Println(appWithMocks.app.Greet())
}

// First print line coded here

var appSet = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
	wire.InterfaceValue(new(timer), realTime{}),
)

type app struct {
	g greeter
}

type greeter struct {
	T timer
}

type timer interface {
	Now() time.Time
}

type realTime struct{}

func (realTime) Now() time.Time {
	return time.Now()
}

func (g greeter) Greet() string {
	return fmt.Sprintf("Good day! It is %v", g.T.Now())
}

func (a app) Greet() string {
	return a.g.Greet()
}

// Approach A code appended

// mockTimer implements timer using a mocked time.
type mockTimer struct {
	T time.Time
}

func newMockTimer() *mockTimer {
	return &mockTimer{}
}

func (m *mockTimer) Now() time.Time {
	return m.T
}

var appSetWithoutMocks = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
)

// Approach B code appended
var mockAppSet = wire.NewSet(
	wire.Struct(new(app), "*"),
	wire.Struct(new(greeter), "*"),
	wire.Struct(new(appWithMocks), "*"),
	newMockTimer,
	wire.Bind(new(timer), new(*mockTimer)),
)

type appWithMocks struct {
	app app
	mt  *mockTimer
}
