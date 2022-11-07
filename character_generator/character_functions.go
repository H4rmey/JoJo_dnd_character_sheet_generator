package character_generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func ExportCharacterToYaml(cs Character_Sheet) {
	file, err := os.OpenFile(strings.ReplaceAll(cs.Name, " ", "_")+".yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(cs)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}
}
func PrintCharacterSheet(character_sheet Character_Sheet) {
	fmt.Printf("============= Character Sheet Level: %d ================\n", character_sheet.Level)
	for i := 0; i < len(character_sheet.Skill_pro); i++ {
		sp := character_sheet.Skill_pro[i]
		fmt.Printf("p:%d - %s (%s): %d \n", sp.Level, sp.Skill_name, sp.Stat_type, sp.Value)
	}

	fmt.Printf("proficiency bonus: %d\n", character_sheet.Proficiency_bonus)
	fmt.Printf("Passions: %s\n", character_sheet.Passion.Name)
	fmt.Printf("Stand Type: %s\n", character_sheet.Stand.Name)
	fmt.Printf("Ability Dice: %s\n", character_sheet.Ability_dice)

	fmt.Printf("Ability Score:\n")
	fmt.Printf("cha: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["cha"], character_sheet.Char_ability_modifiers["cha"])
	fmt.Printf("con: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["con"], character_sheet.Char_ability_modifiers["con"])
	fmt.Printf("dex: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["dex"], character_sheet.Char_ability_modifiers["dex"])
	fmt.Printf("itl: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["itl"], character_sheet.Char_ability_modifiers["itl"])
	fmt.Printf("str: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["str"], character_sheet.Char_ability_modifiers["str"])
	fmt.Printf("wis: %d   \t modifiers: %d \n", character_sheet.Char_ability_scores["wis"], character_sheet.Char_ability_modifiers["wis"])

	fmt.Printf("Saving Throws:\n")
	fmt.Printf("cha: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["cha"].Value, character_sheet.Saving_throws["cha"].Level)
	fmt.Printf("con: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["con"].Value, character_sheet.Saving_throws["con"].Level)
	fmt.Printf("dex: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["dex"].Value, character_sheet.Saving_throws["dex"].Level)
	fmt.Printf("itl: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["itl"].Value, character_sheet.Saving_throws["itl"].Level)
	fmt.Printf("str: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["str"].Value, character_sheet.Saving_throws["str"].Level)
	fmt.Printf("wis: %d   \t is_poficient: %d \n", character_sheet.Saving_throws["wis"].Value, character_sheet.Saving_throws["wis"].Level)

	fmt.Printf("Stand Ability Scores:\n")
	fmt.Printf("pow: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["pow"], character_sheet.Stand_ability_modifiers["pow"])
	fmt.Printf("pre: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["pre"], character_sheet.Stand_ability_modifiers["pre"])
	fmt.Printf("dur: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["dur"], character_sheet.Stand_ability_modifiers["dur"])
	fmt.Printf("ran: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["ran"], character_sheet.Stand_ability_modifiers["ran"])
	fmt.Printf("spe: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["spe"], character_sheet.Stand_ability_modifiers["spe"])
	fmt.Printf("ngy: %d   \t modifier: %d\n", character_sheet.Stand_ability_scores["ngy"], character_sheet.Stand_ability_modifiers["ngy"])

	fmt.Printf("char_AC: %d\n", character_sheet.Armor_class)
	fmt.Printf("stand_AC: %d\n", character_sheet.Stand_ac)
	fmt.Printf("stand_movement: %d\n", character_sheet.Stand_movement)
	fmt.Printf("stand_attack_per_turn: %d\n", character_sheet.Stand_attack_per_turn)
	fmt.Printf("stand_damage_reduction: %d\n", character_sheet.Stand_damage_reduction)
	fmt.Printf("initiative: %d\n", character_sheet.Initiative)
	fmt.Printf("hit_dice: %s\n", character_sheet.Stand.HitDice)
	fmt.Printf("stand_dc: %d\n", character_sheet.Stand_dc)
	fmt.Printf("hit_points: %d\n", character_sheet.Hit_points)
}

func calculate_stand_AC(am map[string]int) int {
	var a = [3]int{-1, -1, -1}
	a[0] = 10 + am["pre"] + am["spe"]
	a[1] = 10 + am["pre"] + am["dur"]
	a[2] = 10 + am["spe"] + am["dur"]

	highest := -99999
	for _, i := range a {
		if i > highest {
			highest = i
		}
	}

	return highest
}

func LoadAbilities() []Abilities {
	file, err := ioutil.ReadFile("../yaml/abilities.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var yaml_abilities Yaml2Abilities
	error := yaml.Unmarshal([]byte(file), &yaml_abilities)
	if error != nil {
		log.Fatal(err)
	}

	return yaml_abilities.Abilities
}

func LoadFeats() []Feat {
	file, err := ioutil.ReadFile("../yaml/feats.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var yaml_feats Yaml2Feats
	error := yaml.Unmarshal([]byte(file), &yaml_feats)
	if error != nil {
		log.Fatal(err)
	}

	return yaml_feats.Feats
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func generate_ability_scores_stands(stand Stand, ability_scores map[string]int) map[string]int {
	stand_ability_scores := map[string]int{
		"pow": stand.ModificationChart.Str * ability_scores["str"],
		"pre": stand.ModificationChart.Dex * ability_scores["dex"],
		"dur": stand.ModificationChart.Con * ability_scores["con"],
		"ran": stand.ModificationChart.Itl * ability_scores["itl"],
		"spe": stand.ModificationChart.Wis * ability_scores["wis"],
		"ngy": stand.ModificationChart.Cha * ability_scores["cha"],
	}
	return stand_ability_scores
}

func generate_ability_modifiers_stand(stand_ability_scores map[string]int) map[string]int {
	stand_ability_modifiers := map[string]int{
		"pow": GenerateModifierStand(stand_ability_scores["pow"]),
		"pre": GenerateModifierStand(stand_ability_scores["pre"]),
		"dur": GenerateModifierStand(stand_ability_scores["dur"]),
		"ran": GenerateModifierStand(stand_ability_scores["ran"]),
		"spe": GenerateModifierStand(stand_ability_scores["spe"]),
		"ngy": GenerateModifierStand(stand_ability_scores["ngy"]),
	}
	return stand_ability_modifiers
}

func LoadStands() Yaml2Stand {
	file, err := ioutil.ReadFile("./yaml/stand_types.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var yaml_stands Yaml2Stand
	error := yaml.Unmarshal([]byte(file), &yaml_stands)
	if error != nil {
		log.Fatal(err)
	}

	return yaml_stands
}

func GenerateModifierChar(value int) int {
	return int(math.Floor((float64(value) - 10.0) / 2.0))
}

func GenerateModifierStand(value int) int {
	return value / 10
}

func generate_ability_modifiers_character(ability_scores map[string]int) map[string]int {
	ability_modifiers_char := make(map[string]int, 6)

	ability_modifiers_char["cha"] = GenerateModifierChar(ability_scores["cha"])
	ability_modifiers_char["con"] = GenerateModifierChar(ability_scores["con"])
	ability_modifiers_char["dex"] = GenerateModifierChar(ability_scores["dex"])
	ability_modifiers_char["itl"] = GenerateModifierChar(ability_scores["itl"])
	ability_modifiers_char["str"] = GenerateModifierChar(ability_scores["str"])
	ability_modifiers_char["wis"] = GenerateModifierChar(ability_scores["wis"])

	return ability_modifiers_char
}

func generate_passion_ability_score(passion Passion, ability_scores map[string]int) map[string]int {
	cas := make(map[string]int, 6)

	cas["cha"] = ability_scores["cha"] + passion.Traits.AbilityScores.Cha
	cas["con"] = ability_scores["con"] + passion.Traits.AbilityScores.Con
	cas["dex"] = ability_scores["dex"] + passion.Traits.AbilityScores.Dex
	cas["itl"] = ability_scores["itl"] + passion.Traits.AbilityScores.Itl
	cas["str"] = ability_scores["str"] + passion.Traits.AbilityScores.Str
	cas["wis"] = ability_scores["wis"] + passion.Traits.AbilityScores.Wis

	return cas
}
