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

type DBInfo struct {
	DBString       string
	DBName         string
	CollectionName string
}

type User struct {
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

type Point struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Polyline struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

type Polygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
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

type Mahasiswa struct {
	NPM          string `json:"npm" bson:"npm"`
	Nama_Lengkap string `json:"nama_lengkap" bson:"nama_lengkap"`
	Alamat       string `json:"alamat" bson:"alamat"`
}

type Dosen struct {
	NIDN       string `json:"nidn" bson:"nidn"`
	Nama_Dosen string `json:"nama_dosen" bson:"nama_dosen"`
	Alamat     string `json:"alamat" bson:"alamat"`
}

type Ruangan struct {
	Kode_Ruangan string `json:"kode_ruangan" bson:"kode_ruangan"`
	Nama_Ruangan string `json:"nama_ruangan" bson:"nama_ruangan"`
	Kapasitas    string `json:"kapasitas" bson:"kapasitas"`
}

type Matakuliah struct {
	Kode_Matakuliah string `json:"kode_matakuliah" bson:"kode_matakuliah"`
	Nama_Matakiliah string `json:"nama_matakuliah" bson:"nama_matakuliah"`
	SKS             string `json:"sks" bson:"sks"`
}
