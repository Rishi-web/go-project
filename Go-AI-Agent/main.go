package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	OpenWeatherMapURL = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getCurrentTime() string {
	now := time.Now()
	return now.Format("Monday, Jan 2, 2006 15:04:05 MST")
}

func getWeather(city string) (string, error) {
	//apiKey := os.Getenv("OPENWEATHER_API_KEY")
	apiKey := "abcd"
	if apiKey == "" {
		return "", fmt.Errorf("OPENWEATHER_API_KEY not set")
	}

	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", OpenWeatherMapURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if len(data.Weather) == 0 {
		return "", fmt.Errorf("could not get weather description")
	}

	return fmt.Sprintf("Weather in %s: %s, %.1f°C, Humidity: %d%%",
		data.Name, data.Weather[0].Description, data.Main.Temp, data.Main.Humidity), nil
}

// transcribeAudio sends an audio file to OpenAI's Whisper API for transcription - Speach to Text
func transcribeAudio(filename string) (string, error) {
	//apiKey := os.Getenv("OPENAI_API_KEY")
	apiKey := "abcd"
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY not set")
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, _ := writer.CreateFormFile("file", filename)
	io.Copy(fw, file)

	writer.WriteField("model", "whisper-1")
	writer.Close()

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", body)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Text string `json:"text"`
	}
	fmt.Println("Transcribing audio...", resp.Status)
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Text, nil
}

// parseIntent determines the intent of the user's input
// It returns the intent type and any relevant data (like city for weather)
func parseIntent(input string) (string, string) {
	if strings.Contains(strings.ToLower(input), "time") {
		return "time", ""
	} else if strings.Contains(strings.ToLower(input), "weather") {
		// Try to extract city
		words := strings.Split(input, " ")
		for i, w := range words {
			if strings.ToLower(w) == "in" && i+1 < len(words) {
				return "weather", words[i+1]
			}
		}
		return "weather", "London" // default
	}
	return "unknown", ""
}

func speak(text string) {
	// Linux (espeak)
	exec.Command("espeak", text).Run()
}

func main() {
	fmt.Println("🧠 AI Agent is ready!")
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("\nAsk something (time, weather <city>, or 'exit'): ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)

	// 	if input == "time" {
	// 		fmt.Println("🕒 Current time is:", getCurrentTime())
	// 	} else if input == "exit" {
	// 		fmt.Println("👋 Goodbye!")
	// 		return
	// 	} else if strings.HasPrefix(input, "weather ") {
	// 		city := strings.TrimSpace(strings.TrimPrefix(input, "weather "))
	// 		if city == "" {
	// 			fmt.Println("⚠️ Please specify a city. Try: weather Tokyo")
	// 			continue
	// 		}
	// 		weather, err := getWeather(city)
	// 		if err != nil {
	// 			fmt.Println("❌ Error:", err)
	// 		} else {
	// 			fmt.Println("🌦️", weather)
	// 		}
	// 	} else {
	// 		fmt.Println("🤖 I didn't understand. Try 'time', 'weather <city>', or 'exit'.")
	// 	}
	// }
	fmt.Println("🎤 Speak now...")

	// Record audio
	exec.Command("bash", "record.sh").Run()

	// Transcribe it
	spokenText, err := transcribeAudio("input.wav")
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("📝 You said:", spokenText)

	// Parse intent
	intent, city := parseIntent(spokenText)

	var response string
	switch intent {
	case "time":
		response = "The current time is " + getCurrentTime()
	case "weather":
		weather, err := getWeather(city)
		if err != nil {
			response = "Sorry, I couldn't get the weather."
		} else {
			response = weather
		}
	default:
		response = "Sorry, I didn't understand that."
	}

	fmt.Println("🤖", response)
	speak(response)
}
