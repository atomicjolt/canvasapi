package models

type ModuleItemSequenceNode struct {
	Prev        *ModuleItem              `json:"prev" url:"prev,omitempty"`                 // The previous ModuleItem in the sequence.
	Current     *ModuleItem              `json:"current" url:"current,omitempty"`           // The ModuleItem being queried.Example: 768, 123, A lonely page, Page
	Next        *ModuleItem              `json:"next" url:"next,omitempty"`                 // The next ModuleItem in the sequence.Example: 769, 127, Project 1, Assignment
	MasteryPath map[string](interface{}) `json:"mastery_path" url:"mastery_path,omitempty"` // The conditional release rule for the module item, if applicable.Example: true, , , false, false, /courses/11/modules, /courses/11/modules/items/9/choose, false
}

func (t *ModuleItemSequenceNode) HasErrors() error {
	return nil
}
