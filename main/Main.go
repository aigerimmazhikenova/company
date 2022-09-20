package main

func main() {
	hhKz := JobSite{
		subscribers: []Person{},
		vacancies:   []string{},
	}
	bob := newPerson("Bob")

	hhKz.addVacancy("Senior HTML Developer")
	hhKz.subscribe(bob)
	hhKz.addVacancy("Java Junior Developer")

	hhKz.removeVacancy("Senior HTML Developer")
}
