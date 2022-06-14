package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Iterator mongo iterator implementation
type Iterator struct {
	cur *mongo.Cursor
}

// Close iterator close
func (i *Iterator) Close(ctx context.Context) error {
	return i.cur.Close(ctx)
}

// Next read next data
func (i *Iterator) Next(ctx context.Context) bool {
	return i.cur.Next(ctx)
}

// Decode decode data
func (i *Iterator) Decode(entity interface{}) error {
	return i.cur.Decode(entity)
}
