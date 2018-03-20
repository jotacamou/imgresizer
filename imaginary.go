package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ImaginaryOps(profileName string, file string) (*http.Response, error) {
	var baseUrl, queryParam string

	profile := config.Profiles[profileName]
	if profile["ops"] == "" {
		err := errors.New("No profile found.")
		return nil, err
	}

	baseUrl = "http://" + config.ImaginHost + ":" +
		strconv.Itoa(config.ImaginPort) + "/" + profile["ops"]

	for k, v := range profile {
		if k != "ops" {
			if queryParam == "" {
				queryParam = k + "=" + v
			} else {
				queryParam = queryParam + "&" + k + "=" + v
			}
		}
	}

	if config.EnableUrlSource == true && strings.Contains(file, "http") {
		queryParam = queryParam + "&url=" + file
	} else {
		queryParam = queryParam + "&file=" + file
	}

	resp, err := http.Get(baseUrl + "?" + strings.ToLower(queryParam))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
