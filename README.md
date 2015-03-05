GoEnigma
========

An enigma machine written in the golang

Build
-----

Build:

    $ go build

Run:

    $ ./GoEnigma


Testing
-------

Unit tests depend on [GoConvey](https://github.com/smartystreets/goconvey), so install the dependencies first:

    $ go get github.com/smartystreets/goconvey
    $ go get github.com/smartystreets/assertions
    
Run the unit tests

    $ go test ./... -v