package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type Contact struct{
	FullName string `json:"fullname"`
	Tel string `json:"tel"`
}

/*
 Using regex to extract a phone number from a current
 current line while vcard is being read
*/
func ExtractPhoneNumber(str string) string{
	// regex to match the contact number only
	reg, err  := regexp.Compile(`:.*\d+$`)
	if err != nil{
		panic(err)
	}

	return reg.FindString(str)[1:]
}

/*
 Writes the contact list (arr of Contact structs)
 to a json file and creates the json file at <writePath>
*/
func WriteContactsToJson(contacts []Contact, writePath string)error{
  // convert contacts arr to json
	json, err := json.MarshalIndent(contacts, " "," ",)
  if err != nil{
		panic(err)
	}

	// open file at writePath, if path already exists, it will be overridden
	file, err := os.Create(writePath)
	if err != nil{
		panic(err)
	}


	writer := bufio.NewWriter(file)
	// write to json file
	_, err = writer.Write(json)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	return nil
}

/*
 Reads the vcard file data from readPath into the contacts array
*/
func ReadVcard(contacts *[]Contact, readPath string)error {
  file, err := os.Open(readPath)

	if err != nil{
		return err
	}

	defer file.Close()

	// helps to format a contact
	var tempContact Contact

	scanner := bufio.NewScanner(file)

	vcardValidated := false

	// read file line by line
	for scanner.Scan(){
		text := scanner.Text()

		// runs once to first validated the vc file
		if !vcardValidated{
			if text != "BEGIN:VCARD"{
				// Invalid file return error
				return fmt.Errorf("Invalid Vcard File")
			}
			// Valid file continue to read file
			vcardValidated = true
			continue
		}

		if text == "END:VCARD"{
			// end of a contact info, save
			*contacts = append(*contacts, tempContact)
			// clear the prev contact info
			tempContact = Contact{}
		}else if len(text) > 3 && text[:3] == "FN:"{
			// get fullname
			tempContact.FullName = text[3:]	
		}else if len(text) > 3 && text[:3] == "TEL"{
			// get phone number
			tempContact.Tel = ExtractPhoneNumber(text)
		}
	}

	return scanner.Err();
}
