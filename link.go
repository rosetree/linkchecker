package main

import (
	"net/http"
	"regexp"
)

// The link struct holds the Target of the link, which is the url that
// should be checked. It keeps the http Status and a list of all
// Documents, that link to Target.
type link struct {
	Target    string
	Status    string
	Documents []string
}

// FetchStatus requests the HEAD of the current link, if l.Status is
// empty. It will not fetch the link again!
func (l *link) FetchStatus() error {
	if l.Status != "" {
		return nil
	}

	resp, err := http.Head(l.Target)
	if err != nil {
		return err
	}

	l.Status = resp.Status

	return nil
}

// StatusSuccess returns a bool indicating if the status of the link
// is one of the sucess status codes (beginning with 2).
func (l link) StatusSuccess() bool {
	is2xx, err := regexp.MatchString("^2.*", l.Status)
	if err != nil {
		return false
	}
	return is2xx
}
