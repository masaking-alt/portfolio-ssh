# SSH Portfolio New Repo Docs

## 目的

ユーザーが自分のターミナルから次のように接続できる、SSH版ポートフォリオを新規リポジトリで作る。

```sh
ssh ssh.msking.net
```

Web版ポートフォリオは今のまま Cloudflare Pages で運用する。新規リポジトリでは、Webページではなく SSH で入れるポートフォリオ体験だけを扱う。

## ざっくり構成

- デプロイ先: Fly.io
- DNS: `ssh.msking.net` のようなサブドメインを Fly.io に向ける
- 接続方法: ユーザーは通常の `ssh` コマンドだけで接続する
- アプリ内容: SSH接続後に、ポートフォリオの自己紹介、作品一覧、作品詳細、連絡先を表示する
- 実装方針: 本物のLinuxシェルを渡すのではなく、ポートフォリオ専用のSSH/TUIアプリにする

## 利用ドメイン

Cloudflareで購入した `msking.net` を使う。

- Web版ポートフォリオ: Cloudflare Pages で運用
- SSH版ポートフォリオ: `ssh.msking.net`

SSH版の接続コマンド:

```sh
ssh ssh.msking.net
```

Cloudflare DNSでは、`ssh` サブドメインだけを Fly.io に向ける。SSHはHTTPではないので、Cloudflareのプロキシは使わず DNS only にする。

レコード例:

```txt
Type   Name   Content              Proxy status
A      ssh    <Fly.io IPv4>        DNS only
AAAA   ssh    <Fly.io IPv6>        DNS only
```

Fly.io側のIPは `fly ips list` で確認する。IPv4がない場合は `fly ips allocate-v4` で割り当てる。専用IPv4は月額課金になる可能性があるので、料金は作成前に確認する。

## 体験イメージ

接続直後にトップ画面を表示する。

```txt
MASAKING PORTFOLIO
AI-Native Full-Stack Developer

Hello,
I'm Masaking.

[1] Works
[2] About
[3] Contact
[q] Quit
```

作品一覧を開くと、現行ポートフォリオの Works をすべて表示する。

```txt
WORKS

10. Marple                              Webアプリ
  Marp形式のスライドをブラウザ上で直感的に編集できるWebアプリです。

9. Ugomemo                              Webアプリ
  アイデアをノードとして配置し、つなげながら発想を広げるためのコラボレーションツール。
```

作品を選ぶと詳細を表示する。

```txt
Marple
Category: Webアプリ
URL: https://marple.hackuniv.club/

Description:
Marp形式のスライドをブラウザ上で直感的に編集できるWebアプリです...

Technologies:
Next.js / React / TypeScript / Tailwind CSS / Marp Core / OpenAI API / Clerk / Cloudflare Workers / Cloudflare D1 / チーム開発
```

## 画面・メニュー

最低限ほしい画面は次の通り。

- Home: ヒーロー相当の自己紹介
- Works: 作品一覧
- Work Detail: 作品詳細
- About: 自己紹介本文
- Contact: メール、SNSリンク
- Help: 操作キーの簡単な一覧

操作はシンプルでよい。

- `↑` / `↓` / `←` / `→`: 選択移動
- `Enter`: 決定
- `Esc` / `Backspace`: 戻る
- `q`: 終了

## 現行サイト基本情報

- サイト名: Portfolio
- 表示名: まさきんぐ_masaking
- 英語名: Masaking
- 肩書き: AI-Native Full-Stack Developer
- ヒーローテキスト:
  - Hello,
  - I'm Masaking.
- 背景大文字テキスト:
  - PORTFOLIO
  - WORKS
  - ABOUT
  - CONTACT
- サイド装飾テキスト:
  - MASAKING PORTFOLIO
  - ///
  - AI-Native Full-Stack Developer
- ナビゲーション:
  - Top
  - Works
  - About
  - Contact
- フッター:
  - © masaking

## SEO・メタ情報

Web版から内容を引き継ぐ場合の参考。

- title: `Portfolio | Masaking`
- description: `まさきんぐのポートフォリオサイト。AIを活用しながら、フロントエンドからバックエンドまでWeb開発に取り組んでいます。`
- theme color: `#143438`
- og:type: `website`
- og:title: `Portfolio | Masaking`
- og:description: `AIを活用しながら、フロントエンドからバックエンドまでWeb開発に取り組むポートフォリオサイト。`
- og:image: `https://masaking.pages.dev/OGP.png`
- og:image:width: `1200`
- og:image:height: `630`
- og:site_name: `Portfolio`
- og:url: `https://masaking.pages.dev/`
- twitter:card: `summary_large_image`
- twitter:title: `Portfolio | Masaking`
- twitter:description: `AIを活用しながら、フロントエンドからバックエンドまでWeb開発に取り組むポートフォリオサイト。`
- twitter:image: `https://masaking.pages.dev/OGP.png`

## About

### リード文

高知大学で情報科学を専攻。  
バイブコーディングに没頭する日々。

モバイル版では次の改行で表示している。

```txt
高知大学で
情報科学を専攻。
バイブコーディングに
没頭する日々。
```

