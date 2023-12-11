package befous

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}

type User struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type CredentialUser struct {
	Status bool `json:"status" bson:"status"`
	Data   struct {
		Name     string `json:"name" bson:"name"`
		Username string `json:"username" bson:"username"`
		Role     string `json:"role" bson:"role"`
	} `json:"data" bson:"data"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

type FullGeoJson struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   Geometry           `json:"geometry" bson:"geometry"`
}

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type GeoJsonPoint struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates []float64 `json:"coordinates" bson:"coordinates"`
		Type        string    `json:"type" bson:"type"`
	} `json:"geometry" bson:"geometry"`
}

type GeoJsonLineString struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
		Type        string      `json:"type" bson:"type"`
	} `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   struct {
		Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
		Type        string        `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"geometry" bson:"geometry"`
}

type LongLat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// --------------------------------------------------------------------- Projek 3 ---------------------------------------------------------------------

type Kegiatan struct {
	ID      int    `json:"id" bson:"id"`
	Nama    string `json:"nama" bson:"nama"`
	Note    string `json:"note" bson:"note"`
	Tanggal string `json:"tanggal" bson:"tanggal"`
}

type Jadwal struct {
	ID   int    `json:"id" bson:"id"`
	Nama string `json:"nama" bson:"nama"`
	Hari string `json:"hari" bson:"hari"`
	Jam  string `json:"jam" bson:"jam"`
}
