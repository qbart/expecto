# expecto

![example spec](https://github.com/user-attachments/assets/f7e4a512-59d4-4aa7-abd2-6d509b56a74a)

Idea is to provide simple wrapper for testing library that immediately exits when assertion fails
and prints colored output with diff.

Most matchers follow the pattern: `matcher(t *testing.T, msg string, value any, expected any)`

## Matchers

```go
expecto.NoErr(t, "parsing config", err)
expecto.Eq(t, "number of files", len(c.Files), 1)
expecto.Eq(t, "number of chunks", len(c.Files[0].Chunks), 2)
expecto.Eq(t, "number of migrations", len(c.Files[0].Migrations), 1)
expecto.Nil(t, "transaction", c.Transaction)
expecto.Contains(t, "notification message", str, "error message")
expecto.Map(t, anyMap).HasKey("has cache key", "cache-1")
```

See files for the full list of assertions.

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
