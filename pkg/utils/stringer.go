package utils

type StringerFunc func() string

func (sf StringerFunc) String() string {
	return sf()
}
