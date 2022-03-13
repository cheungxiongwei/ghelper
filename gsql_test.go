package ghelper

import (
	"testing"
	"time"
)

// go test -v
// go test -bench .

type Post struct {
	Id        int64
	AuthorId  int64     `pgsql:"where"`
	Title     string    `pgsql:"set"`
	Content   string    `pgsql:"set"`
	CreatedAt time.Time `pgsql:"set"`
	UpdatedAt time.Time `pgsql:"set"`
}

func TestGInsert(t *testing.T) {
	got := GInsert(Post{})
	want := `INSERT INTO "post" ("author_id","title","content","created_at","updated_at") VALUES ($1,$2,$3,$4,$5);`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGUpdate(t *testing.T) {
	got := GUpdate(Post{})
	want := `UPDATE "post" SET "title" = $1,"content" = $2,"created_at" = $3,"updated_at" = $4 WHERE "author_id" = $5;`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGDelete(t *testing.T) {
	got := GDelete(Post{})
	want := `DELETE FROM "post"  WHERE "author_id" = $1;`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkGInsert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GInsert(Post{})
	}
}

func BenchmarkGUpdate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GUpdate(Post{})
	}
}

func BenchmarkGDelete(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GDelete(Post{})
	}
}
