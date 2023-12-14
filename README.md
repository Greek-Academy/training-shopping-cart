# Hans-on

イマイチなコードを作成しました。
みなさんの力で「良いコード」へ書き換えましょう。

※ 余裕があれば、ユニットテストの実行・修正にもチャレンジしましょう。

## Problem

[ショッピングカートのクラス](src/shoppingCart.ts)と、その関係処理をリファクタリングしましょう。

## Documents

[Design Document](docs/design-document.md)

## Setup

下記言語の環境を用意しました。
他言語で実装したい方は、仕様書に沿って「良いコード」を作成してみてください。

- TypeScript
- Go

### TypeScript

`Node.js` をインストールしておいてください
https://nodejs.org/en

1. 作業ディレクトリへ移動
```bash
cd typescript
```

2. パッケージのインストール（下記いずれか）

- npm
- [pnpm](https://pnpm.io/installation)
- [yarn](https://classic.yarnpkg.com/lang/en/docs/install/)

```bash
# npm
npm install

# pnpm
pnpm install

# yarn
yarn install
```

### Go

`Go` をインストールしておいてください
https://go.dev/doc/install

1. 作業ディレクトリへ
```bash
cd go
```

## Unit Test

メソッド名、引数、返り値等を変更した場合、テストが失敗する可能性があります。
その際は、変更に合わせてテストコードを修正していください。

### TypeScript

※ `typescript` ディレクトリで実行してください

```bash
# npm
npx jest

# pnpm
pnpm jest

# yarn
yarn jest
```

### Go

※ `go` ディレクトリで実行してください

```bash
go test
```

## Note

- [説明スライド Good Code, Bad Code...](https://docs.google.com/presentation/d/1v-eZWwXChFjBmbYOXALoMWh9gasn53Pd9-gr0lWzCUE)
