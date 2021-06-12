package models

type RubricAssessment struct {
	ID                  int64    `json:"id"`                    // the ID of the rubric.Example: 1
	RubricID            int64    `json:"rubric_id"`             // the rubric the assessment belongs to.Example: 1
	RubricAssociationID int64    `json:"rubric_association_id"` // Example: 2
	Score               int64    `json:"score"`                 // Example: 5.0
	ArtifactType        string   `json:"artifact_type"`         // the object of the assessment.Example: Submission
	ArtifactID          int64    `json:"artifact_id"`           // the id of the object of the assessment.Example: 3
	ArtifactAttempt     int64    `json:"artifact_attempt"`      // the current number of attempts made on the object of the assessment.Example: 2
	AssessmentType      string   `json:"assessment_type"`       // the type of assessment. values will be either 'grading', 'peer_review', or 'provisional_grade'.Example: grading
	AssessorID          int64    `json:"assessor_id"`           // user id of the person who made the assessment.Example: 6
	Data                string   `json:"data"`                  // (Optional) If 'full' is included in the 'style' parameter, returned assessments will have their full details contained in their data hash. If the user does not request a style, this key will be absent..
	Comments            []string `json:"comments"`              // (Optional) If 'comments_only' is included in the 'style' parameter, returned assessments will include only the comments portion of their data hash. If the user does not request a style, this key will be absent..
}

func (t *RubricAssessment) HasError() error {
	return nil
}
