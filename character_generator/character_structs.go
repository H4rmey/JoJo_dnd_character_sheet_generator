package character_generator

type Proficiency struct {
	Level      int    `yaml:"level"`
	Value      int    `yaml:"value"`
	Stat_type  string `yaml:"stat_type"`
	Skill_name string `yaml:"skill_name"`
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
	Name                    string                 `yaml:"name"`
	Ability_dice            string                 `yaml:"ability_dice"`
	Hit_dice                string                 `yaml:"hit_dice"`
	Proficiency_bonus       int                    `yaml:"proficiency_bonus"`
	Stand_damage_reduction  int                    `yaml:"stand_damage_reduction"`
	Stand_movement          int                    `yaml:"stand_movment"`
	Stand_attack_per_turn   int                    `yaml:"stand_attack_per_turn"`
	Stand_dc                int                    `yaml:"stand_dc"`
	Stand_ac                int                    `yaml:"stand_ac"`
	Hit_points              int                    `yaml:"hit_points"`
	Armor_class             int                    `yaml:"armor_class"`
	Initiative              int                    `yaml:"initiative"`
	Passive_perception      int                    `yaml:"passive_perception"`
	Movement                int                    `yaml:"movement"`
	Level                   int                    `yaml:"level"`
	Feats_list              []string               `yaml:"feats_list"`
	Abilities_list          []string               `yaml:"abilities_list"`
	Stand                   Stand                  `yaml:"stand"`
	Passion                 Passion                `yaml:"passion"`
	Skill_pro               []Proficiency          `yaml:"skill_pro"`
	Saving_throws           map[string]Proficiency `yaml:"saving_throws"`
	Char_ability_scores     map[string]int         `yaml:"char_ability_scores"`
	Char_ability_modifiers  map[string]int         `yaml:"char_ability_modifiers"`
	Stand_ability_scores    map[string]int         `yaml:"stand_ability_scores"`
	Stand_ability_modifiers map[string]int         `yaml:"stand_ability_modifiers"`
	Hit_points_rolls        []int                  `yaml:"hit_points_rolls"`
	Feats                   []Feat                 `yaml:"feats"`
	Abilities               []Abilities            `yaml:"abilities"`
	Stands                  []Stand                `yaml:"stands"`
	Passions                []Passion              `yaml:"passions"`
}

type Export_Character_Details struct {
	// Base
	Name      string `yaml:"name"`
	Languages string `yaml:"languages"`
	Age       string `yaml:"age"`
	Height    string `yaml:"height"`
	Weight    string `yaml:"weight"`
	SkinTone  string `yaml:"skin_tone"`
	EyeColor  string `yaml:"eye_color"`
	HairColor string `yaml:"hair_color"`
	// Details
	Description       string `yaml:"description"`
	Backstory         string `yaml:"backstory"`
	Relations         string `yaml:"relations"`
	Inventory         string `yaml:"inventory"`
	Weapons           string `yaml:"Weapons"`
	Image             string `yaml:"image_path"`
	Flaws             string `yaml:"flaws"`
	Ideals            string `yaml:"ideals"`
	Personality       string `yaml:"personality"`
	AttackDescription string `yaml:"attack_description"`
	ExtraInfo         string `yaml:"extra_info"`
}

// Character_Sheet
type Export_Character_Sheet struct {
	Name              string                   `yaml:"name"`
	Passion           string                   `yaml:"passion"`
	Standtype         string                   `yaml:"stand_type"`
	HitPoints         int                      `yaml:"hit_points"`
	Level             int                      `yaml:"level"`
	HitDice           string                   `yaml:"hit_dice"`
	Movement          int                      `yaml:"movement"`
	Proficiency_bonus int                      `yaml:"proficiency_bonus"`
	ArmorClass        int                      `yaml:"armor_class"`
	Initiative        int                      `yaml:"initiative"`
	Perception        int                      `yaml:"passive_perception"`
	Feats_list        []string                 `yaml:"feats_list"`
	Abilities_list    []string                 `yaml:"abilities_list"`
	Stand             Export_Stand             `yaml:"stand"`
	SavingThrows      map[string]Proficiency   `yaml:"saving_throws"`
	AbilityScores     map[string]int           `yaml:"char_ability_scores"`
	AbilityModifiers  map[string]int           `yaml:"char_ability_modifiers"`
	Skill_pro         []Proficiency            `yaml:"skill_pro"`
	CharDetails       Export_Character_Details `yaml:"character_details"`
}

type Export_Stand struct {
	Name                      string         `yaml:"name"`
	AbilityDice               string         `yaml:"ability_dice"`
	DamageReduction           int            `yaml:"damage_reduction"`
	Movement                  int            `yaml:"movement"`
	AttackPerTurn             int            `yaml:"attack_per_turn"`
	DC                        int            `yaml:"stand_dc"`
	AC                        int            `yaml:"stand_ac"`
	AttackDice                string         `yaml:"attack_dice"`
	AttackDiceHigherLevel     string         `yaml:"attack_dice_higher_level"`
	AttackDicePastLevelEleven string         `yaml:"attack_dice_past_level_eleven"`
	HitDice                   string         `yaml:"hit_dice"`
	OnLevelUp                 string         `yaml:"on_level_up"`
	Description               string         `yaml:"description"`
	Note                      string         `yaml:"note"`
	AbilityScores             map[string]int `yaml:"ability_scores"`
	AbilityModifiers          map[string]int `yaml:"ability_modifiers"`
}
