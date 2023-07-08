package consts

type TErrors struct {
	INTERNAL      string
	UNAUTHORIZED  string
	UNPROCESSABLE string
	NOT_FOUND     string
	BAD_REQUEST   string
}

var Errors = TErrors{
	INTERNAL:      "internal",
	UNAUTHORIZED:  "unauthorized",
	UNPROCESSABLE: "unprocessable",
	NOT_FOUND:     "not_found",
	BAD_REQUEST:   "bad_request",
}
