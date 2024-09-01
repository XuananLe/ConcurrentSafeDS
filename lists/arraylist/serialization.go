package arraylist

import (
	"encoding/json"

	"github.com/XuananLe/ConcurrentSafeDS/containers"
)

// Check if implement the List interface
var _ containers.JsonSerialze[int] = &ArrayList[int]{}
var _ containers.JsonDeserialize[int] = &ArrayList[int]{}


func (l *ArrayList[T]) MarshalJson() ([]byte, error) {
	return json.Marshal(l.slice);
}

func (l *ArrayList[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.slice);
}