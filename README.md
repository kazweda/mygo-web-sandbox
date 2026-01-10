# go-web-sandbox

A minimal sandbox repository for hackathons.

## Purpose
- Static HTML generation using Go
- Designed for GitHub Pages deployment
- Interactive features are intentionally left unimplemented

## License
MIT

## Go Web Sandbox

Go + templ を使った、ハッカソン向け Web サンドボックスです。
擬似 SSG で静的 HTML を生成し、GitHub Pages にデプロイします。

### Requirements

- Go 1.22+
- templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Getting Started (Safe Setup)

このリポジトリは **sandbox 配布用**です。  
**直接このリポジトリに push しないでください。**

必ず **自分用のリポジトリを作成してから**作業してください。

---

### 方法 A: 新しいリポジトリを作成して clone（最も安全）

#### 1. GitHub で新しいリポジトリを作成
- 任意の名前（例: `my-go-web-sandbox`）
- Public / Private は自由

#### 2. このリポジトリを clone
```bash
git clone https://github.com/<original-owner>/go-web-sandbox.git
cd go-web-sandbox
```

#### 3. origin を自分のリポジトリに付け替える
```bash
git remote remove origin
git remote add origin https://github.com/<your-name>/my-go-web-sandbox.git
git push -u origin main
```

これ以降の push は **自分のリポジトリ**にのみ反映されます。

---

### 方法 B: Fork（GitHub に慣れている場合）

1. GitHub の **Fork** ボタンを押す
2. fork 先のリポジトリを clone
3. fork 先で自由に Issue / Push を行う

---

> ⚠️ このリポジトリはテンプレート用途です  
> 作業は必ず自分のリポジトリで行ってください
