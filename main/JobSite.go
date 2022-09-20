package main

type JobSite struct {
	subscribers []Person
	vacancies   []string
}

func (job *JobSite) addVacancy(vacancy string) {
	job.vacancies = append(job.vacancies, vacancy)
	job.sendAll()
}

func (job *JobSite) removeVacancy(vacancy string) {
	tempVacancies := make([]string, 0)
	for _, value := range job.vacancies {
		if value != vacancy {
			tempVacancies = append(tempVacancies, value)
		}
	}
	job.vacancies = tempVacancies
	job.sendAll()
}

func (job *JobSite) sendAll() {
	for _, value := range job.subscribers {
		value.handleEvent(job.vacancies)
	}
}

func (job *JobSite) subscribe(p *Person) {
	job.subscribers = append(job.subscribers, *p)
}

func (job *JobSite) unsubscribe(p Person) {
	tempSubscribers := make([]Person, 0)
	for _, value := range job.subscribers {
		if value != p {
			tempSubscribers = append(tempSubscribers, value)
		}
	}
	job.subscribers = tempSubscribers
}
