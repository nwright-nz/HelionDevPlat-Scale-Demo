package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudfoundry-samples/lattice-app/helpers"
)

type Hello struct {
	Time time.Time
}

func (p *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, _ := helpers.FetchIndex()
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
	})
	styledTemplate.Execute(w, Body{Body: fmt.Sprintf(`
<div class="hello">
	
	Helion Platform Scaling Test - WWAS 2016
</div>

<div class="my-index">My instance number is</div>

<div class="index">%d</div>
<div class="mid-color">Uptime: %s</div>
<div class="bottom-color"></div>
    `, index, time.Since(p.Time))})
}

