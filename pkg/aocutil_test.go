package aocutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testFile = "../testdata/file.txt"

func TestMustBitSliceToInt(t *testing.T) {
	bits := []int{0, 1, 1, 0}
	want := 6
	got := MustBitSliceToInt(bits)
	if got != want {
		t.Errorf("MustBitSliceToInt(%v) = %d; want %d", bits, got, want)
	}
}

func TestMustFileToStrings(t *testing.T) {
	got := MustFileToStrings(testFile)
	wantI := 3
	if len(got) != wantI {
		t.Errorf("len(MustFileToStrings(%s)) = %d; want %d", testFile, len(got), wantI)
	}
	wantS := "111"
	if got[0] != wantS {
		t.Errorf("MustFileToStrings(%s) = %s; want %s", testFile, got[0], wantS)
	}
	wantS = "33333"
	if got[2] != wantS {
		t.Errorf("MustFileToStrings(%s) = %s; want %s", testFile, got[2], wantS)
	}
}

func TestMustFileToOneString(t *testing.T) {
	got := MustFileToOneString(testFile)
	want := "111\n2222\n33333"
	if got != want {
		t.Errorf("MustFileToStrings(%s) = %s; want %s", testFile, got, want)
	}
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
	s := "f2HacbJELfls88"
	s = SortString(s)
	require.Equal("288EHJLabcffls", s)
	s = SortString("X7m2ax9")
	require.Equal("279Xamx", s)
}
