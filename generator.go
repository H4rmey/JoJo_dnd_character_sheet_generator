package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Proficiency struct {
	level      int    `yaml:"level"`
	value      int    `yaml:"value"`
	stat_type  string `yaml:"stat_type"`
	skill_name string `yaml:"skill_name"`
}

// Stand structs
type Yaml2Stand struct {
	Stands []Stand `yaml:"stands"`
}

type Stand struct {
	AttackDice                string            `yaml:"attack_dice"`
	AttackDiceHigherLevel     string            `yaml:"attack_dice_higher_level"`
	HitDice                   string            `yaml:"hit_dice"`
	OnLevelUp                 string            `yaml:"on_level_up"`
	Name                      string            `yaml:"name"`
	ModificationChart         ModificationChart `yaml:"modification_chart"`
	Description               string            `yaml:"description"`
	Note                      string            `yaml:"note"`
	Special                   []int             `yaml:"special"`
	AttackDicePastLevelEleven string            `yaml:"attack_dice_past_level_eleven"`
	LevelChart                []string          `yaml:"levels"`
	Abilities                 []string          `yaml:"abilities"`
}

type ModificationChart struct {
	Str int `yaml:"str"`
	Dex int `yaml:"dex"`
	Con int `yaml:"con"`
	Itl int `yaml:"itl"`
	Wis int `yaml:"wis"`
	Cha int `yaml:"cha"`
}

// Passion Structs
type Yaml2Passion struct {
	Passions []Passion `yaml:"Passions"`
}

type Passion struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Examples    string `yaml:"examples"`
	Traits      Traits `yaml:"traits"`
}

type Traits struct {
	Proficiencies      []string      `yaml:"proficiencies"`
	ProficienciesExtra string        `yaml:"proficiencies_extra"`
	Languages          int           `yaml:"languages"`
	Extra              string        `yaml:"extra"`
	SavingThrows       []string      `yaml:"saving_throws"`
	AbilityScores      AbilityScores `yaml:"ability_scores"`
}

type AbilityScores struct {
	Str int `yaml:"str"`
	Dex int `yaml:"dex"`
	Con int `yaml:"con"`
	Itl int `yaml:"itl"`
	Wis int `yaml:"wis"`
	Cha int `yaml:"cha"`
}

// Feat Structs
type Yaml2Feats struct {
	Feats []Feat `yaml:"feats"`
}

type Feat struct {
	Name        string   `yaml:"name"`
	Prereq      string   `yaml:"prereq"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effects"`
}

// Abilities Structs
type Yaml2Abilities struct {
	Abilities []Abilities `yaml:"abilities"`
}

type Abilities struct {
	Name    string   `yaml:"name"`
	Text    string   `yaml:"text"`
	Allowed []string `yaml:"allowed"`
}

// Character_Sheet
type Character_Sheet struct {
	name                    string                 `yaml:"name"`
	passion                 Passion                `yaml:"passion"`
	char_ability_scores     map[string]int         `yaml:"char_ability_scores"`
	char_ability_modifiers  map[string]int         `yaml:"char_ability_modifiers"`
	skill_pro               []Proficiency          `yaml:"skill_pro"`
	saving_throws           map[string]Proficiency `yaml:"saving_throws"`
	ability_dice            string                 `yaml:"ability_dice"`
	hit_points              int                    `yaml:"hit_points"`
	armor_class             int                    `yaml:"armor_class"`
	hit_dice                string                 `yaml:"hit_dice"`
	initiative              int                    `yaml:"initiative"`
	passive_perception      int                    `yaml:"passive_perception"`
	movement                int                    `yaml:"movement"`
	level                   int                    `yaml:"level"`
	proficiency_bonus       int                    `yaml:"proficiency_bonus"`
	stand_ability_scores    map[string]int         `yaml:"stand_ability_scores"`
	stand_ability_modifiers map[string]int         `yaml:"stand_ability_modifiers"`
	stand_damage_reduction  int                    `yaml:"stand_damage_reduction"`
	stand_movement          int                    `yaml:"stand_damage_reduction"`
	stand_attack_per_turn   int                    `yaml:"stand_damage_reduction"`
	stand_dc                int                    `yaml:"stand_dc"`
	stand_ac                int                    `yaml:"stand_ac"`
	stand                   Stand                  `yaml:"stand"`
	feats_list              []string               `yaml:"feats_list"`
	abilities_list          []string               `yaml:"abilities_list"`
	feats                   []Feat                 `yaml:"feats"`
	abilities               []Abilities            `yaml:"abilities"`
}

