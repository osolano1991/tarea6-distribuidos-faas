package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type DoctorRef struct {
	DoctorId int    `json:"doctor_id"`
	Nombre  string `json:"nombre"`
}

type Paciente struct {
	Id              int       `json:"_id"`
	Paciente        string    `json:"paciente"`
	Edad            string    `json:"edad"`
	Enfermedad      string    `json:"enfermedad"`
	Nacionalidad    string    `json:"nacionalidad"`
	Tratamientos    string    `json:"tratamientos"`
	Doctores        []DoctorRef `json:"doctores"`
}

var items []Paciente

var jsonData string = `[
	{
		"_id": 1,
		"paciente": "Luis Solano Mora",
		"edad": "26",
		"enfermedad": "Hipertension",
		"nacionalidad": "Estados Unidos",
		"tratamientos": "Pastillas",
		"doctores": [
			{
				"doctor_id": 1,
				"nombre": "Oscar Solano Mora"
			},
			{
				"doctor_id": 2,
				"nombre": "Tatiana Chaves Salazar"
			}
		]
	},
	{
		"_id": 2,
		"paciente": "Floribeth Mora JIménez",
		"edad": "52",
		"enfermedad": "Miopia",
		"nacionalidad": "Inglaterra",
		"tratamientos": "Lentes",
		"doctores": [
			{
				"doctor_id": 3,
				"nombre": "Angie Segura Solano"
			},
			{
				"doctor_id": 4,
				"nombre": "Ronald Solano López"
			}
		]
	},
	{
		"_id": 3,
		"paciente": "Nadia Solano Mora",
		"edad": "27",
		"enfermedad": "Colesterol",
		"nacionalidad": "España",
		"tratamientos": "Bajar de peso",
		"doctores": [
			{
				"doctor_id": 2,
				"nombre": "Tatiana Chaves Salazar"
			},
			{
				"doctor_id": 3,
				"nombre": "Angie Segura Solano"
			}
		]
	},
	{
		"_id": 4,
		"paciente": "Jouser Thomas Acuña",
		"edad": "15",
		"enfermedad": "Diabetes",
		"nacionalidad": "Francia",
		"tratamientos": "Inyecciones",
		"doctores": [
			{
				"doctor_id": 1,
				"nombre": "Oscar Solano Mora"
			},
			{
				"doctor_id": 4,
				"nombre": "Ronald Solano López"
			}
		]
	}
]`

func FindItem(id int) *Paciente {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
