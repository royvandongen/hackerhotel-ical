package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/alexflint/go-arg"
	ics "github.com/arran4/golang-ical"
	"github.com/gorilla/mux"
)

type config struct {
	ScheduleURL   string `arg:"env:SCHEDULE_URL, -s, --schedule" help:"URL to the ical schedule" placeholder:"<url>" default:"https://pretalx.hackerhotel.nl/2025/schedule/export/schedule.ics"`
	Token         string `arg:"required,env:TOKEN, -t, --token" help:"Authentication Token" placeholder:"<token>"`
	ListenAddress string `arg:"env:LISTEN_ADDRESS, -l, --listen" help:"Port to listen on" placeholder:"<bind address>" default:"0.0.0.0:5000"`
}

var Config config

// downloadICal downloads an iCal file from a given URL
func downloadICal(url string) ([]byte, error) {
	var client = http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token "+Config.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// filterICalByLocation filters events in an iCal data by location and creates a new calendar with those events
func filterICalByLocation(icalData []byte, locationFilter string) (*ics.Calendar, error) {
	cal, err := ics.ParseCalendar(strings.NewReader(string(icalData)))
	if err != nil {
		return nil, err
	}

	newCal := ics.NewCalendar()
	for _, event := range cal.Events() {
		if strings.Contains(event.GetProperty(ics.ComponentPropertyLocation).Value, locationFilter) {
			newCal.Components = append(newCal.Components, event)
		}
	}
	return newCal, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
	fmt.Println("Endpoint Hit: /")
}

func returnSingleLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location, err := url.QueryUnescape(vars["id"])
	if err != nil {
		http.Error(w, "Failed to decode location id", http.StatusInternalServerError)
		return
	}

	fmt.Printf("returnSingleLocation: %s\n", location)

	icalData, err := downloadICal(Config.ScheduleURL)
	if err != nil {
		http.Error(w, "Failed to download iCal file", http.StatusInternalServerError)
		return
	}

	newCal, err := filterICalByLocation(icalData, location)
	if err != nil {
		http.Error(w, "Failed to filter events", http.StatusInternalServerError)
		return
	}

	newICalContent := newCal.Serialize()
	// Set the correct header to serve an iCal file
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=schedule.ics")
	_, err = w.Write([]byte(newICalContent))
	if err != nil {
		http.Error(w, "Failed to write iCal file", http.StatusInternalServerError)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/location/{id}", returnSingleLocation)
	log.Fatal(http.ListenAndServe(Config.ListenAddress, myRouter))
}

func main() {
	arg.MustParse(&Config)
	fmt.Printf("Listening on %s, Using schedule: %s\n", Config.ListenAddress, Config.ScheduleURL)
	fmt.Println()
	fmt.Println("Endpoints available:")
	fmt.Printf("http://%s/\n", Config.ListenAddress)
	fmt.Printf("http://%s/location/<name>\n", Config.ListenAddress)
	fmt.Println("Ready and waiting for requests...")
	handleRequests()
}
