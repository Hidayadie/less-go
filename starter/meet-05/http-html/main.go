package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Article struct {
	Slug    string   `json:"slug"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

type pageData struct {
	Title    string
	Articles []Article
	Article  *Article // nil = halaman daftar
}

var pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="id">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
</head>
<body>
  <header><a href="/">HTTP Wiki</a></header>
  <main>
  {{with .Article}}
    <article>
      <h1>{{.Title}}</h1>
      <p>{{.Summary}}</p>
      <p>{{range .Tags}}<span>#{{.}}</span> {{end}}</p>
    </article>
    <p><a href="/articles">&larr; Semua artikel</a></p>
  {{else}}
    <h1>{{.Title}}</h1>
    <ul>
    {{range .Articles}}
      <li>
        <a href="/articles/{{.Slug}}">{{.Title}}</a> - {{.Summary}}
        {{range .Tags}}<span>#{{.}}</span> {{end}}
      </li>
    {{end}}
    </ul>
  {{end}}
  </main>
</body>
</html>
`))

func main() {
	articles, err := loadArticles("data/articles.json")
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + getenv("PORT", "8000")
	srv := &http.Server{
		Addr:              addr,
		Handler:           logRequests(newMux(articles)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("server listening on", addr)
	log.Fatal(srv.ListenAndServe())
}

func newMux(articles []Article) http.Handler {
	mux := http.NewServeMux()

	// {$} = hanya cocok persis dengan "/", path lain otomatis 404
	mux.HandleFunc("GET /{$}", listHandler("HTTP Wiki", articles))
	mux.HandleFunc("GET /articles", listHandler("Artikel", articles))

	mux.HandleFunc("GET /articles/{slug}", func(w http.ResponseWriter, r *http.Request) {
		article, ok := findArticle(articles, r.PathValue("slug"))
		if !ok {
			http.NotFound(w, r)
			return
		}
		render(w, http.StatusOK, pageData{Title: article.Title, Article: article})
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "ok")
	})

	return mux
}

func listHandler(title string, articles []Article) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, http.StatusOK, pageData{Title: title, Articles: articles})
	}
}

func loadArticles(path string) ([]Article, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()

	var articles []Article
	if err := json.NewDecoder(file).Decode(&articles); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	if len(articles) == 0 {
		return nil, errors.New("articles data is empty")
	}

	seen := make(map[string]bool, len(articles))
	for i, a := range articles {
		if a.Slug == "" {
			return nil, fmt.Errorf("article #%d has empty slug", i)
		}
		if seen[a.Slug] {
			return nil, fmt.Errorf("duplicate slug %q", a.Slug)
		}
		seen[a.Slug] = true
	}

	return articles, nil
}

func findArticle(articles []Article, slug string) (*Article, bool) {
	for i := range articles {
		if articles[i].Slug == slug {
			return &articles[i], true
		}
	}
	return nil, false
}

// render menulis ke buffer dulu, supaya kalau template gagal
// klien tidak menerima halaman setengah jadi dengan status 200.
func render(w http.ResponseWriter, status int, data pageData) {
	var buf bytes.Buffer
	if err := pageTemplate.Execute(&buf, data); err != nil {
		log.Println("render template:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
