gost
====

Simple static file server. Current version is 0.0.0.

## Install

~~~text
$ go install github.com/golang-id/gost
~~~

## Usage

Running static file server on `http://localhost:8080` with current directory as document root:

~~~text
$ gost
~~~

You can specify optional `port` and/or `path` (to be served as the document root):

~~~text
$ gost -port=12345 -path="/home/gedex/my_static_web"
~~~

See the help:

~~~text
$ gost --help
Usage of gost:
  -path="./": Path served as document root.
  -port=8080: Port to listen
~~~

## Roadmap

Provides binaries for major platforms using [goxc](https://github.com/laher/goxc).

Good reference: https://github.com/jingweno/gh

## Credits

* [Serving Static Files with HTTP](https://code.google.com/p/go-wiki/wiki/HttpStaticFiles)
* [Node.JS static file web server](https://gist.github.com/rpflorence/701407)

## License

See [LICENSE](./LICENSE) file.
