package infra

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoConsultation struct {
	collection *mongo.Collection
}

// FindByID(id string) (*Consultation, error)
// Create(consultation *Consultation) error
// Update(consultation *Consultation) error

func NewMongoConsultation(collection *mongo.Collection) *mongoConsultation {
	return &mongoConsultation{collection: collection}
}

// mongo driver find by id
func (m *mongoConsultation) FindByID(ctx context.Context, id string) (*domain.Consultation, error) {
	// Implementation for finding a consultation by ID
	var consultation domain.Consultation
	err := m.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&consultation)
	if err != nil {
		return nil, err
	}
	return &consultation, nil
}

func (m *mongoConsultation) Create(ctx context.Context, consultation *domain.Consultation) error {
	// Implementation for creating a new consultation
	_, err := m.collection.InsertOne(ctx, consultation)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoConsultation) Update(ctx context.Context, consultation *domain.Consultation) error {
	// Implementation for updating an existing consultation
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": consultation.ID}, bson.M{"$set": consultation})
	if err != nil {
		return err
	}
	return nil
}