// Generates the main stats of a character
func generate_ability_scores_character() map[string]int {
	stats_char := make(map[string]int, 6)

	base_value := 8

	stats_char["cha"] = base_value
	stats_char["con"] = base_value
	stats_char["dex"] = base_value
	stats_char["itl"] = base_value
	stats_char["str"] = base_value
	stats_char["wis"] = base_value

	point_pool_max := 27
	point_pool := point_pool_max

	min := -2
	max := 9

	rand.Seed(time.Now().UnixNano())

	values := [6]int{0, 0, 0, 0, 0, 0}
	for point_pool > 0 {
		for i := 0; i < 6; i++ {
			value := rand.Intn(max) + min
			if point_pool >= value && values[i]+value <= max && values[i]+value >= min {
				point_pool -= value
				values[i] += value
			}
		}
	}

	stats_char["cha"] += values[0]
	stats_char["con"] += values[1]
	stats_char["dex"] += values[2]
	stats_char["itl"] += values[3]
	stats_char["str"] += values[4]
	stats_char["wis"] += values[5]

	result := 0
	for _, v := range values {
		result += v
	}

	if result > point_pool_max {
		log.Panic("point pool was depleted?")
	}

	return stats_char
}

func generate_modifier_char(value int) int {
	return int(math.Floor((float64(value) - 10.0) / 2.0))
}

func generate_modifier_stand(value int) int {
	return value / 10
}

func generate_ability_modifiers_character(ability_scores map[string]int) map[string]int {
	ability_modifiers_char := make(map[string]int, 6)

	ability_modifiers_char["cha"] = generate_modifier_char(ability_scores["cha"])
	ability_modifiers_char["con"] = generate_modifier_char(ability_scores["con"])
	ability_modifiers_char["dex"] = generate_modifier_char(ability_scores["dex"])
	ability_modifiers_char["itl"] = generate_modifier_char(ability_scores["itl"])
	ability_modifiers_char["str"] = generate_modifier_char(ability_scores["str"])
	ability_modifiers_char["wis"] = generate_modifier_char(ability_scores["wis"])

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

func generate_passion_saving_throws(passion Passion, ability_modifiers_char map[string]int, proficiency_bonus int) map[string]Proficiency {
	saving_throws := map[string]Proficiency{
		"cha": Proficiency{level: 0, value: ability_modifiers_char["cha"], stat_type: "cha"},
		"con": Proficiency{level: 0, value: ability_modifiers_char["con"], stat_type: "con"},
		"dex": Proficiency{level: 0, value: ability_modifiers_char["dex"], stat_type: "dex"},
		"itl": Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "itl"},
		"str": Proficiency{level: 0, value: ability_modifiers_char["str"], stat_type: "str"},
		"wis": Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis"},
	}

	saving_throws_new := passion.Traits.SavingThrows
	saving_throws_amount := len(saving_throws_new)
	ability_scores_list := []string{"cha", "con", "dex", "itl", "str", "wis"}
	for i := 0; i < saving_throws_amount; i++ {
		saving_throw := saving_throws_new[i]

		if strings.Contains(saving_throw, "|") {
			temp := strings.Split(saving_throw, "|")

			rand.Seed(time.Now().UnixNano())
			temp_len := len(temp) - 1
			temp_id := rand.Intn(temp_len)

			saving_throw = temp[temp_id]
		}

		if strings.Contains(saving_throw, "any") {
			rand.Seed(time.Now().UnixNano())
			temp_len := len(ability_scores_list) - 1
			temp_id := rand.Intn(temp_len - 1)

			saving_throw = ability_scores_list[temp_id]

			//remove entry from slice
			ability_scores_list[temp_id] = ability_scores_list[len(ability_scores_list)-1]
			ability_scores_list[len(ability_scores_list)-1] = ""
			ability_scores_list = ability_scores_list[:len(ability_scores_list)-1]
		}

		if entry, ok := saving_throws[saving_throw]; ok {
			if strings.Contains(saving_throw, "+") {
				entry.level = 2
			} else {
				entry.level = 1
			}
			entry.value += proficiency_bonus * saving_throws["cha"].level
			saving_throws[saving_throw] = entry
		}

	}

	return saving_throws
}

