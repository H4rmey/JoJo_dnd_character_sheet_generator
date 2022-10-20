package character_generator

import (
	"io/ioutil"
	"log"

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

func StartBuildingCharacter() {
	app := tview.NewApplication()

	stands_list := GetNameArrayStands(GetAllStands().Stands)

	form := tview.NewForm().AddDropDown("Class", stands_list, 0, nil)

	form.AddInputField("olivier", "", 20, nil, nil).
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
