package response

type Extensions map[string]interface{}

func (e Extensions) SetCode(code int) Extensions {
	e["code"] = code
	return e
}
