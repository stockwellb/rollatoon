package rollatoon

type Race struct {
	Name    string
	Faction string
	Classes Classes
}

type Result struct {
	Race     string
	Class    string
	Greeting string
}

type Classes []string

var allianceRaces = []Race{
	Race{"Human", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Paladin", "Monk", "Death Knight"}},
	Race{"Dwarf", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Shaman", "Monk", "Death Knight"}},
	Race{"Night Elf", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Druid", "Monk", "Death Knight"}},
	Race{"Gnome", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Monk", "Death Knight"}},
	Race{"Draenei", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Paladin", "Monk", "Death Knight"}},
	Race{"Worgen", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Druid", "Death Knight"}},
	Race{"Pandaren", "Alliance", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Shaman", "Monk", "Death Knight"}},
}

var hordeRaces = []Race{
	Race{"Orc", "Horde", []string{"Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Shaman", "Monk", "Death Knight"}},
	Race{"Tauren", "Horde", []string{"Priest", "Warrior", "Druid", "Hunter", "Paladin", "Shaman", "Monk", "Death Knight"}},
	Race{"Goblin", "Horde", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Shaman", "Death Knight"}},
	Race{"Troll", "Horde", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Druid", "Shaman", "Monk", "Death Knight"}},
	Race{"Undead", "Horde", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Monk", "Death Knight"}},
	Race{"Blood Elf", "Horde", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Warlock", "Monk", "Death Knight"}},
	Race{"Pandaren", "Horde", []string{"Priest", "Rogue", "Warrior", "Mage", "Hunter", "Shaman", "Monk", "Death Knight"}},
}
