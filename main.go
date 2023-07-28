package main

import (
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/job"
	"github.com/marqstree/gstep/route"
	"github.com/marqstree/gstep/util/db"
)

func main() {
	config.Setup()
	util_db.Setup()
	route.Setup()
	job.Setup()
}
