package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ScoringController (c *gin.Context) {

	var controller Controller

	if err := c.ShouldBindJSON(&controller); err == nil {
		//Write data to reddis for memcache

		//Set a body
		//body_sicd := `{"Name":"` + controller.Name + `", "Personal_number":"` + controller.Personal_number + `", "Birth_date":"` + controller.Birth_date + `", "Branch_code":"` + controller.Branch_code + `"}`
		//body_dhn := `{"Name":"` + controller.Name + `", "Birth_date":"` + controller.Birth_date + `"}`
		//body_kemendagri := `{"NIK":"` + controller.NIK + `"}`
		
		ProduceKafka(controller, KAFKA_PUB_URL_PRESCREENING_DATA) // Produce data to Kafka topic
		c.JSON(http.StatusBadRequest, gin.H{"status": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}
}
/*	//Do concuren function
	var sicd Sicd
	var url string
	if err := c.ShouldBindJSON(&sicd); err == nil {
		//url = "http://api.briconnect.bri.co.id/sid/sicd/" + sicd.Name + "/" + sicd.Birth_date + "/" + sicd.Personal_number + "/" + sicd.Branch_code
		url = "http://172.18.41.129/sid/sicd/" + sicd.Name + "/" + sicd.Birth_date + "/" + sicd.Personal_number + "/" + sicd.Branch_code
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	READ_WS_SICD: resp, err := resty.R().
		SetHeader("Authorization", auth).
		Get(url)
	if err==nil {
		if resp.Status() == "200 OK" {	//Auth token is valid
			c.String(http.StatusOK, resp.String()) 	//Send SICD data as response in JSON format
		} else {
			var resp_body map[string]interface{}
			json.Unmarshal(resp.Body(), &resp_body)
			if resp_body["error"] == "invalid_token" || resp_body["error"] == "missing_token" {
				//fmt.Println("Token "+ auth +" is not valid") // Auth token is not valid or expired 
				auth,_ = RequestToken()
				goto READ_WS_SICD
			} else {	//Other error
				//fmt.Printf("Error: %v\n",resp_body["message"])
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url or JSON format"})
			}
		}
	} else {
		// Request Error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to connect to prescreening service"})
	}*/
//}