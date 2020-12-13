package domain

type Error struct {
	Message string
	Code int
	Error string
}

const(
	dataValidateError=400
	repositoryError=501
)

var statusText = map[int]string{
	dataValidateError: "Data Validate Error",
	repositoryError: "Repository Error",
}

func errorText(code int) string {
	return statusText[code]
}