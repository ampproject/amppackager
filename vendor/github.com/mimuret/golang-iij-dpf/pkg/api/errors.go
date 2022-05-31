package api

import "errors"

func IsBadResponse(err error, f func(b *BadResponse) bool) bool {
	bad := &BadResponse{}
	if !errors.As(err, &bad) {
		return false
	}
	if f == nil {
		return true
	}
	return f(bad)
}

func IsStatusCode(err error, code int) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsStatusCode(code)
	})
}

func IsErrType(err error, name string) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsErrType(name)
	})
}

func IsErrMsg(err error, msg string) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsErrMsg(msg)
	})
}

func IsErrorCode(err error, code string) (bool, string) {
	var (
		res       bool
		attribute string
	)
	IsBadResponse(err, func(bad *BadResponse) bool {
		res, attribute = bad.IsErrorCode(code)
		return res
	})
	return res, attribute
}

func IsErrorCodeAttribute(err error, code string, attribute string) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsErrorCodeAttribute(code, attribute)
	})
}

func IsAuthError(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsAuthError()
	})
}

func IsRequestFormatError(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsRequestFormatError()
	})
}

func IsParameterError(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsParameterError()
	})
}

func IsNotFound(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsNotFound()
	})
}

func IsTooManyRequests(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsTooManyRequests()
	})
}

func IsSystemError(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsSystemError()
	})
}

func IsGatewayTimeout(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsGatewayTimeout()
	})
}

func IsInvalidSchema(err error) bool {
	return IsBadResponse(err, func(bad *BadResponse) bool {
		return bad.IsInvalidSchema()
	})
}
