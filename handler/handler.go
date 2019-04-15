package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/4ndr3ye/DemoServer/model"
	"github.com/4ndr3ye/DemoServer/security"
	"github.com/4ndr3ye/DemoServer/static"
)

func ErrorHandler(e error) {
	if e != nil {
		log.Println(e)
	}
}

func Base(w http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "This is Demo server"}
	json.NewEncoder(w).Encode(message)
}

func Secret(w http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Welcome"}
	json.NewEncoder(w).Encode(message)
}

func GetLogin(w http.ResponseWriter, req *http.Request) {
	p, _ := static.LoadPage("login")
	fmt.Fprint(w, string(p.Body))
	return
}

func Login(w http.ResponseWriter, req *http.Request) {
	var t model.Credential

	body, err := ioutil.ReadAll(req.Body)
	ErrorHandler(err)

	err = json.Unmarshal(body, &t)
	ErrorHandler(err)

	h := sha1.New()
	h.Write([]byte(t.Password))
	sha1Password := hex.EncodeToString(h.Sum(nil))

	result, err := model.CheckUser(t.Username, sha1Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if result.ID != "" {
		security.AddCookie(w, "DemoCookie", security.GenerateToken(result.Email))
		http.Redirect(w, req, "/", 302)
		return
	}

	http.Error(w, "", http.StatusUnauthorized)
	message := map[string]string{"message": "You are hacker"}
	json.NewEncoder(w).Encode(message)
}

func SubmitCustumer(w http.ResponseWriter, req *http.Request) {
	var customer model.Customer
	var err error
	req.ParseForm()
	customer.Firstname = req.Form.Get("firstname")
	customer.Lastname = req.Form.Get("lastname")
	customer.Email = req.Form.Get("email")
	http.Redirect(w, req, "addcustumer", 301)
	err = model.AddCustomer(customer)
	ErrorHandler(err)
}

func GetFilms(w http.ResponseWriter, req *http.Request) {
	films, err := model.ListFilms(req.FormValue("query"))
	ErrorHandler(err)
	p, _ := static.LoadPage("top")
	fmt.Fprint(w, string(p.Body))
	p, _ = static.LoadPage("films")
	fmt.Fprint(w, string(p.Body))
	if req.FormValue("query") != "" {
		fmt.Fprintf(w, "You serach for: "+req.FormValue("query")+"<br>")
	}
	for _, film := range films {
		fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%s</td><td>%d</td><td>%.2f</td><td>%d</td></tr>",
			film.ID, film.Title, film.Description, film.Year, film.Rate, film.Length)
	}
	p, _ = static.LoadPage("bottom")
	fmt.Fprint(w, string(p.Body))
}

func GetAddCustumers(w http.ResponseWriter, req *http.Request) {
	p, _ := static.LoadPage("top")
	fmt.Fprint(w, string(p.Body))
	p, _ = static.LoadPage("addcustumer")
	fmt.Fprint(w, string(p.Body))
	p, _ = static.LoadPage("bottom")
	fmt.Fprint(w, string(p.Body))
}

func GetCustumers(w http.ResponseWriter, req *http.Request) {
	custumers, err := model.ListCustumers(req.FormValue("query"))
	ErrorHandler(err)
	p, _ := static.LoadPage("top")
	fmt.Fprint(w, string(p.Body))
	p, _ = static.LoadPage("custumers")
	fmt.Fprint(w, string(p.Body))
	if req.FormValue("query") != "" {
		fmt.Fprintf(w, "You serach for: "+req.FormValue("query")+"<br>")
	}
	for _, custumer := range custumers {
		fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			custumer.ID, custumer.Firstname, custumer.Lastname, custumer.Email)
	}
	p, _ = static.LoadPage("bottom")
	fmt.Fprint(w, string(p.Body))
}

func Welcome(w http.ResponseWriter, req *http.Request) {
	p, _ := static.LoadPage("top")
	fmt.Fprint(w, string(p.Body))
	fmt.Fprintf(w, "<h1>Welcome</h1>")
	p, _ = static.LoadPage("bottom")
	fmt.Fprint(w, string(p.Body))
}

func GetCheckAlive(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	p, _ := static.LoadPage("top")
	fmt.Fprint(w, string(p.Body))
	if add, ok := req.Form["add"]; !ok {
		p, _ = static.LoadPage("alive")
		fmt.Fprint(w, string(p.Body))
	} else {
		output, err := exec.Command("bash", "-c", "ping -c 4 "+add[0]+" &> /dev/null && echo Success || echo Fail").CombinedOutput()
		ErrorHandler(err)
		fmt.Fprintf(w, "<h1>%s</h1>", output)
	}

	p, _ = static.LoadPage("bottom")
	fmt.Fprint(w, string(p.Body))
}