### 本文

昔からプログラミングに興味があり、AIを利用することでそのハードルが下がったため、web制作に挑戦しています。ユーザーにとって魅力的で使いやすいサイトを作れるようになりたいです。

好きなゲームはモンハンで、特に3DSのシリーズが好きです。好きなアーティストはVaundy。アニメやマンガを見るのも好きです。


### Built with

見出し:

- Built with

説明:

- このサイトは以下の技術を使用して制作しました

技術タグ:

- HTML
- CSS
- JavaScript (React)
- Vite
- React Router
- ESLint
- Node.js
- npm

## Contact

表示文:

- Let's create something together.

メール:

- `banbenjianggui@gmail.com`

SNS:

- Instagram: `https://www.instagram.com/masa.ki8904?igsh=b3hqMjd6aHdnazJ5`
- X (Twitter): `https://x.com/masaking_alt`
- BeReal: `https://bere.al/masaki9876`

## Works

現行サイトでは ID の降順で表示している。

### 10. Marple

- category: Webアプリ
- image: `/marple.jpeg`
- externalUrl: `https://marple.hackuniv.club/`
- technologies:
  - Next.js
  - React
  - TypeScript
  - Tailwind CSS
  - Marp Core
  - OpenAI API
  - Clerk
  - Cloudflare Workers
  - Cloudflare D1
  - チーム開発

description:

Marp形式のスライドをブラウザ上で直感的に編集できるWebアプリです。Markdown編集とスライドプレビューを同時に確認でき、AIによるスライド編集支援、gif風バージョン管理機能、PDF/PPTX形式でのエクスポートまで一貫して行えます。VSCode拡張やローカル環境構築なしで、プレゼン資料作成を効率化できる点が特徴です。

### 9. Ugomemo

- category: Webアプリ
- image: `/ugomemo.jpeg`
- externalUrl: `https://ugomemo.ugomemo.workers.dev/`
- technologies:
  - Next.js
  - React
  - Clerk
  - Cloudflare Workers
  - Cloudflare D1
  - SQLite
  - チーム開発

description:

アイデアをノードとして配置し、つなげながら発想を広げるためのコラボレーションツール。直感的なキャンバス、AIによる視点提案、学生・クリエイター・プランナー向けのユースケースを1ページで伝えるランディング構成にしています。

### 8. MESI-KO

- category: Webアプリ
- image: `/mesi-ko.jpeg`
- externalUrl: `https://mesi-ko.vercel.app/`
- technologies:
  - Next.js
  - React
  - NextAuth
  - Vercel
  - Neon
  - PostgreSQL
  - チーム開発

description:

ごはんの予定調整に特化した日程調整アプリ。月間カレンダーから日付ごとのイベント作成へ進め、参加候補数や回答率をサイドパネルで確認できます。招待コード参加やログイン導線も備え、少人数の予定を軽くまとめられる構成にしました。

### 7. Selection Charcount

- category: Chrome拡張
- image: `/selection_charcount.jpeg`
- externalUrl: `https://chromewebstore.google.com/detail/selection-charcount/kmjbolcinnpdiokkieehndnmnbhnmokk`
- technologies:
  - Chrome Extension (Manifest V3)
  - JavaScript
  - Chrome Scripting API

description:

右クリックメニューから呼び出せる選択文字数カウンター。選択中テキストをコンテキストメニューイベントで受け取り、余分な空白や絵文字を除外した上で文字数を計算し、画面右下に浮遊パネルとして描画します。Service Worker ベースの background と scripting API を使って任意ページにインジェクトできるので、Web 上のライティングやレビューで即座にカウントできます。

### 6. Moodle Enhancer for Kochi University

- category: Chrome拡張
- image: `/moodle-enhancer.jpeg`
- externalUrl: `https://chromewebstore.google.com/detail/moodle-dashboard-tweaks/jmnmogonkjhmhgcncebieodddgbaamfg`
- technologies:
  - Chrome Extension (Manifest V3)
  - JavaScript
  - Content Script
  - CSS

description:

高知大学 Moodle を快適に使うための Chrome 拡張。ダッシュボードやコース画面をダークテーマで再構成し、課題・イベントを色分けカードで強調。リソースを自動収集する「資料」タブやメディア保存ボタン、クイズ画面の視認性改善など、学習体験を底上げする施策をコンテンツスクリプトでまとめて提供しています。

### 5. 総合映像研究会ホームページ

- category: Webサイト
- image: `/soueiken-hp.jpeg`
- externalUrl: `https://soueiken-hp.pages.dev/`
- technologies:
  - React
  - Fetch API
  - CSS
  - 自作 Markdown レンダラー
  - Cloudflare Pages
  - Firestore
  - チーム開発

description:

高知大学・総合映像研究会の公式サイトです。サークルの雰囲気を一気に伝えるヒーローセクションやサイドバー、活動予定・BBS・コンタクトなどの複数セクションを整備し、Markdown で管理する活動レポートをトップの一覧と詳細ページに分離して運用できる構成にしています。

### 4. Part-time Shift

