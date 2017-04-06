package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		response.JSON(w, response.Result{Result: "Error Creating"}, http.StatusNotAcceptable)
		return
	}

	handler.DB.Create(&doc)
	response.JSON(w, response.Result{Result: "Document Created"}, http.StatusCreated)
}

// GetByID fetches a document given the ID
func (handler DocHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	docID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var doc models.Doc

	if err := handler.DB.First(&doc, docID).Error; err != nil {
		response.JSON(w, response.Result{Result: "Document not found"}, http.StatusNotFound)
		return
	}

	response.JSON(w, doc, http.StatusOK)
}

// Exists checks the document title to see if it exists
func (handler DocHandler) Exists(w http.ResponseWriter, r *http.Request) {
	docTitle := chi.URLParam(r, "title")

	var doc models.Doc
	result := true

	if err := handler.DB.Where("title = ?", docTitle).First(&doc).Error; err != nil {
		result = false
	}

	response.JSON(w, response.Result{Result: result}, http.StatusOK)
}
