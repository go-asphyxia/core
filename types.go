package core

type (
	Byte interface {
		~byte
	}

	Rune interface {
		~rune
	}

	String interface {
		~string
	}
)

type (
	Reader ReaderOf[byte]

	ReaderOf[B Byte] interface {
		ReadOf(p []B) (n int, err error)
	}
)
