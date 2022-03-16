package characters

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

var LionTurtleDB *gorm.DB

type Character struct {
	gorm.Model
	Name                 string
	Playbook             string
	Training             string
	FightingStyle        string
	BackgroundMilitary   bool
	BackgroundMonastic   bool
	BackgroundOutlaw     bool
	BackgroundPrivileged bool
	BackgroundUrban      bool
	BackgroundWilderness bool
	Hometown             string
	// demeanorOne    string
	// demeanorTwo    string
	// demeanorThree  string
	// demeanorFour   string
	// demeanorFive   string
	// appearance     string
	// historyOne     string
	// historyTwo     string
	// historyThree   string
	// historyFour    string
	// historyFive    string
	// connectionsOne string
	// connectionsTwo string

	// STATS
	// creativity int
	// focus      int
	// harmony    int
	// passion    int

	// CONDITIONS
	// fatigue  int
	// afraid   bool
	// angry    bool
	// guilty   bool
	// insecure bool
	// troubled bool

	// BALANCE
	// center                   int
	// balance                  int
	// momentOfBalanceAvailable bool

	// feature???

	// GROWTH
	// growth                      int
	// growthYourPlaybook          int
	// growthAnotherPlaybook       int
	// growthRaiseStat             int
	// growthShiftCenter           int
	// growthUnlockMomentOfBalance int

	// ABILITIES
	// moves      []playbooks.Move
	// techniques []techniques.Technique
}

func (c *Character) Save() {
	if c.ID == 0 {
		result := LionTurtleDB.Create(c)
		if result.Error != nil {
			log.Printf("Error inserting character: %s", result.Error.Error())
		}
	}

	result := LionTurtleDB.Save(c)
	if result.Error != nil {
		log.Printf("Error saving character ID=%d: %s", c.ID, result.Error.Error())
	}
}

func (c *Character) Delete() {
	result := LionTurtleDB.Delete(c)
	if result.Error != nil {
		log.Printf("Error deleting character ID=%d: %s", c.ID, result.Error.Error())
	}
}

func (c *Character) GetFormData(r *http.Request) error {
	c.Name = r.FormValue("Name")
	c.Playbook = r.FormValue("Playbook")
	c.Training = r.FormValue("Training")
	c.FightingStyle = r.FormValue("FightingStyle")
	c.BackgroundMilitary = r.FormValue("BackgroundMilitary") == "1"
	c.BackgroundMonastic = r.FormValue("BackgroundMonastic") == "1"
	c.BackgroundOutlaw = r.FormValue("BackgroundOutlaw") == "1"
	c.BackgroundPrivileged = r.FormValue("BackgroundPrivileged") == "1"
	c.BackgroundUrban = r.FormValue("BackgroundUrban") == "1"
	c.BackgroundWilderness = r.FormValue("BackgroundWilderness") == "1"
	c.Hometown = r.FormValue("Hometown")
	return nil
}

func GetAllCharacters() []Character {
	allCharacters := []Character{}
	result := LionTurtleDB.Find(&allCharacters)
	if result.Error != nil {
		log.Printf("Error retrieving characters: %s", result.Error.Error())
	}
	return allCharacters
}

func GetCharacterByID(id int) (Character, error) {
	log.Printf("Fetching character with ID=%d", id)
	char := Character{}
	result := LionTurtleDB.First(&char, id)
	if result.Error != nil {
		log.Printf("Error retrieving character: %s", result.Error.Error())
		return Character{}, result.Error
	}
	return char, nil
}

func GetCharacterByIDString(id string) (Character, error) {
	log.Printf("Fetching character with ID=%s", id)
	idInt, convErr := strconv.Atoi(id)
	if convErr != nil {
		log.Printf("Could not convert '%s' to int", id)
		return Character{}, convErr
	}
	return GetCharacterByID(idInt)
}

func RandomName() string {
	names := []string{
		"Aditi",
		"Akash",
		"Anil",
		"Batsal",
		"Chaha",
		"Chang",
		"Chimini",
		"Devna",
		"Ehani",
		"Hayate",
		"Idha",
		"Imay",
		"Mukta",
		"Sanani",
		"Soma",
		"Sora",
		"Tau",
		"Toofan",
		"Unnat",
		"Yawen",
		"Binh",
		"Bowen",
		"Caihong",
		"Chia-Hao",
		"Dae",
		"Diu",
		"Hanna",
		"Heng",
		"Kim",
		"Kyung",
		"Minh",
		"Nuan",
		"Qiang",
		"Quiyue",
		"Shufen",
		"Thi",
		"Woong",
		"Xiaobo",
		"Ya-Ting",
		"Zixin",
		"Achak",
		"Aklaq",
		"Aputi",
		"Atka",
		"Hanta",
		"Kallik",
		"Kanti",
		"Kitchi",
		"Makwa",
		"Meeka",
		"Miki",
		"Niimi",
		"Noodin",
		"Siqniq",
		"Tapisa",
		"Thaki",
		"Ukiuk",
		"Waaseyaa",
		"Yuka",
		"Ziibi",
		"Asayo",
		"Ayami",
		"Bashira",
		"Davaa",
		"Erdene",
		"Ganzaya",
		"Hanako",
		"Jaw Long",
		"Kayo",
		"Keisuke",
		"Kenshin",
		"Manami",
		"Mayu",
		"Qacha",
		"Qudan",
		"Satsuki",
		"Saya",
		"Tuguslar",
		"Yuka",
		"Zolzaya",
		"Bo",
		"He",
		"Jia",
		"Ju",
		"Mu",
		"Shi",
		"Yan",
		"Zan",
		"Bai",
		"Hua",
	}
	random := rand.New(rand.NewSource(time.Now().Unix()))
	return names[random.Intn(len(names))]
}
