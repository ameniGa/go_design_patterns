package future

type SuccessFunc func(string)
type FailureFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failureFunc FailureFunc
}

func (ms *MaybeString) Success(f SuccessFunc) *MaybeString {
	ms.successFunc = f
	return ms
}

func (ms *MaybeString) Fail(f FailureFunc) *MaybeString {
	ms.failureFunc = f
	return ms
}

func (ms *MaybeString) Execute(f ExecuteStringFunc) {
	go func(m *MaybeString) {
		str, err := f()
		if err != nil {
			m.failureFunc(err)
			return
		}
		m.successFunc(str)
	}(ms)
}
