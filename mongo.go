package peda

import (
	"context"
	"log"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func SetConnectionTest(mongostring, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: mongostring,
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllBangunanLineString(mongoenv *mongo.Database, collname string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoenv, collname)
	return lokasi
}

func PostPoint(mongoconn *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, pointdata)
}

func PostLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, linestringdata)
}

func PostPolygon(mongoconn *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, polygondata)
}

func MemasukkanKoordinat(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(Coordinate)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}

func GeoIntersects(mongoconn *mongo.Database, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("geojson")
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
			},
		},
	}
	var lokasi Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoIntersects: %v\n", err)
	}
	return lokasi.Properties.Name

}

func GeoWithin(mongoconn *mongo.Database, coordinates [][][]float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("geojson")
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates,
				},
			},
		},
	}
	var lokasi Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoWithin: %v\n", err)
	}
	return lokasi.Properties.Name

}

// --------------------------------------------------------------------- Projek 3 ---------------------------------------------------------------------

// ---------------------------------------------------------------------- User

// Create

func InsertUserdata(mongoenv *mongo.Database, collname, name, email, username, password string, admin, author bool) (InsertedID interface{}) {
	req := new(User)
	req.Name = name
	req.Email = email
	req.Username = username
	req.Password = password
	req.Role.Admin = admin
	req.Role.Author = author
	req.Role.User = true
	return atdb.InsertOneDoc(mongoenv, collname, req)
}

// Read

func GetAllUser(mongoenv *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mongoenv, collname)
	return user
}

func FindUser(mongoenv *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoenv, collname, filter)
}

func IsPasswordValid(mongoenv *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoenv, collname, filter)
	hashChecker := CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func usernameExists(mongoenv, dbname string, userdata User) bool {
	mconn := SetConnection(mongoenv, dbname).Collection("user")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

// Update

func EditUser(mongoenv *mongo.Database, collname, name, email, username, password string, admin, author, user bool) interface{} {

	req := new(User)
	req.Name = name
	req.Email = email
	req.Username = username
	req.Password = password
	req.Role.Admin = admin
	req.Role.Author = author
	req.Role.User = user
	filter := bson.M{"username": username}
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, req)
}

// Delete

func DeleteUser(mongoenv *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mongoenv, collname, filter)
}

//-------------------------------------------------------------------- Berita

// Create

func InsertBerita(mongoenv *mongo.Database, collname string, databerita Berita) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, databerita)
}

// Read

func GetAllBerita(mongoenv *mongo.Database, collname string) []Berita {
	berita := atdb.GetAllDoc[[]Berita](mongoenv, collname)
	return berita
}

func FindBerita(mongoenv *mongo.Database, collname string, databerita Berita) Berita {
	filter := bson.M{"id": databerita.ID}
	return atdb.GetOneDoc[Berita](mongoenv, collname, filter)
}

func idBeritaExists(mongoenv, dbname string, databerita Berita) bool {
	mconn := SetConnection(mongoenv, dbname).Collection("berita")
	filter := bson.M{"id": databerita.ID}

	var berita Berita
	err := mconn.FindOne(context.Background(), filter).Decode(&berita)
	return err == nil
}

// Update

func EditBerita(mongoenv *mongo.Database, collname string, databerita Berita) interface{} {
	filter := bson.M{"id": databerita.ID}
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, databerita)
}

// Delete

func DeleteBerita(mongoenv *mongo.Database, collname string, databerita Berita) interface{} {
	filter := bson.M{"id": databerita.ID}
	return atdb.DeleteOneDoc(mongoenv, collname, filter)
}

// -------------------------------------------------------------------- Pemrograman --------------------------------------------------------------------

func GetAllKegiatan(mongoenv *mongo.Database, collname string) []Kegiatan {
	kegiatan := atdb.GetAllDoc[[]Kegiatan](mongoenv, collname)
	return kegiatan
}

func GetAllJadwal(mongoenv *mongo.Database, collname string) []Jadwal {
	lokasi := atdb.GetAllDoc[[]Jadwal](mongoenv, collname)
	return lokasi
}
