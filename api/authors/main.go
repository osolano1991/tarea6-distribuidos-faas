package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type DoctorRef struct {
	DoctorId int    `json:"doctor_id"`
	Title  string `json:"nombre"`
}

type Cita struct {
	Id            int       `json:"_id"`
	Cita          string    `json:"cita"`
	Dia           string    `json:"dia"`
	Hora          int       `json:"hora"`
	Ubicacion     string    `json:"ubicacion"`
	Especialidad  string    `json:"especialidad"`
	Doctores       []DoctorRef `json:"doctores"`
}

var items []Cita

var jsonData string = `[
	{
		"_id": 1,
		"cita": "CITA 1",
		"dia": "Lunes",
		"hora": "04:00 p.m",
		"ubicacion": "Heredia",
		"especialidad": "Medicina General",
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
		"cita": "CITA 2",
		"dia": "Martes",
		"hora": "10:00 a.m",
		"ubicacion": "San José",
		"especialidad": "Oftalmología",
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
		"cita": "CITA 3",
		"dia": "Miércoles",
		"hora": "08:00 a.m",
		"ubicacion": "Puntarenas",
		"especialidad": "Neurología",
		"doctores": [
			{
				"doctor_id": 7,
				"nombre": "Keilyn Chaves Salazar"
			},
			{
				"doctor_id": 6,
				"nombre": "Arcelio Chaves Villarevia"
			}
			,
			{
				"doctor_id": 1,
				"nombre": "Oscar Solano Mora"
			}
		]
	},
	{
		"_id": 4,
		"cita": "CITA 4",
		"dia": "Jueves",
		"hora": "12:00 p.m",
		"ubicacion": "Alajuela",
		"especialidad": "Pediatría",
		"doctores": [
			{
				"doctor_id": 1,
				"nombre": "Oscar Solano Mora"
			},
			{
				"doctor_id": 5,
				"nombre": "Lorena Salazar Mora"
			}
			,
			{
				"doctor_id": 2,
				"nombre": "Tatiana Chaves Salazar"
			}
		]
	}
]`

func FindItem(id int) *Cita {
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
