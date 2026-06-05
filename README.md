# SSH Portfolio

Masaking のポートフォリオを、ブラウザではなく SSH で表示するためのアプリです。

```sh
ssh ssh.msking.net
```

Web版ポートフォリオは Cloudflare Pages で運用し、このリポジトリでは SSH 接続後のポートフォリオ体験だけを管理します。

## 開発

Go、Wish、Bubble Tea、Lip Gloss、Bubbles で実装しています。

```sh
go run ./cmd/server
```

別ターミナルから接続します。

```sh
ssh -p 2222 localhost
```

## 環境変数

| 名前 | 既定値 | 内容 |
| --- | --- | --- |
| `HOST` | `0.0.0.0` | SSHサーバの待ち受けホスト |
| `PORT` | `2222` | SSHサーバの内部待ち受けポート |
| `HOST_KEY_PATH` | `.ssh/ssh_host_ed25519_key` | SSHホスト鍵の保存先 |

## デプロイ

Fly.io にTCPサービスとしてデプロイし、外部ポート `22` をアプリ内部の `2222` に割り当てます。

```sh
fly deploy
```

Cloudflare DNS では `ssh.msking.net` を Fly.io のIPに向け、プロキシは使わず DNS only にします。