func generate_passion_proficiencies(passion Passion, ability_modifiers_char map[string]int, proficiency_bonus int) []Proficiency {

	var skill_pro = []Proficiency{
		Proficiency{level: 0, value: ability_modifiers_char["dex"], stat_type: "dex", skill_name: "Acrobatics"},
		Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis", skill_name: "AnimalHandling"},
		Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "int", skill_name: "Arcana"},
		Proficiency{level: 0, value: ability_modifiers_char["str"], stat_type: "str", skill_name: "Athlete"},
		Proficiency{level: 0, value: ability_modifiers_char["cha"], stat_type: "cha", skill_name: "Deception"},
		Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "int", skill_name: "History"},
		Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis", skill_name: "Insight"},
		Proficiency{level: 0, value: ability_modifiers_char["cha"], stat_type: "cha", skill_name: "Intimidation"},
		Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "int", skill_name: "Investigation"},
		Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis", skill_name: "Medicine"},
		Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "int", skill_name: "Nature"},
		Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis", skill_name: "Perception"},
		Proficiency{level: 0, value: ability_modifiers_char["cha"], stat_type: "cha", skill_name: "Performer"},
		Proficiency{level: 0, value: ability_modifiers_char["cha"], stat_type: "cha", skill_name: "Persuasion"},
		Proficiency{level: 0, value: ability_modifiers_char["itl"], stat_type: "int", skill_name: "Religion"},
		Proficiency{level: 0, value: ability_modifiers_char["dex"], stat_type: "dex", skill_name: "SlightOfHand"},
		Proficiency{level: 0, value: ability_modifiers_char["dex"], stat_type: "dex", skill_name: "Stealth"},
		Proficiency{level: 0, value: ability_modifiers_char["wis"], stat_type: "wis", skill_name: "Survival"},
	}

	nof_proficiencies := len(passion.Traits.Proficiencies)
	for i := 0; i < nof_proficiencies; i++ {
		prof := passion.Traits.Proficiencies[i]

		if strings.Contains(prof, "|") {
			temp := strings.Split(prof, "|")

			rand.Seed(time.Now().UnixNano())
			temp_len := len(temp) - 1
			temp_id := rand.Intn(temp_len)

			prof = temp[temp_id]
		}

		for p := 0; p < len(skill_pro); p++ {
			if prof == skill_pro[p].skill_name {
				if strings.Contains(prof, "+") {
					skill_pro[p].level = 2
				} else {
					skill_pro[p].level = 1
				}

				skill_pro[p].value += skill_pro[p].level * proficiency_bonus
			}
		}
	}
	return skill_pro
}

func select_passion() Passion {
	file, err := ioutil.ReadFile("./yaml/passions.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var yaml_passions Yaml2Passion
	error := yaml.Unmarshal([]byte(file), &yaml_passions)
	if error != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	nof_passions := len(yaml_passions.Passions)
	char_passion := rand.Intn(nof_passions)

	char_passion = 0 //TODO: remove this line!

	return yaml_passions.Passions[char_passion]
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
		"pow": generate_modifier_stand(stand_ability_scores["pow"]),
		"pre": generate_modifier_stand(stand_ability_scores["pre"]),
		"dur": generate_modifier_stand(stand_ability_scores["dur"]),
		"ran": generate_modifier_stand(stand_ability_scores["ran"]),
		"spe": generate_modifier_stand(stand_ability_scores["spe"]),
		"ngy": generate_modifier_stand(stand_ability_scores["ngy"]),
	}
	return stand_ability_modifiers
}

