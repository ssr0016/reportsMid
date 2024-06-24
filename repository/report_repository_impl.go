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
			average_attendance
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
			average_attendance
		FROM reports
		WHERE 
			id = $1
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL, reportId)
	helper.ErrorPanic(errQuery)
	defer result.Close()

	report := model.Report{}
	var worshipServiceJSON []byte

	if result.Next() {
		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.WorkerName,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
			&report.CreatedAt,
			&report.UpdatedAt,
			&worshipServiceJSON,
			&report.AverageAttendance,
		)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.New("report not found")
			}

			return nil, err
		}

		// Unmarshal worshipServiceJSON into []int
		err = json.Unmarshal(worshipServiceJSON, &report.WorshipService)
		if err != nil {
			return nil, err
		}

	}

	return &report, nil
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
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
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
			average_attendance = $6,
			updated_at = $7
		WHERE 
			id = $8
	`

	worshipServiceJSON, err := json.Marshal(report.WorshipService)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.WorkerName,
		report.AreaOfAssignment,
		report.NameOfChurch,
		worshipServiceJSON,
		report.AverageAttendance,
		now,
		report.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
