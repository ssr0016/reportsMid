CREATE TABLE reports (
    id SERIAL PRIMARY KEY,
    month_of VARCHAR(100) NOT NULL,
    worker_name VARCHAR(100) NOT NULL,
    area_of_assignment VARCHAR(100) NOT NULL,
    name_of_church VARCHAR(100) NOT NULL,
    worship_service JSONB NOT NULL,
    sunday_school JSONB NOT NULL,
    prayer_meetings JSONB,
    bible_studies JSONB,
    mens_fellowships JSONB,
    womens_fellowships JSONB,
    youth_fellowships JSONB,
    child_fellowships JSONB,
    outreach JSONB,
    training_or_seminars JSONB,
    leadership_conferences JSONB,
    leadership_training JSONB,
    others JSONB,
    family_days JSONB,
    tithes_and_offerings JSONB,
    average_attendance FLOAT8 NOT NULL,
    worship_service_avg FLOAT8 DEFAULT NULL,
    sunday_school_avg FLOAT8 DEFAULT NULL,
    prayer_meetings_avg FLOAT8 DEFAULT NULL,
    bible_studies_avg FLOAT8 DEFAULT NULL,
    mens_fellowships_avg FLOAT8 DEFAULT NULL,
    womens_fellowships_avg FLOAT8 DEFAULT NULL,
    youth_fellowships_avg FLOAT8 DEFAULT NULL,
    child_fellowships_avg FLOAT8 DEFAULT NULL,
    outreach_avg FLOAT8 DEFAULT NULL,
    training_or_seminars_avg FLOAT8 DEFAULT NULL,
    leadership_conferences_avg FLOAT8 DEFAULT NULL,
    leadership_training_avg FLOAT8 DEFAULT NULL,
    others_avg FLOAT8 DEFAULT NULL,
    family_days_avg FLOAT8 DEFAULT NULL,
    tithes_and_offerings_avg FLOAT8 DEFAULT NULL,
    home_visited JSONB,
    bible_study_or_group_led JSONB,
    sermon_or_message_preached JSONB,
    person_newly_contacted JSONB,
    person_followed_up JSONB,
    person_led_to_christ JSONB,
    names JSONB,
    home_visited_avg FLOAT8 DEFAULT NULL,
    bible_study_or_group_led_avg FLOAT8 DEFAULT NULL,
    sermon_or_message_preached_avg FLOAT8 DEFAULT NULL,
    person_newly_contacted_avg FLOAT8 DEFAULT NULL,
    person_followed_up_avg FLOAT8 DEFAULT NULL,
    person_led_to_christ_avg FLOAT8 DEFAULT NULL,
    narrative_report TEXT,  
    challenges_and_problem_encountered TEXT,
    prayer_request TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
