package database

import (
	"github.com/Basileus1990/CrowdCompute.git/dataStructures"
)

// temporary
var db map[string]dataStructures.Task

func init() {
	db = make(map[string]dataStructures.Task)
}
