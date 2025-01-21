package aocutil

import (
	"testing"
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
	wantI := 3
	got := MustFileToStrings(testFile)
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
	want := "111\n2222\n33333"
	got := MustFileToOneString(testFile)
	if got != want {
		t.Errorf("MustFileToStrings(%s) = %s; want %s", testFile, got, want)
	}
}

func TestMustFileToInts(t *testing.T) {
	want := 3
	got := MustFileToInts(testFile)
	if len(got) != want {
		t.Errorf("len(MustFileToStrings(%s)) = %d; want %d", testFile, len(got), want)
	}

	want = 111
	if got[0] != want {
		t.Errorf("MustFileToStrings(%s) = %d; want %d", testFile, got[0], want)
	}

	want = 33333
	if got[2] != want {
		t.Errorf("MustFileToStrings(%s) = %d; want %d", testFile, got[2], want)
	}
}

func TestSortStringInPlace(t *testing.T) {
	s := "f2HacbJELfls88"
	want := "288EHJLabcffls"
	got := SortString(s)
	if got != want {
		t.Errorf("SortString(%s)) = %s; want %s", s, got, want)
	}

	s = "X7m2ax9"
	want = "279Xamx"
	got = SortString(s)
	if got != want {
		t.Errorf("SortString(%s) = %s; want %s", s, got, want)
	}
}
