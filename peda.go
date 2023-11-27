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

func MembuatGeojsonPoint(publickey, mongoenv, dbname, colluser, collgeojson string, r *http.Request) string {
	var response CredentialGeojson
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpoint GeoJsonPoint

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						PostPoint(mconn, collgeojson, geojsonpoint)
						response.Message = "data point berhasil masuk"
					}
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}

	err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
	if err != nil {
		return err.Error()
	}

	return ReturnStruct(response)
}

func MembuatGeojsonPolyline(publickey, mongoenv, dbname, colluser, collgeojson string, r *http.Request) string {
	var response CredentialGeojson
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonline GeoJsonLineString

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					err := json.NewDecoder(r.Body).Decode(&geojsonline)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						PostLinestring(mconn, collgeojson, geojsonline)
						response.Message = "data polyline berhasil masuk"
					}
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func MembuatGeojsonPolygon(publickey, mongoenv, dbname, colluser, collgeojson string, r *http.Request) string {
	var response CredentialGeojson
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpolygon GeoJsonPolygon

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						PostPolygon(mconn, collgeojson, geojsonpolygon)
						response.Message = "data polygon berhasil masuk"
					}
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func AmbilDataGeojson(publickey, mongoenv, dbname, colluser, collgeojson string, r *http.Request) string {
	var response CredentialGeojson
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					datagedung := GetAllBangunanLineString(mconn, collgeojson)
					return ReturnStruct(datagedung)
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func PostGeoIntersects(mongoenv, dbname string, r *http.Request) string {
	var longlat LongLat
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = GeoIntersects(mconn, longlat.Longitude, longlat.Latitude)
	}
	return ReturnStruct(response)
}

func PostGeoWithin(mongoenv, dbname string, r *http.Request) string {
	var coordinate GeometryPolygon
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = GeoWithin(mconn, coordinate.Coordinates)
	}
	return ReturnStruct(response)
}

// --------------------------------------------------------------------- Projek 3 ---------------------------------------------------------------------

func Authorization(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	response.Status = false

	mconn := SetConnection(mongoenv, dbname)

	var userdata User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		userdata.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, userdata) {
				response.Message = "berhasil decode token"
				datauser := FindUser(mconn, collname, userdata)
				response.Status = true
				response.Data.Username = datauser.Username
				response.Data.Name = datauser.Name
				response.Data.Email = datauser.Email
				response.Data.Role = datauser.Role
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func Registrasi(mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if usernameExists(mongoenv, dbname, datauser) {
		response.Message = "Username telah dipakai"
	} else {
		if err != nil {
			response.Message = "error parsing application/json: " + err.Error()
		} else {
			hash, hashErr := HashPassword(datauser.Password)
			if hashErr != nil {
				response.Message = "Gagal Hash Password" + err.Error()
			}
			InsertUserdata(mconn, collname, datauser.Name, datauser.Email, datauser.Username, hash, datauser.Role.Admin, datauser.Role.Author)
			response.Status = true
			response.Message = "Berhasil Input data"
		}
	}
	return ReturnStruct(response)
}

func Login(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		if usernameExists(mongoenv, dbname, datauser) {
			if IsPasswordValid(mconn, collname, datauser) {
				user := FindUser(mconn, collname, datauser)
				tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(privatekey))
				if err != nil {
					return ReturnStruct(response.Message == "gagal encode token :"+err.Error())
				} else {
					response.Status = true
					response.Data.Name = user.Name
					response.Data.Email = user.Email
					response.Data.Username = user.Username
					response.Data.Role = user.Role
					response.Message = "user berhasil login"
					response.Token = tokenstring
					return ReturnStruct(response)
				}
			} else {
				response.Message = "password salah"
			}
		} else {
			response.Message = "akun tidak ditemukan"
		}

	}
	return ReturnStruct(response)
}

