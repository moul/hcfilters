# hcfilters

:smile: filters for [github.com/gregjones/httpcache](https://github.com/gregjones/httpcache)

[![CircleCI](https://circleci.com/gh/moul/hcfilters.svg?style=shield)](https://circleci.com/gh/moul/hcfilters)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/moul.io/hcfilters)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/hcfilters/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/hcfilters.svg)](https://github.com/moul/hcfilters/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/hcfilters)](https://goreportcard.com/report/moul.io/hcfilters)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/hcfilters/badge)](https://www.codefactor.io/repository/github/moul/hcfilters)
[![codecov](https://codecov.io/gh/moul/hcfilters/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/hcfilters)
[![GolangCI](https://golangci.com/badges/github.com/moul/hcfilters.svg)](https://golangci.com/r/github.com/moul/hcfilters)
[![Sourcegraph](https://sourcegraph.com/github.com/moul/hcfilters/-/badge.svg)](https://sourcegraph.com/github.com/moul/hcfilters?badge)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```console
$ go get -u moul.io/hcfilters
```

```golang
// before
client := &http.Client{
	Transport: httpcache.NewTransport(
		diskcache.New(diskcachePath),
	),
}

// after
client := &http.Client{
	Transport: httpcache.NewTransport(
		hcfilters.MaxSize( // skip caching results > 2Mb
			diskcache.New(diskcachePath),
			2*1024*1024,
		),
	),
}
```

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
