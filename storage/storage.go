package storage

import "io"

type Storage interface {
	Save(p *Page)error
	PickRandom(userName string) (*Page,error)
	Remove(p *Page) error 
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL string
	UserName string
}

func (p Page) Hash() (string, error) {
	h:=sha1.New 

	io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
}