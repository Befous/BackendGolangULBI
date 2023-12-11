package befous

import (
	"fmt"
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
	mconn := SetConnection("mongoenv", dbname)
	datagedung := GetAllUser(mconn, "user")
	fmt.Println(datagedung)
}

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	datagedung := GeoIntersects(mconn, collname, 103.60768133536988, -1.628526295003084)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := [][][]float64{
		{
			{103.62892373959272, -1.616812371154296},
			{103.62890068598779, -1.616866839799556},
			{103.62896041578165, -1.616890931699615},
			{103.62898556516905, -1.6168364630550514},
			{103.62892373959272, -1.616812371154296},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("mongoenv", dbname)
	datagedung := Near(mconn, collname, 103.6037314895799, -1.632582001101999)
	fmt.Println(datagedung)
}

func TestFindUser(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	datagedung := FindUser(mconn, collname, User{Username: "ibrohim"})
	fmt.Println(datagedung)
}

func TestGetAllBangunan(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	x := GetAllBangunan(mconn, collname)
	fmt.Println(x)
}
