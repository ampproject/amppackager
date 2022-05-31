package api

import "net/http"

type Action string

const (
	ActionCreate Action = "Create"
	ActionRead   Action = "Read"
	ActionList   Action = "List"
	ActionUpdate Action = "Update"
	ActionDelete Action = "Delete"
	ActionCount  Action = "Count"
	ActionCancel Action = "Cancel"
	ActionApply  Action = "Apply"
)

func Actions() []Action {
	return []Action{
		ActionCreate,
		ActionRead,
		ActionList,
		ActionUpdate,
		ActionDelete,
		ActionCount,
		ActionCancel,
		ActionApply,
	}
}

func (a Action) ToMethod() string {
	switch a {
	case ActionCreate:
		return http.MethodPost
	case ActionRead:
		return http.MethodGet
	case ActionList:
		return http.MethodGet
	case ActionUpdate:
		return http.MethodPatch
	case ActionDelete:
		return http.MethodDelete
	case ActionCount:
		return http.MethodGet
	case ActionCancel:
		return http.MethodDelete
	case ActionApply:
		return http.MethodPatch
	}
	return ""
}
