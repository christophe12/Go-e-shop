## Go-E-Shop

The goal is to build a simple e-commerce Rest api with go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development.

### Prerequisites

To run this app you need to have below dependencies:

##### Make Sure You Have:

- [goorm](https://github.com/jinzhu/gorm)
- [go mysql driver package](https://github.com/go-sql-driver/mysql)

### Configuring DB

Under ``` main.go ``` you can fill in the db constants

### Launch and view the APP

If you are on unix you can use below script to run all go files excluding test files.
```
go run $(ls -1 *.go | grep -v _test.go)
```

## Authors

* **Christophe Mutabazi** - *Full stack engineer* - [Personalwebsite](http://orbit.surge.sh/)

## License
Copyright {2018}
