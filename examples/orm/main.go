package main

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	bigquery "github.com/HeliosData/go-sql-bigquery"
	_ "github.com/HeliosData/go-sql-bigquery/dialects/bigquery"
	"os"
	"time"
)

func main() {
	var err error
	var db *gorm.DB
	logrus.SetLevel(logrus.DebugLevel)

	// Get the Connection String from the Environment Variable of BIGQUERY_CONNECTION_STRING
	uri := os.Getenv(bigquery.ConnectionStringEnvKey)
	if db, err = gorm.Open("bigquery", uri); err != nil {
		logrus.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&Animal{})

	// Add an animal
	django := &Animal{
		Name: "Django",
		Size: 1,
		Born: time.Now(),
		Uuid: uuid.NewV4(),
		Blob: []byte("Hello World"),
		Kind: 4,
	}
	err = db.Save(django).Error
	if err != nil {
		logrus.Fatal(err)
	}

}

type Animal struct {
	Name       string
	Size       int64
	Born       time.Time
	Blob       []byte
	InstanceId uuid.UUID `gorm:"type:uuid"`
	Kind       int32
	Parent     uuid.UUID `gorm:"type:uuid"`
	Uuid       uuid.UUID `gorm:"type:uuid"`
}



//func (a *Animal) TableName() string {
//	return "animals"
//}
