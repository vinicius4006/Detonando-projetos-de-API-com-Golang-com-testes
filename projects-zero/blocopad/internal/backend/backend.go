package backend

import (
	"errors"
	"log"

	"com.blocopad/blocopad/internal/db"
)

func GetKey(key string) (string, error) {
	if len(key) == 0 || len(key) > 36 {
		return "", errors.New("Key with wrong sizee")
	}

	oneTime, data, err := db.GetNote(key)

	if err != nil {
		return "", err
	}

	if oneTime {
		if err := db.DeleteNote(key); err != nil {
			log.Panic("Cannot delete onetime note!")
		}
	}

	return data, nil
}

func SaveKey(data string, oneTime bool) (string, error) {
	byteSize := len([]rune(data))

	if byteSize == 0 || byteSize > (32*1024) {
		return "", errors.New(("Invalid note size"))

	}

	uuidCode, err := db.SaveNote(data, oneTime)
	if err != nil {
		return "", errors.New((err.Error()))
	}

	return uuidCode, nil
}
