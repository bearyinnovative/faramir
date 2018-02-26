package email

import (
	"fmt"

	s "github.com/bearyinnovative/faramir/string"
)

type options struct {
	domain  string
	length  uint
	letters []rune
}

var defaultOptions = options{
	length:  32,
	letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	domain:  "bearyinnovative.com",
}

type Option func(*options)

func Domain(domain string) Option {
	return func(o *options) {
		o.domain = domain
	}
}

// Length of the generated email
func Length(length uint) Option {
	return func(o *options) {
		o.length = length
	}
}

// Letters the generated email should contains
func Letters(letters []rune) Option {
	return func(o *options) {
		if len(letters) == 0 {
			panic("letters should not be empty")
		}
		o.letters = letters
	}
}

func Email(opts ...Option) string {

	opt := defaultOptions

	for _, o := range opts {
		o(&opt)
	}

	email := fmt.Sprintf("%s@%s",
		s.String(
			s.Letters(opt.letters),
			s.Length(opt.length)),
		opt.domain)

	return email
}
