package saver
import (
	"os"
	"errors"
//	"fmt"
)

type Saver interface {
	Save(content []byte, name string) error
}

type FileSaver struct {
}

func (fs FileSaver) Save(content []byte, name string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
//		fmt.Printf("openfile failed, %s", name)
		return err
	}
	n, err := f.Write(content)
	if n != len(content) || err != nil {
		return errors.New("Write content to file failed")
	}
	return nil
}
