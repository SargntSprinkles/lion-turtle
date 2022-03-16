package playbooks

import (
	"testing"
)

func TestLoadPlaybooks(t *testing.T) {
	playbooks := LoadPlaybooks("")
	for _, pbn := range PlaybookNames {
		pb, found := playbooks[pbn]
		if !found {
			t.Errorf("Playbook '%s' not found", pbn)
		}
		if pb.Playbook != pbn {
			t.Errorf("Expected '%s', got '%s'", pbn, pb.Playbook)
		}
	}
}
