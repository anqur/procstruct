package edsl

type Filer interface {
	Header(text string) Filer
}
