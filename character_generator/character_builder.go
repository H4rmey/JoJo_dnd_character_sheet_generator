package character_generator

import (
	"github.com/rivo/tview"
)

func GetAllPassions() {
}

func GetAllStands() {
}

func GetAllAbilities() {
}

func GetAllFeats() {
}

func StartBuildingCharacter() {
	app := tview.NewApplication()

	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}