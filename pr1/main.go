package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main(){
	elliot:=&Person{
		Name: "Hatsker",
		Age: 18,
		SocialFollowers: &SocialFollowers{
			Youtobe: 120,
			Facebook: 130,
		},
		
	}
	data,err:=proto.Marshal(elliot)
	if err!=nil{
		fmt.Println("error while Marshaling",err)
	}
	fmt.Println(data)

	newElliot:=&Person{}
	err = proto.Unmarshal(data,newElliot)
	if err!=nil{
		fmt.Println("error while unmarshiling: ",err)
	}


	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetYoutobe())
	fmt.Println("facebook: ",newElliot.SocialFollowers.GetFacebook())

}