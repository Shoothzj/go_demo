package demo_ucloud

import "testing"

func TestListUlb(t *testing.T) {
	describeElb()
}

func TestCleanUlbServer(t *testing.T) {
	cleanElb([]string {"uhost-233c1yfr", "uhost-vcby4yuw","uhost-hsro4cz5"})
}

func TestChangeUlbServer(t *testing.T) {
	changeUlbServer([]string {"uhost-233c1yfr", "uhost-vcby4yuw","uhost-hsro4cz5"})
}
