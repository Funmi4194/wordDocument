package helper

import (
	"fmt"
	"testing"
)

func Test_Content(t *testing.T) {
	result := Content("You can chat with me by creating a new\n Word document with the message in the content  \n- Walter White! Happy Valentines Day")

	if result != 23 {
		t.Error("incorrect result: expected 24, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}

func Test_TotalCharacters(t *testing.T) {
	result := TotalCharacters("You can chat with me by creating a new\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day")

	if result != 125 {
		t.Error("incorrect result: expected 124, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}

func Test_CharWithoutSpace(t *testing.T) {
	result := CharWithoutSpace("You can chat with me by creating a new\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day")

	if result != 102 {
		t.Error("incorrect result: expected 102, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}

func Test_Sentence(t *testing.T) {
	result := Sentence("You can chat with me by creating a new\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day")

	if result != 2 {
		t.Error("incorrect result: expected 124, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}

func Test_Paragraphs(t *testing.T) {
	result := Paragraphs("You can chat with me by creating a new\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day")

	if result != 1 {
		t.Error("incorrect result: expected 1, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}

func Test_Line(t *testing.T) {
	result := Line("You can chat with me by creating a new\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day")

	if result != 3 {
		t.Error("incorrect result: expected 124, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}
func Test_Page(t *testing.T) {
	result := Pages("You can chat with me by creating a new\n\n Word document with the message in the content.  \n- Walter White! Happy Valentines Day\n funmi is a dog\n she loves to eat like a cow\n I dont likw her for any reson")

	if result != 1 {
		t.Error("incorrect result: expected 1, got", result)
	} else {
		fmt.Println("Result correctly predicted")
	}
}
