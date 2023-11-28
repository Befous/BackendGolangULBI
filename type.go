package peda

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeometryPolygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string        `json:"type,omitempty" bson:"type,omitempty"`
}

type GeometryLineString struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeometryPoint struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

type GeoJsonPoint struct {
	Type       string        `json:"type" bson:"type"`
	Properties Properties    `json:"properties" bson:"properties"`
	Geometry   GeometryPoint `json:"geometry" bson:"geometry"`
}

type GeoJsonLineString struct {
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   GeometryLineString `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string          `json:"type" bson:"type"`
	Properties Properties      `json:"properties" bson:"properties"`
	Geometry   GeometryPolygon `json:"geometry" bson:"geometry"`
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

type Lokasi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Properties Name               `bson:"properties,omitempty"`
	Geometry   Geometry           `bson:"geometry,omitempty"`
	Kategori   string             `bson:"kategori,omitempty"`
}

type Name struct {
	Name string `bson:"name,omitempty"`
}

type LongLat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// --------------------------------------------------------------------- Projek 3 ---------------------------------------------------------------------

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type Pesan struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

type Coordinate struct {
	Type        string    `json:"type" bson:"type"`
	Name        string    `json:"name" bson:"name"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

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

type User struct {
	Name     string    `json:"name,omitempty" bson:"name,omitempty"`
	Email    string    `json:"email,omitempty" bson:"email,omitempty"`
	Username string    `json:"username" bson:"username"`
	Password string    `json:"password" bson:"password"`
	Role     SemuaRole `json:"role,omitempty" bson:"role,omitempty"`
	Nomor_hp string    `json:"nomor_hp,omitempty" bson:"nomor_hp,omitempty"`
}

type SemuaRole struct {
	Admin  bool `json:"admin" bson:"admin"`
	Author bool `json:"author" bson:"author"`
	User   bool `json:"user" bson:"user"`
}

type CredentialUser struct {
	Status  bool   `json:"status" bson:"status"`
	Data    User   `json:"data,omitempty" bson:"data,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

type CredentialBerita struct {
	Status  bool   `json:"status" bson:"status"`
	Data    Berita `json:"data,omitempty" bson:"data,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

type CredentialGeojson struct {
	Status  bool    `json:"status" bson:"status"`
	Data    GeoJson `json:"data,omitempty" bson:"data,omitempty"`
	Message string  `json:"message,omitempty" bson:"message,omitempty"`
	Token   string  `json:"token,omitempty" bson:"token,omitempty"`
}

type Berita struct {
	ID       string   `json:"id" bson:"id"`
	Kategori string   `json:"kategori" bson:"kategori"`
	Judul    string   `json:"judul" bson:"judul"`
	Preview  string   `json:"preview" bson:"preview"`
	Konten   Paragraf `json:"konten" bson:"konten"`
}

type Paragraf struct {
	Paragraf1  string `json:"paragraf1,omitempty" bson:"paragraf1,omitempty"`
	Paragraf2  string `json:"paragraf2,omitempty" bson:"paragraf2,omitempty"`
	Paragraf3  string `json:"paragraf3,omitempty" bson:"paragraf3,omitempty"`
	Paragraf4  string `json:"paragraf4,omitempty" bson:"paragraf4,omitempty"`
	Paragraf5  string `json:"paragraf5,omitempty" bson:"paragraf5,omitempty"`
	Paragraf6  string `json:"paragraf6,omitempty" bson:"paragraf6,omitempty"`
	Paragraf7  string `json:"paragraf7,omitempty" bson:"paragraf7,omitempty"`
	Paragraf8  string `json:"paragraf8,omitempty" bson:"paragraf8,omitempty"`
	Paragraf9  string `json:"paragraf9,omitempty" bson:"paragraf9,omitempty"`
	Paragraf10 string `json:"paragraf10,omitempty" bson:"paragraf10,omitempty"`
}
