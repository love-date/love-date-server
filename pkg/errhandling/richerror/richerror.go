package richerror

type Kind uint

const (
	KindForbidden Kind = iota + 1
	KindNotFound
	KindUnexpected
	KindBadRequest
	KindUnauthorized
)

type Op string

type RichError struct {
	operation    Op
	message      string
	wrappedError error
	kind         Kind
	meta         map[string]interface{}
}

func New(op Op) RichError {

	return RichError{operation: op}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op

	return r
}

func (r RichError) WithWrapError(err error) RichError {
	r.wrappedError = err

	return r
}

func (r RichError) WithMessage(msg string) RichError {
	r.message = msg

	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Error() string {
	return r.message
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {

		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {

		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {

		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {

		return r.wrappedError.Error()
	}

	return re.Message()
}
