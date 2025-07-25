package http

import (
	"appointment-queue-service/internal/appointment/app"
	"appointment-queue-service/internal/appointment/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AppointmentHandler struct {
	App app.AppointmentApp
}

func NewAppointmentHandler(router *httprouter.Router, app app.AppointmentApp) {
	handler := &AppointmentHandler{App: app}

	router.GET("/appointments/:id", handler.FindByID)
	router.GET("/appointments/user/:user_id", handler.FindByUserID)
	router.POST("/appointments", handler.Create)
	router.PUT("/appointments/:id", handler.Update)
	router.PUT("/appointments/:id/mark-paid", handler.MarkAsPaid)
}

// GET /appointments/:id
func (h *AppointmentHandler) FindByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid appointment id", http.StatusBadRequest)
		return
	}

	a, err := h.App.FindByIDAppointment(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(a)
}

// GET /appointments/user/:user_id
func (h *AppointmentHandler) FindByUserID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	list, err := h.App.FindByUserID(r.Context(), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

// POST /appointments
func (h *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var a domain.AppointmentModel
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := h.App.CreateAppointment(r.Context(), &a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// PUT /appointments/:id
func (h *AppointmentHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var a domain.AppointmentModel
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	updated, err := h.App.UpdateAppointment(r.Context(), uint(id), &a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

// PUT /appointments/:id/mark-paid
func (h *AppointmentHandler) MarkAsPaid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.App.MarkAsPaid(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
