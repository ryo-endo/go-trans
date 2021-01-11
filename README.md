go-trans
=======

[![CircleCI](https://circleci.com/gh/ryo-endo/go-trans.svg?style=svg)](https://circleci.com/gh/ryo-endo/go-trans)
[![Coverage Status](https://coveralls.io/repos/github/ryo-endo/go-trans/badge.svg)](https://coveralls.io/github/ryo-endo/go-trans)

## Description

CLI tool for translation.

## Installation

```sh
% go get github.com/ryo-endo/go-trans
```

## Usage

```sh
> go-trans Hello World.
ハローワールド。

> go-trans -from ja -to en ハローワールド。
Hello world.

> go-trans -i
Hello
こんにちは
```

## Options
Usage of ./go-trans:
```
  -from string
    	Set the language code. "en" "ja" "vn" (default "en")
  -to string
    	Set the language code. "en" "ja" "vn" (default "ja")
  -i	interactive mode.
```

## Author

[ryo-endo](https://github.com/ryo-endo)