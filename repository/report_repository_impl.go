package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"reports/helper"
	"reports/model"
	"time"
)

type ReportRepositoryImpl struct {
	Db *sql.DB
}

func NewReportRepository(Db *sql.DB) ReportRepository {
	return &ReportRepositoryImpl{Db: Db}
}

// Delete implements BookRepository
func (r *ReportRepositoryImpl) Delete(ctx context.Context, reportId int) error {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		DELETE FROM reports
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, rawSQL, reportId)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements BookRepository
func (r *ReportRepositoryImpl) FindAll(ctx context.Context) ([]model.Report, error) {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
        SELECT 
            id,
            month_of,
            worker_name,
            area_of_assignment,
            name_of_church,
            created_at,
            updated_at,
            worship_service,
            sunday_school,
            prayer_meetings,
            bible_studies,
            mens_fellowships,
            womens_fellowships,
            youth_fellowships,
            child_fellowships,
            outreach,
            training_or_seminars,
            leadership_conferences,
            leadership_training,
            others,
            family_days,
            tithes_and_offerings,
            average_attendance,
            home_visited,
            bible_study_or_group_led,
            sermon_or_message_preached,
            person_newly_contacted,
            person_followed_up,
            person_led_to_christ,
            names,
			narrative_report,
			challenges_and_problem_encountered,
			prayer_request
        FROM reports
    `
	result, err := tx.QueryContext(ctx, rawSQL)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var reports []model.Report

	for result.Next() {
		var report model.Report
		var (
			worshipServiceJSON        []byte
			sundaySchoolJSON          []byte
			prayerMeetingsJSON        []byte
			bibleStudiesJSON          []byte
			mensFellowshipsJSON       []byte
			womensFellowshipsJSON     []byte
			youthFellowshipsJSON      []byte
			childFellowshipsJSON      []byte
			outreachJSON              []byte
			trainingOrSeminarsJSON    []byte
			leadershipConferencesJSON []byte
			leadershipTrainingJSON    []byte
			othersJSON                []byte
			familyDaysJSON            []byte
			tithesAndOfferingsJSON    []byte
			homeVisitedJSON           []byte
			bibleStudyOrGroupLedJSON  []byte
			sermonOrMessageJSON       []byte
			personNewlyContactedJSON  []byte
			personFollowedUpJSON      []byte
			personLedToChristJSON     []byte
			namesJSON                 []byte
		)

		// Scan the row into variables
		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.WorkerName,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
			&report.CreatedAt,
			&report.UpdatedAt,
			&worshipServiceJSON,
			&sundaySchoolJSON,
			&prayerMeetingsJSON,
			&bibleStudiesJSON,
			&mensFellowshipsJSON,
			&womensFellowshipsJSON,
			&youthFellowshipsJSON,
			&childFellowshipsJSON,
			&outreachJSON,
			&trainingOrSeminarsJSON,
			&leadershipConferencesJSON,
			&leadershipTrainingJSON,
			&othersJSON,
			&familyDaysJSON,
			&tithesAndOfferingsJSON,
			&report.AverageAttendance,
			&homeVisitedJSON,
			&bibleStudyOrGroupLedJSON,
			&sermonOrMessageJSON,
			&personNewlyContactedJSON,
			&personFollowedUpJSON,
			&personLedToChristJSON,
			&namesJSON,
			&report.NarrativeReport,
			&report.ChallengesAndProblemEncountered,
			&report.PrayerRequest,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSONB fields into their respective slices
		if worshipServiceJSON != nil {
			if err := json.Unmarshal(worshipServiceJSON, &report.WorshipService); err != nil {
				return nil, err
			}
		}

		if sundaySchoolJSON != nil {
			if err := json.Unmarshal(sundaySchoolJSON, &report.SundaySchool); err != nil {
				return nil, err
			}
		}

		if prayerMeetingsJSON != nil {
			if err := json.Unmarshal(prayerMeetingsJSON, &report.PrayerMeetings); err != nil {
				return nil, err
			}
		}

		if bibleStudiesJSON != nil {
			if err := json.Unmarshal(bibleStudiesJSON, &report.BibleStudies); err != nil {
				return nil, err
			}
		}

		if mensFellowshipsJSON != nil {
			if err := json.Unmarshal(mensFellowshipsJSON, &report.MensFellowships); err != nil {
				return nil, err
			}
		}

		if womensFellowshipsJSON != nil {
			if err := json.Unmarshal(womensFellowshipsJSON, &report.WomensFellowships); err != nil {
				return nil, err
			}
		}

		if youthFellowshipsJSON != nil {
			if err := json.Unmarshal(youthFellowshipsJSON, &report.YouthFellowships); err != nil {
				return nil, err
			}
		}

		if childFellowshipsJSON != nil {
			if err := json.Unmarshal(childFellowshipsJSON, &report.ChildFellowships); err != nil {
				return nil, err
			}
		}

		if outreachJSON != nil {
			if err := json.Unmarshal(outreachJSON, &report.Outreach); err != nil {
				return nil, err
			}
		}

		if trainingOrSeminarsJSON != nil {
			if err := json.Unmarshal(trainingOrSeminarsJSON, &report.TrainingOrSeminars); err != nil {
				return nil, err
			}
		}

		if leadershipConferencesJSON != nil {
			if err := json.Unmarshal(leadershipConferencesJSON, &report.LeadershipConferences); err != nil {
				return nil, err
			}
		}

		if leadershipTrainingJSON != nil {
			if err := json.Unmarshal(leadershipTrainingJSON, &report.LeadershipTraining); err != nil {
				return nil, err
			}
		}

		if othersJSON != nil {
			if err := json.Unmarshal(othersJSON, &report.Others); err != nil {
				return nil, err
			}
		}

		if familyDaysJSON != nil {
			if err := json.Unmarshal(familyDaysJSON, &report.FamilyDays); err != nil {
				return nil, err
			}
		}

		if tithesAndOfferingsJSON != nil {
			if err := json.Unmarshal(tithesAndOfferingsJSON, &report.TithesAndOfferings); err != nil {
				return nil, err
			}
		}

		// Unmarshal JSONB fields into []int for additional fields
		if homeVisitedJSON != nil {
			if err := json.Unmarshal(homeVisitedJSON, &report.HomeVisited); err != nil {
				return nil, err
			}
		}

		if bibleStudyOrGroupLedJSON != nil {
			if err := json.Unmarshal(bibleStudyOrGroupLedJSON, &report.BibleStudyOrGroupLed); err != nil {
				return nil, err
			}
		}

		if sermonOrMessageJSON != nil {
			if err := json.Unmarshal(sermonOrMessageJSON, &report.SermonOrMessagePreached); err != nil {
				return nil, err
			}
		}

		if personNewlyContactedJSON != nil {
			if err := json.Unmarshal(personNewlyContactedJSON, &report.PersonNewlyContacted); err != nil {
				return nil, err
			}
		}

		if personFollowedUpJSON != nil {
			if err := json.Unmarshal(personFollowedUpJSON, &report.PersonFollowedUp); err != nil {
				return nil, err
			}
		}

		if personLedToChristJSON != nil {
			if err := json.Unmarshal(personLedToChristJSON, &report.PersonLedToChrist); err != nil {
				return nil, err
			}
		}

		if namesJSON != nil {
			if err := json.Unmarshal(namesJSON, &report.Names); err != nil {
				return nil, err
			}
		}

		// Append the populated report to the slice
		reports = append(reports, report)
	}

	// Check for any error during result iteration
	if err := result.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

// FindById implements BookRepository
func (r *ReportRepositoryImpl) FindById(ctx context.Context, reportId int) (*model.Report, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT 
			id,
			month_of,
			worker_name,
			area_of_assignment,
			name_of_church,
			created_at,
			updated_at,
			worship_service,
			sunday_school,
			prayer_meetings,
			bible_studies,
			mens_fellowships,
			womens_fellowships,
			youth_fellowships,
			child_fellowships,
			outreach,
			training_or_seminars,
			leadership_conferences,
			leadership_training,
			others,
			family_days,
			tithes_and_offerings,
			average_attendance,
			home_visited,
			bible_study_or_group_led,
			sermon_or_message_preached,
			person_newly_contacted,
			person_followed_up,
			person_led_to_christ,
			names,
			narrative_report,
			challenges_and_problem_encountered,
			prayer_request
		FROM reports
		WHERE 
			id = $1
	`

	// Query the database
	row := tx.QueryRowContext(ctx, rawSQL, reportId)

	// Initialize a new Report struct to hold the retrieved data
	report := &model.Report{}

	// Variables to hold JSONB data
	var (
		worshipServiceJSON              []byte
		sundaySchoolJSON                []byte
		prayerMeetingsJSON              []byte
		bibleStudiesJSON                []byte
		mensFellowshipsJSON             []byte
		womensFellowshipsJSON           []byte
		youthFellowshipsJSON            []byte
		childFellowshipsJSON            []byte
		outreachJSON                    []byte
		trainingOrSeminarsJSON          []byte
		leadershipConferencesJSON       []byte
		leadershipTrainingJSON          []byte
		othersJSON                      []byte
		familyDaysJSON                  []byte
		tithesAndOfferingsJSON          []byte
		homeVisitedJSON                 []byte
		bibleStudyOrGroupLedJSON        []byte
		sermonOrMessagePreachedJSON     []byte
		personNewlyContactedJSON        []byte
		personFollowedUpJSON            []byte
		personLedToChristJSON           []byte
		namesJSON                       []byte
		narrativeReportString           sql.NullString
		challengesAndProblemEncountered sql.NullString
		prayerRequestString             sql.NullString
	)

	// Scan the row into variables
	err = row.Scan(
		&report.Id,
		&report.MonthOf,
		&report.WorkerName,
		&report.AreaOfAssignment,
		&report.NameOfChurch,
		&report.CreatedAt,
		&report.UpdatedAt,
		&worshipServiceJSON,
		&sundaySchoolJSON,
		&prayerMeetingsJSON,
		&bibleStudiesJSON,
		&mensFellowshipsJSON,
		&womensFellowshipsJSON,
		&youthFellowshipsJSON,
		&childFellowshipsJSON,
		&outreachJSON,
		&trainingOrSeminarsJSON,
		&leadershipConferencesJSON,
		&leadershipTrainingJSON,
		&othersJSON,
		&familyDaysJSON,
		&tithesAndOfferingsJSON,
		&report.AverageAttendance,
		&homeVisitedJSON,
		&bibleStudyOrGroupLedJSON,
		&sermonOrMessagePreachedJSON,
		&personNewlyContactedJSON,
		&personFollowedUpJSON,
		&personLedToChristJSON,
		&namesJSON,
		&narrativeReportString,
		&challengesAndProblemEncountered,
		&prayerRequestString,
	)

	// Handle potential errors from scanning
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("report not found")
		}
		return nil, err
	}

	// Unmarshal JSONB data into respective fields
	if err := json.Unmarshal(worshipServiceJSON, &report.WorshipService); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(sundaySchoolJSON, &report.SundaySchool); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(prayerMeetingsJSON, &report.PrayerMeetings); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bibleStudiesJSON, &report.BibleStudies); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mensFellowshipsJSON, &report.MensFellowships); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(womensFellowshipsJSON, &report.WomensFellowships); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(youthFellowshipsJSON, &report.YouthFellowships); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(childFellowshipsJSON, &report.ChildFellowships); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(outreachJSON, &report.Outreach); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(trainingOrSeminarsJSON, &report.TrainingOrSeminars); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(leadershipConferencesJSON, &report.LeadershipConferences); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(leadershipTrainingJSON, &report.LeadershipTraining); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(othersJSON, &report.Others); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(familyDaysJSON, &report.FamilyDays); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(tithesAndOfferingsJSON, &report.TithesAndOfferings); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(homeVisitedJSON, &report.HomeVisited); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bibleStudyOrGroupLedJSON, &report.BibleStudyOrGroupLed); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(sermonOrMessagePreachedJSON, &report.SermonOrMessagePreached); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(personNewlyContactedJSON, &report.PersonNewlyContacted); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(personFollowedUpJSON, &report.PersonFollowedUp); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(personLedToChristJSON, &report.PersonLedToChrist); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(namesJSON, &report.Names); err != nil {
		return nil, err
	}

	// Handle nullable strings
	if narrativeReportString.Valid {
		report.NarrativeReport = narrativeReportString.String
	}
	if challengesAndProblemEncountered.Valid {
		report.ChallengesAndProblemEncountered = challengesAndProblemEncountered.String
	}
	if prayerRequestString.Valid {
		report.PrayerRequest = prayerRequestString.String
	}

	return report, nil
}

