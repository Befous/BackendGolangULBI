package befous

import (
	"fmt"
	"os"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "befous"
var collname = "geojson"

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("Private Key: " + privateKey)
	fmt.Println("Public Key: " + publicKey)
}

func TestEncode(t *testing.T) {
	name := "Test Nama"
	username := "Test Username"
	role := "Test Role"

	tokenstring, err := Encode(name, username, role, privatekey)
	fmt.Println("error : ", err)
	fmt.Println("token : ", tokenstring)
}

func TestDecode(t *testing.T) {
	pay, err := Decode(publickey, encode)
	name := DecodeGetName(publickey, encode)
	username := DecodeGetUsername(publickey, encode)
	role := DecodeGetRole(publickey, encode)

	fmt.Println("name :", name)
	fmt.Println("username :", username)
	fmt.Println("role :", role)
	fmt.Println("err : ", err)
	fmt.Println("payload : ", pay)
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", dbname)
	datagedung := GetAllUser(mconn, "user")
	fmt.Println(datagedung)
}
func TestCobaUsernameExists(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", dbname)
	var user User
	user.Username = "befous44th"
	datagedung := UsernameExists(mconn, "user", user)
	fmt.Println(datagedung)
}

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := Point{
		Coordinates: []float64{
			103.60768133536988, -1.628526295003084,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{103.62892373959272, -1.616812371154296},
				{103.62890068598779, -1.616866839799556},
				{103.62896041578165, -1.616890931699615},
				{103.62898556516905, -1.6168364630550514},
				{103.62892373959272, -1.616812371154296},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOCONNSTRINGENV", dbname, collname)
	coordinates := Point{
		Coordinates: []float64{
			103.6037314895799, -1.632582001101999,
		},
	}
	datagedung := Near(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", dbname)
	coordinates := Point{
		Coordinates: []float64{
			103.6037314895799, -1.632582001101999,
		},
	}
	datagedung := NearSphere(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", dbname)
	coordinates := Polyline{
		Coordinates: [][]float64{
			{102.6037314895799, -1.732582001101999},
			{104.6037314895799, -1.532582001101999},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestFindUser(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	datagedung := FindUser(mconn, collname, User{Username: "ibrohim"})
	tokenstring, tokenerr := Encode(datagedung.Name, datagedung.Username, datagedung.Role, os.Getenv("privatekey"))
	if tokenerr != nil {
		fmt.Println("Gagal encode token: " + tokenerr.Error())
	}
	fmt.Println(tokenstring)
}

func TestGetAllBangunan(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	x := GetAllBangunan(mconn, collname)
	fmt.Println(x)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	var user User
	user.Name = "test"
	InsertUser(mconn, "test", user)
}
