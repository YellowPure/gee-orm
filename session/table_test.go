package session

import "testing"

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	if err1 != nil || err2 != nil {
		t.Fatal("Failed to create table")
	}
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
