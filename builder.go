package gopherpanic

type Builder interface {
	New() ErrorBuilder
	Default() ErrorBuilder
	WithMessage(message string) ErrorBuilder
	WithPosition(position Position) ErrorBuilder
	WithCause(causes ...Error) ErrorBuilder
	Build() Error
}

type ErrorBuilder struct {
	code     Code
	message  string
	position Position
	traces   []Trace
}

// Create a new empty Error
func (builder ErrorBuilder) New() ErrorBuilder {
	return ErrorBuilder{}
}

// Initialize the default values
//
// - Code: UnknownError
//
// - Message: "an enexpected error occurred"
//
// - Position: *the current possition of Default call in your code*
func (builder ErrorBuilder) Default() ErrorBuilder {
	return ErrorBuilder{code: UnknownError, message: "an unexpected error occurred", position: Position{}.spawn(2), traces: nil}
}

func (builder ErrorBuilder) WithCode(code Code) ErrorBuilder {
	builder.code = code
	return builder
}

func (builder ErrorBuilder) WithMessage(message string) ErrorBuilder {
	builder.message = message
	return builder
}

func (builder ErrorBuilder) WithPosition(position Position) ErrorBuilder {
	builder.position = position
	return builder
}

func (builder ErrorBuilder) WithTraces(traces ...Trace) ErrorBuilder {
	builder.traces = traces
	return builder
}

func (builder ErrorBuilder) Build() Error {
	return Error{
		Code:     builder.code,
		Message:  builder.message,
		Position: builder.position,
		Traces:   builder.traces,
	}
}
