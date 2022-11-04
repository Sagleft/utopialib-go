package utopia

type queryFieldsBuilder struct {
	m map[string]interface{}
}

func newMapBuilder() *queryFieldsBuilder {
	return &queryFieldsBuilder{
		m: make(map[string]interface{}),
	}
}

func (q *queryFieldsBuilder) getMap() map[string]interface{} {
	return q.m
}

func (q *queryFieldsBuilder) add(
	value interface{},
	unsetValue interface{},
	fieldTag string,
) *queryFieldsBuilder {
	if value != unsetValue {
		q.m[fieldTag] = value
	}
	return q
}
