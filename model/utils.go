package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"time"
)

type AppError struct {
	DetailedError string `json:"detailed_error"`        // Internal error string to help the developer
	RequestId     string `json:"request_id,omitempty"`  // The RequestId that's also set in the header
	StatusCode    int    `json:"status_code,omitempty"` // The http status code
	Where         string `json:"-"`                     // The function where it happened in the form of Struct.Func
	params        map[string]interface{}
}

func (appErr *AppError) Error() string {
	return appErr.Where + ": "  + appErr.DetailedError
}

func (appErr *AppError) ToJson() string {
	b, _ := json.Marshal(appErr)
	return string(b)
}


func NewAppError(where string, params map[string]interface{}, details string, status int) *AppError {
	appErr := &AppError{}
	appErr.params = params
	appErr.Where = where
	appErr.DetailedError = details
	appErr.StatusCode = status
	return appErr
}


func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func NewId() string {
	uuid := uuid.New()
	return uuid.String()
}

func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}
