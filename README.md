# 🔥 Teargas - Simple REST API tester made in Go with JWT support. 🔍🕵️‍♂️🌐

[![Go](https://github.com/VentGrey/teargas/actions/workflows/go.yml/badge.svg)](https://github.com/VentGrey/teargas/actions/workflows/go.yml)
[![CodeQL](https://github.com/VentGrey/teargas/actions/workflows/codeql.yml/badge.svg)](https://github.com/VentGrey/teargas/actions/workflows/codeql.yml)

My personal REST API consuming software written in Go. Teargas is a powerful yet lightweight tool designed for testing and consuming REST APIs with ease. With built-in JWT support, it's versatile and perfect for modern web development.


## 🌟 Features

- Simple & intuitive CLI interface.
- JWT Authentication support.
- Simple response statistics for quick evaluations.
- Binary is **NOT** small.
- Normally compiled binary size is ~7.3MB (Yes, fuck static binaries)
- Stripped binary is ~5.0MB (Yep...)
  

## 🚨 Copyright disclaimer

This software name and thumbnail don't intend to infringe copyright laws by illegally copying or claiming content that isn't mine. This is made solely as a tribute to the artists/bands I like, the original idea + artwork concept and registered trademarks are property of said artist and their registered trademark holders. 🚫👮‍♂️🔒💼

Source inspiration taken from: [Teargas - Katatonia](https://www.youtube.com/watch?v=90NkngiWgqU). 🎵🎤🎧

## 🚀 Getting started

### Prerequisites

- Go 1.16 or higher 🚀
- A simple REST API to consume 🌐

### Installation

Download the latest version from the [releases page](https://github.com/VentGrey/teargas/releases)

or

1. Clone the repo
```sh
git clone https://github.com/VentGrey/teargas.git
```

2. Build the binary

```sh
cd teargas

# Build a normally compiled binary
go build -o teargas teargas.go

# Build a "stripped" binary
go build -o teargas -ldflags="-s -w" teargas.go
```

3. Run the program

```sh
./teargas -url <URL> -output [output file]
```

4. (Optionally build a Debian Package)

> Since `dh-make-golang` is walking crap. The method for building a Debian Package here is similar to the one used in Linux Mint.

``` sh

```

## 🤖 Usage

Basic usage:

```sh
teargas -url <URL> -output [output file]
```

Make a request using JWT authentication:

``` sh
teargas -url <URL> -output [output file] -username <USERNAME> -password <YOURPASSWORD> -authurl <AUTH_URL>
```

## 🤝 Contributing

Contributions are always welcome! 🤗...however I think this doesn't need much tho. Feel free to propose any changes you'd like :)

## 🐛 Bugs

Please report any bugs to omar@laesquinagris.com 📧🐛

## 📜 License

Distributed under the GPL-2+ License. See LICENSE for more information. 📜📝

## 📖 References
- [GitHub Actions](https://github.com/features/actions) 🤖🔍
- [JSON-iterator for Go](https://github.com/json-iterator/go) MIT licensed. 📃🔍
