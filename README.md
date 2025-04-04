# expecto

![example spec](https://github.com/user-attachments/assets/f7e4a512-59d4-4aa7-abd2-6d509b56a74a)

## Matchers

```go
should := expecto.New(t)
should.NoErr("parsing config", err)
should.Eq("number of files", len(c.Files), 1)
should.Eq("number of chunks", len(c.Files[0].Chunks), 2)
should.Eq("number of migrations", len(c.Files[0].Migrations), 1)
should.Nil("transaction", c.Transaction)
```

## Temporary file system

Random dir in system TempDir is created.

```go
fs, dir, cleanup := expecto.TempFS(
   "src/path.txt",
   "content1",
   "src/b/2.txt",
   "content2,
)
defer cleanup()

```
