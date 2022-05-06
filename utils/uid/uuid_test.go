package uid

import "testing"

func TestUuid(t *testing.T) {
	uid := NewId()
	t.Log(uid)
}
