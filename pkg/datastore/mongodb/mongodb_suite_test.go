package mongodb

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

	"github.com/oam-dev/kubevela/pkg/apiserver/infrastructure/datastore"
)

func TestMongodb(t *testing.T) {
	_, err := New(context.TODO(), datastore.Config{
		URL:      "mongodb://localhost:27017",
		Database: "kubevela",
	})
	if err != nil {
		t.Fatal(err)
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongodb Suite")
}