func select_stand() Stand {
	file, err := ioutil.ReadFile("./yaml/stand_types.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var yaml_stands Yaml2Stand
	error := yaml.Unmarshal([]byte(file), &yaml_stands)
	if error != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	nof_stands := len(yaml_stands.Stands)
	stand_id := rand.Intn(nof_stands)

	stand_id = 1 //TODO: remove this line!

	return yaml_stands.Stands[stand_id]
}

func select_ability() string {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(6)
	ability_dices := [6]string{"0d0", "1d4", "1d6", "1d8", "1d10", "1d12"}

	return ability_dices[id]
}

func calculate_char_AC(am map[string]int) int {
	var a = [3]int{-1, -1, -1}
	a[0] = 10 + am["dex"] + am["wis"]
	a[1] = 10 + am["dex"] + am["con"]
	a[2] = 10 + am["wis"] + am["con"]

	highest := -99999
	for _, i := range a {
		if i > highest {
			highest = i
		}
	}

	return highest
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

func load_abilities() []Abilities {
	file, err := ioutil.ReadFile("./yaml/abilities.yaml")

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

func load_feats() []Feat {
	file, err := ioutil.ReadFile("./yaml/feats.yaml")

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

func add_feat(cs *Character_Sheet, feat_name string) bool {
	feats := cs.feats

	var feat Feat
	index_feat := 0

	if feat_name == "" {
		fmt.Printf("Adding random feat...\n")
		rand.Seed(time.Now().UnixNano())
		nof_feats := len(feats)
		index_feat = rand.Intn(nof_feats)
		feat = feats[index_feat]
		fmt.Printf("Adding random feat: %s\n", feat.Name)
	} else {
		fmt.Printf("Adding feat by name: %s\n", feat_name)
		for i := 0; i < len(feats); i++ {
			if feats[i].Name == feat_name {
				index_feat = i
				break
			}
		}
		fmt.Printf("Adding feat: %s\n", feat_name)
	}

	//check if feat is not already added
	if stringInSlice(feat.Name, cs.feats_list) {
		fmt.Printf("Feat: %s already in character, trying again...\n", feat.Name)
		return false
	}

	//check if prereq is possible
	prereq := feat.Prereq

	var pr []string // get list of prereqs
	if strings.Contains(prereq, "|") {
		pr = strings.Split(prereq, "|")
	} else {
		pr = append(pr, prereq)
	}

	allow_addition := false
	for i := 0; i < len(pr); i++ {
		preq := pr[i]
		allow_addition = (preq == "")
		if strings.Contains(preq, ">") {
			a := strings.Split(preq, ">")
			b, _ := strconv.Atoi(a[1])
			allow_addition = cs.char_ability_scores[a[0]] >= b
			break
		} else if strings.Contains(preq, "passion") {
			a := strings.Split(preq, "=")
			allow_addition = cs.passion.Name == a[1]
			break
		} else if strings.Contains(preq, "feat") {
			a := strings.Split(preq, "=")
			if strings.Contains(a[1], "!") {
				allow_addition = !stringInSlice(a[1], cs.feats_list)
				break
			}
			allow_addition = stringInSlice(a[1], cs.feats_list)
		}

	}

	if allow_addition {
		cs.feats_list = append(cs.feats_list, feat.Name)
		fmt.Printf("Added feat: %s\n", feat.Name)
		return true
	}

	return false
}

func improve_ability_score_random(cs *Character_Sheet) {
	picked := ""
	stop_loop := false
	for stop_loop != true {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(6)

		asl := []string{"cha", "con", "dex", "itl", "str", "wis"}
		abs := asl[index]
		if picked != abs {
			cs.char_ability_scores[abs] += 1
			fmt.Printf("Improved: %s\n", abs)

			if picked != "" {
				break
			}

			picked = abs
		}
	}
}

func improve_ability_score_optimized(cs *Character_Sheet) {
	//TODO: this one just copies random function for now, am lazy :)
	picked := ""
	stop_loop := false
	for stop_loop != true {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(6)

		asl := []string{"cha", "con", "dex", "itl", "str", "wis"}

		abs := asl[index]

		if picked != abs {
			cs.char_ability_scores[abs] += 1
			fmt.Printf("Improved: %s \n", abs)

			if picked != "" {
				break
			}

			picked = abs
		}
	}
}

func add_new_ability(cs *Character_Sheet, ability_name string) bool {

	if strings.Contains(ability_name, "|") {
		fmt.Printf("Found multiple abilities, picking random from: %s\n", ability_name)
		abilities := strings.Split(ability_name, "|")

		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(abilities))

		ability_name = abilities[index]
	}

	if strings.Contains(ability_name, "feat") {
		feat := strings.Split(ability_name, "=")[1]
		fmt.Printf("Instead of ability, adding feat: %s\n", feat)
		return add_feat(cs, feat)
	}

	fmt.Printf("Adding Ability: %s\n", ability_name)

	// check if feat is not already added
	if !stringInSlice(ability_name, cs.stand.Abilities) {
		fmt.Printf("Ability: %s - not configure for this stand, what went wrong???\n", ability_name)
		return false
	}

	//check if prereq is possible
	if !stringInSlice(cs.stand.Name, cs.abilities_list) {
		cs.abilities_list = append(cs.abilities_list, ability_name)
		fmt.Printf("Added ability: %s\n", ability_name)
		return true
	} else {
		fmt.Printf("Ability: %s already in sheet retrying...\n", ability_name)
		return false
	}
}

func add_class_features(cs *Character_Sheet) {
	// luc = level up chart
	luc := cs.stand.LevelChart
	level := cs.level - 1

	level_row := strings.Split(luc[level], ";")

	additions := strings.Split(level_row[3], "&")
	for i := 0; i < len(additions); i++ {
		addition := additions[i]
		succes := false
		for succes != true {
			if addition == "Ability_Score_Improvement" {
				fmt.Printf("Improving Ability Scores...\n")
				improve_ability_score_random(cs)
				succes = true
			} else if addition == "Custom_Ability" {
				fmt.Printf("Please discuss a self made ability with your dm\n")
				succes = true
			} else if strings.Contains(addition, "attack_dice_inc") {
				//TODO: increase the attack dice
				succes = true
			} else if strings.Contains(addition, "hit_rate_inc") {
				//TODO: increase hit rate
				//} else if strings.Contains(addition, "feat") {
				//TODO: add extra feat
				succes = true
			} else if strings.Contains(addition, "choose") {
				fmt.Printf("Choosing ability, from one of the following levels: %s\n", addition)

				options := strings.Split(addition, "=")[1]
				levels := strings.Split(options, "|")

				rand.Seed(time.Now().UnixNano())
				index := rand.Intn(len(levels))
				level_to_pick_from, _ := strconv.Atoi(levels[index])

				fmt.Printf("Picking Ability from level: %d\n", level_to_pick_from)
				tlr := strings.Split(cs.stand.LevelChart[level_to_pick_from-1], ";")

				ability := strings.Split(tlr[3], "|")

				rand.Seed(time.Now().UnixNano())
				index = rand.Intn(len(ability))

				fmt.Printf("Attempting to add ability: %s\n", ability[index])
				succes = add_new_ability(cs, ability[index])
			} else {
				succes = add_new_ability(cs, addition)
			}
		}
		generate_character_statistics(cs, false)
	}

	//if allow_addition {
	//	cs.feats_list = append(cs.feats_list, feat.Name)
	//}
}

func increase_hit_points(cs *Character_Sheet) {
	fmt.Printf("Increasing hit points of character!\n")

	fmt.Printf("Using hit dice: %s\n", cs.hit_dice)
	current_hp := cs.hit_points
	hp_dice := strings.Split(cs.hit_dice, "d")
	nof_dice, _ := strconv.Atoi(hp_dice[0])
	nof_faces, _ := strconv.Atoi(hp_dice[1])

	for i := 0; i < nof_dice; i++ {
		rand.Seed(time.Now().UnixNano())
		hp_increase := rand.Intn(nof_faces)

		if hp_increase == 0 {
			hp_increase = 1
		}

		fmt.Printf("Roll nr%d: %d\n", i+1, hp_increase)

		cs.hit_points += hp_increase
	}

	fmt.Printf("Adding the Constitution modifier of: %d\n", cs.char_ability_modifiers["con"])
	cs.hit_points += cs.char_ability_modifiers["con"]

	fmt.Printf("Changed hit points from: %d to: %d\n", current_hp, cs.hit_points)

}

func increase_stand_points(cs *Character_Sheet) {
	fmt.Printf("Increasing hit points of character!\n")

	fmt.Printf("Using level up dice: %s\n", cs.stand.OnLevelUp)
	level_up_dice := strings.Split(cs.stand.OnLevelUp, "d")
	nof_dice, _ := strconv.Atoi(level_up_dice[0])
	nof_faces, _ := strconv.Atoi(level_up_dice[1])

	total := 0
	for i := 0; i < nof_dice; i++ {
		rand.Seed(time.Now().UnixNano())
		add_this := rand.Intn(nof_faces)

		if add_this == 0 {
			add_this = 1
		}

		fmt.Printf("Roll nr%d: %d\n", i+1, add_this)

		total += add_this
	}

	fmt.Printf("Adding the level of the stand user: %d\n", cs.level)
	total += cs.level

	fmt.Printf("Total points to spend: %d\n", total)

	point_pool_max := total
	point_pool := total

	min := 1
	max := total

	values := [6]int{0, 0, 0, 0, 0, 0}
	for point_pool > 0 {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(5)
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(max) + min
		if point_pool >= value && values[i]+value <= max && values[i]+value >= min {
			point_pool -= value
			values[i] += value
		}
	}

	cs.stand_ability_scores["pow"] += values[0]
	cs.stand_ability_scores["pre"] += values[1]
	cs.stand_ability_scores["dur"] += values[2]
	cs.stand_ability_scores["ran"] += values[3]
	cs.stand_ability_scores["spe"] += values[4]
	cs.stand_ability_scores["ngy"] += values[5]

	result := 0
	for _, v := range values {
		result += v
	}

	if result > point_pool_max {
		log.Panic("point pool was depleted?")
	}

	cs.stand_ability_modifiers = generate_ability_modifiers_stand(cs.stand_ability_scores)
	fmt.Printf("NEW: Stand Ability Scores:\n")
	fmt.Printf("pow: %d   \t modifier: %d\n", cs.stand_ability_scores["pow"], cs.stand_ability_modifiers["pow"])
	fmt.Printf("pre: %d   \t modifier: %d\n", cs.stand_ability_scores["pre"], cs.stand_ability_modifiers["pre"])
	fmt.Printf("dur: %d   \t modifier: %d\n", cs.stand_ability_scores["dur"], cs.stand_ability_modifiers["dur"])
	fmt.Printf("ran: %d   \t modifier: %d\n", cs.stand_ability_scores["ran"], cs.stand_ability_modifiers["ran"])
	fmt.Printf("spe: %d   \t modifier: %d\n", cs.stand_ability_scores["spe"], cs.stand_ability_modifiers["spe"])
	fmt.Printf("ngy: %d   \t modifier: %d\n", cs.stand_ability_scores["ngy"], cs.stand_ability_modifiers["ngy"])
}

func level_up(cs *Character_Sheet) {
	fmt.Print("============= level up! ================\n")
	cs.level++

	// luc = level up chart
	luc := cs.stand.LevelChart
	level := cs.level - 1

	println(luc[level])

	level_row := strings.Split(luc[level], ";")
	cs.proficiency_bonus, _ = strconv.Atoi(level_row[1])
	nof_feat, _ := strconv.Atoi(level_row[2])

	fmt.Printf("level: %d | prof_bonus: %d | nof_feat: %s | class_features: %s | ab_dice: %s\n", cs.level, cs.proficiency_bonus, level_row[2], level_row[3], level_row[4])

	cs.skill_pro = generate_passion_proficiencies(cs.passion, cs.char_ability_modifiers, cs.proficiency_bonus)

	//add feats
	nof_feat_to_add := nof_feat - len(cs.feats_list)
	for i := 0; i < nof_feat_to_add; i++ {
		add_feat(cs, "")
	}

	// parse the table kinda :)
	if cs.level < 15 {
		add_class_features(cs)
	}

	// set new attack dice on level 11
	if cs.level == 11 {
		cs.stand.AttackDice = cs.stand.AttackDicePastLevelEleven
	}

	// increase hit points
	increase_hit_points(cs)

	// increase stand ability scores
	increase_stand_points(cs)

	//change ability dice
	ab_dice_new := level_row[4]
	ab_dice_old := cs.ability_dice
	nof_faces := strings.Split(cs.ability_dice, "d")[1]
	cs.ability_dice = strings.ReplaceAll(ab_dice_new, "x", nof_faces)
	fmt.Printf("Ability Dice Increased from: %s to: %s\n", ab_dice_old, cs.ability_dice)

	print_character_sheet(*cs)
}

func print_character_sheet(character_sheet Character_Sheet) {
	fmt.Printf("============= Character Sheet Level: %d ================\n", character_sheet.level)
	for i := 0; i < len(character_sheet.skill_pro); i++ {
		sp := character_sheet.skill_pro[i]
		fmt.Printf("p:%d - %s (%s): %d \n", sp.level, sp.skill_name, sp.stat_type, sp.value)
	}

	fmt.Printf("proficiency bonus: %d\n", character_sheet.proficiency_bonus)
	fmt.Printf("Passions: %s\n", character_sheet.passion.Name)
	fmt.Printf("Stand Type: %s\n", character_sheet.stand.Name)
	fmt.Printf("Ability Dice: %s\n", character_sheet.ability_dice)

	fmt.Printf("Ability Score:\n")
	fmt.Printf("cha: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["cha"], character_sheet.char_ability_modifiers["cha"])
	fmt.Printf("con: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["con"], character_sheet.char_ability_modifiers["con"])
	fmt.Printf("dex: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["dex"], character_sheet.char_ability_modifiers["dex"])
	fmt.Printf("itl: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["itl"], character_sheet.char_ability_modifiers["itl"])
	fmt.Printf("str: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["str"], character_sheet.char_ability_modifiers["str"])
	fmt.Printf("wis: %d   \t modifiers: %d \n", character_sheet.char_ability_scores["wis"], character_sheet.char_ability_modifiers["wis"])

	fmt.Printf("Saving Throws:\n")
	fmt.Printf("cha: %d   \t is_poficient: %d \n", character_sheet.saving_throws["cha"].value, character_sheet.saving_throws["cha"].level)
	fmt.Printf("con: %d   \t is_poficient: %d \n", character_sheet.saving_throws["con"].value, character_sheet.saving_throws["con"].level)
	fmt.Printf("dex: %d   \t is_poficient: %d \n", character_sheet.saving_throws["dex"].value, character_sheet.saving_throws["dex"].level)
	fmt.Printf("itl: %d   \t is_poficient: %d \n", character_sheet.saving_throws["itl"].value, character_sheet.saving_throws["itl"].level)
	fmt.Printf("str: %d   \t is_poficient: %d \n", character_sheet.saving_throws["str"].value, character_sheet.saving_throws["str"].level)
	fmt.Printf("wis: %d   \t is_poficient: %d \n", character_sheet.saving_throws["wis"].value, character_sheet.saving_throws["wis"].level)

	fmt.Printf("Stand Ability Scores:\n")
	fmt.Printf("pow: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["pow"], character_sheet.stand_ability_modifiers["pow"])
	fmt.Printf("pre: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["pre"], character_sheet.stand_ability_modifiers["pre"])
	fmt.Printf("dur: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["dur"], character_sheet.stand_ability_modifiers["dur"])
	fmt.Printf("ran: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["ran"], character_sheet.stand_ability_modifiers["ran"])
	fmt.Printf("spe: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["spe"], character_sheet.stand_ability_modifiers["spe"])
	fmt.Printf("ngy: %d   \t modifier: %d\n", character_sheet.stand_ability_scores["ngy"], character_sheet.stand_ability_modifiers["ngy"])

	fmt.Printf("char_AC: %d\n", character_sheet.armor_class)
	fmt.Printf("stand_AC: %d\n", character_sheet.stand_ac)
	fmt.Printf("stand_movement: %d\n", character_sheet.stand_movement)
	fmt.Printf("stand_attack_per_turn: %d\n", character_sheet.stand_attack_per_turn)
	fmt.Printf("stand_damage_reduction: %d\n", character_sheet.stand_damage_reduction)
	fmt.Printf("initiative: %d\n", character_sheet.initiative)
	fmt.Printf("hit_dice: %s\n", character_sheet.stand.AttackDice)
	fmt.Printf("stand_dc: %d\n", character_sheet.stand_dc)
	fmt.Printf("hit_points: %d\n", character_sheet.hit_points)
}

func generate_character_statistics(character_sheet *Character_Sheet, recalc_starters bool) {
	fmt.Printf("============= Generate Character Sheet ================\n")
	// Calculate generate ability modifiers again :)
	character_sheet.char_ability_modifiers = generate_ability_modifiers_character(character_sheet.char_ability_scores)
	character_sheet.saving_throws = generate_passion_saving_throws(character_sheet.passion, character_sheet.char_ability_modifiers, character_sheet.proficiency_bonus)

	//generate_character_skill_proficiencies()
	character_sheet.skill_pro = generate_passion_proficiencies(character_sheet.passion, character_sheet.char_ability_modifiers, character_sheet.proficiency_bonus)

	if recalc_starters == true {
		// Calculate stant ability scores
		character_sheet.stand_ability_scores = generate_ability_scores_stands(character_sheet.stand, character_sheet.char_ability_scores)
		// Calculate stant ability modifiers
		character_sheet.stand_ability_modifiers = generate_ability_modifiers_stand(character_sheet.stand_ability_scores)
	}

	// Calculate some small things like AC
	character_sheet.armor_class = calculate_char_AC(character_sheet.char_ability_modifiers)
	character_sheet.stand_ac = calculate_stand_AC(character_sheet.stand_ability_modifiers)
	character_sheet.stand_movement = character_sheet.stand_ability_modifiers["spe"] * 2
	character_sheet.stand_attack_per_turn = character_sheet.stand_ability_scores["spe"]/50 + 1
	character_sheet.stand_damage_reduction = character_sheet.stand_ability_scores["dur"] / 10
	character_sheet.initiative = character_sheet.char_ability_modifiers["dex"] + character_sheet.char_ability_modifiers["wis"]
	character_sheet.hit_dice = character_sheet.stand.AttackDice
	character_sheet.stand_dc = 8 + character_sheet.proficiency_bonus + character_sheet.char_ability_modifiers["cha"]

	if recalc_starters == true {
		// Calculate starting hit points
		character_sheet.hit_points = 10 + character_sheet.char_ability_modifiers["con"]
	}
}

func main() {
	// Create character sheet
	var character_sheet Character_Sheet
	character_sheet.level = 0
	// set base proficiency_bonus
	character_sheet.proficiency_bonus = 2

	////////////////////////////////////////////////////////////////////////////////////////
	//START OF RANDOM STUFFS
	////////////////////////////////////////////////////////////////////////////////////////

	// Roll and assign Stats
	character_sheet.char_ability_scores = generate_ability_scores_character()
	// pick the Characterʼs Passion
	character_sheet.passion = select_passion()
	// pick a class/stand type
	character_sheet.stand = select_stand()
	// pick ability
	character_sheet.ability_dice = select_ability()
	//load all feats
	character_sheet.feats = load_feats()
	//load all abilities
	character_sheet.abilities = load_abilities()
	// Calculate passion bonusses
	character_sheet.char_ability_scores = generate_passion_ability_score(character_sheet.passion, character_sheet.char_ability_scores)

	////////////////////////////////////////////////////////////////////////////////////////
	//END OF RANDOM STUFFS
	////////////////////////////////////////////////////////////////////////////////////////

	generate_character_statistics(&character_sheet, true)
	print_character_sheet(character_sheet)
	for i := 0; i < 20; i++ {
		level_up(&character_sheet)
	}

	// 4. Find what Class(es) your character will be playing, add the abilities
	// 5. Determine your characterʼs Maximum Hit Points, Armor Class, and Stand Armor Class (If Applicable)
	// 6. Talk with your DM regarding your Starting Equipment
	// 7. Find your Characterʼs Stand Score and Modifiers, as well as your Standʼs Ability, and how it works (Ignore this step if your character is not a Stand User)
}
