package prototype

type Resume struct {
	name      string
	sex       string
	age       int
	education map[string]string
	work      []WorkExperience
}

type WorkExperience struct {
	TimeStart string
	TimeEnd   string
	Company   string
}

func NewResume(name, sex string, age int) *Resume {
	return &Resume{
		name:      name,
		sex:       sex,
		age:       age,
		education: make(map[string]string),
		work:      make([]WorkExperience, 0, 3),
	}
}

func (r *Resume) SetEducation(schoolName, schoolAddr string) {
	r.education["SchoolName"] = schoolName
	r.education["SchoolAddr"] = schoolAddr
}

func (r *Resume) AddWorkExperience(company, start, end string) {
	r.work = append(r.work, WorkExperience{
		TimeStart: start,
		TimeEnd:   end,
		Company:   company,
	})

}

func (r *Resume) Clone() *Resume {
	clone := *r

	clone.work = make([]WorkExperience, len(r.work))
	copy(clone.work, r.work)

	clone.education = make(map[string]string, len(r.education))
	for k, v := range r.education {
		clone.education[k] = v
	}

	return &clone

}
