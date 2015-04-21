package sessions

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/astaxie/beego/session"
)

type pageContent struct {
	Session  string
	Username string
	Picture  string
}

var (
	globalSessions *session.Manager //global memory space for the session manager
	rootTmpl       = template.Must(template.ParseFiles("templates/outer.html", "templates/inner.html"))
)

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid",
														"enableSetCookie,omitempty": true,
														"gclifetime":180,
														"maxLifetime": 180,
														"secure": false,
														"sessionIDHashFunc": "sha1",
														"sessionIDHashKey": "",
														"cookieLifeTime": 180,
														"providerConfig": ""}`) //since we are using memory, we don't need a provider config
	go globalSessions.GC()

	//Serve static files/content. ex. css and favicon
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("public/static"))))

	//Serve the pages
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/createsession", createSessionHandler)
	http.HandleFunc("/otherpage", otherPageHandler)
	http.HandleFunc("/destroysession", destroySessionHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//open the session
	sess, _ := globalSessions.SessionStart(w, r)

	//check if the session is already created
	var pc pageContent
	sessID := sess.Get("Session")
	if sessID != nil {
		pc.Session = sess.Get("Session").(string)
		pc.Username = sess.Get("Username").(string)
		pc.Picture = sess.Get("Picture").(string)
	}

	rootTmpl.Execute(w, pc)
}

func otherPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Okay, you've had your fun, now move along!</h1>")
	fmt.Fprint(w, "<a href='/'>Go back</a>")
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {

	sess, _ := globalSessions.SessionStart(w, r)

	sess.Set("Session", sess.SessionID())
	sess.Set("Picture", r.FormValue("picture"))
	sess.Set("Username", r.FormValue("username"))

	http.Redirect(w, r, "/", 302)
}

func destroySessionHandler(w http.ResponseWriter, r *http.Request) {

	//get session info
	sess, _ := globalSessions.SessionStart(w, r)

	//destroy the session if it actually exists
	sessID := sess.Get("Session")
	if sessID != nil {
		globalSessions.SessionDestroy(w, r)
	}

	http.Redirect(w, r, "/", 302)
}
