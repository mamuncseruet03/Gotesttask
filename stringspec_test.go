package main

import "testing"

func TestTestValidity(t *testing.T) {

	got := testValidity("23-ab-48-caba-56-haha")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAverageNumber(t *testing.T) {

	got := averageNumber("23-ab-48-caba-56-haha")
	want := 42

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestWholeStory(t *testing.T) {

	got := wholeStory("1-hello-2-world")
	want := "hello world"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestStoryStats(t *testing.T) {

	got1, got2, got3, got4 := storyStats("23-ab-48-caba-56-haha")
	want1, want2, want3, want4 := "ab", "caba", 3, []string{"caba", "haha"}

	if got1 != want1 || got2 != want2 || got3 != want3 || len(got4) != len(want4) {
		t.Errorf("got %v, %v, %v, %v wanted %v, %v, %v, %v", got1, got2, got3, got4, want1, want2, want3, want4)
	}
}
