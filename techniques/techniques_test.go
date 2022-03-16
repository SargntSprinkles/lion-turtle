package techniques

import (
	"testing"
)

func TestLoadTechniques(t *testing.T) {
	techniques := LoadTechniques("")
	for _, techName := range []string{"Pinpoint Aim", "Tag Team"} {
		tech, found := techniques[techName]
		if !found {
			t.Errorf("Technique '%s' not found", techName)
		}
		if tech.Name != techName {
			t.Errorf("Expected '%s', got '%s'", techName, tech.Name)
		}
	}
}
