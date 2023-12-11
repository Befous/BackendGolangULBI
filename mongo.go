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

func EditUser(mongoenv *mongo.Database, collname string, datauser User) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, datauser)
}

// Delete

func DeleteUser(mongoenv *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mongoenv, collname, filter)
}

// ---------------------------------------------------------------------- Geojson ----------------------------------------------------------------------

// Create

func PostPoint(mongoconn *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, pointdata)
}

func PostLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, linestringdata)
}

func PostPolygon(mongoconn *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, polygondata)
}

// Read

func GetAllBangunan(mongoenv *mongo.Database, collname string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoenv, collname)
	return lokasi
}

// Update

// Delete

func DeleteGeojson(mongoenv *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mongoenv, collname, filter)
}

func GeoIntersects(mongoconn *mongo.Database, collname string, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection(collname)
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
	var lokasi FullGeoJson
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoIntersects: %v\n", err)
	}
	return lokasi.Properties.Name

}

func GeoWithin(mongoconn *mongo.Database, collname string, coordinates [][][]float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection(collname)
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
	var lokasi FullGeoJson
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoWithin: %v\n", err)
	}
	return lokasi.Properties.Name

}

func Near(mongoconn *mongo.Database, collname string, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection(collname)
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "LineString",
					"coordinates": []float64{long, lat},
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

func GetAllKegiatan(mongoenv *mongo.Database, collname string) []Kegiatan {
	kegiatan := atdb.GetAllDoc[[]Kegiatan](mongoenv, collname)
	return kegiatan
}

func GetAllJadwal(mongoenv *mongo.Database, collname string) []Jadwal {
	lokasi := atdb.GetAllDoc[[]Jadwal](mongoenv, collname)
	return lokasi
}
