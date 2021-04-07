package asana

import (
	asana "bitbucket.org/mikehouston/asana-go"
)

func NewClient(pat string) *asana.Client {
	return asana.NewClientWithAccessToken(pat)
}
