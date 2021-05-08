package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)


func MakeProtoGetRequest(requestUrl string) []byte {
	client := &http.Client{}
	resp, err := client.Get(requestUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data

}

func MakeProtoPostRequest(requestUrl string, buffData []byte) {
	client := &http.Client{}
	_, err := client.Post(requestUrl, "", bytes.NewBuffer(buffData))
	if err != nil {
		log.Fatalln(err)
	}

}

func getProtoUsers() {
	data := MakeProtoGetRequest("http://127.0.0.1:5000/proto")

	users := &UsersList{}
	if err := proto.Unmarshal(data, users); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	log.Println("!!!Grabbing whole users list!!!!!")
	log.Println(users.Users)
}

func getProtoUser() {
	data := MakeProtoGetRequest("http://127.0.0.1:5000/proto-solo")

	user := &User{}
	if err := proto.Unmarshal(data, user); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	log.Println("!!!Grabbing solo user info!!!!!")
	log.Println("ID:", user.Id, "NAME:", user.Name, "PHONE:", user.Phone)
}

func addProtoUser(id int) {
	user := &User{}
	user.Id = int32(id)
	user.Name = "Go Client testing"
	user.Phone = "9823723923923"

	buffOut, err := proto.Marshal(user)
	if err != nil {
		log.Fatalln("Failed to encode user info:", err)
	}
	MakeProtoPostRequest("http://localhost:5000/proto_create", buffOut)
	log.Println("!!!Added new solo user info!!!!!")
}

func main() {

	for i:=0; i < 100; i++ {
		addProtoUser(i)
	}

	log.Print("\n\n")
	log.Println("================================")
	log.Println("================================")
	log.Print("\n\n")

	getProtoUsers()

	log.Print("\n\n")
	log.Println("================================")
	log.Println("================================")
	log.Print("\n\n")

	getProtoUser()

	log.Print("\n\n")

}

