# iaas-api-go

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/iaas-api-go.svg)](https://pkg.go.dev/github.com/sacloud/iaas-api-go)
[![Tests](https://github.com/sacloud/iaas-api-go/workflows/Tests/badge.svg)](https://github.com/sacloud/iaas-api-go/actions/workflows/tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/iaas-api-go)](https://goreportcard.com/report/github.com/sacloud/iaas-api-go)

Go言語向けのさくらのクラウドIaaS APIライブラリ

## 概要

iaas-api-goは[sacloud/libsacloud v2](https://github.com/sacloud/libsacloud)の後継プロジェクトで、さくらのクラウド APIのうちのIaaS部分を担当します。

概要/設計/実装方針: [docs/overview.md](https://github.com/sacloud/iaas-api-go/blob/main/docs/design/overview.md)

### libsacloudとiaas-api-goのバージョン対応表

| libsacloud | iaas-api-go | Note/Status                       |
|------------|-------------|-----------------------------------|
| v1         | -           | libsacloud v1系はiaas-api-goへの移植対象外 |
| v2         | v1          | 開発中                               |
| v3(未リリース)  | v2          | 未リリース/未着手                         |


### 関連プロジェクト

- [sacloud/iaas-service-go](https://github.com/sacloud/iaas-service-go): sacloud/iaas-api-goを用いた高レベルAPIライブラリ
- [sacloud/api-client-go](https://github.com/sacloud/api-client-go): sacloudプロダクト向けHTTP/APIクライアントライブラリ(環境変数やプロファイルの処理など)
- [sacloud/packages-go](https://github.com/sacloud/packages-go): sacloudプロダクト向けの汎用パッケージ群

## License

`sacloud/iaas-api-go` Copyright (C) 2022-2023 [The sacloud/iaas-api-go Authors](AUTHORS).

This project is published under [Apache 2.0 License](LICENSE).
