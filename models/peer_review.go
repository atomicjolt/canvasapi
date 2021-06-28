package models

type PeerReview struct {
	AssessorID         int64  `json:"assessor_id" url:"assessor_id,omitempty"`                 // The assessors user id.Example: 23
	AssetID            int64  `json:"asset_id" url:"asset_id,omitempty"`                       // The id for the asset associated with this Peer Review.Example: 13
	AssetType          string `json:"asset_type" url:"asset_type,omitempty"`                   // The type of the asset.Example: Submission
	ID                 int64  `json:"id" url:"id,omitempty"`                                   // The id of the Peer Review.Example: 1
	UserID             int64  `json:"user_id" url:"user_id,omitempty"`                         // The user id for the owner of the asset.Example: 7
	WorkflowState      string `json:"workflow_state" url:"workflow_state,omitempty"`           // The state of the Peer Review, either 'assigned' or 'completed'.Example: assigned
	User               string `json:"user" url:"user,omitempty"`                               // the User object for the owner of the asset if the user include parameter is provided (see user API) (optional).Example: User
	Assessor           string `json:"assessor" url:"assessor,omitempty"`                       // The User object for the assessor if the user include parameter is provided (see user API) (optional).Example: User
	SubmissionComments string `json:"submission_comments" url:"submission_comments,omitempty"` // The submission comments associated with this Peer Review if the submission_comment include parameter is provided (see submissions API) (optional).Example: SubmissionComment
}

func (t *PeerReview) HasError() error {
	return nil
}
