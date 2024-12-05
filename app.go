package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetInfo(addresses []string, location string) string {
	return ""
}

func (a *App) Atlas(address string) (AtlasDocuments, error) {
	baseUrl := "https://api.phila.gov/ais/v1/search"
	url := fmt.Sprintf("%s/%s", baseUrl, address)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return AtlasDocuments{}, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return AtlasDocuments{}, err
	}

	var data AtlasDeedResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return AtlasDocuments{}, err
	}

	return GetAtlasDocuments(data)
}

func GetAtlasDocuments(deed AtlasDeedResponse) (AtlasDocuments, error) {
	url := buildQueryWithSQL(deed)
	resp, err := http.Get(url)

	var documents AtlasDocuments

	if err != nil {
		return documents, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return documents, err
	}
	err = json.Unmarshal(b, &documents)
	return documents, err
}

func buildQueryWithSQL(response AtlasDeedResponse) string {
	baseURL := "https://services.arcgis.com/fLeGjb7u4uXqeF9q/ArcGIS/rest/services/RTT_SUMMARY/FeatureServer/0/query"
	var sqlBuilder strings.Builder
	if len(response.Features) > 0 {
		feature := response.Features[0]
		props := feature.Properties

		sqlBuilder.WriteString(fmt.Sprintf("(ADDRESS_LOW >= %d AND ADDRESS_HIGH <= %d)", props.AddressLow, props.AddressLow))

		if props.StreetName != "" {
			sqlBuilder.WriteString(fmt.Sprintf(" AND STREET_NAME = '%s'", props.StreetName))
		}
		if props.StreetSuffix != "" {
			sqlBuilder.WriteString(fmt.Sprintf(" AND STREET_SUFFIX = '%s'", props.StreetSuffix))
		}
		if props.StreetPredir != "" {
			sqlBuilder.WriteString(fmt.Sprintf(" AND STREET_PREDIR = '%s'", props.StreetPredir))
		}
	}
	params := url.Values{}
	params.Set("where", sqlBuilder.String())
	params.Set("outFields", "DOCUMENT_ID, DISPLAY_DATE, DOCUMENT_TYPE, GRANTORS, GRANTEES, UNIT_NUM")
	params.Set("returnDistinctValues", "true")
	params.Set("returnGeometry", "false")
	params.Set("f", "json")
	params.Set("sqlFormat", "standard")
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	return fullURL
}
