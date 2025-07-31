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

	router.GET("/queues/:id", handler.FindByIDQueue)
	router.GET("/queues-today/:doctor_id", handler.FindTodayByDoctor)
	router.POST("/queues", handler.CreateQueue)
	router.PUT("/queues/:id", handler.UpdateQueue)
	router.POST("/queues/generate", handler.GenerateNextQueue)
}

// @Summary Get queue by ID
// @Description Get queue detail by ID
// @Tags queues
// @Produce json
// @Param id path int true "Queue ID"
// @Success 200 {object} domain.QueueModel
// @Failure 400 {string} string "invalid id"
// @Failure 404 {string} string "queue not found"
// @Router /queues/{id} [get]
func (h *QueueHandler) FindByIDQueue(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// @Summary Get today's queues by doctor
// @Description Get all today's queues for a specific doctor
// @Tags queues
// @Produce json
// @Param doctor_id path int true "Doctor ID"
// @Success 200 {array} domain.QueueModel
// @Failure 400 {string} string "invalid doctor_id"
// @Failure 500 {string} string "internal server error"
// @Router /queues-today/{doctor_id} [get]
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

// @Summary Create a queue
// @Description Create a new queue record
// @Tags queues
// @Accept json
// @Produce json
// @Param queue body domain.QueueModel true "Queue data"
// @Success 201 {object} domain.QueueModel
// @Failure 400 {string} string "invalid JSON or validation error"
// @Router /queues [post]
func (h *QueueHandler) CreateQueue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var q domain.QueueModel
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	// Ambil value dari header
	if userIDStr := r.Header.Get("ts-user-id"); userIDStr != "" {
		q.UserID = userIDStr // UUID sebagai string
	}
	q.UserName = r.Header.Get("ts-user-name")
	q.UserRole = r.Header.Get("ts-user-role")
	q.UserEmail = r.Header.Get("ts-user-email")

	if err := h.App.CreateQueue(r.Context(), &q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

// @Summary Update queue
// @Description Update an existing queue by ID
// @Tags queues
// @Accept json
// @Produce json
// @Param id path int true "Queue ID"
// @Param queue body domain.QueueModel true "Updated queue data"
// @Success 200 {object} domain.QueueModel
// @Failure 400 {string} string "invalid input"
// @Router /queues/{id} [put]
func (h *QueueHandler) UpdateQueue(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// @Summary Generate next queue number
// @Description Automatically generate the next queue number for a doctor and create a queue
// @Tags queues
// @Accept json
// @Produce json
// @Param data body dto.GenerateQueueRequest true "Queue generation input"
// @Success 201 {object} domain.QueueModel
// @Failure 400 {string} string "invalid request or generation failed"
// @Router /queues/generate [post]
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
		req.UserEmail,
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
