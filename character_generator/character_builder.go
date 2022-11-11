package character_generator

import (
	"strconv"

	"github.com/rivo/tview"
)

type Character_Builder struct {
	pages     *tview.Pages
	buttons   *tview.Form
	grid      *tview.Grid
	pageIndex int
}

func (cb *Character_Builder) AddPage(id int, primitive *tview.Primitive) {
	cb.pages.AddPage(strconv.Itoa(id), *primitive, true, true)
}

func (cb *Character_Builder) Init() {
	cb.pageIndex = 0
	cb.buttons.AddButton("<-- Previous", func() {
		cb.pageIndex = (cb.pageIndex - 1) % cb.pages.GetPageCount()
		cb.pages.SwitchToPage(strconv.Itoa(cb.pageIndex))
	})
	cb.buttons.AddButton("Next -->", func() {
		cb.pageIndex = (cb.pageIndex + 1) % cb.pages.GetPageCount()
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

func SetAbilityModifiers(cs *Character_Sheet) tview.Primitive {
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

func CreateStandPage(cs *Character_Sheet) {
	gridStandScores := tview.NewForm()
	gridStandStats := tview.NewForm()
	gridAbilityScores := tview.NewForm()
	gridAbilityStats := tview.NewForm()
	gridCharInfo := tview.NewForm()
	gridCharDetail := tview.NewForm()
	gridAbilities := tview.NewForm()
	gridSkills := tview.NewForm()
	gridSaving := tview.NewForm()

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

	cb := Character_Builder{
		pages:     tview.NewPages(),
		buttons:   tview.NewForm(),
		grid:      tview.NewGrid(),
		pageIndex: 0,
	}
	cb.Init()
	//1. pick stand
	//2. pick passion
	//3. generate ability scores/modifiers (also calc other stuff)
	pickPassion := PickStand(&cs)
	pickStand := PickPassion(&cs)
	cb.AddPage(0, &pickStand)
	cb.AddPage(1, &pickPassion)

	app.SetRoot(cb.grid, true)
	app.SetFocus(cb.grid)
	app.EnableMouse(true)

	uiRun(app)
}
