package peda

type GeometryPolygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string        `json:"type" bson:"type"`
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

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type User struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
	Token    string `json:"token,omitempty" bson:"token,omitempty"`
	Private  string `json:"private,omitempty" bson:"private,omitempty"`
	Publick  string `json:"publick,omitempty" bson:"publick,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type ResponseDataUser struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty"`
}

type Response struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
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

type Token struct {
	Key    string
	Values string
}

type PostToken struct {
	Response string `json:"response"`
}

type Jaja struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
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

//-----------------------test buat project, nanti hapus kalo udah

type Berita struct {
	ID       string `json:"id" bson:"id"`
	Kategori string `json:"kategori" bson:"kategori"`
	Judul    string `json:"judul" bson:"judul"`
	Preview  string `json:"preview" bson:"preview"`
	Konten   string `json:"konten" bson:"konten"`
}

type AuthorizationStruct struct {
	Status  bool   `json:"status" bson:"status"`
	Data    User   `json:"data,omitempty" bson:"data,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
