package utopia

type uMap map[string]interface{}

func (u uMap) add(
	value interface{},
	unsetValue interface{},
	fieldTag string,
) uMap {
	if value != unsetValue {
		u[fieldTag] = value
	}
	return u
}
