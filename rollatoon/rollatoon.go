package rollatoon

import (
	"github.com/gorilla/mux"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

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

var index = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/index.html",
))

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func allianceHandler(rw http.ResponseWriter, req *http.Request) {

	var race = allianceRaces[randomNumber(0, len(allianceRaces))]
	var class = race.Classes[randomNumber(0, len(race.Classes))]
	var result = Result{
		race.Name,
		class,
		"Well met, I have chosen a random alliance race/class for you.",
	}

	if err := index.Execute(rw, &result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func wowHandler(rw http.ResponseWriter, req *http.Request) {
	var all = make([]Race, len(hordeRaces)+len(allianceRaces))
	copy(all, allianceRaces)
	copy(all[len(allianceRaces):], hordeRaces)
	var race = all[randomNumber(0, len(all))]
	var class = race.Classes[randomNumber(0, len(race.Classes))]
	var result = Result{
		race.Name,
		class,
		"Greetings, I have chosen a random WoW race/class for you.",
	}

	if err := index.Execute(rw, &result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func hordeHandler(rw http.ResponseWriter, req *http.Request) {

	var race = hordeRaces[randomNumber(0, len(hordeRaces))]
	var class = race.Classes[randomNumber(0, len(race.Classes))]
	var result = Result{
		race.Name,
		class,
		"Mok'ra, I have chosen a random horde race/class for you.",
	}

	if err := index.Execute(rw, &result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", wowHandler).
		Methods("GET")
	r.HandleFunc("/horde", hordeHandler).
		Methods("GET")
	r.HandleFunc("/alliance", allianceHandler).
		Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":1968", nil)
}
