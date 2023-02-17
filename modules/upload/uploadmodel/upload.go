package uploadmodel

import (
	"errors"
	"lesson-5-goland/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SqlModel `json:",inline"`
	common.Image    `json:",inline"`
}

var (
	ErrFileTooLarge = common.NewCustomError(errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge")
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err,
		"file is not image",
		"ErrFileIsNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "can not save file", "ErrCannotSaveFile")
}
