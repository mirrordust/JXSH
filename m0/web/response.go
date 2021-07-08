package web

// ********** response **********

func newErrorResponse(s ...string) errorResponse {
	if len(s) == 0 {
		return errorResponse{Msg: "param error"}
	}
	return errorResponse{Msg: s[0]}
}

type errorResponse struct {
	Msg string
}

// ********** error **********

func newDBError(msg string) error {
	return &dBError{msg}
}

type dBError struct {
	msg string
}

func (e *dBError) Error() string {
	return "DB error: " + e.msg
}
