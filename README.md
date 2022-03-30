# junit-merger

This is a fork of the original [junit-merger](github.com/imsky/junit-merger) project, updated to support [Go modules](https://go.dev/ref/mod) with the simplified project structure.

`junit-merger` merges many JUnit reports into one report.

## Usage

```text
$ junit-merger *.xml > merged.xml
$ junit-merger -o merged.xml a.xml b.xml c.xml
```

## Installation

`go install github.com/nezorflame/junit-merger@latest`

## License

`junit-merger` is provided under the [MIT license](https://opensource.org/licenses/MIT).

## Credit

Made by [Ivan Malopinsky](http://imsky.co).
Modified by [Ilya Danilkin](https://github.com/nezorflame).
