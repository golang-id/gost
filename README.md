gost
====

Simple static file server. Current version is 0.1.1.

## Install

~~~text
$ go install github.com/golang-id/gost
~~~

## Usage

Running static file server on `http://localhost:8080` with current
directory as document root:

~~~text
$ gost
~~~

You can specify optional `listen` address and/or `path` (to be served
as the document root) flags.

~~~text
$ gost -listen=:12345 -path="/home/gedex/my_static_web"
~~~

You can also specify an IP address if needed. In this example the server will
bind to IP address 10.0.1.5 on port 9000.

~~~text
$ gost -listen 10.0.1.5:9000
~~~

See the help:

~~~text
$ gost --help
Usage of gost:
  -listen=":8080": Address to bind to
  -path="./": Path served as document root
~~~

## Roadmap

Provides binaries for major platforms using [goxc](https://github.com/laher/goxc).

Good reference: https://github.com/jingweno/gh

## Credits

* [Serving Static Files with HTTP](https://code.google.com/p/go-wiki/wiki/HttpStaticFiles)
* [Node.JS static file web server](https://gist.github.com/rpflorence/701407)

## License

See [LICENSE](./LICENSE) file.
