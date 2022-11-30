package utopia

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Value float64
}

type testStruct2 struct {
	Value string
}

func TestUMap(t *testing.T) {
	u := uMap{}

	// when value is not set
	v := testStruct{}
	key1 := "test"
	key2 := "test2"

	// then
	u = u.add(key1, v.Value)
	_, isExists := u[key1]
	require.Equal(t, false, isExists)

	// when value is set
	v.Value = 1

	// then
	u = uMap{}
	u = u.add(key1, v.Value)
	_, isExists = u[key1]
	require.Equal(t, true, isExists)

	// when set key-value
	u = uMap{}
	u = u.set(key2, 2)
	_, isExists = u[key2]
	require.Equal(t, true, isExists)
}

func TestUMap2(t *testing.T) {
	u := uMap{}

	// when value is not set
	v := testStruct2{}
	key1 := "test"
	key2 := "test2"

	// then
	u = u.add(key1, v.Value)
	_, isExists := u[key1]
	require.Equal(t, false, isExists)

	// when value is set
	v.Value = "val"

	// then
	u = uMap{}
	u = u.add(key1, v.Value)
	_, isExists = u[key1]
	require.Equal(t, true, isExists)

	// when set key-value
	u = uMap{}
	u = u.set(key2, "str")
	_, isExists = u[key2]
	require.Equal(t, true, isExists)
}
