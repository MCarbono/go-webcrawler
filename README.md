<h1 align="center">Go WebCrawler</h1>


## About

This is a rss webcrawler written in GO. The application receives feed urls from news websites, 
retrieves some data from it and saves it in the database.

---

## Libs/Dependencies

| Name        | Description | Documentation | Installation |
| ----------- | ----------- | ------------- | ----------- |     
| gofeed      | RSS webcrawler       |  github.com/mmcdole/gofeed  |  go get github.com/mmcdole/gofeed       |
| go-cmp   | Test library        | github.com/google/go-cmp     |   go get github.com/google/go-cmp          |
|  sqlite3  |   sqlite3 database driver     | github.com/mattn/go-sqlite3    |   go get github.com/mattn/go-sqlite3          |

---

## Run

Inside the root folder, run one of the commands below:

```bash
    # Makefile command
    $ make
```

```bash
    # Go command
    $ go run main.go
```

---

## Run Tests

Inside the root folder, run one of the commands below:

```bash
    # Makefile command
    $ make tests
```

```bash
    # Go command
    $ go test ./... -v
```