package gopresence

import "testing"

func TestPresenceErr(t *testing.T) {
	_, err := New("invalid client id")

	if err != nil {
		t.Fail()
	}
}
