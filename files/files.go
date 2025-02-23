package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb{
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
		return data, err
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}