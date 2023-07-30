package main

import (
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/job"
	"github.com/marqstree/gstep/route"
	"github.com/marqstree/gstep/util/db/DbUtil"
)

func main() {
	config.Setup()
	DbUtil.Setup()
	route.Setup()
	job.Setup()
}
