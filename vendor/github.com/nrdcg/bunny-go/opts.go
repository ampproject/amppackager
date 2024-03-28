package bunny

// Option is a type for Client options.
type Option func(*Client)

// WithHTTPRequestLogger is an option to log all sent out HTTP-Request via a log function.
func WithHTTPRequestLogger(logger Logf) Option {
	return func(clt *Client) {
		clt.httpRequestLogf = logger
	}
}

// WithHTTPResponseLogger is an option to log all received HTTP-Responses via a log function.
func WithHTTPResponseLogger(logger Logf) Option {
	return func(clt *Client) {
		clt.httpResponseLogf = logger
	}
}

// WithUserAgent is an option to specify the value of the User-Agent HTTP
// Header.
func WithUserAgent(userAgent string) Option {
	return func(clt *Client) {
		clt.userAgent = userAgent
	}
}

// WithLogger is an option to set a log function to which informal and warning
// messages will be logged.
func WithLogger(logger Logf) Option {
	return func(clt *Client) {
		clt.logf = logger
	}
}