func HapusUser(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var auth User
	var datauser User

	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, collname, auth)
				if auth2.Role.Admin == true {
					err := json.NewDecoder(r.Body).Decode(&datauser)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						if datauser.Username == "" {
							response.Message = "parameter dari function ini adalah username"
						} else {
							if usernameExists(mongoenv, dbname, datauser) {
								DeleteUser(mconn, collname, datauser)
								response.Status = true
								response.Message = "berhasil hapus " + datauser.Username + " dari database"
							} else {
								response.Message = "akun yang ingin dihapus tidak ditemukan"
							}
						}
					}
				} else {
					response.Message = "anda bukan admin jadi tidak diizinkan"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func UpdateUser(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var auth User
	var datauser User

	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, collname, auth)
				if auth2.Role.Admin == true {
					err := json.NewDecoder(r.Body).Decode(&datauser)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						if datauser.Username == "" {
							response.Message = "parameter dari function ini adalah username"
						} else {
							hash, hashErr := HashPassword(datauser.Password)
							if hashErr != nil {
								response.Message = "Gagal Hash Password" + err.Error()
							}
							if usernameExists(mongoenv, dbname, datauser) {
								EditUser(mconn, collname, datauser.Name, datauser.Email, datauser.Username, hash, datauser.Role.Admin, datauser.Role.Author, datauser.Role.User)
								response.Status = true
								response.Message = "berhasil update " + datauser.Username + " dari database"
							} else {
								response.Message = "akun yang ingin diedit tidak ditemukan"
							}
						}
					}
				} else {
					response.Message = "anda bukan admin jadi tidak diizinkan"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

// ---------------------------------------------------------------------- Berita

func TambahBerita(publickey, mongoenv, dbname, colluser, collberita string, r *http.Request) string {
	var response CredentialBerita
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var databerita Berita
	err := json.NewDecoder(r.Body).Decode(&databerita)

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.Author == true {
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						if idBeritaExists(mongoenv, dbname, databerita) {
							response.Message = "ID telah ada"
						} else {
							response.Status = true
							if err != nil {
								response.Message = "error parsing application/json: " + err.Error()
							} else {
								response.Status = true
								InsertBerita(mconn, collberita, databerita)
								response.Message = "berhasil Input data"
							}
						}
					}
				} else {
					response.Message = "anda bukan author ataupun admin jadi tidak diizinkan"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func AmbilDataBerita(publickey, mongoenv, dbname, colluser, collberita string, r *http.Request) string {
	var response CredentialBerita
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					databerita := GetAllBerita(mconn, collberita)
					return ReturnStruct(databerita)
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func AmbilSatuBerita(publickey, mongoenv, dbname, colluser, collberita string, r *http.Request) string {
	var response CredentialBerita
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	var databerita Berita

	var auth User
	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.User == true {
					idberita := r.URL.Query().Get("page")
					databerita.ID = idberita
					if idBeritaExists(mongoenv, dbname, databerita) {
						berita := FindBerita(mconn, collberita, databerita)
						return ReturnStruct(berita)
					} else {
						response.Message = "berita tidak ditemukan"
					}
				} else {
					response.Message = "akun anda tidak aktif"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func HapusBerita(publickey, mongoenv, dbname, colluser, collberita string, r *http.Request) string {
	var response CredentialBerita
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var auth User
	var databerita Berita

	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.Author == true || auth2.Role.Admin == true {
					err := json.NewDecoder(r.Body).Decode(&databerita)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						if databerita.ID == "" {
							response.Message = "parameter dari function ini adalah id"
						} else {
							if idBeritaExists(mongoenv, dbname, databerita) {
								DeleteBerita(mconn, collberita, databerita)
								response.Status = true
								response.Message = "berhasil hapus " + databerita.ID + " dari database"
							} else {
								response.Message = "berita tidak ditemukan"
							}
						}
					}
				} else {
					response.Message = "anda bukan author ataupun admin jadi tidak diizinkan"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

func UpdateBerita(publickey, mongoenv, dbname, colluser, collberita string, r *http.Request) string {
	var response CredentialBerita
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var auth User
	var databerita Berita

	goblok := r.Header.Get("token")

	if goblok == "" {
		response.Message = "header login tidak ditemukan"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), goblok)

		auth.Username = checktoken //userdata.Username dibuat menjadi checktoken agar userdata.Username dapat digunakan sebagai filter untuk menggunakan function FindUser

		if checktoken == "" {
			response.Message = "hasil decode tidak ditemukan"
		} else {
			if usernameExists(mongoenv, dbname, auth) {
				auth2 := FindUser(mconn, colluser, auth)
				if auth2.Role.Admin == true {
					err := json.NewDecoder(r.Body).Decode(&databerita)
					if err != nil {
						response.Message = "error parsing application/json: " + err.Error()
					} else {
						if databerita.ID == "" {
							response.Message = "parameter dari function ini adalah id"
						} else {
							if idBeritaExists(mongoenv, dbname, databerita) {
								EditBerita(mconn, collberita, databerita)
								response.Status = true
								response.Message = "berhasil update " + databerita.ID + " dari database"
							} else {
								response.Message = "berita tidak ditemukan"
							}
						}
					}
				} else {
					response.Message = "anda bukan admin jadi tidak diizinkan"
				}
			} else {
				response.Message = "akun tidak ditemukan"
			}
		}
	}
	return ReturnStruct(response)
}

// -------------------------------------------------------------------- Pemrograman --------------------------------------------------------------------

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
