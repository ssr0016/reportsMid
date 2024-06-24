package response

import "time"

type ReportResponse struct {
	Id                       int       `json:"id"`
	MonthOf                  string    `json:"month_of"`
	WorkerName               string    `json:"worker_name"`
	AreaOfAssignment         string    `json:"area_of_assignment"`
	NameOfChurch             string    `json:"name_of_church"`
	WorshipService           []int     `json:"worship_service,omitempty"`
	SundaySchool             []int     `json:"sunday_school,omitempty"`
	PrayerMeetings           []int     `json:"prayer_meetings,omitempty"`
	BibleStudies             []int     `json:"bible_studies,omitempty"`
	MensFellowships          []int     `json:"mens_fellowships,omitempty"`
	WomensFellowships        []int     `json:"womens_fellowships,omitempty"`
	YouthFellowships         []int     `json:"youth_fellowships,omitempty"`
	ChildFellowships         []int     `json:"child_fellowships,omitempty"`
	Outreach                 []int     `json:"outreach,omitempty"`
	TrainingOrSeminars       []int     `json:"training_or_seminars,omitempty"`
	LeadershipConferences    []int     `json:"leadership_conferences,omitempty"`
	LeadershipTraining       []int     `json:"leadership_training,omitempty"`
	Others                   []int     `json:"others,omitempty"`
	FamilyDays               []int     `json:"family_days,omitempty"`
	TithesAndOfferings       []int     `json:"tithes_and_offerings,omitempty"`
	WorshipServiceAvg        float64   `json:"worship_service_average"`
	SundaySchoolAvg          float64   `json:"sunday_school_average"`
	PrayerMeetingsAvg        float64   `json:"prayer_meetings_average"`
	BibleStudiesAvg          float64   `json:"bible_studies_average"`
	MensFellowshipsAvg       float64   `json:"mens_fellowships_average"`
	WomensFellowshipsAvg     float64   `json:"womens_fellowships_average"`
	YouthFellowshipsAvg      float64   `json:"youth_fellowships_average"`
	ChildFellowshipsAvg      float64   `json:"child_fellowships_average"`
	OutreachAvg              float64   `json:"outreach_average"`
	TrainingOrSeminarsAvg    float64   `json:"training_or_seminars_average"`
	LeadershipConferencesAvg float64   `json:"leadership_conferences_average"`
	LeadershipTrainingAvg    float64   `json:"leadership_training_average"`
	OthersAvg                float64   `json:"others_average"`
	FamilyDaysAvg            float64   `json:"family_days_average"`
	TithesAndOfferingsAvg    float64   `json:"tithes_and_offerings_average"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}
