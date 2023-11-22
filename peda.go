package peda

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func MembuatGeojsonPoint(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
	if err != nil {
		return err.Error()
	}
	PostPoint(mconn, collname, geojsonpoint)
	return ReturnStruct(geojsonpoint)
}

func MembuatGeojsonPolyline(mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		var geojsonline GeoJsonLineString
		err := json.NewDecoder(r.Body).Decode(&geojsonline)
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			PostLinestring(mconn, collname, geojsonline)
			response.Message = "Data polyline berhasil masuk"
		}
	} else {
		response.Message = "Token Salah"
	}
	return ReturnStruct(response)
}

func MembuatGeojsonPolygon(mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		var geojsonpolygon GeoJsonPolygon
		err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			PostPolygon(mconn, collname, geojsonpolygon)
			response.Message = "Data polygon berhasil masuk"
		}
	} else {
		response.Message = "Token Salah"
	}
	return ReturnStruct(response)
}

func AmbilDataGeojson(mongoenv, dbname, collname string) string {
	mconn := SetConnection(mongoenv, dbname)
	datagedung := GetAllBangunanLineString(mconn, collname)
	return ReturnStruct(datagedung)
}

func RegistrasiUser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if usernameExists(mongoenv, dbname, datauser) {
		response.Status = false
		response.Message = "Username telah dipakai"
	} else {
		response.Status = true
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			response.Status = true
			hash, hashErr := HashPassword(datauser.Password)
			if hashErr != nil {
				response.Message = "Gagal Hash Password" + err.Error()
			}
			InsertUserdata(mconn, collname, datauser.Name, datauser.Email, datauser.Username, hash, datauser.Role.Admin, datauser.Role.Author)
			response.Message = "Berhasil Input data"
		}
	}
	return ReturnStruct(response)
}

func LoginUser(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collname, datauser) {
			response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(privatekey))
			if err != nil {
				response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				response.Message = "Selamat Datang"
				response.Token = tokenstring
			}
		} else {
			response.Message = "Password Salah"
		}
	}
	return ReturnStruct(response)
}

func HapusUser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		DeleteUser(mconn, collname, datauser)
		response.Message = "Berhasil Delete data"
	}
	return ReturnStruct(response)
}

func MembuatGeojsonPointHeader(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)

	var datapoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&datapoint)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		err := PostPoint(mconn, collname, datapoint)
		if err != nil {
			return ReturnStruct(CreateResponse(true, "Success: Point created", datapoint))
		} else {
			return ReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return ReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}
}

func MembuatGeojsonPolylineHeader(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)

	var datapolyline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&datapolyline)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		err := PostLinestring(mconn, collname, datapolyline)
		if err != nil {
			return ReturnStruct(CreateResponse(true, "Success: LineString created", datapolyline))
		} else {
			return ReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return ReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}
}

func MembuatGeojsonPolygonHeader(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)

	var datapolygon GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapolygon)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		err := PostPolygon(mconn, collname, datapolygon)
		if err != nil {
			return ReturnStruct(CreateResponse(true, "Success: Polygon created", datapolygon))
		} else {
			return ReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return ReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}
}

func AmbilDataGeojsonHeader(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)

	if r.Header.Get("token") == os.Getenv("token") {
		datagedung := GetAllBangunanLineString(mconn, collname)
		err := json.NewDecoder(r.Body).Decode(&datagedung)
		if err != nil {
			return ReturnStruct(CreateResponse(true, "Success: GeoJson shown", datagedung))
		} else {
			return ReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return ReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}
}

func AmbilDataKegiatan(mongoenv, dbname, collname string) string {
	mconn := SetConnection(mongoenv, dbname)
	datakegiatan := GetAllKegiatan(mconn, collname)
	return ReturnStruct(datakegiatan)
}

func AmbilDataJadwal(mongoenv, dbname, collname string) string {
	mconn := SetConnection(mongoenv, dbname)
	datajadwal := GetAllJadwal(mconn, collname)
	return ReturnStruct(datajadwal)
}

//-----------------------test buat project, nanti hapus kalo udah

func TambahBerita(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var databerita Berita
	err := json.NewDecoder(r.Body).Decode(&databerita)
	if idBeritaExists(mongoenv, dbname, databerita) {
		response.Status = false
		response.Message = "ID telah ada"
	} else {
		response.Status = true
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			response.Status = true
			InsertBerita(mconn, collname, databerita.ID, databerita.Kategori, databerita.Judul, databerita.Preview, databerita.Konten)
			response.Message = "Berhasil Input data"
		}
	}
	return ReturnStruct(response)
}

func AmbilDataBerita(mongoenv, dbname, collname string) string {
	mconn := SetConnection(mongoenv, dbname)
	databerita := GetAllBerita(mconn, collname)
	return ReturnStruct(databerita)
}

func AmbilSatuBerita(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	var databerita Berita
	berita := FindBerita(mconn, collname, databerita)

	return ReturnStruct(berita)
}

// -------------coba decode

func CobaCobaAja(publickey string, r *http.Request) string {
	req := new(ResponseDataUser)
	req.Status = false

	// Read cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		req.Message = "error parsing cookie: " + err.Error()
	}

	checktoken := watoken.DecodeGetId(os.Getenv(publickey), cookie.Value)
	if checktoken == "" {
		req.Message = "hasil decode tidak ada"
	} else {
		req.Message = checktoken
	}

	return ReturnStruct(req)
}

func Authorization(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var req AuthorizationStruct
	req.Status = false

	mconn := SetConnection(mongoenv, dbname)

	var userdata User

	// Read cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		req.Message = "error parsing cookie: " + err.Error()
	}
	checktoken := watoken.DecodeGetId(os.Getenv(publickey), cookie.Value)

	userdata.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

	if checktoken == "" {
		req.Message = "hasil decode tidak ada"
	} else {
		req.Message = "hasil decode token"
		datauser := FindUser(mconn, collname, userdata)
		req.Status = true
		req.Data.Username = datauser.Username
		req.Data.Name = datauser.Name
		req.Data.Email = datauser.Email
		req.Data.Role = datauser.Role
	}

	return ReturnStruct(req)
}

func AuthorizationHeaders(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var req AuthorizationStruct
	req.Status = false

	mconn := SetConnection(mongoenv, dbname)

	var userdata User
	goblok := r.Header.Get("token")

	checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

	userdata.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

	if checktoken == "" {
		req.Message = "hasil decode tidak ada"
	} else {
		req.Message = "hasil decode token"
		datauser := FindUser(mconn, collname, userdata)
		req.Status = true
		req.Data.Username = datauser.Username
		req.Data.Name = datauser.Name
		req.Data.Email = datauser.Email
		req.Data.Role = datauser.Role
	}

	return ReturnStruct(req)
}
