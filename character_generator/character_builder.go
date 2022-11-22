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

func PageCreateStand(ecs *Export_Character_Sheet, app *tview.Application) *tview.Grid {
	// gridStandScores := tview.NewGrid()
	// gridStandStats := tview.NewGrid()
	// gridAbilityScores := tview.NewGrid()
	// gridAbilityStats := tview.NewGrid()
	// gridCharInfo := tview.NewGrid()
	// gridAbilities := tview.NewGrid()
	// gridSkills := tview.NewGrid()
	// gridSaving := tview.NewGrid()

	//define the character building sheet
	CBS := tview.NewGrid()

	//define the grid for character details
	formCharDetail := tview.NewForm()
	formCharDetail.SetBorder(true)
	textLength := 24
	formCharDetail.AddInputField("Name: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.Name = text
	}).AddInputField("Languages: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.Languages = text
	}).AddInputField("Age: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.Age = text
	}).AddInputField("Height: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.Height = text
	}).AddInputField("Weight: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.Weight = text
	}).AddInputField("SkinTone: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.SkinTone = text
	}).AddInputField("EyeColor: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.EyeColor = text
	}).AddInputField("HairColor: ", "", textLength, nil, func(text string) {
		ecs.CharDetails.HairColor = text
	})

	formAbiltiyScores := tview.NewForm()
	formAbiltiyScores.SetBorder(true)
	abilityScores := []string{"7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}
	formAbiltiyScores.AddDropDown("Strength", abilityScores, 3, func(option string, optionIndex int) {

	})
	formAbiltiyScores.AddDropDown("Dexterity", abilityScores, 3, func(option string, optionIndex int) {

	})
	formAbiltiyScores.AddDropDown("Constitution", abilityScores, 3, func(option string, optionIndex int) {

	})
	formAbiltiyScores.AddDropDown("Intelligence", abilityScores, 3, func(option string, optionIndex int) {

	})
	formAbiltiyScores.AddDropDown("Wisdom", abilityScores, 3, func(option string, optionIndex int) {

	})
	formAbiltiyScores.AddDropDown("Charisma", abilityScores, 3, func(option string, optionIndex int) {

	})

	CBS.AddItem(formAbiltiyScores, 0, 0, 1, 1, 10, 10, true)
	CBS.AddItem(formCharDetail, 0, 1, 1, 1, 10, 10, false)

	return CBS
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
	var ecs Export_Character_Sheet

	cs.Stands = LoadStands().Stands

	//flex := PickStand(&cs)

	cb := Character_Builder{
		pages:     tview.NewPages(),
		buttons:   tview.NewForm(),
		grid:      tview.NewGrid(),
		pageIndex: 0,
	}
	cb.Init()
	cb.grid = PageCreateStand(&ecs, app)

	//pickPassion := PickStand(&cs)
	//pickStand := PickPassion(&cs)
	//cb.AddPage(0, &pickStand)
	//cb.AddPage(1, &pickPassion)

	app.SetRoot(cb.grid, true)
	app.SetFocus(cb.grid)
	app.EnableMouse(true)

	uiRun(app)
}
