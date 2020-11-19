package encoding

type MediaEncoder interface {
	Encode(input string, output string) error
	CanEncode(extension string) bool
}
