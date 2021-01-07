package groutine

import "testing"

func Test_rwLock(t *testing.T) {
	rwLock()
}

// test lock
func Test_wLock(t *testing.T) {
	Lock()
}

// test rwlock and lock
func Test_Lock_sub(t *testing.T) {
	t1 := Lock()
	t2 := rwLock()
	t3 := t1 - t2
	if t3 > 0 {
		t.Log("t3", t3)
		return
	}
	t.Error("t3", 2)

}

// test critical resource value unexpected
func Test_criticalResource(t *testing.T) {
	 criticalResource()
}
