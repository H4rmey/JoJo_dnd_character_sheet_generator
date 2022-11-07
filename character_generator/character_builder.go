package character_generator

import (
	"strconv"

	"github.com/rivo/tview"
)

type Character_Builder struct {
	pages     *tview.Pages
	buttons   *tview.Form
	grid      *tview.Grid
	pageList  []tview.Grid
	pageIndex int
}

func AddPage(cb *Character_Builder, id int, primitive *tview.Primitive) {
	cb.pages.AddPage(strconv.Itoa(id), *primitive, true, true)
}

func Init(cb *Character_Builder) {
	cb.pageIndex = 0
	cb.buttons.AddButton("<-- Previous", func() {
		cb.pageIndex = (cb.pageIndex - 1) % len(cb.pageList)
		cb.pages.SwitchToPage(strconv.Itoa(cb.pageIndex))
	})
	cb.buttons.AddButton("Next -->", func() {
		cb.pageIndex = (cb.pageIndex + 1) % len(cb.pageList)
		cb.pages.SwitchToPage(strconv.Itoa(cb.pageIndex))
	})

	cb.grid = tview.NewGrid()
	cb.grid.AddItem(cb.pages, 0, 0, 1, 1, 10, 10, true)
	cb.grid.AddItem(cb.buttons, 1, 0, 1, 1, 10, 10, true)
}

func PickPassion(cs *Character_Sheet) tview.Primitive {
	passionList := GetNameArrayPassions(GetAllPassions().Passions)

	textView := tview.NewTextView()
	textView.SetBorderPadding(1, 1, 1, 1)
	textView.SetWordWrap(true)

	form := tview.NewForm()
	form.AddDropDown("Passion:", passionList, 0, func(firstName string, value int) {
		textView.SetText(cs.Stands[value].Description)
	})

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Select your stand type").SetTitleAlign(tview.AlignCenter)

	grid := tview.NewGrid()
	grid.AddItem(form, 0, 0, 1, 1, 10, 10, true)
	grid.AddItem(textView, 1, 0, 1, 1, 10, 10, true)

	return grid
}

func PickStand(cs *Character_Sheet) tview.Primitive {
	stands_list := GetNameArrayStands(GetAllStands().Stands)

	textView := tview.NewTextView()
	textView.SetBorderPadding(1, 1, 1, 1)
	textView.SetWordWrap(true)

	form := tview.NewForm()
	form.AddDropDown("Class", stands_list, 0, func(firstName string, value int) {
		textView.SetText(cs.Stands[value].Description)
	})

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Select your stand type").SetTitleAlign(tview.AlignCenter)

	grid := tview.NewGrid()
	grid.AddItem(form, 0, 0, 1, 1, 10, 10, true)
	grid.AddItem(textView, 1, 0, 1, 1, 10, 10, true)

	return grid
}

func uiRun(app *tview.Application) {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func StartBuildingCharacter() {
	app := tview.NewApplication()
	var cs Character_Sheet

	cs.Stands = LoadStands().Stands

	//flex := PickStand(&cs)

	var cb Character_Builder
	Init(&cb)

	app.SetRoot(cb.grid, true)
	app.SetFocus(cb.grid)
	app.EnableMouse(true)

	uiRun(app)
}
