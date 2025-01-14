package countries

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	domaincounties "github.com/firdasafridi/parkinglot/internal/entity/countries"
	"github.com/pkg/errors"
)

const (
	url     = "https://restcountries.com"
	version = "v3.1"
	path    = "name"
)

var (
	getCountriesURL = fmt.Sprintf("%s/%s/%s", url, version, path)
)

func RequestCountry(ctx context.Context, country string) (detailCountries []domaincounties.ResponseCountry, err error) {
	url := fmt.Sprintf("%s/%s", getCountriesURL, country)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "RequestCountry.NewRequestWithContext")
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "RequestCountry.ioutil.ReadAll")
	}

	detailCountries = make([]domaincounties.ResponseCountry, 0)
	if err = json.Unmarshal(body, &detailCountries); err != nil {
		return nil, errors.Wrap(err, "RequestCountry.Unmarshal")
	}

	return detailCountries, nil
}
