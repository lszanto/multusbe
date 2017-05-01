package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/models"
	"github.com/lszanto/multusbe/response"
	"github.com/pressly/chi"
	validator "gopkg.in/go-playground/validator.v9"
)

// DocHandler handles docs
type DocHandler struct {
	DB *gorm.DB
}

// NewDocHandler creates a new doc handler
func NewDocHandler(db *gorm.DB) DocHandler {
	return DocHandler{DB: db}
}

// Create lets us create a new document
func (handler DocHandler) Create(w http.ResponseWriter, r *http.Request) {
	var doc models.Doc
	json.NewDecoder(r.Body).Decode(&doc)

	validate := validator.New()

	err := validate.Struct(doc)
	if err != nil {
		response.JSON(w, response.Result{Error: "Error creating document"}, http.StatusNotAcceptable)
		return
	}

	handler.DB.Create(&doc)
	response.JSON(w, response.Result{Result: "Document Created"}, http.StatusCreated)
}

// Get fetches a document given the title
func (handler DocHandler) Get(w http.ResponseWriter, r *http.Request) {
	docTitle, _ := url.QueryUnescape(chi.URLParam(r, "title"))
	docTitle = strings.ToLower(docTitle)

	var doc models.Doc

	if err := handler.DB.Where("title = ?", docTitle).First(&doc).Error; err != nil {
		response.JSON(w, response.Result{Error: "Document not found"}, http.StatusNotFound)
		return
	}

	response.JSON(w, doc, http.StatusOK)
}
