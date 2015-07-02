package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"sort"
	"sync"
	"time"

	"github.com/alexcesaro/quotedprintable"
	r "github.com/dancannon/gorethink"
	"github.com/dchest/uniuri"
	"github.com/eaigner/dkim"
	"github.com/lavab/api/models"
	"github.com/namsral/flag"
)

var (
	rethinkAddress  = flag.String("rethinkdb_address", "172.16.0.1:28015", "Address of the RethinkDB server")
	rethinkDatabase = flag.String("rethinkdb_database", "prod", "RethinkDB database to use")

	apiURL       = flag.String("api_url", "https://api.lavaboom.com", "URL of the API to use")
	smtpdAddress = flag.String("smtpd_address", "172.16.0.1:2525", "Address of the forwarding email server")
	dkimKey      = flag.String("dkim_key", "./dkim.key", "Path of the DKIM key")
)

func main() {
	flag.Parse()

	key, err := ioutil.ReadFile(*dkimKey)
	if err != nil {
		log.Fatal(err)
	}

	dc, err := dkim.NewConf("lavaboom.com", "mailer")
	if err != nil {
		log.Fatal(err)
	}

	dk, err := dkim.New(dc, key)
	if err != nil {
		log.Fatal(err)
	}

	session, err := r.Connect(r.ConnectOpts{
		Address:  *rethinkAddress,
		Database: *rethinkDatabase,
	})
	if err != nil {
		log.Fatal(err)
	}

	r.DB(*rethinkDatabase).TableCreate("onboarding").Exec(session)
	r.Table("onboarding").IndexCreate("time").Exec(session)

	var stateLock sync.Mutex

	// Load the hub state from RethinkDB
	cursor, err := r.Table("onboarding").OrderBy(r.OrderByOpts{
		Index: "time",
	}).Run(session)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close()
	var state State
	if err := cursor.All(&state); err != nil {
		log.Fatal(err)
	}

	sort.Sort(state)
	log.Printf("%+v", state)

	change := make(chan struct{})

	go func() {
		for {
			log.Print("hub loop")

			stateLock.Lock()
			timersToDelete := []int{}
			for id, timer := range state {
				if timer.Time.Before(time.Now()) {
					email := &bytes.Buffer{}
					if err := emtpl.Execute(email, map[string]interface{}{
						"from":       timer.From,
						"to":         timer.To,
						"subject":    timer.Subject,
						"body":       quotedprintable.EncodeToString([]byte(timer.Body)),
						"message_id": "onboarding-" + uniuri.NewLen(uniuri.UUIDLen) + "@lavaboom.com",
						"date":       time.Now().Format(time.RubyDate),
					}); err != nil {
						log.Print(err)
						continue
					}

					body := bytes.Replace(email.Bytes(), []byte("\n"), []byte("\r\n"), -1)
					sbody, err := dk.Sign(body)
					if err != nil {
						log.Print(err)
						continue
					}

					if err := smtp.SendMail(*smtpdAddress, nil, timer.From, timer.To, sbody); err != nil {
						log.Print(err)
						continue
					}

					// Delete it from RDB and the state
					r.Table("onboarding").Get(timer.ID).Delete().Exec(session)
					timersToDelete = append(timersToDelete, id)
				} else {
					break
				}
			}
			for y, x := range timersToDelete {
				i := x - y
				copy(state[i:], state[i+1:])
				state[len(state)-1] = nil
				state = state[:len(state)-1]
			}
			stateLock.Unlock()

			if len(state) > 0 {
				select {
				case <-time.After(state[0].Time.Sub(time.Now())):
					break
				case <-change:
					break
				}
			} else {
				<-change
			}
		}
	}()

	http.HandleFunc("/onboarding", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}

		var event struct {
			Account string `json:"account"`
		}
		if err := json.Unmarshal(body, &event); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}

		cursor, err := r.Table("accounts").Get(event.Account).Run(session)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}
		defer cursor.Close()
		var account *models.Account
		if err := cursor.One(&account); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}

		x1, ok := account.Settings.(map[string]interface{})
		if !ok {
			http.Error(w, "Account misconfigured #1", 500)
			return
		}

		x2, ok := x1["firstName"]
		if !ok {
			http.Error(w, "Account misconfigured #2", 500)
			return
		}

		firstName, ok := x2.(string)
		if !ok {
			http.Error(w, "Account misconfigured #3", 500)
			return
		}

		stateLock.Lock()
		defer stateLock.Unlock()

		// Render the email contents
		o1buf := &bytes.Buffer{}
		if err := o1tpl.Execute(o1buf, map[string]interface{}{
			"first_name": firstName,
		}); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}
		o2buf := &bytes.Buffer{}
		if err := o2tpl.Execute(o2buf, map[string]interface{}{
			"first_name": firstName,
		}); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}
		o3buf := &bytes.Buffer{}
		if err := o3tpl.Execute(o3buf, map[string]interface{}{
			"first_name": firstName,
		}); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}
		o4buf := &bytes.Buffer{}
		if err := o4tpl.Execute(o4buf, map[string]interface{}{
			"first_name": firstName,
		}); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}

		// Four emails in total
		timers := []*Timer{
			// 1. Welcome to Lavaboom
			&Timer{
				ID:      uniuri.NewLen(uniuri.UUIDLen),
				Time:    time.Now().Add(time.Second * 55),
				From:    "Felix from Lavaboom <hello@lavaboom.com>",
				To:      []string{account.StyledName + "@lavaboom.com"},
				Subject: "Welcome to Lavaboom",
				Body:    o1buf.String(),
			},
			// 2. Getting started
			&Timer{
				ID:      uniuri.NewLen(uniuri.UUIDLen),
				Time:    time.Now().Add(time.Second * 60),
				From:    "Julie from Lavaboom <hello@lavaboom.com>",
				To:      []string{account.StyledName + "@lavaboom.com"},
				Subject: "Getting started with Lavaboom",
				Body:    o2buf.String(),
			},
			// 3. Security information
			&Timer{
				ID:      uniuri.NewLen(uniuri.UUIDLen),
				Time:    time.Now().Add(time.Minute * 2),
				From:    "Andrei from Lavaboom <hello@lavaboom.com>",
				To:      []string{account.StyledName + "@lavaboom.com"},
				Subject: "Important security information",
				Body:    o3buf.String(),
			},
			// 4. How's it going?
			&Timer{
				ID:      uniuri.NewLen(uniuri.UUIDLen),
				Time:    time.Now().Add(time.Minute * 360),
				From:    "Lavabot from Lavaboom <hello@lavaboom.com>",
				To:      []string{account.StyledName + "@lavaboom.com"},
				Subject: "How's it going?",
				Body:    o4buf.String(),
			},
		}

		state = append(state, timers...)

		if err := r.Table("onboarding").Insert(timers).Exec(session); err != nil {
			http.Error(w, err.Error(), 500)
			log.Print(err)
			return
		}

		// Sort it and ping the worker
		sort.Sort(state)
		change <- struct{}{}

		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8000", nil)
}
