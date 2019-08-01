package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	dd "github.com/hugozhu/godingtalk"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

var c *dd.DingTalkClient
var calendarId string
var staffId string
var timezone string

func init() {
	c = dd.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	calendarId = os.Getenv("calendar_id")
	staffId = os.Getenv("staff_id")
	timezone = os.Getenv("timezone")
	if timezone == "" {
		timezone = "Asia/Shanghai"
	}
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func eventsFromFile(file string) (map[string]dd.Event, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cache map[string]dd.Event
	err = json.NewDecoder(f).Decode(&cache)
	return cache, err
}

func saveEvents(path string, cache map[string]dd.Event) {
	fmt.Printf("Saving events map to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(cache)
}

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	from := time.Now()
	to := time.Now().AddDate(0, 0, 1)
	log.Println(from.Format("2006-01-02") + " " + to.Format("2006-01-02"))
	events, _ := c.ListEvents(staffId, from, to)
	cache, _ := eventsFromFile("events.json")
	if cache == nil {
		cache = make(map[string]dd.Event)
	}
	for _, event := range events {
		log.Println(event.Summary)
		reminders := &calendar.EventReminders{
			UseDefault: true,
		}
		if _, exist := cache[event.Id]; !exist {
			googleEvent := &calendar.Event{
				Summary:     event.Summary,
				Location:    event.Location,
				Description: event.Description,
				Start: &calendar.EventDateTime{
					DateTime: event.Start.DateTime,
					TimeZone: timezone,
				},
				End: &calendar.EventDateTime{
					DateTime: event.End.DateTime,
					TimeZone: timezone,
				},
				Reminders: reminders,
			}
			cache[event.Id] = event
			// log.Println(srv, googleEvent)
			googleEvent, err = srv.Events.Insert(calendarId, googleEvent).Do()
			if err != nil {
				log.Fatalf("Unable to create event. %v\n", err)
			}
			fmt.Printf("Event created: %s\n", googleEvent.HtmlLink)
		}
	}
	saveEvents("events.json", cache)
}
