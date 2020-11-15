package main

import (
	"net/http"

	AdvertisementAdd "../Modules/Advertisement/Add"
	AdvertisementDelete "../Modules/Advertisement/Delete"
	AdvertisementGet "../Modules/Advertisement/Get"
	AdvertisementUpdate "../Modules/Advertisement/Update"
	Login "../Modules/Auth/Login"
	NewPassword "../Modules/Auth/NewPassword"
	Reset "../Modules/Auth/ResetPassword"
	Signup "../Modules/Auth/SignUp"
	Comment "../Modules/Comment"
	PhotoUpload "../Modules/PhotoUploader"

	fb "../Modules/Firebase"
)

func main() {
	go fb.ConnectFirebase()
	createServer()
}

func createServer() {
	go http.HandleFunc("/signup", Signup.HandleSignUp)
	go http.HandleFunc("/login", Login.HandleLogin)
	go http.HandleFunc("/resetPassword", Reset.ResetPasswordHandler)
	go http.HandleFunc("/newPassword", NewPassword.NewPassword)
	go http.HandleFunc("/advertisement/add", AdvertisementAdd.AdvertisementAddHandler)
	go http.HandleFunc("/advertisement/get", AdvertisementGet.AdvertisementGetHandler)
	go http.HandleFunc("/advertisement/get/mine", AdvertisementGet.GetSelfAdvertisementHandler)
	go http.HandleFunc("/advertisement/update", AdvertisementUpdate.AdvertisementUpdateHandler)
	go http.HandleFunc("/advertisement/delete", AdvertisementDelete.AdvertisementDeleteHandler)
	go http.HandleFunc("/upload-photo", PhotoUpload.HandleUpload)
	go http.HandleFunc("/comment", Comment.CommentHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err)
	}
}
