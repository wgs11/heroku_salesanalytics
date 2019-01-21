package main

import (
	"fmt"
	"net/http"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type ManagerForm struct {
	ID int `json:"employee_id", db:"employee_id"`
	First string `json:"fname", db:"fname"`
	Last string `json:"lname", db:"lname"`
}

type NewUser struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
	First	string `json:"fname", db:"fname"`
	Last 	string `json:"fname", db:"lname"`
	Home	string `json:"location_name", db:"location_name"`
	Position string `json:"position", db:"position"`
}

type QuestionSet struct {
	Questions []string
}

func SeedDB(w http.ResponseWriter, r *http.Request) {
	questions := []string{
		"Is the parking lot clean?",
		"Is the patio & furniture clean?",
		"Are the outdoor lights working?",
		"Are weeds pulled & landscaping in good condition?",
		"Is the area around the dumpster clean?",
		"Are the sidewalk clean?",
		"Are the exteriors of the windows and window ledges clean?",
		"Is the team wearing clean uniforms?",
		"Are their uniforms in good condition?",
		"Are aprons clean?",
		"Do they have on the proper slip-resistant shoes?",
		"Is hair properly restrained?",
		"Is the store awareness binder being completed daily?",
		"Is the scoop weighing system in place and being used daily?",
		"Are the tables clean (tops/sides/bases/legs)?",
		"Are the chairs clean (back/cushions/legs)?",
		"Are the undersides of the tables and chairs gum free?",
		"Are the floors and baseboards clean?",
		"Are the walls clean?",
		"Are the trash recepticals clean? ",
		"Are the interiors of the windows and window ledges clean?",
		"Are the fans and light fixtures clean?",
		"Are the fans and light fixtures clean?",
		"Are the highchairs clean and in good repair ? ",
		"Are all lights in working condition ?",
		"Are the doors and entryways clean?",
		"Are the napkin dispensers clean and stocked?",
		"Is the water fountain or water dispenser clean?",
		"Is the front counter clean and free of trash/spills?",
		"Is the area around registers clean?",
		"Are the tips jars clean and creative?",
		"Is the store decorated appropriately for season?",
		"Are there displays for minor holidays?",
		"Are the correct promotional signs up in store ?",
		"Are the displays organized?",
		"Are the Seasonal displays well merchandised?",
		"Are the displays on counters well maintained/stocked?",
		"Are the front count displays stocked ?",
		"Is the area around menu board attractive?",
		"Do the menu boards have the current specials up?",
		"Are the everyday candy tables stocked?",
		"Do the displays look appealing ?",
		"Are all cases / displays stocked ?",
		"Are the case decorated and clean ? ",
		"Is the pint case stocked and faced properly?",
		"Is the area around pint case clean and organized ?",
		"Are all ice cream display freezers stocked/proper rotation?",
		"Are the dipping cabinets clean and defrosted ?",
		"Are the countertops clean and organized?",
		"Are the cabinets & shelves clean and organized?",
		"Is the soda fountain/drain/ice holder clean?",
		"Are the hand sinks and faucets clean?",
		"Are the paper towel and soap dispensers clean?",
		"Are the syrup and topping holders/pumps clean?",
		"Are the spoon holders clean?",
		"Is the microwave clean?",
		"Is the cup/lid area clean/organized/stocked?",
		"Is the sundae bar clean (inside/outside)?",
		"Is the icemaker clean?",
		"Are the menu boards clean? ",
		"Are all lights in working condition ?",
		"Is the ice cream properly rotated?",
		"Is the ice cream scraped well?",
		"Is the candy properly rotated?",
		"Is the candy all Julian dated?",
		"Are the dates good on dairy products ?0",
		"Proper amount of cones for resale displayed ? ",
		"Outdated items have been pulled from shelves ?",
		"Proper amount of bananas displayed ?",
		"Proper amount of resale items in store ? ",
		"Is the order level appropriate for the time of year?",
		"Are seasonal products stocked well?",
		"Correct amount of paper products in store ?",
		"Correct amount of chemicals in the store ?",
		"Is there a limited amount of store bought chemicals ?",
		"Are all orders put away and organized?",
		"Are the floors clean?",
		"Are the prep tables clean (top/legs/shelves)?",
		"Are the food products in the backroom properly rotated?",
		"Are the information boards clean, organized, and up to date?",
		"Is the mixer clean?",
		"Is the office/desk area clean and organized?",
		"Is the laptop clean ?",
		"Is the cart clean ?",
		"Is the 3 compartment sink clean?",
		"Is the walk-in / holding freezers clean and organized ?",
		"Are the trashcans clean ? ",
		"Is the mop sink & mop bucket clean ?",
		"Are all lights in working condition ?",
		"Are the floors clean?",
		"Are the walls clean?",
		"Are the light fixtures clean?",
		"Are all lights in working condition ?",
		"Are the vents/fans clean and dust free?",
		"Is the toilet clean (inside/under lid/base/handle)?",
		"Is the urinal clean (inside/underneath/handle)?",
		"Is the sink clean (inside/underside/pipes/handles)?",
		"Are the paper towel dispensers stocked and clean?",
		"Are the soap dispensers stocked and clean?",
		"Are the air fresheners working?",
		"Is the changing table clean?"}
	for _,question := range questions {
		store.DBSeed(question)
	}
}

func IsSignedIn(w http.ResponseWriter, r *http.Request) bool {
	session,_ := cache.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	} else {return true}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	creds := &Credentials{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	if store.CheckUser(creds) != nil {
		fmt.Println("there was no issue signing in")
		session,_ := cache.Get(r, "cookie-name")
		session.Values["authenticated"] = true
		session.Values["user"] = creds.Username
		if (creds.Username == "Sheppy"){
			session.Values["admin"] = true
		} else {
			session.Values["admin"] = false
		}
		session.Save(r,w)
	} else{
	}
	http.Redirect(w,r,"/", 302)

}


func Signup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	creds := &NewUser{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	creds.First = r.FormValue("fname")
	creds.Last = r.FormValue("lname")
	creds.Position = r.FormValue("position")
	creds.Home = r.FormValue("store")
	fmt.Println(creds.Username, creds.Password, creds.First,creds.Last,creds.Position,creds.Home)
	store.CreateUser(creds)
	http.Redirect(w, r, "/", http.StatusFound)
}