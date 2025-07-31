package http

import (
	"appointment-queue-service/internal/appointment/app"
	"appointment-queue-service/internal/appointment/domain"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AppointmentHandler struct {
	App app.AppointmentApp
}

func NewAppointmentHandler(router *httprouter.Router, app app.AppointmentApp) {
	handler := &AppointmentHandler{App: app}

	router.GET("/appointments", handler.FindAll)
	router.GET("/appointments/:id", handler.FindByID)
	router.GET("/appointments-by-user/:user_id", handler.FindByUserID)
	router.POST("/appointments", handler.CreateAppointment)
	router.PATCH("/appointments/:id", handler.PatchAppointment)
	router.PUT("/appointments/:id/mark-paid", handler.MarkAsPaid)
}

func (h *AppointmentHandler) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	appointments, err := h.App.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(appointments)
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
func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

// PATCH /appointments/:id
func (h *AppointmentHandler) PatchAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	existing, err := h.App.FindByIDAppointment(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "appointment not found", http.StatusNotFound)
		return
	}

	// Decode only provided fields
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Apply updates
	for k, v := range updates {
		switch k {
		case "user_full_name":
			existing.UserFullName = v.(string)
		case "doctor_name":
			existing.DoctorName = v.(string)
		case "doctor_specialization":
			existing.DoctorSpecialization = v.(string)
		case "appointment_at":
			if t, err := time.Parse(time.RFC3339, v.(string)); err == nil {
				existing.AppointmentAt = t
			}
		case "status":
			existing.Status = v.(string)
		case "is_paid":
			existing.IsPaid = v.(bool)
		}
	}

	updated, err := h.App.UpdateAppointment(r.Context(), existing.ID, existing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	updated, err := h.App.FindByIDAppointment(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
