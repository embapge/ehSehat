package http

import (
	"appointment-queue-service/internal/queue/app"
	"appointment-queue-service/internal/queue/delivery/dto"
	"appointment-queue-service/internal/queue/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type QueueHandler struct {
	App app.QueueApp
}

func NewQueueHandler(router *httprouter.Router, app app.QueueApp) {
	handler := &QueueHandler{App: app}

	router.GET("/queues/:id", handler.FindByID)
	router.GET("/queues/today/:doctor_id", handler.FindTodayByDoctor)
	router.POST("/queues", handler.Create)
	router.PUT("/queues/:id", handler.Update)
	router.POST("/queues/generate", handler.GenerateNextQueue)
}

func (h *QueueHandler) FindByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	q, err := h.App.FindByIDQueue(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(q)
}

func (h *QueueHandler) FindTodayByDoctor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	doctorID, err := strconv.Atoi(ps.ByName("doctor_id"))
	if err != nil {
		http.Error(w, "invalid doctor_id", http.StatusBadRequest)
		return
	}

	queues, err := h.App.FindTodayByDoctor(r.Context(), uint(doctorID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(queues)
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var q domain.QueueModel
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.App.CreateQueue(r.Context(), &q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *QueueHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var q domain.QueueModel
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	q.ID = uint(id)

	if err := h.App.UpdateQueue(r.Context(), &q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func (h *QueueHandler) GenerateNextQueue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req dto.GenerateQueueRequest // ⬅️ pakai struct dari dto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	q, err := h.App.GenerateNextQueue(
		r.Context(),
		req.DoctorID,
		req.UserID,
		req.UserName,
		req.UserRole,
		req.AppointmentID,
		req.PatientID,
		req.PatientName,
		req.DoctorName,
		req.DoctorSpecialization,
		req.Type,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}
