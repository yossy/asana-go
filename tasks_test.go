package asana

import "testing"

func TestPickUpTaskID(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: "[Link Asana Task](https://app.asana.com/0/1196738113624713/1200167372563640/f)", want: "1200167372563640"},
		{body: "[Link Asana Task](https://app.asana.com/0/0/1200167372563640)", want: "1200167372563640"},
		{body: "[Link Asana Task]()", want: ""},
		{body: "body", want: ""},
		{body: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.body, func(t *testing.T) {
			got, _ := PickUpTaskID(tt.body)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
