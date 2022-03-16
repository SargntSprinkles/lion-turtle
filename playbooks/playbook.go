package playbooks

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SargntSprinkles/lion-turtle/techniques"
	"github.com/SargntSprinkles/lion-turtle/util"
)

var Playbooks map[string]Playbook = LoadPlaybooks("playbooks/")

var PlaybookNames = []string{
	"The Adamant",
	"The Bold",
}

type Playbook struct {
	Playbook           string
	Source             string
	Principles         []string
	StartingCreativity int
	StartingFocus      int
	StartingHarmony    int
	StartingPassion    int
	DemeanorOptions    []string
	HistoryQuestions   []string
	Connections        []string
	MomentOfBalance    string
	GrowthQuestion     string
	Feature            Feature
	Moves              []Move
	Technique          techniques.Technique
}

type Feature struct {
	Name        string
	Description string
}

type Move struct {
	Name        string
	Description string
}

func LoadPlaybooks(dir string) map[string]Playbook {
	playbooks := map[string]Playbook{}
	for _, pb := range PlaybookNames {
		tempPlaybook := Playbook{}
		filename := dir + pb + ".json"
		playbookFile, fileErr := os.ReadFile(filename)
		util.CheckError(fileErr, fmt.Sprintf("Error reading file %s", filename))
		jsonErr := json.Unmarshal(playbookFile, &tempPlaybook)
		util.CheckError(jsonErr, fmt.Sprintf("Error unmarshalling json %v", playbookFile))
		playbooks[pb] = tempPlaybook
	}
	return playbooks
}
