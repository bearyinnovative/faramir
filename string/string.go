package string

import (
	"math/rand"
	"time"
)

type options struct {
	length  uint
	letters []rune
}

// Option for generating string
type Option func(*options)

var defaultOptions = options{
	length:  32,
	letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
}

// Length of the generated string
func Length(length uint) Option {
	return func(o *options) {
		o.length = length
	}
}

// Letters the generated string should contains
func Letters(letters []rune) Option {
	return func(o *options) {
		if len(letters) == 0 {
			panic("letters should not be empty")
		}
		o.letters = letters
	}
}

// String generate random string
func String(opts ...Option) string {
	opt := defaultOptions

	for _, o := range opts {
		o(&opt)
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, opt.length)

	for i := range b {
		b[i] = opt.letters[rand.Intn(len(opt.letters))]
	}
	return string(b)
}
