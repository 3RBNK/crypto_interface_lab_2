package app

const (
	SaltSize   = 32
	HMACSize   = 32
	iVSize     = 16
	Iterations = 600_000
	KeySize    = 32

	headerSize = SaltSize + HMACSize + iVSize
)
