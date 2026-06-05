package portfolio

type Profile struct {
	SiteName       string
	DisplayName    string
	EnglishName    string
	Title          string
	HeroLines      []string
	AboutLead      string
	AboutBody      []string
	BuiltWithIntro string
	BuiltWith      []string
	ContactLead    string
	Email          string
	Socials        []Social
	Works          []Work
}

type Social struct {
	Name string
	URL  string
}

type Work struct {
	ID           int
	Title        string
	Category     string
	ImagePath    string
	ExternalURL  string
	Technologies []string
	Description  string
}

func DefaultProfile() Profile {
	return Profile{
		SiteName:    "Portfolio",
		DisplayName: "まさきんぐ_masaking",
		EnglishName: "Masaking",
		Title:       "AI-Native Full-Stack Developer",
		HeroLines:   []string{"Hello,", "I'm Masaking."},
		AboutLead:   "高知大学で情報科学を専攻。バイブコーディングに没頭する日々。",
		AboutBody: []string{
			"昔からプログラミングに興味があり、AIを利用することでそのハードルが下がったため、web制作に挑戦しています。ユーザーにとって魅力的で使いやすいサイトを作れるようになりたいです。",
			"好きなゲームはモンハンで、特に3DSのシリーズが好きです。好きなアーティストはVaundy。アニメやマンガを見るのも好きです。",
		},
		BuiltWithIntro: "このサイトは以下の技術を使用して制作しました",
		BuiltWith: []string{
			"HTML",
			"CSS",
			"JavaScript (React)",
			"Vite",
			"React Router",
			"ESLint",
			"Node.js",
			"npm",
		},
		ContactLead: "Let's create something together.",
		Email:       "banbenjianggui@gmail.com",
		Socials: []Social{
			{Name: "Instagram", URL: "https://www.instagram.com/masa.ki8904?igsh=b3hqMjd6aHdnazJ5"},
			{Name: "X (Twitter)", URL: "https://x.com/masaking_alt"},
			{Name: "BeReal", URL: "https://bere.al/masaki9876"},
		},
		Works: []Work{
			{
				ID:          10,
				Title:       "Marple",
				Category:    "Webアプリ",
				ImagePath:   "/marple.jpeg",
				ExternalURL: "https://marple.hackuniv.club/",
				Technologies: []string{
					"Next.js",
					"React",
					"TypeScript",
					"Tailwind CSS",
					"Marp Core",
					"OpenAI API",
					"Clerk",
					"Cloudflare Workers",
					"Cloudflare D1",
					"チーム開発",
				},
				Description: "Marp形式のスライドをブラウザ上で直感的に編集できるWebアプリです。Markdown編集とスライドプレビューを同時に確認でき、AIによるスライド編集支援、gif風バージョン管理機能、PDF/PPTX形式でのエクスポートまで一貫して行えます。VSCode拡張やローカル環境構築なしで、プレゼン資料作成を効率化できる点が特徴です。",
			},
			{
				ID:          9,
				Title:       "Ugomemo",
				Category:    "Webアプリ",
				ImagePath:   "/ugomemo.jpeg",
				ExternalURL: "https://ugomemo.ugomemo.workers.dev/",
				Technologies: []string{
					"Next.js",
					"React",
					"Clerk",
					"Cloudflare Workers",
					"Cloudflare D1",
					"SQLite",
					"チーム開発",
				},
				Description: "アイデアをノードとして配置し、つなげながら発想を広げるためのコラボレーションツール。直感的なキャンバス、AIによる視点提案、学生・クリエイター・プランナー向けのユースケースを1ページで伝えるランディング構成にしています。",
			},
			{
				ID:          8,
				Title:       "MESI-KO",
				Category:    "Webアプリ",
				ImagePath:   "/mesi-ko.jpeg",
				ExternalURL: "https://mesi-ko.vercel.app/",
				Technologies: []string{
					"Next.js",
					"React",
					"NextAuth",
					"Vercel",
					"Neon",
					"PostgreSQL",
					"チーム開発",
				},
				Description: "ごはんの予定調整に特化した日程調整アプリ。月間カレンダーから日付ごとのイベント作成へ進め、参加候補数や回答率をサイドパネルで確認できます。招待コード参加やログイン導線も備え、少人数の予定を軽くまとめられる構成にしました。",
			},
			{
				ID:          7,
				Title:       "Selection Charcount",
				Category:    "Chrome拡張",
				ImagePath:   "/selection_charcount.jpeg",
				ExternalURL: "https://chromewebstore.google.com/detail/selection-charcount/kmjbolcinnpdiokkieehndnmnbhnmokk",
				Technologies: []string{
					"Chrome Extension (Manifest V3)",
					"JavaScript",
					"Chrome Scripting API",
				},
				Description: "右クリックメニューから呼び出せる選択文字数カウンター。選択中テキストをコンテキストメニューイベントで受け取り、余分な空白や絵文字を除外した上で文字数を計算し、画面右下に浮遊パネルとして描画します。Service Worker ベースの background と scripting API を使って任意ページにインジェクトできるので、Web 上のライティングやレビューで即座にカウントできます。",
			},
			{
				ID:          6,
				Title:       "Moodle Enhancer for Kochi University",
				Category:    "Chrome拡張",
				ImagePath:   "/moodle-enhancer.jpeg",
				ExternalURL: "https://chromewebstore.google.com/detail/moodle-dashboard-tweaks/jmnmogonkjhmhgcncebieodddgbaamfg",
				Technologies: []string{
					"Chrome Extension (Manifest V3)",
					"JavaScript",
					"Content Script",
					"CSS",
				},
				Description: "高知大学 Moodle を快適に使うための Chrome 拡張。ダッシュボードやコース画面をダークテーマで再構成し、課題・イベントを色分けカードで強調。リソースを自動収集する「資料」タブやメディア保存ボタン、クイズ画面の視認性改善など、学習体験を底上げする施策をコンテンツスクリプトでまとめて提供しています。",
			},
			{
				ID:          5,
				Title:       "総合映像研究会ホームページ",
				Category:    "Webサイト",
				ImagePath:   "/soueiken-hp.jpeg",
				ExternalURL: "https://soueiken-hp.pages.dev/",
				Technologies: []string{
					"React",
					"Fetch API",
					"CSS",
					"自作 Markdown レンダラー",
					"Cloudflare Pages",
					"Firestore",
					"チーム開発",
				},
				Description: "高知大学・総合映像研究会の公式サイトです。サークルの雰囲気を一気に伝えるヒーローセクションやサイドバー、活動予定・BBS・コンタクトなどの複数セクションを整備し、Markdown で管理する活動レポートをトップの一覧と詳細ページに分離して運用できる構成にしています。",
			},
			{
				ID:          4,
				Title:       "Part-time Shift",
				Category:    "Webアプリ",
				ImagePath:   "/parttime_shift.jpeg",
				ExternalURL: "https://parttime-shift.pages.dev/",
				Technologies: []string{
					"React",
					"TypeScript",
					"Vite",
					"Flask",
					"SQLite",
					"Cloudflare Pages",
					"Render",
				},
				Description: "スマホに最適化されたフルスタックのシフト管理アプリ。日付と勤務時間を登録すると、ダッシュボードで次の勤務や今月の労働時間・給与見込みを自動集計。カレンダー表示や csv エクスポートにも対応し、自分の予定をまとめて管理できます。",
			},
			{
				ID:          3,
				Title:       "Memo Pad",
				Category:    "Webアプリ",
				ImagePath:   "/memoapp.jpeg",
				ExternalURL: "https://memoapp-5rh.pages.dev/",
				Technologies: []string{
					"React",
					"Testing Library",
					"Jest",
					"Cloudflare Pages",
				},
				Description: "このメモ帳アプリは、シンプルな操作でメモを作成・編集できるのはもちろん、Firebase 認証を活用して端末をまたいだ同期まで自動で行ってくれるクラウド対応メモツールです。メールアドレスを使わずにユーザー名でログインでき、さらにワンクリックの Google ログインにも対応。ログインすると異なるデバイス間でリアルタイムにメモが反映されるため、クリップボードとしても使うことができます。軽快な UI に加え、安心の認証と自動バックアップ機能を備えた、日常使いにぴったりの Web メモアプリです。",
			},
			{
				ID:          2,
				Title:       "Todo App",
				Category:    "Webアプリ",
				ImagePath:   "/todo_app.jpeg",
				ExternalURL: "https://todo-app-d6h.pages.dev/",
				Technologies: []string{
					"React",
					"TypeScript",
					"Testing Library",
					"Cloudflare Pages",
					"CSS",
				},
				Description: "シンプルなTodoアプリ。非常にシンプルながら欲しい機能は揃っています。フィルター機能も搭載し、全てのタスク、未完了のタスク、完了済みのタスクを簡単に切り替えられます。",
			},
			{
				ID:          1,
				Title:       "Gemini謹製 夏のひまつぶしコレクション",
				Category:    "ゲーム",
				ImagePath:   "/gemini_hima.jpeg",
				ExternalURL: "https://himatubusi.pages.dev/",
				Technologies: []string{
					"HTML",
					"CSS",
					"JavaScript",
					"Cloudflare Pages",
					"チーム開発",
				},
				Description: "夏休みといえば？そう、暇。とにかく暇。その暇潰しの一助となるサイト。数多くのミニゲームを遊ぶことができます。楽しさを求めず、ただ時間をつぶしたい人向け。本当にしょうもないゲームをたくさん制作しました。",
			},
		},
	}
}
