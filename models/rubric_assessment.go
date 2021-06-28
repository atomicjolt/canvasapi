package models

type RubricAssessment struct {
	ID                  int64    `json:"id" url:"id,omitempty"`                                       // the ID of the rubric.Example: 1
	RubricID            int64    `json:"rubric_id" url:"rubric_id,omitempty"`                         // the rubric the assessment belongs to.Example: 1
	RubricAssociationID int64    `json:"rubric_association_id" url:"rubric_association_id,omitempty"` // Example: 2
	Score               int64    `json:"score" url:"score,omitempty"`                                 // Example: 5.0
	ArtifactType        string   `json:"artifact_type" url:"artifact_type,omitempty"`                 // the object of the assessment.Example: Submission
	ArtifactID          int64    `json:"artifact_id" url:"artifact_id,omitempty"`                     // the id of the object of the assessment.Example: 3
	ArtifactAttempt     int64    `json:"artifact_attempt" url:"artifact_attempt,omitempty"`           // the current number of attempts made on the object of the assessment.Example: 2
	AssessmentType      string   `json:"assessment_type" url:"assessment_type,omitempty"`             // the type of assessment. values will be either 'grading', 'peer_review', or 'provisional_grade'.Example: grading
	AssessorID          int64    `json:"assessor_id" url:"assessor_id,omitempty"`                     // user id of the person who made the assessment.Example: 6
	Data                string   `json:"data" url:"data,omitempty"`                                   // (Optional) If 'full' is included in the 'style' parameter, returned assessments will have their full details contained in their data hash. If the user does not request a style, this key will be absent..
	Comments            []string `json:"comments" url:"comments,omitempty"`                           // (Optional) If 'comments_only' is included in the 'style' parameter, returned assessments will include only the comments portion of their data hash. If the user does not request a style, this key will be absent..
}

func (t *RubricAssessment) HasError() error {
	return nil
}
