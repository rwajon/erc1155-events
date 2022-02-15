package utils

import (
	"fmt"
	"os"
)

func WriteToFile(fileName string, text string) error {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file error: ", err)
		return err
	}
	l, err := f.WriteString(text)
	if err != nil {
		fmt.Println("write to file error: ", err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println("close file error: ", err)
		return err
	}
	return nil
}
