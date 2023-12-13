package befous

import (
	"context"
	"log"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ------------------------------------------------------------------ Set Connection ------------------------------------------------------------------

func SetConnection(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func SetConnection2dsphere(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	db := atdb.MongoConnect(DBmongoinfo)

	// Create a geospatial index if it doesn't exist
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "geometry", Value: "2dsphere"},
		},
	}

	_, err := db.Collection("geojson").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Printf("Error creating geospatial index: %v\n", err)
	}

	return db
}

// ----------------------------------------------------------------------- User -----------------------------------------------------------------------

// Create

func InsertUser(mongoenv *mongo.Database, collname string, datauser User) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, datauser)
}

// Read

func GetAllUser(mconn *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mconn, collname)
	return user
}

func FindUser(mconn *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mconn, collname, filter)
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, collname, filter)
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

func EditUser(mconn *mongo.Database, collname string, datauser User) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datauser)
}

// Delete

func DeleteUser(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

// ---------------------------------------------------------------------- Geojson ----------------------------------------------------------------------

// Create

func PostPoint(mconn *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return atdb.InsertOneDoc(mconn, collection, pointdata)
}

func PostLinestring(mconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mconn, collection, linestringdata)
}

func PostPolygon(mconn *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mconn, collection, polygondata)
}

// Read

func GetAllBangunan(mconn *mongo.Database, collname string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mconn, collname)
	return lokasi
}

// Update

// Delete

func DeleteGeojson(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func GeoIntersects(mconn *mongo.Database, collname string, coordinates Point) (namalokasi string) {
	lokasicollection := mconn.Collection(collname)
	filter := bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var lokasi FullGeoJson
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoIntersects: %v\n", err)
	}
	return lokasi.Properties.Name

}

func GeoWithin(mconn *mongo.Database, collname string, coordinates Polygon) (namalokasi string) {
	lokasicollection := mconn.Collection(collname)
	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates.Coordinates,
				},
			},
		},
	}
	var lokasi FullGeoJson
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoWithin: %v\n", err)
	}
	return lokasi.Properties.Name

}

func Near(mconn *mongo.Database, collname string, coordinates Point) (namalokasi string) {
	lokasicollection := mconn.Collection(collname)
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "LineString",
					"coordinates": coordinates.Coordinates,
				},
				"$maxDistance": 1000,
			},
		},
	}
	var lokasi FullGeoJson
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("Near: %v\n", err)
	}
	return lokasi.Properties.Name

}

// -------------------------------------------------------------------- Pemrograman --------------------------------------------------------------------

func GetAllKegiatan(mconn *mongo.Database, collname string) []Kegiatan {
	kegiatan := atdb.GetAllDoc[[]Kegiatan](mconn, collname)
	return kegiatan
}

func GetAllJadwal(mconn *mongo.Database, collname string) []Jadwal {
	lokasi := atdb.GetAllDoc[[]Jadwal](mconn, collname)
	return lokasi
}

func InsertMahasiswa(mongoenv *mongo.Database, collname string, datamahasiswa Mahasiswa) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, datamahasiswa)
}

func GetAllMahasiswa(mconn *mongo.Database, collname string) []Mahasiswa {
	mahasiswa := atdb.GetAllDoc[[]Mahasiswa](mconn, collname)
	return mahasiswa
}

func InsertDosen(mongoenv *mongo.Database, collname string, datadosen Dosen) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, datadosen)
}

func GetAllDosen(mconn *mongo.Database, collname string) []Dosen {
	dosen := atdb.GetAllDoc[[]Dosen](mconn, collname)
	return dosen
}

func InsertRuangan(mongoenv *mongo.Database, collname string, dataruangan Ruangan) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, dataruangan)
}

func GetAllRuangan(mconn *mongo.Database, collname string) []Ruangan {
	ruangan := atdb.GetAllDoc[[]Ruangan](mconn, collname)
	return ruangan
}

func InsertMatakuliah(mongoenv *mongo.Database, collname string, datamatakuliah Matakuliah) interface{} {
	return atdb.InsertOneDoc(mongoenv, collname, datamatakuliah)
}

func GetAllMatakuliah(mconn *mongo.Database, collname string) []Matakuliah {
	matakuliah := atdb.GetAllDoc[[]Matakuliah](mconn, collname)
	return matakuliah
}
