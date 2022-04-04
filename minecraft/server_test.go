package minecraft

import (
	"discraft/collection"
	"reflect"
	"testing"
)

func TestServer_CurrentPlayerList_GracefullyHandlesErrors(t *testing.T) {
	// This server will fail a TCP connection with "connection refused"
	s := Server{Hostname: "tcpbin.com", Port: 1234}

	want := collection.Strings{}

	if got := s.CurrentPlayerList(); !reflect.DeepEqual(got, want) {
		t.Errorf("CurrentPlayerList() = %v, want %v", got, want)
	}
}
