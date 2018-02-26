package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkGenWithDefaultOption(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		String()
	}
}

func benchmarkGenWithLength(length uint, b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		String(Length(length))
	}
}

func BenchmarkGenWithLength32(b *testing.B) {
	benchmarkGenWithLength(32, b)
}

func BenchmarkGenWithLength128(b *testing.B) {
	benchmarkGenWithLength(128, b)
}

func BenchmarkGenWithLength256(b *testing.B) {
	benchmarkGenWithLength(256, b)
}

func BenchmarkGenWithLength4096(b *testing.B) {
	benchmarkGenWithLength(4096, b)
}

func BenchmarkGenWithLength65536(b *testing.B) {
	benchmarkGenWithLength(65536, b)
}

func TestGen(t *testing.T) {
	Convey("generate random string", t, func() {
		Convey("with default options", func() {
			s := String()

			So(s, ShouldHaveLength, defaultStringOptions.length)

			for _, c := range s {
				So(defaultStringOptions.letters, ShouldContain, c)
			}
		})

		Convey("with length", func() {
			lens := []uint{
				1, 2, 3, 4, 8, 16,
			}

			for _, l := range lens {
				s := String(Length(l))

				So(s, ShouldHaveLength, l)
				for _, c := range s {
					So(defaultStringOptions.letters, ShouldContain, c)
				}
			}
		})

		Convey("with empty letters", func() {
			So(func() {
				String(Letters([]rune{}))
			}, ShouldPanic)
		})

		Convey("with letters", func() {
			letters := [][]rune{
				[]rune("123"),
				[]rune("1"),
				[]rune("abc"),
				[]rune("abc123"),
			}

			var length uint = 8

			for _, l := range letters {
				So(func() {
					s := String(Letters(l), Length(length))

					So(s, ShouldHaveLength, length)
					for _, c := range s {
						So(l, ShouldContain, c)
					}

				}, ShouldNotPanic)
			}
		})
	})
}
