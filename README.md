# ghelper ![workflow](https://github.com/cheungxiongwei/ghelper/actions/workflows/go.yml/badge.svg)

example:
```go
package main
import (
    "github.com/cheungxiongwei/ghelper"
)

type Post struct {
    Id        int64
    AuthorId  int64     `pgsql:"where"`
    Title     string    `pgsql:"set"`
    Content   string    `pgsql:"set"`
    CreatedAt time.Time `pgsql:"set"`
    UpdatedAt time.Time `pgsql:"set"`
}

func main() {
    println(GInsert(Post{}))
    println(GUpdate(Post{}))
    println(GDelete(Post{}))
}
```

output:
```sql
INSERT INTO "post" ("author_id","title","content","created_at","updated_at") VALUES ($1,$2,$3,$4,$5);
UPDATE "post" SET "title" = $1,"content" = $2,"created_at" = $3,"updated_at" = $4 WHERE "author_id" = $5;
DELETE FROM "post"  WHERE "author_id" = $1; 
```
