package mysql

import "testing"

type TestReplacer struct{}

func (t TestReplacer) Replace(name string) string {
	return ""
}

func TestBase(t *testing.T) {
	p := GormMysqlPlugin{}
	if p.Enabled() != defaultConfig["enabled"] {
		t.Fatalf("enabled should be %v", defaultConfig["enabled"])
	}
	if p.Order() != defaultConfig["order"] {
		t.Fatalf("enabled should be %v", defaultConfig["order"])
	}
	if p.Load() == nil {
		t.Fatal("username, host and database should be null")
	}

}

func TestRepalce(t *testing.T) {
	SetReplacer(&TestReplacer{})
}
