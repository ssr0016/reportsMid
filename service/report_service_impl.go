package service

import (
	"context"
	"fmt"
	"reports/data/request"
	"reports/data/response"
	"reports/helper"
	"reports/model"
	"reports/repository"
	"time"
)

type ReportServiceImpl struct {
	reportRepository repository.ReportRepository
}

func NewReportServiceImpl(reportRepository repository.ReportRepository) ReportService {
	return &ReportServiceImpl{reportRepository: reportRepository}
}

func (r *ReportServiceImpl) Create(ctx context.Context, request *request.ReportCreateRequest) error {
	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)

	report := model.Report{
		MonthOf:               request.MonthOf,
		WorkerName:            request.WorkerName,
		AreaOfAssignment:      request.AreaOfAssignment,
		NameOfChurch:          request.NameOfChurch,
		WorshipService:        request.WorshipService,
		SundaySchool:          request.SundaySchool,
		PrayerMeetings:        request.PrayerMeetings,
		BibleStudies:          request.BibleStudies,
		MensFellowships:       request.MensFellowships,
		WomensFellowships:     request.WomensFellowships,
		YouthFellowships:      request.YouthFellowships,
		ChildFellowships:      request.ChildFellowships,
		Outreach:              request.Outreach,
		TrainingOrSeminars:    request.TrainingOrSeminars,
		LeadershipConferences: request.LeadershipConferences,
		LeadershipTraining:    request.LeadershipTraining,
		Others:                request.Others,
		FamilyDays:            request.FamilyDays,
		TithesAndOfferings:    request.TithesAndOfferings,
		AverageAttendance:     request.AverageAttendance,
		CreatedAt:             now,
		UpdatedAt:             now,
	}

	// Save the report using the repository
	err = r.reportRepository.Save(ctx, &report)
	if err != nil {
		return fmt.Errorf("failed to save report: %w", err)
	}

	return nil
}

func (r *ReportServiceImpl) Delete(ctx context.Context, reportId int) error {
	// Retrieve the report by its ID
	report, err := r.reportRepository.FindById(ctx, reportId)
	if err != nil {
		return err // Return error if FindById fails
	}

	// Delete the report using its ID
	err = r.reportRepository.Delete(ctx, report.Id)
	if err != nil {
		return err // Return error if Delete fails
	}

	return nil
}

func (r *ReportServiceImpl) FindAll(ctx context.Context) ([]response.ReportResponse, error) {
	reports, err := r.reportRepository.FindAll(ctx)
	if err != nil {
		return nil, err // Return error if FindAll fails
	}

	var reportResp []response.ReportResponse

	for _, value := range reports {
		report := response.ReportResponse{
			Id:                    value.Id,
			MonthOf:               value.MonthOf,
			WorkerName:            value.WorkerName,
			AreaOfAssignment:      value.AreaOfAssignment,
			NameOfChurch:          value.NameOfChurch,
			WorshipService:        value.WorshipService,
			SundaySchool:          value.SundaySchool,
			PrayerMeetings:        value.PrayerMeetings,
			BibleStudies:          value.BibleStudies,
			MensFellowships:       value.MensFellowships,
			WomensFellowships:     value.WomensFellowships,
			YouthFellowships:      value.YouthFellowships,
			ChildFellowships:      value.ChildFellowships,
			Outreach:              value.Outreach,
			TrainingOrSeminars:    value.TrainingOrSeminars,
			LeadershipConferences: value.LeadershipConferences,
			LeadershipTraining:    value.LeadershipTraining,
			Others:                value.Others,
			FamilyDays:            value.FamilyDays,
			TithesAndOfferings:    value.TithesAndOfferings,
			CreatedAt:             value.CreatedAt,
			UpdatedAt:             value.UpdatedAt,
		}

		// Calculate average attendance for each type
		report.WorshipServiceAvg = model.CalculateAverage(value.WorshipService)
		report.SundaySchoolAvg = model.CalculateAverage(value.SundaySchool)
		report.PrayerMeetingsAvg = model.CalculateAverage(value.PrayerMeetings)
		report.BibleStudiesAvg = model.CalculateAverage(value.BibleStudies)
		report.MensFellowshipsAvg = model.CalculateAverage(value.MensFellowships)
		report.WomensFellowshipsAvg = model.CalculateAverage(value.WomensFellowships)
		report.YouthFellowshipsAvg = model.CalculateAverage(value.YouthFellowships)
		report.ChildFellowshipsAvg = model.CalculateAverage(value.ChildFellowships)
		report.OutreachAvg = model.CalculateAverage(value.Outreach)
		report.TrainingOrSeminarsAvg = model.CalculateAverage(value.TrainingOrSeminars)
		report.LeadershipConferencesAvg = model.CalculateAverage(value.LeadershipConferences)
		report.LeadershipTrainingAvg = model.CalculateAverage(value.LeadershipTraining)
		report.OthersAvg = model.CalculateAverage(value.Others)
		report.FamilyDaysAvg = model.CalculateAverage(value.FamilyDays)
		report.TithesAndOfferingsAvg = model.CalculateAverage(value.TithesAndOfferings)

		reportResp = append(reportResp, report)
	}

	return reportResp, nil
}

func (r *ReportServiceImpl) FindById(ctx context.Context, reportId int) (response.ReportResponse, error) {
	report, err := r.reportRepository.FindById(ctx, reportId)
	helper.ErrorPanic(err)

	reportResp := response.ReportResponse{
		Id:               report.Id,
		MonthOf:          report.MonthOf,
		WorkerName:       report.WorkerName,
		AreaOfAssignment: report.AreaOfAssignment,
		NameOfChurch:     report.NameOfChurch,
		WorshipService:   report.WorshipService,
		// AverageAttendance: report.AverageAttendance,
		CreatedAt: report.CreatedAt,
		UpdatedAt: report.UpdatedAt,
	}

	return reportResp, nil
}

func (r *ReportServiceImpl) Update(ctx context.Context, request *request.ReportUpdateRequest) error {
	report, err := r.reportRepository.FindById(ctx, request.Id)
	if err != nil {
		return err
	}

	report.MonthOf = request.MonthOf
	report.WorkerName = request.WorkerName
	report.AreaOfAssignment = request.AreaOfAssignment
	report.NameOfChurch = request.NameOfChurch

	err = r.reportRepository.Update(ctx, report)
	if err != nil {
		return err
	}

	return nil
}
