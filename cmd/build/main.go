package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/kazweda/mygo-web-sandbox/templates"
)

func main() {
	ctx := context.Background()

	loc, err := fetchLocation(ctx)
	if err != nil {
		panic(err)
	}

	weather, err := fetchWeather(ctx, loc.Lat, loc.Lon)
	if err != nil {
		panic(err)
	}

	outDir := "docs"
	outFile := filepath.Join(outDir, "index.html")

	if err := os.MkdirAll(outDir, 0755); err != nil {
		panic(err)
	}

	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := templates.Page(loc, weather).Render(ctx, f); err != nil {
		panic(err)
	}
}

const (
	issueOwner = "kazweda"
	issueRepo  = "mygo-web-sandbox"
	issueID    = 1
)

type issuePayload struct {
	Location templates.Location `json:"location"`
}

func fetchLocation(ctx context.Context) (templates.Location, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", issueOwner, issueRepo, issueID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return templates.Location{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return templates.Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
		return templates.Location{}, fmt.Errorf("github issue fetch failed: status=%d body=%s", res.StatusCode, string(b))
	}

	var payload struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return templates.Location{}, err
	}

	parsed, err := parseIssueBody(payload.Body)
	if err != nil {
		return templates.Location{}, err
	}

	return parsed.Location, nil
}

var jsonBlockRE = regexp.MustCompile("(?s)```json\\s*(.*?)\\s*```")

func parseIssueBody(body string) (issuePayload, error) {
	blocks := jsonBlockRE.FindAllStringSubmatch(body, -1)
	if len(blocks) == 0 {
		// fallback to whole body if it might be plain JSON
		return decodeIssuePayload(body)
	}

	for _, m := range blocks {
		if len(m) < 2 {
			continue
		}
		if p, err := decodeIssuePayload(m[1]); err == nil {
			return p, nil
		}
	}

	return issuePayload{}, errors.New("no valid json block found in issue body")
}

func decodeIssuePayload(raw string) (issuePayload, error) {
	var p issuePayload
	if err := json.Unmarshal([]byte(strings.TrimSpace(raw)), &p); err != nil {
		return issuePayload{}, err
	}
	return p, nil
}

// Open-Meteo: Fetch current weather for given lat/lon
func fetchWeather(ctx context.Context, lat, lon float64) (templates.Weather, error) {
	base := "https://api.open-meteo.com/v1/forecast"
	q := url.Values{}
	q.Set("latitude", fmt.Sprintf("%f", lat))
	q.Set("longitude", fmt.Sprintf("%f", lon))
	q.Set("current_weather", "true")
	u := base + "?" + q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return templates.Weather{}, err
	}

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return templates.Weather{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
		return templates.Weather{}, fmt.Errorf("open-meteo fetch failed: status=%d body=%s", res.StatusCode, string(b))
	}

	var om struct {
		CurrentWeather struct {
			Temperature   float64 `json:"temperature"`
			Windspeed     float64 `json:"windspeed"`
			Winddirection float64 `json:"winddirection"`
			Weathercode   int     `json:"weathercode"`
			IsDay         int     `json:"is_day"`
			Time          string  `json:"time"`
		} `json:"current_weather"`
	}

	if err := json.NewDecoder(res.Body).Decode(&om); err != nil {
		return templates.Weather{}, err
	}

	return templates.Weather{
		Temperature:   om.CurrentWeather.Temperature,
		Windspeed:     om.CurrentWeather.Windspeed,
		Winddirection: om.CurrentWeather.Winddirection,
		Weathercode:   om.CurrentWeather.Weathercode,
		IsDay:         om.CurrentWeather.IsDay,
		Time:          om.CurrentWeather.Time,
	}, nil
}
