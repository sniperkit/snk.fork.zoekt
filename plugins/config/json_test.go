/*
Sniperkit-Bot
- Date: 2018-08-12 12:11:26.372554071 +0200 CEST m=+0.045728207
- Status: analyzed
*/

package mirror

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type configEntry struct {
	GithubUser string
	GitilesURL string
	CGitURL    string
	Name       string
	Exclude    string
}

func TestJSON(t *testing.T) {
	j, err := ioutil.ReadFile("../../shared/config/examples/mirror.json")
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	result := []configEntry{}
	if err := json.Unmarshal(j, &result); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
}
