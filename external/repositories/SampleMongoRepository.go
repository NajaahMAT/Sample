package repositories

import (
	"Sample/domain/entities"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sample struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"userId"`
	URL    string             `bson:"url"`
	Title  string             `bson:"title"`
}

type SampleRepository struct {}


func (r SampleRepository) GetSample(ctx context.Context) (resp []*entities.Sample, err error) {
	//cur, err := r.db.Find(ctx, bson.M{})
	//defer cur.Close(ctx)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//out := make([]*Sample, 0)
	//
	//for cur.Next(ctx) {
	//	user := new(Sample)
	//	err := cur.Decode(user)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	out = append(out, user)
	//}
	//if err := cur.Err(); err != nil {
	//	return nil, err
	//}
	//
	//return toSamples(out), nil
	return resp, nil
}


//func NewSampleRepository(db *mongo.Database, collection string) *SampleRepository {
//	return &SampleRepository{
//		db: db.Collection(collection),
//	}
//}

//func (r SampleRepository) CreateSample(ctx context.Context, sample *entities.Sample) error {
//
//	model := toModel(sample)
//
//	res, err := db.Collection().InsertOne(ctx, model)
//	if err != nil {
//		return err
//	}
//
//	sample.ID = res.InsertedID.(primitive.ObjectID).Hex()
//
//	return nil
//}
//
//func (r SampleRepository) GetSample(ctx context.Context) ([]*entities.Sample, error) {
//	//uid, _ := primitive.ObjectIDFromHex(sample.ID)
//	cur, err := r.db.Find(ctx, bson.M{})
//	defer cur.Close(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	out := make([]*Sample, 0)
//
//	for cur.Next(ctx) {
//		user := new(Sample)
//		err := cur.Decode(user)
//		if err != nil {
//			return nil, err
//		}
//
//		out = append(out, user)
//	}
//	if err := cur.Err(); err != nil {
//		return nil, err
//	}
//
//	return toSamples(out), nil
//}

//func (r SampleRepository) DeleteSample(ctx context.Context, user *entities.Sample, id string) error {
//	objID, _ := primitive.ObjectIDFromHex(id)
//	uID, _ := primitive.ObjectIDFromHex(user.ID)
//
//	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
//	return err
//}
//
//func toModel(b *entities.Sample) *Sample {
//	uid, _ := primitive.ObjectIDFromHex(b.UserID)
//
//	return &Sample{
//		UserID: uid,
//		URL:    b.URL,
//		Title:  b.Title,
//	}
//}

//func toSample(b *Sample) *entities.Sample {
//	return &entities.Sample{
//		ID:     b.ID.Hex(),
//		UserID: b.UserID.Hex(),
//		URL:    b.URL,
//		Title:  b.Title,
//	}
//}
//
//func toSamples(bs []*Sample) []*entities.Sample {
//	out := make([]*entities.Sample, len(bs))
//
//	for i, b := range bs {
//		out[i] = toSample(b)
//	}
//
//	return out
//}
