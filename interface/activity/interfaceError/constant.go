package interfaceError

import "errors"

var (
	InvalidRequestBody           = errors.New("invalid request body data")
	FailedCreateNewActivity      = errors.New("failed to create new activity")
	FailedDeleteExistingActivity = errors.New("failed to delete existing activity")
	FailedUpdateExistingActivity = errors.New("failed update existing activity")
	FailedGetListActivity        = errors.New("failed get list activity")
	DataNotFound                 = errors.New("sorry, activity data does not exist")
)
