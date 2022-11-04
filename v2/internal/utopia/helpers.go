package utopia

type uMap map[string]interface{}

func (u uMap) add(
	fieldTag string,
	value interface{},
	unsetValue interface{},
) uMap {
	if value != unsetValue {
		u[fieldTag] = value
	}
	return u
}

func (u uMap) set(fieldTag string, value interface{}) uMap {
	u[fieldTag] = value
	return u
}
