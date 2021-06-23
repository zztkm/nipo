# nipo

日報

## Feature

- 日報リポジトリの初期化
	- nipo.jsonに独自のnipo書式を追加
- [sink](https://veltiosoft.dev/sink/)を開く
- 任意の文字列を置換
	- Markdown -> nipo
	- `## ` -> `●`

```json
{
    "header": {
        "# ": "★",
        "## ": "●"
    }
}
```

## Installation

```console
go install github.com/zztkm/nipo@latest
```

## Usage

- Initializing the `日報(Daily Report)` Repository

```console
nipo init

# create `nipo.json` file
```

- ヘッダーの変換を行う
	- 変換の内容は`nipo.json`に従う

```console
nipo converte 2021-04-08.md
```
