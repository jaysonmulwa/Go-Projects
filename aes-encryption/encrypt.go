package main

import (

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

		text := []byte("Hiki ndio kitext")
		key := []byte("kipassword hivi")


		// generate a new aes cipher using our 32 byte long key
		c, err := aes.NewCipher(key)


				if err != nil {
        			fmt.Println(err)
    			}

		// gcm or Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers

		
		gcm, err := cipher.NewGCM(c)


    			if err != nil {
				        fmt.Println(err)
				    }

		// creates a new byte array the size of the nonce
    	// which must be passed to Seal
		nonce := make([]byte, gcm.Noncesize())


			// populates our nonce with a cryptographically secure
		    // random sequence
		    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		        fmt.Println(err)
		    }





		//fmt.Println(gcm.Seal(nonce, nonce, text, nil))

		// the WriteFile method returns an error if unsuccessful
			err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
			// handle this error
			if err != nil {
			  // print it out
			  fmt.Println(err)
			}






}