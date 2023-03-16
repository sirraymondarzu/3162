// Filename: cmd/web/data.go
package main

import (
	"github.com/lewisdalwin/poll/internal/models"
)

type templateData struct {
	Question *models.Question
}

