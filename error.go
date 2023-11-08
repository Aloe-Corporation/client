package client

import "fmt"

type FailRequestError struct {
	Code         int
	ResponseBody []byte
}

func (e *FailRequestError) Error() string {
	responseBodyAsStr := string(e.ResponseBody)
	if responseBodyAsStr == "" {
		return fmt.Sprintf("%d fail request", e.Code)
	}
	return fmt.Sprintf("%d fail request, error message: %s", e.Code, responseBodyAsStr)
}