// Save implements BookRepository
func (r *ReportRepositoryImpl) Save(ctx context.Context, report *model.Report) error {
	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	// Marshal arrays to JSONB
	worshipServiceJSON, err := json.Marshal(report.WorshipService)
	if err != nil {
		return err
	}

	sundaySchoolJSON, err := json.Marshal(report.SundaySchool)
	if err != nil {
		return err
	}

	prayerMeetingsJSON, err := json.Marshal(report.PrayerMeetings)
	if err != nil {
		return err
	}

	bibleStudiesJSON, err := json.Marshal(report.BibleStudies)
	if err != nil {
		return err
	}

	mensFellowshipsJSON, err := json.Marshal(report.MensFellowships)
	if err != nil {
		return err
	}

	womensFellowshipsJSON, err := json.Marshal(report.WomensFellowships)
	if err != nil {
		return err
	}

	youthFellowshipsJSON, err := json.Marshal(report.YouthFellowships)
	if err != nil {
		return err
	}

	childFellowshipsJSON, err := json.Marshal(report.ChildFellowships)
	if err != nil {
		return err
	}

	outreachJSON, err := json.Marshal(report.Outreach)
	if err != nil {
		return err
	}

	trainingOrSeminarsJSON, err := json.Marshal(report.TrainingOrSeminars)
	if err != nil {
		return err
	}

	leadershipConferencesJSON, err := json.Marshal(report.LeadershipConferences)
	if err != nil {
		return err
	}

	leadershipTrainingJSON, err := json.Marshal(report.LeadershipTraining)
	if err != nil {
		return err
	}

	othersJSON, err := json.Marshal(report.Others)
	if err != nil {
		return err
	}

	familyDaysJSON, err := json.Marshal(report.FamilyDays)
	if err != nil {
		return err
	}

	tithesAndOfferingsJSON, err := json.Marshal(report.TithesAndOfferings)
	if err != nil {
		return err
	}

	homeVisitedJSON, err := json.Marshal(report.HomeVisited)
	if err != nil {
		return err
	}

	bibleStudyOrGroupLedJSON, err := json.Marshal(report.BibleStudyOrGroupLed)
	if err != nil {
		return err
	}

	sermonOrMessagePreachedJSON, err := json.Marshal(report.SermonOrMessagePreached)
	if err != nil {
		return err
	}

	personNewlyContactedJSON, err := json.Marshal(report.PersonNewlyContacted)
	if err != nil {
		return err
	}

	personFollowedUpJSON, err := json.Marshal(report.PersonFollowedUp)
	if err != nil {
		return err
	}

	personLedToChristJSON, err := json.Marshal(report.PersonLedToChrist)
	if err != nil {
		return err
	}

	namesJSON, err := json.Marshal(report.Names)
	if err != nil {
		return err
	}

	rawSQL := `
		INSERT INTO reports (
			month_of,
			worker_name,
			area_of_assignment,
			name_of_church,
			worship_service,
			sunday_school,
			prayer_meetings,
			bible_studies,
			mens_fellowships,
			womens_fellowships,
			youth_fellowships,
			child_fellowships,
			outreach,
			training_or_seminars,
			leadership_conferences,
			leadership_training,
			others,
			family_days,
			tithes_and_offerings,
			average_attendance,
			home_visited,
			bible_study_or_group_led,
			sermon_or_message_preached,
			person_newly_contacted,
			person_followed_up,
			person_led_to_christ,
			names,
			narrative_report,
			challenges_and_problem_encountered,
			prayer_request,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32)
	`

	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.WorkerName,
		report.AreaOfAssignment,
		report.NameOfChurch,
		worshipServiceJSON,
		sundaySchoolJSON,
		prayerMeetingsJSON,
		bibleStudiesJSON,
		mensFellowshipsJSON,
		womensFellowshipsJSON,
		youthFellowshipsJSON,
		childFellowshipsJSON,
		outreachJSON,
		trainingOrSeminarsJSON,
		leadershipConferencesJSON,
		leadershipTrainingJSON,
		othersJSON,
		familyDaysJSON,
		tithesAndOfferingsJSON,
		report.AverageAttendance,
		homeVisitedJSON,
		bibleStudyOrGroupLedJSON,
		sermonOrMessagePreachedJSON,
		personNewlyContactedJSON,
		personFollowedUpJSON,
		personLedToChristJSON,
		namesJSON,
		report.NarrativeReport,
		report.ChallengesAndProblemEncountered,
		report.PrayerRequest,
		report.CreatedAt,
		report.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// Update implements BookRepository
func (r *ReportRepositoryImpl) Update(ctx context.Context, report *model.Report) error {
	tx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		helper.ErrorPanic(err)
	}

	now := time.Now().In(loc)

	rawSQL := `
		UPDATE reports SET
			month_of = $1,
			worker_name = $2,
			area_of_assignment = $3,
			name_of_church = $4,
			worship_service = $5,
			sunday_school = $6,
			prayer_meetings = $7,
			bible_studies = $8,
			mens_fellowships = $9,
			womens_fellowships = $10,
			youth_fellowships = $11,
			child_fellowships = $12,
			outreach = $13,
			training_or_seminars = $14,
			leadership_conferences = $15,
			leadership_training	= $16,
			others = $17,
			family_days = $18,
			tithes_and_offerings = $19,
			average_attendance = $20,
			home_visited = $21,
			bible_study_or_group_led = $22,
			sermon_or_message_preached = $23,
			person_newly_contacted = $24,
			person_followed_up = $25,
			person_led_to_christ = $26,
			names = $27,
			narrative_report = $28,
			challenges_and_problem_encountered = $29,
			prayer_request = $30,
			updated_at = $31
		WHERE 
			id = $32
	`

	// Marshal arrays to JSON
	worshipServiceJSON, err := json.Marshal(report.WorshipService)
	if err != nil {
		return err
	}
	sundaySchoolJSON, err := json.Marshal(report.SundaySchool)
	if err != nil {
		return err
	}
	prayerMeetingsJSON, err := json.Marshal(report.PrayerMeetings)
	if err != nil {
		return err
	}
	bibleStudiesJSON, err := json.Marshal(report.BibleStudies)
	if err != nil {
		return err
	}
	mensFellowshipsJSON, err := json.Marshal(report.MensFellowships)
	if err != nil {
		return err
	}
	womensFellowshipsJSON, err := json.Marshal(report.WomensFellowships)
	if err != nil {
		return err
	}
	youthFellowshipsJSON, err := json.Marshal(report.YouthFellowships)
	if err != nil {
		return err
	}
	childFellowshipsJSON, err := json.Marshal(report.ChildFellowships)
	if err != nil {
		return err
	}
	outreachJSON, err := json.Marshal(report.Outreach)
	if err != nil {
		return err
	}
	trainingOrSeminarsJSON, err := json.Marshal(report.TrainingOrSeminars)
	if err != nil {
		return err
	}
	leadershipConferencesJSON, err := json.Marshal(report.LeadershipConferences)
	if err != nil {
		return err
	}
	leadershipTrainingJSON, err := json.Marshal(report.LeadershipTraining)
	if err != nil {
		return err
	}
	othersJSON, err := json.Marshal(report.Others)
	if err != nil {
		return err
	}
	familyDaysJSON, err := json.Marshal(report.FamilyDays)
	if err != nil {
		return err
	}
	tithesAndOfferingsJSON, err := json.Marshal(report.TithesAndOfferings)
	if err != nil {
		return err
	}
	homeVisitedJSON, err := json.Marshal(report.HomeVisited)
	if err != nil {
		return err
	}
	bibleStudyOrGroupLedJSON, err := json.Marshal(report.BibleStudyOrGroupLed)
	if err != nil {
		return err
	}
	sermonOrMessagePreachedJSON, err := json.Marshal(report.SermonOrMessagePreached)
	if err != nil {
		return err
	}
	personNewlyContactedJSON, err := json.Marshal(report.PersonNewlyContacted)
	if err != nil {
		return err
	}
	personFollowedUpJSON, err := json.Marshal(report.PersonFollowedUp)
	if err != nil {
		return err
	}
	personLedToChristJSON, err := json.Marshal(report.PersonLedToChrist)
	if err != nil {
		return err
	}
	namesJSON, err := json.Marshal(report.Names)
	if err != nil {
		return err
	}

	// Execute the update query
	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.WorkerName,
		report.AreaOfAssignment,
		report.NameOfChurch,
		worshipServiceJSON,
		sundaySchoolJSON,
		prayerMeetingsJSON,
		bibleStudiesJSON,
		mensFellowshipsJSON,
		womensFellowshipsJSON,
		youthFellowshipsJSON,
		childFellowshipsJSON,
		outreachJSON,
		trainingOrSeminarsJSON,
		leadershipConferencesJSON,
		leadershipTrainingJSON,
		othersJSON,
		familyDaysJSON,
		tithesAndOfferingsJSON,
		report.AverageAttendance,
		homeVisitedJSON,
		bibleStudyOrGroupLedJSON,
		sermonOrMessagePreachedJSON,
		personNewlyContactedJSON,
		personFollowedUpJSON,
		personLedToChristJSON,
		namesJSON,
		report.NarrativeReport,
		report.ChallengesAndProblemEncountered,
		report.PrayerRequest,
		now,
		report.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
