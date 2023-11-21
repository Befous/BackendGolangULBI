package peda

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	datagedung := GetAllUser(mconn, "user")
	fmt.Println(datagedung)
}

// 	result := GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname, datauser)
// 	fmt.Println(result)
// 	// You can add assertions here to validate the result, or check the database for the created user.
// }

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"
	userdata.Role = "admin"
	mconn := SetConnection("MONGOCONNSTRINGENV", "befous")
	CreateNewUserRole(mconn, "user", userdata)
}

func TestCreateNewUserToken(t *testing.T) {
	// Create a User struct
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"
	userdata.Role = "admin"

	// Generate private and public keys using watoken.GenerateKey
	privateKey, publicKey := watoken.GenerateKey()

	// Store the private and public keys in the userdata
	userdata.Private = privateKey
	userdata.Publick = publicKey // Corrected the field name from Publick to Public

	// Encode a token using the privateKey
	hasil, err := watoken.Encode("befous", privateKey)
	fmt.Println(hasil, err)
	if err != nil {
		t.Errorf("Failed to create user and token: %v", err)
	} else {
		t.Logf("User and token created successfully")

		// Assuming you have a MongoDB client and a database connection, use the client and connection to insert the userdata
		// Replace "yourDatabaseName" with your actual database name
		client, err := mongo.NewClient(options.Client().ApplyURI("mongoenv"))
		if err != nil {
			t.Errorf("Failed to create MongoDB client: %v", err)
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err = client.Connect(ctx)
			if err != nil {
				t.Errorf("Failed to connect to MongoDB: %v", err)
			} else {
				// Use the database name and collection name where you want to insert the user data
				db := client.Database("befous")
				collection := db.Collection("user")

				_, err = collection.InsertOne(ctx, userdata)
				if err != nil {
					t.Errorf("Failed to insert user data into MongoDB: %v", err)
				} else {
					t.Logf("User data inserted into MongoDB successfully")
				}
			}
		}
	}
}

func TestDeleteUser(t *testing.T) {

	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	DeleteUser(mconn, "user", userdata)
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestFunciionUser(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestGeneratePasswordHashh(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestHashFunctionn(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}
func TestFindUser(t *testing.T) {
	var userdata User
	userdata.Username = "befous"
	mconn := SetConnection("mongoenv", "befous")
	res := FindUser(mconn, "user", userdata)
	fmt.Println(res)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "befous"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	// privateKey, publicKey := watoken.GenerateKey()
	privateKey := "535ff96a3ffc289eccda837a9e323e66e00855a1918fcb9a20889f77d85bf04d9bac8917faf1a9457de01191ce77048661073449cbd6e79e4885128379db0623"
	publicKey := "9bac8917faf1a9457de01191ce77048661073449cbd6e79e4885128379db0623"
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("ibrohim", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("mongoenv", "befous")
	var userdata User
	userdata.Username = "befous"
	userdata.Password = "befous"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestWatoken(t *testing.T) {
	privateKey := "535ff96a3ffc289eccda837a9e323e66e00855a1918fcb9a20889f77d85bf04d9bac8917faf1a9457de01191ce77048661073449cbd6e79e4885128379db0623"
	userid := "ibrohim"

	tokenstring, err := watoken.EncodeforMinutes(userid, privateKey, 1)
	fmt.Println("error : ", err)
	fmt.Println("token : ", tokenstring)
}

func TestDecode1Menit(t *testing.T) {
	publicKey := "9bac8917faf1a9457de01191ce77048661073449cbd6e79e4885128379db0623"

	tokenstring := "v4.public.eyJleHAiOiIyMDIzLTExLTIyVDAxOjU4OjA1KzA3OjAwIiwiaWF0IjoiMjAyMy0xMS0yMlQwMTo1NzowNSswNzowMCIsImlkIjoiaWJyb2hpbSIsIm5iZiI6IjIwMjMtMTEtMjJUMDE6NTc6MDUrMDc6MDAifQHf-cyyCpWvUMD4OJhdaWIKrKgtgja0eBbzOol0JXLMu5AqlwvEKPDS7bhmBLF0isDd954cAR3fjRu6dnIiCA4"
	body := watoken.DecodeGetId(publicKey, tokenstring)
	fmt.Println("isi : ", body)
}
