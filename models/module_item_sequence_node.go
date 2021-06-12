package models

type ModuleItemSequenceNode struct {
	Prev        *ModuleItem `json:"prev"`         // The previous ModuleItem in the sequence.
	Current     *ModuleItem `json:"current"`      // The ModuleItem being queried.Example: 768, 123, A lonely page, Page
	Next        *ModuleItem `json:"next"`         // The next ModuleItem in the sequence.Example: 769, 127, Project 1, Assignment
	MasteryPath string      `json:"mastery_path"` // The conditional release rule for the module item, if applicable.Example: true, , , false, false, /courses/11/modules, /courses/11/modules/items/9/choose, false
}

func (t *ModuleItemSequenceNode) HasError() error {
	return nil
}
