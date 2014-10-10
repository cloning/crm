package auth

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

type FileTokenRepository struct {
	directory string
}

func (this *FileTokenRepository) Save(token *Token) error {
	b, err := this.serialize(token)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(this.filename(token.Key), b, 0644)
}

func (this *FileTokenRepository) Get(key string) (*Token, error) {
	fileContent, err := ioutil.ReadFile(this.filename(key))

	if os.IsNotExist(err) {
		return nil, nil
	}

	t, err := this.deserialize(fileContent)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (this *FileTokenRepository) serialize(t *Token) ([]byte, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err := e.Encode(t)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return b.Bytes(), nil
}

func (this *FileTokenRepository) filename(key string) string {
	return this.directory + "/" + key
}

func (this *FileTokenRepository) deserialize(b []byte) (*Token, error) {
	var t Token
	d := gob.NewDecoder(bytes.NewReader(b))
	err := d.Decode(&t)
	return &t, err
}
