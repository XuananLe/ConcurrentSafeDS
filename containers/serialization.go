package containers

type JsonSerialze[T comparable] interface {
	MarshalJson() ([]byte, error)  
}

type JsonDeserialize[T comparable] interface {
	UnmarshalJSON([]byte) error
}