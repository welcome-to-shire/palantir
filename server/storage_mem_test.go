package main

import "testing"

func TestCreate(t *testing.T) {
	store := makeStorage(t)

	t1, err := store.Create("test", "test title", "test content")
	if err != nil {
		t.Fatal(err)
	}
	t2, err := store.Create("test", "test title", "test content")
	if err != nil {
		t.Fatal(err)
	}
	if t1 == t2 {
		t.Error("should create new ticket for identical message")
	}
}

func TestGetOne(t *testing.T) {
	store := makeStorage(t)

	subject := "test"
	title := "title"
	content := "content"

	_, err := store.GetOne(subject)
	if err == nil {
		t.Error("should not get from empty list")
	}

	_, err = store.Create(subject, title, content)
	if err != nil {
		t.Fatal(err)
	}
	m, err := store.GetOne(subject)
	if err != nil {
		t.Fatal(err)
	}
	if m.Title != title {
		t.Errorf("title expected: %s got: %s", title, m.Title)
	}
	if m.Content != content {
		t.Errorf("content expected: %s got: %s", content, m.Content)
	}

	_, err = store.GetOne(subject)
	if err == nil {
		t.Error("should be empty list after read")
	}

}

func makeStorage(t *testing.T) Storage {
	store := MakeMemStorage()
	if store == nil {
		t.Fatal("cannot make storage")
	}

	return store
}
