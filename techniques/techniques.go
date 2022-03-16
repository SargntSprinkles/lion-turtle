package techniques

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SargntSprinkles/lion-turtle/util"
)

var Techniques map[string]Technique = LoadTechniques("techniques/")

type Technique struct {
	Name        string
	Approach    string
	Description string
	Playbook    string
}

func LoadTechniques(dir string) map[string]Technique {
	techniques := map[string]Technique{}
	filename := dir + "techniques.json"
	techniquesFile, fileErr := os.ReadFile(filename)
	util.CheckError(fileErr, fmt.Sprintf("Error reading file %s", filename))
	jsonErr := json.Unmarshal(techniquesFile, &techniques)
	util.CheckError(jsonErr, fmt.Sprintf("Error unmarshalling json %v", techniquesFile))
	return techniques
}
