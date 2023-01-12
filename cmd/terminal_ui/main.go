package main

import (
	"context"

	"github.com/rivo/tview"
)

var (
	ctx context.Context
)

func getWorlds(right *tview.List) {

	for i, world := range worlds {
		worldText := world
		right = right.AddItem(worldText, "Not Ready", rune(i+97), func() {
			modal := tview.NewModal().
				SetText(worldText).
				AddButtons([]string{"Home"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Home" {
						home()
					}
				})
			app := ctx.Value("app").(*tview.Application)
			app.SetRoot(modal, true)
		})
	}

}

func home() {
	app := ctx.Value("app").(*tview.Application)
	flex := ctx.Value("flex").(*tview.Flex)
	app.SetRoot(flex, true)
}

func main() {

	ctx = context.Background()
	app := tview.NewApplication()
	ctx = context.WithValue(ctx, "app", app)

	newWorld := func() {
		app := ctx.Value("app").(*tview.Application)
		form := ctx.Value("form").(*tview.Form)
		app.SetRoot(form, true)

	}
	refreshWorlds := func() {
		right := ctx.Value("right").(*tview.List)
		right.Clear()
		getWorlds(right)
	}
	renderWorld := func() {
		right := ctx.Value("right").(*tview.List)
		app := ctx.Value("app").(*tview.Application)
		app.SetFocus(right)
	}
	createWorld := func() {

	}

	form := tview.NewForm().
		AddInputField("World name", "", 20, nil, nil).
		AddInputField("Sim Pop", "", 3, tview.InputFieldInteger, nil).
		AddInputField("Sim Size", "", 3, tview.InputFieldInteger, nil).
		AddButton("Save", createWorld).
		AddButton("Home", home)

	ctx = context.WithValue(ctx, "form", form)

	opList := tview.NewList().
		AddItem("Create World", "Create a new creature simulation", 'a', newWorld).
		AddItem("Refresh Worlds", "Refreshes world list", 'b', refreshWorlds).
		AddItem("Render World", "Renders the simulation", 'c', renderWorld).
		AddItem("Quit", "Press to exit", 'q', func() {
			app := ctx.Value("app").(*tview.Application)
			app.Stop()
		})
	ctx = context.WithValue(ctx, "opList", opList)

	right := tview.NewList()
	getWorlds(right)
	ctx = context.WithValue(ctx, "right", right)
	flex := tview.NewFlex().
		AddItem(opList, 0, 1, true).
		AddItem(right, 40, 1, true)
	ctx = context.WithValue(ctx, "flex", flex)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
