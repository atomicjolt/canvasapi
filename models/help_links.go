package models

type HelpLinks struct {
	HelpLinkName     string      `json:"help_link_name" url:"help_link_name,omitempty"`         // Help link button title.Example: Help And Policies
	HelpLinkIcon     string      `json:"help_link_icon" url:"help_link_icon,omitempty"`         // Help link button icon.Example: help
	CustomHelpLinks  []*HelpLink `json:"custom_help_links" url:"custom_help_links,omitempty"`   // Help links defined by the account. Could include default help links..Example: {'id'=>'link1', 'text'=>'Custom Link!', 'subtext'=>'Something something.', 'url'=>'https://google.com', 'type'=>'custom', 'available_to'=>['user', 'student', 'teacher', 'admin', 'observer', 'unenrolled'], 'is_featured'=>true, 'is_new'=>false, 'feature_headline'=>'Check this out!'}
	DefaultHelpLinks []*HelpLink `json:"default_help_links" url:"default_help_links,omitempty"` // Default help links provided when account has not set help links of their own..Example: {'available_to'=>['student'], 'text'=>'Ask Your Instructor a Question', 'subtext'=>'Questions are submitted to your instructor', 'url'=>'#teacher_feedback', 'type'=>'default', 'id'=>'instructor_question', 'is_featured'=>false, 'is_new'=>true, 'feature_headline'=>''}, {'available_to'=>['user', 'student', 'teacher', 'admin', 'observer', 'unenrolled'], 'text'=>'Search the Canvas Guides', 'subtext'=>'Find answers to common questions', 'url'=>'https://community.canvaslms.com/t5/Guides/ct-p/guides', 'type'=>'default', 'id'=>'search_the_canvas_guides', 'is_featured'=>false, 'is_new'=>false, 'feature_headline'=>''}, {'available_to'=>['user', 'student', 'teacher', 'admin', 'observer', 'unenrolled'], 'text'=>'Report a Problem', 'subtext'=>'If Canvas misbehaves, tell us about it', 'url'=>'#create_ticket', 'type'=>'default', 'id'=>'report_a_problem', 'is_featured'=>false, 'is_new'=>false, 'feature_headline'=>''}
}

func (t *HelpLinks) HasError() error {
	return nil
}
