package infra

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var consultation domain.Consultation
	fmt.Println("id mongo consultation:", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&consultation)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	return &consultation, nil
}

func (m *mongoConsultation) Create(ctx context.Context, consultation *domain.Consultation) error {
	result, err := m.collection.InsertOne(ctx, consultation)
	if err != nil {
		return err
	}
	consultation.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (m *mongoConsultation) Update(ctx context.Context, consultation *domain.Consultation) error {
	objID, err := primitive.ObjectIDFromHex(consultation.ID)
	if err != nil {
		return err
	}

	updateData := bson.M{}
	// Convert consultation to BSON and remove _id
	bsonBytes, err := bson.Marshal(consultation)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, &updateData)
	if err != nil {
		return err
	}
	delete(updateData, "_id")

	_, err = m.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		fmt.Println("error updating consultation:", err)
		return err
	}
	return nil
}
