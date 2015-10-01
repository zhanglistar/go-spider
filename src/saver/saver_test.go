package saver

import (
	"testing"
	"bytes"
	"fmt"
)

func Test_SaverShouldSuccessWhenEveryThingRight(t *testing.T) {
	file := FileSaver{}
	err := file.Save(bytes.NewBufferString("hello world").Bytes(), "testfile")
	if err != nil {
		t.Error("file.save failed")
	}
}

func Test_SaverShouldFailWhenNoPermission(t *testing.T) {
	file := FileSaver{}
	err := file.Save(bytes.NewBufferString("hello world").Bytes(), "/abc")
	if err == nil {
		t.Error("file.save failed")
	}
	fmt.Println(err)
}
