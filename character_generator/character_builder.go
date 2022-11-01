package character_generator

import (
	"io/ioutil"
	"log"

	"strconv"

	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
)

func GetAllPassions() Yaml2Passion {
	file, err := ioutil.ReadFile("./yaml/passions.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data Yaml2Passion
	error := yaml.Unmarshal([]byte(file), &data)
	if error != nil {
		log.Fatal(err)
	}

	return data
}

func GetAllStands() Yaml2Stand {
	file, err := ioutil.ReadFile("./yaml/stand_types.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data Yaml2Stand
	error := yaml.Unmarshal([]byte(file), &data)
	if error != nil {
		log.Fatal(err)
	}

	return data
}

func GetAllAbilities() Yaml2Abilities {
	file, err := ioutil.ReadFile("./yaml/abilities.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data Yaml2Abilities
	error := yaml.Unmarshal([]byte(file), &data)
	if error != nil {
		log.Fatal(err)
	}

	return data
}

func GetAllFeats() Yaml2Feats {
	file, err := ioutil.ReadFile("./yaml/feats.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data Yaml2Feats
	error := yaml.Unmarshal([]byte(file), &data)
	if error != nil {
		log.Fatal(err)
	}

	return data
}

func GetNameArrayPassions(data []Passion) []string {
	ret := []string{}
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i].Name)
	}

	return ret
}

func GetNameArrayStands(data []Stand) []string {
	ret := []string{}
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i].Name)
	}

	return ret
}

func GetNameArrayAbilities(data []Abilities) []string {
	ret := []string{}
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i].Name)
	}

	return ret
}

func GetNameArrayFeats(data []Feat) []string {
	ret := []string{}
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i].Name)
	}

	return ret
}

func StartCharacterBuildingGUI(cs *Character_Sheet, app *tview.Application) *tview.Pages {
	pages := tview.NewPages()

	pickStand := PickStand(cs)
	pickPassion := PickPassion(cs)

	pageList := []tview.Grid{
		*pickStand,
		*pickPassion,
	}

	for pageIndex := 0; pageIndex < len(pageList); pageIndex++ {
		i := strconv.Itoa(pageIndex)
		pages.AddPage(i, &pageList[pageIndex], false, pageIndex == 0)

		buttonNext := tview.NewButton("Next -->").SetSelectedFunc(func() {
			i = strconv.Itoa((pageIndex + 1) % len(pageList))
			pages.SwitchToPage(i)
		})

		pageList[pageIndex].AddItem(buttonNext, 2, 0, 1, 1, 10, 10, true)
	}

	return pages
}

func PickPassion(cs *Character_Sheet) *tview.Grid {
	stands_list := GetNameArrayStands(GetAllStands().Stands)

	textView := tview.NewTextView()
	textView.SetBorderPadding(1, 1, 1, 1)
	textView.SetWordWrap(true)

	form := tview.NewForm()
	form.AddDropDown("Passion:", stands_list, 0, func(firstName string, value int) {
		textView.SetText(cs.Stands[value].Description)
	})

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Select your stand type").SetTitleAlign(tview.AlignCenter)

	grid := tview.NewGrid()
	grid.AddItem(form, 0, 0, 1, 1, 10, 10, true)
	grid.AddItem(textView, 1, 0, 1, 1, 10, 10, true)

	return grid
}

func PickStand(cs *Character_Sheet) *tview.Grid {
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

	flex := PickStand(&cs)

	//flex := StartCharacterBuildingGUI(&cs, app)
	app.SetRoot(flex, true)
	app.SetFocus(flex)
	app.EnableMouse(true)

	uiRun(app)
}
