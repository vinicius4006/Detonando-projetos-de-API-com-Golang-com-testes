package tests

import (
	"errors"
	"strings"
	"testing"

	"com.blocopad/blocopad/internal/backend"
	"com.blocopad/blocopad/internal/db"
)

var (
	deleteInvoked bool
	deleteKey     string
)

func TestGetKeyOk(t *testing.T) {
	//Given - pre-information about the test
	db.GetNote = func(key string) (bool, string, error) {
		return false, "OK", nil
	}

	// When - what i'm doing to start my test
	data, err := backend.GetKey("key1")

	// Then my test "running"
	if err != nil {
		t.Fatal("TestGetKeyOk should not return error")
	}

	if data != "OK" {
		t.Fatal("TestGetKeyOk Invalid return")
	}
}

func TestGetKeyDbError(t *testing.T) {
	// Given
	db.GetNote = func(key string) (bool, string, error) {
		return false, "OK", errors.New("Error")
	}

	// When
	_, err := backend.GetKey("key1")

	if err == nil {
		t.Fatalf("TestGetKeyDbError should not return nil")
	}

}

func TestGetKeyErrorSize(t *testing.T) {
	// Given
	keyString := ""
	keyString2 := strings.Repeat("a", 4000)

	// When
	_, errEmpty := backend.GetKey(keyString)

	// Then

	if errEmpty == nil {
		t.Fatalf("TestGetKeyErrorSize should return error zero length key")
	}

	// When
	_, errOverflow := backend.GetKey(keyString2)

	// Then
	if errOverflow == nil {
		t.Fatalf("TestGetKeyErrorSize should return error key_length > 36")
	}

}

func TestGetKeyDeleteDbError(t *testing.T) {

	// Given
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestGetKeyDeleteDbError the code did not panic")
		}
	}()

	deleteInvoked = false
	deleteKey = ""

	db.GetNote = func(key string) (bool, string, error) {
		return true, "OK", nil
	}

	db.DeleteNote = func(key string) error {
		deleteInvoked = true
		deleteKey = key
		return errors.New("Error")
	}

	// When

	_, err := backend.GetKey("key1")

	// Then

	if err != nil {
		t.Fatalf("TestGetKeyDeleteDbError should not return error")
	}
}

func TestSaveKeyOk(t *testing.T) {

	// Given
	db.SaveNote = func(data string, oneTime bool) (string, error) {
		return "123456", nil
	}

	// When
	uuid, err := backend.SaveKey("blablabla", false)

	// Then
	if err != nil {
		t.Fatalf("TestSaveKey OK Should not return error")
	}

	if uuid != "123456" {
		t.Fatalf("TestSaveKeyOK Invalid return")
	}
}

func TestSaveInvalidSize(t *testing.T) {

	// Given
	dataZeroLength := ""
	dataTooBig := strings.Repeat("a", 33000)

	// Then
	_, errZeroLength := backend.SaveKey(dataZeroLength, false)

	// When
	if errZeroLength == nil {
		t.Fatalf("TestSaveInvalidSize should not accepted length zero")
	}

	// Then
	_, errBigKey := backend.SaveKey(dataTooBig, false)

	// When
	if errBigKey == nil {
		t.Fatal("TestSaveInvalidSize should not accpeted too big key")
	}

}

func TestSaveKeyDbError(t *testing.T) {
	// Given
	db.SaveNote = func(data string, oneTime bool) (string, error) {
		return "123456", errors.New("Error")
	}

	// When
	_, err := backend.SaveKey("blablabla", false)

	// Then
	if err == nil {
		t.Fatalf("TestSaveKeyDbError Should return error")
	}
}