- category: Webアプリ
- image: `/parttime_shift.jpeg`
- externalUrl: `https://parttime-shift.pages.dev/`
- technologies:
  - React
  - TypeScript
  - Vite
  - Flask
  - SQLite
  - Cloudflare Pages
  - Render

description:

スマホに最適化されたフルスタックのシフト管理アプリ。日付と勤務時間を登録すると、ダッシュボードで次の勤務や今月の労働時間・給与見込みを自動集計。カレンダー表示や csv エクスポートにも対応し、自分の予定をまとめて管理できます。

### 3. Memo Pad

- category: Webアプリ
- image: `/memoapp.jpeg`
- externalUrl: `https://memoapp-5rh.pages.dev/`
- technologies:
  - React
  - Testing Library
  - Jest
  - Cloudflare Pages

description:

このメモ帳アプリは、シンプルな操作でメモを作成・編集できるのはもちろん、Firebase 認証を活用して端末をまたいだ同期まで自動で行ってくれるクラウド対応メモツールです。メールアドレスを使わずにユーザー名でログインでき、さらにワンクリックの Google ログインにも対応。ログインすると異なるデバイス間でリアルタイムにメモが反映されるため、クリップボードとしても使うことができます。軽快な UI に加え、安心の認証と自動バックアップ機能を備えた、日常使いにぴったりの Web メモアプリです。

### 2. Todo App

- category: Webアプリ
- image: `/todo_app.jpeg`
- externalUrl: `https://todo-app-d6h.pages.dev/`
- technologies:
  - React
  - TypeScript
  - Testing Library
  - Cloudflare Pages
  - CSS

description:

シンプルなTodoアプリ。非常にシンプルながら欲しい機能は揃っています。フィルター機能も搭載し、全てのタスク、未完了のタスク、完了済みのタスクを簡単に切り替えられます。

### 1. Gemini謹製 夏のひまつぶしコレクション

- category: ゲーム
- image: `/gemini_hima.jpeg`
- externalUrl: `https://himatubusi.pages.dev/`
- technologies:
  - HTML
  - CSS
  - JavaScript
  - Cloudflare Pages
  - チーム開発

description:

夏休みといえば？そう、暇。とにかく暇。その暇潰しの一助となるサイト。数多くのミニゲームを遊ぶことができます。楽しさを求めず、ただ時間をつぶしたい人向け。本当にしょうもないゲームをたくさん制作しました。

## 画像・公開ファイル

新規SSHリポでは画像表示は必須ではない。必要なら作品詳細で画像URLだけ表示する。

現行ポートフォリオの公開ファイル:

- `/profile.jpg`
- `/OGP.png`
- `/favicon.png`
- `/apple-touch-icon.png`
- `/marple.jpeg`
- `/ugomemo.jpeg`
- `/mesi-ko.jpeg`
- `/selection_charcount.jpeg`
- `/moodle-enhancer.jpeg`
- `/soueiken-hp.jpeg`
- `/parttime_shift.jpeg`
- `/memoapp.jpeg`
- `/todo_app.jpeg`
- `/gemini_hima.jpeg`
- `/robots.txt`
- `/sitemap.xml`
- `/_redirects`

## 現行ルーティング

Web版のルート:

- `/`: トップ、作品一覧、About、Contact
- `/work/:id`: 作品詳細
- `/#hero`: Top
- `/#works`: Works
- `/#about`: About
- `/#contact`: Contact

SSH版ではURLルーティングは不要。内部画面として Home / Works / Work Detail / About / Contact を持てばよい。

## 表示上の優先順位

SSH版で最初に見せたい情報:

1. `MASAKING PORTFOLIO`
2. `AI-Native Full-Stack Developer`
3. `Hello, I'm Masaking.`
4. Works への導線
5. Contact への導線

作品一覧では、各作品について次を表示する。

1. ID
2. タイトル
3. カテゴリ
4. 説明の冒頭
5. 使用技術の先頭1〜3個

作品詳細では、各作品について次を表示する。

1. タイトル
2. カテゴリ
3. URL
4. 説明全文
5. 使用技術全文

## 実装メモ

言語は Go か TypeScript のどちらかにする。

Goを触ってみたいので、第一候補は Go。

- Go: Wish + Bubble Tea + Lip Gloss + Bubbles
- TypeScript: ssh2 + Ink または terminal-kit

大事なのは、ログイン後に本物のシェルを開かせないこと。ポートフォリオ表示専用のアプリとして動かす。

Fly.ioでは外向きに SSH 用のTCPポートを開ける。アプリ内部では `2222` のような非特権ポートで待ち受け、Fly.io側で外部 `22` に割り当てる。

DNSは `ssh.msking.net` を Fly.io のIPに向ける。Cloudflare DNSを使う場合、SSH用サブドメインはプロキシせず DNS only にする。

## READMEに最初に載せる文案

````md
# SSH Portfolio

Masaking のポートフォリオを、ブラウザではなく SSH で表示するためのアプリです。

```sh
ssh ssh.msking.net
```

Web版ポートフォリオは Cloudflare Pages で運用し、このリポジトリでは SSH 接続後のポートフォリオ体験だけを管理します。
````
