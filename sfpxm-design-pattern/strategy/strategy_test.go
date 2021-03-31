package strategy

import (
	"log"
	"testing"
)

func Test_strategy(t *testing.T) {
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	fileSaver, err := NewStorageStrategy(strategyType)
	if err != nil {
		log.Println(err)
		return
	}
	err = fileSaver.save("strategy", data)
	if err != nil {
		t.Error(err)
		return
	}

}

func getData() ([]byte, bool) {
	return []byte("test data"), false
}
