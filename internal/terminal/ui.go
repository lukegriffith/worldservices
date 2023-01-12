package terminal

import (
	"context"
	"strconv"

	"github.com/faiface/mainthread"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/worldservices/internal/render"
	"github.com/lukegriffith/worldservices/internal/world"
	"github.com/rivo/tview"
)

var (
	ctx context.Context
)

type ctxKey int

const (
	ctxAppKey ctxKey = iota
	ctxFormKey
	ctxRightKey
	ctxOpListKey
	ctxFlexKey
)

func SetupAndRun(passedContext context.Context) {
	ctx = passedContext
	app := tview.NewApplication()
	ctx = context.WithValue(ctx, ctxAppKey, app)

	newWorld := func() {
		app := ctx.Value(ctxAppKey).(*tview.Application)
		form := ctx.Value(ctxFormKey).(*tview.Form)
		app.SetRoot(form, true)

	}
	refreshWorlds := func() {
		right := ctx.Value(ctxRightKey).(*tview.List)
		right.Clear()
		getWorlds(right)
	}
	setRenderWorldMenu := func() {
		right := ctx.Value(ctxRightKey).(*tview.List)
		app := ctx.Value(ctxAppKey).(*tview.Application)
		app.SetFocus(right)
	}
	createWorld := func() {
		form := ctx.Value(ctxFormKey).(*tview.Form)
		wNameInput := form.GetFormItem(0).(*tview.InputField)
		wPopInput := form.GetFormItem(1).(*tview.InputField)
		wSizeInput := form.GetFormItem(2).(*tview.InputField)

		wName := wNameInput.GetText()
		wPop, err := strconv.Atoi(wPopInput.GetText())
		if err != nil {
			panic(err)
		}
		wSize, err := strconv.Atoi(wSizeInput.GetText())
		if err != nil {
			panic(err)
		}
		register := func() {
			world.RegisterWorld(wName)
			w := world.NewWorld(wSize, wPop)
			world.SetWorld(wName, w)
		}
		go register()
	}

	form := tview.NewForm().
		AddInputField("World name", "", 20, nil, nil).
		AddInputField("Sim Pop", "", 3, tview.InputFieldInteger, nil).
		AddInputField("Sim Size", "", 3, tview.InputFieldInteger, nil).
		AddButton("Save", createWorld).
		AddButton("Home", home)

	ctx = context.WithValue(ctx, ctxFormKey, form)

	opList := tview.NewList().
		AddItem("Create World", "Create a new creature simulation", 'a', newWorld).
		AddItem("Refresh Worlds", "Refreshes world list", 'b', refreshWorlds).
		AddItem("Render World", "Renders the simulation", 'c', setRenderWorldMenu).
		AddItem("Quit", "Press to exit", 'q', func() {
			app := ctx.Value(ctxAppKey).(*tview.Application)
			app.Stop()
		})
	ctx = context.WithValue(ctx, ctxOpListKey, opList)

	right := tview.NewList()
	getWorlds(right)
	ctx = context.WithValue(ctx, ctxRightKey, right)
	flex := tview.NewFlex().
		AddItem(opList, 0, 1, true).
		AddItem(right, 40, 1, true)
	ctx = context.WithValue(ctx, ctxFlexKey, flex)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func getWorlds(right *tview.List) {
	i := 0
	worldsStatus := world.Worlds.GetWorldStatuses()
	for worldName, status := range worldsStatus {
		right = right.AddItem(worldName, status, rune(i+97), func() {
			world.GetWorld(worldName)

			runRender := func() {
				render.Render(worldName)
			}

			runOpenGL := func() {
				pixelgl.Run(runRender)
			}
			/// :( didnt work
			mainthread.Run(runOpenGL)

			app := ctx.Value(ctxAppKey).(*tview.Application)
			flex := ctx.Value(ctxFlexKey).(*tview.Flex)
			app.SetRoot(flex, true)

		})
		i++
	}

}

func home() {
	app := ctx.Value(ctxAppKey).(*tview.Application)
	flex := ctx.Value(ctxFlexKey).(*tview.Flex)
	app.SetRoot(flex, true)
}
