package rollatoon

import (
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/index.html",
))

var about = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/about.html",
))

func aboutHandler(rw http.ResponseWriter, req *http.Request) {
	if err := about.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
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
