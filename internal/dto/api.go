package dto

type ErrorOutput struct {
	Error string `json:"error"`
}

func GetErrorOutput(err string) ErrorOutput {
	return ErrorOutput{
		Error: err,
	}
}

type MessageOutput struct {
	Message string `json:"message"`
}

func GetMessageOutput(msg string) MessageOutput {
	return MessageOutput{
		Message: msg,
	}
}
