package files

import (
	"encoding/gob"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)
	if err != nil {
		return nil
	}

	filePath := filepath.Join(fPath, fName)

	file,err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _=file.Close()}()

	if err:=gob.NewEncoder(file).Encode(page){
		return err
	} 
		return nil 
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
