package krandom

import (
	"testing"

	"github.com/google/uuid"
)

func TestKrand(t *testing.T) {
	//var wg sync.WaitGroup

	number := 10000000
	mapClientId := make(map[string]struct{})
	for i := 0; i < number; i++ {
		//
		//clientId := Krand(16, 3)
		//var id = fmt.Sprintf("%x", md5.Sum([]byte("127.0.0.1"+clientId)))
		//newUUID, _ := uuid.NewUUID()
		//id := newUUID.String()
		id := uuid.New().String()
		mapClientId[id] = struct{}{}
	}

	if len(mapClientId) != number {
		t.Log("actual length: ", len(mapClientId), " expect length: ", number)
		t.Errorf("failed")
	}

}
