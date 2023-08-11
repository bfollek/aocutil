package aocutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testFile = "../testdata/file.txt"

func TestMustBitSliceToInt(t *testing.T) {
	require := require.New(t)
	bits := []int{0, 1, 1, 0}
	i := MustBitSliceToInt(bits)
	require.Equal(i, 6)
}

func TestMustFileToStrings(t *testing.T) {
	require := require.New(t)
	ss := MustFileToStrings(testFile)
	require.Equal(3, len(ss))
	require.Equal(ss[0], "111")
	require.Equal(ss[2], "33333")
}

func TestMustFileToOneString(t *testing.T) {
	require := require.New(t)
	s, err := MustFileToOneString(testFile)
	require.Nil(err)
	require.Equal(s, "111\n2222\n33333")
}

func TestMustFileToInts(t *testing.T) {
	require := require.New(t)
	ii := MustFileToInts(testFile)
	require.Equal(3, len(ii))
	require.Equal(ii[0], 111)
	require.Equal(ii[2], 33333)
}

func TestSortStringInPlace(t *testing.T) {
	require := require.New(t)
	st := "f2HacbJELfls88"
	st = SortStringInPlace(st)
	//SortStringInPlace(st)
	require.Equal("288EHJLabcffls", st)
}
