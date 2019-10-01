# 💉 Orninject – Helps to easily replace environment variables in file

[![CircleCI](https://badgen.net/circleci/github/ornikar/orninject/master)](https://circleci.com/gh/ornikar/orninject)
[![Go Report Card](https://goreportcard.com/badge/github.com/ornikar/orninject)](https://goreportcard.com/report/github.com/ornikar/orninject)
[![Codebeat badge](https://codebeat.co/badges/844e69d7-7534-445d-9b42-564fc49cd99e
)](https://codebeat.co/projects/github-com-ornikar-orninject-master)
[![GoDoc](https://godoc.org/github.com/ornikar/orninject?status.svg)](http://godoc.org/github.com/ornikar/orninject)
[![GitHub tag](https://img.shields.io/github/tag/ornikar/orninject.svg)](Tag)
[![Go version](https://img.shields.io/badge/go-v1.13-blue)](https://golang.org/dl/#stable)

Orninject is a tool that helps to easily replace environment variables in file.
The project is based on three philosophies *KISS*, *DRY* and *YAGNI*.

---

## Installation

Orninject is available on Linux, OSX and Windows platforms.

* Binaries for Mac OS, Linux and Windows are available as tarballs in the [release](https://github.com/ornikar/orninject/releases) page.

* Via Homebrew (Mac OS) or LinuxBrew (Linux)

   ```shell
   brew tap ornikar/orninject
   brew install orninject
   ```

* Building from source
   orninject was built using go 1.13 or above. In order to build orninject from source you must:
   1. Clone this repository
   2. Add the following command in your go.mod file

      ```text
      replace (
        github.com/ornikar/orninject => CLONED_GIT_REPOSITORY
      )
      ```

   3. Build the executable

        ```shell
        go build .
        ```

   4. Use it

        ```shell
        ./orninject
        ```
