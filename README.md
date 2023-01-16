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
|  cobra  |    library to create CLI applications   | github.com/spf13/cobra   |   go get github.com/spf13/cobra          |

---

## Run

Inside the root folder, run one of the commands below passing a flag called input with the urls separated by comma:

```bash
    $ go run main.go --input=https://catracalivre.com.br/feed/,https://www.infomoney.com.br/feed/
```

```bash
    # Shorthand version
    $ go run main.go -i=https://catracalivre.com.br/feed/,https://www.infomoney.com.br/feed/
```

If you don't wanna pass a flag, run one of the commands below and five default urls will be scrapp:

```bash
    # Makefile command
    $ make
```

```bash
    # Go command
    $ go run main.go
```

## CLI Documentation

To read more about this project CLI documentation, execute the command below:

```bash
    $ go run main.go --help
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