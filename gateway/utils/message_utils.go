package utils

import "fmt"

var UnableProcess = "Unable to process the request."
var ErrorOccured = "Unexpected error occured."
var TryAgain = "Please try again later."

func LoginRegisterErrorMessage(typeForm *string) string {
	return fmt.Sprintf("Please check again your %s form and try again.", *typeForm)
}

func GetGeneralErrorMessage() string {
	return fmt.Sprintf("%s %s", ErrorOccured, TryAgain)
}