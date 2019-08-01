# SQLite Manual

## Environment

- macOS Mojave 10.14.6
- Homebrew 2.1.8

## Install List

- sqlite
- gcc(Xcode)
- package(go-sqlite3)

## Setup

### sqlite

```$xslt
$ brew list | grep sqlite

$ brew install sqlite

$ sqlite3
SQLite version 3.24.0 2018-06-04 14:10:15
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> 
```

### gcc(Xcode)

- [Xcode](https://developer.apple.com/xcode/)

```$xslt
$ xcode-select --install

$ gcc --version
Configured with: --prefix=/Library/Developer/CommandLineTools/usr --with-gxx-include-dir=/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk/usr/include/c++/4.2.1
Apple LLVM version 10.0.1 (clang-1001.0.46.4)
Target: x86_64-apple-darwin18.7.0
Thread model: posix
InstalledDir: /Library/Developer/CommandLineTools/usr/bin
```

### package(go-sqlite3)

- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

```$xslt
$ go get github.com/mattn/go-sqlite3

```