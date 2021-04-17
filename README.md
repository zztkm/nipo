# nipo

日報


## Feature

- 日報リポジトリの初期化
	- config.jsonに独自のnipo書式を追加
	- jsonはhttps://github.com/zztkm/dynamic-jsonでパースされる
- [sink](https://veltiosoft.dev/sink/)を開く
- Markdown書式をnipo書式に変換

```json
{
	"## ": "●"
}
```

## Installation

```console
go get github.landscape.co.jp/landbox/nipo/cmd/nipo
```

## Usage

- Initializing the `日報(Daily Report)` Repository

```console
nipo init
```

- converteした内容をファイルに書き込む

```console
nipo converte 2021-04-08.md > 2021-04-08.nipo
```
