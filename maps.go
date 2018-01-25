package main

import "fmt"
import "net/http"
import "strings"
import "github.com/kelvins/geocoder/structs"
import "errors"
import "strconv"
import "encoding/json"

const (
	geocodeApiUrl = "https://maps.googleapis.com/maps/api/geocode/json?"
)

var ApiKey string

// Address structure used in the Geocoding and GeocodingReverse functions
// Note: The FormattedAddress field should be used only for the GeocodingReverse
// to get the formatted address from the Google Geocoding API. It is not used in
// the Geocoding function.
type Address struct {
	Street           string
	Number           int
	Neighborhood     string
	District         string
	City             string
	County           string
	State            string
	Country          string
	PostalCode       string
	FormattedAddress string
	Types            string
}

// Location structure used in the Geocoding and GeocodingReverse functions
type Location struct {
	Latitude  float64
	Longitude float64
}

func callGoogleMaps(address string) Locationterminal {
  ApiKey = "AIzaSyCO61_FHmrUaK1eGlcLwR8VKFCEkaEkelE"
  var result Locationterminal
  address = strings.Replace(address, " ", "+", -1)

  // Create the URL based on the formated address
  url := geocodeApiUrl + "address=" + address

  // Use the API Key if it was set
  if ApiKey != "" {
    url += "&key=" + ApiKey
  }
  // Send the HTTP request and get the results
  resultsReq, err := httpRequest(url)
  if err != nil {
    fmt.Println("Error Lho 2 " + err.Error())
  }

  if strings.ToUpper(resultsReq.Status) != "OK" {
    fmt.Println("Address :"+address+" Not Found")
    result.statusInformation = "NOT FOUND"
  } else {
    latitude := resultsReq.Results[0].Geometry.Location.Lat
    longitude := resultsReq.Results[0].Geometry.Location.Lng

    result.latitudeTerminal = strconv.FormatFloat(latitude, 'f', 6, 64)
    result.longitudeTerminal = strconv.FormatFloat(longitude, 'f', 6, 64)
    result.statusInformation = "FOUND"
  }

  return result
}

// httpRequest function send the HTTP request, decode the JSON
// and return a Results structure
func httpRequest(url string) (structs.Results, error) {

	var results structs.Results

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}

	// For control over HTTP client headers, redirect policy, and other settings, create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}

	// Callers should close resp.Body when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	// The "OK" status indicates that no error has occurred, it means
	// the address was analyzed and at least one geographic code was returned
	if strings.ToUpper(results.Status) != "OK" {
		// If the status is not "OK" check what status was returned
		switch strings.ToUpper(results.Status) {
		case "ZERO_RESULTS":
			err = errors.New("No results found.")
			break
		case "OVER_QUERY_LIMIT":
			err = errors.New("You are over your quota.")
			break
		case "REQUEST_DENIED":
			err = errors.New("Your request was denied.")
			break
		case "INVALID_REQUEST":
			err = errors.New("Probably the query is missing.")
			break
		case "UNKNOWN_ERROR":
			err = errors.New("Server error. Please, try again.")
			break
		default:
			break
		}
	}

	return results, err
}
