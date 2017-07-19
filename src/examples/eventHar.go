package main

// convert log entries captured into logJSON.json to a HAR file
import (
	"encoding/json"
	"events"
	"fmt"
	"httpArchive"
	"io/ioutil"
	"log"
	"os"

	"github.com/fedesog/webdriver"
)

type PerfTimings struct {
	NavigationStart            int64 `json:"navigationStart"`
	UnloadEventStart           int64 `json:"unloadEventStart"`
	UnloadEventEnd             int64 `json:"unloadEventEnd"`
	RedirectStart              int64 `json:"redirectStart"`
	RedirectEnd                int64 `json:"redirectEnd"`
	FetchStart                 int64 `json:"fetchStart"`
	DomainLookupStart          int64 `json:"domainLookupStart"`
	DomainLookupEnd            int64 `json:"domainLookupEnd"`
	ConnectStart               int64 `json:"connectStart"`
	ConnectEnd                 int64 `json:"connectEnd"`
	SecureConnectionStart      int64 `json:"secureConnectionStart"`
	RequestStart               int64 `json:"requestStart"`
	ResponseStart              int64 `json:"responseStart"`
	ResponseEnd                int64 `json:"responseEnd"`
	DomLoading                 int64 `json:"domLoading"`
	DomInteractive             int64 `json:"domInteractive"`
	DomContentLoadedEventStart int64 `json:"domContentLoadedEventStart"`
	DomContentLoadedEventEnd   int64 `json:"domContentLoadedEventEnd"`
	DomComplete                int64 `json:"domComplete"`
	LoadEventStart             int64 `json:"loadEventStart"`
	LoadEventEnd               int64 `json:"loadEventEnd"`
}

type LoggedJSON struct {
	Timings  PerfTimings          `json:"timings"`
	PerfLogs []webdriver.LogEntry `json:"perfLogs"`
}

func main() {

	file, e := ioutil.ReadFile("./logJSON.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var logJSON LoggedJSON
	json.Unmarshal(file, &logJSON)

	convertToHar(logJSON)
}

func convertToHar(logJSON LoggedJSON) {

	e, err := events.NewFromLogEntries(logJSON.PerfLogs)
	if err != nil {
		log.Fatal(err)
	}

	har, err := httpArchive.CreateHARFromEvents(e)
	if err != nil {
		log.Fatal(err)
	}
	harJSON, err := json.Marshal(har)

	// write out the HAR file and the RAW chromeEvens to file
	ioutil.WriteFile("./chromdriver.har", harJSON, 0644)
}
