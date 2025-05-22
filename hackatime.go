package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type HackatimeStats struct {
	Data struct {
		IsCodingActivityVisible bool      `json:"is_coding_activity_visible"`
		IsOtherUsageVisible    bool      `json:"is_other_usage_visible"`
		Status                 string    `json:"status"`
		Start                  time.Time `json:"start"`
		End                    time.Time `json:"end"`
		Range                  string    `json:"range"`
		HumanReadableRange     string    `json:"human_readable_range"`
		TotalSeconds           int       `json:"total_seconds"`
		DailyAverage          float64   `json:"daily_average"`
		HumanReadableTotal     string    `json:"human_readable_total"`
		HumanReadableDailyAvg  string    `json:"human_readable_daily_average"`
		Languages              []struct {
			Name          string  `json:"name"`
			TotalSeconds int     `json:"total_seconds"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Percent      float64 `json:"percent"`
			Digital      string  `json:"digital"`
		} `json:"languages"`
	} `json:"data"`
}

func hackatimeStats() HackatimeStats {
	username := os.Getenv("HACKATIME_USERNAME")
	if username == "" {
		fmt.Println("HACKATIME_USERNAME environment variable not set")
		return HackatimeStats{}
	}

	stats, err := getHackatimeStats(username)
	if err != nil {
		fmt.Printf("Error getting Hackatime stats: %v\n", err)
		return HackatimeStats{}
	}

	return *stats
}

func getHackatimeStats(username string) (*HackatimeStats, error) {
	apiKey := os.Getenv("HACKATIME_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("HACKATIME_API_KEY environment variable not set")
	}

	url := fmt.Sprintf("https://hackatime.hackclub.com/api/v1/users/%s/stats", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stats HackatimeStats
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func formatProgressBar(percent float64, width int) string {
	filled := int((percent / 100) * float64(width))
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func wakatimeSingleCategoryBar(title string, data []interface{}, limit int) string {
	result := fmt.Sprintf("\n%s\n", title)

	// Format category
	for i := 0; i < limit && i < len(data); i++ {
		item := data[i].(map[string]interface{})
		name := item["name"].(string)
		text := item["text"].(string)
		percent := item["percent"].(float64)
		result += fmt.Sprintf("%-15s %-12s %s %6.2f%%\n",
			name, text, formatProgressBar(percent, 25), percent)
	}

	return result
}

func wakatimeDoubleCategoryBar(title1 string, data1 []interface{}, title2 string, data2 []interface{}, limit int) string {
	result := fmt.Sprintf("\n%s\n", title1)

	// Format first category
	for i := 0; i < limit && i < len(data1); i++ {
		item := data1[i].(map[string]interface{})
		name := item["name"].(string)
		text := item["text"].(string)
		percent := item["percent"].(float64)
		result += fmt.Sprintf("%-15s %-12s %s %6.2f%%\n",
			name, text, formatProgressBar(percent, 25), percent)
	}

	result += fmt.Sprintf("\n%s\n", title2)

	// Format second category
	for i := 0; i < limit && i < len(data2); i++ {
		item := data2[i].(map[string]interface{})
		name := item["name"].(string)
		text := item["text"].(string)
		percent := item["percent"].(float64)
		result += fmt.Sprintf("%-15s %-12s %s %6.2f%%\n",
			name, text, formatProgressBar(percent, 25), percent)
	}

	return result
}
func wakatimeLanguages(title string, data []struct {
	Name          string  `json:"name"`
	TotalSeconds int     `json:"total_seconds"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
}, limit int, totalTime string) string {
	result := fmt.Sprintf("\n%s\n", title)

	// Format languages
	for i := 0; i < limit && i < len(data); i++ {
		item := data[i]
		name := item.Name
		text := item.Text
		percent := item.Percent
		result += fmt.Sprintf("%-15s %-12s %s %6.2f%%\n",
			name, text, formatProgressBar(percent, 25), percent)
	}

	result += fmt.Sprintf("\nTotal: %s\n", totalTime)
	return result
}
