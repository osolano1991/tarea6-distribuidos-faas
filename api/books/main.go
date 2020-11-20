package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type Doctor struct {
	Id            int    `json:"_id"`
	Nombre        string `json:"nombre"`
	Edad          string `json:"edad"`
	Nacionalidad  int    `json:"nacionalidad"`
	Especialidad  string `json:"especialidad"`
	CentroMedico  int    `json:"centromedico"`
	Cita          string `json:"cita"`
	Cita_Id       int    `json:"cita_id"`
	Paciente      string `json:"paciente"`
	Paciente_Id   int    `json:"paciente_id"`
}

var doctores []Doctor

var jsonData string = `[
	{
		"_id": 1,
		"nombre": "Oscar Solano Mora",
		"edad": "29",
		"nacionalidad": "Costa Rica",
		"especialidad": "Medicina General",
		"centromedico": "San José",
		"cita": "CITA 1",
		"cita_id": 1,
		"paciente": "Luis Solano Mora",
		"paciente_id": 1
	},
	{
		"_id": 2,
		"nombre": "Tatiana Chaves Salazar",
		"edad": "30",
		"nacionalidad": "Costa Rica",
		"especialidad": "Odontologia",
		"centromedico": "Cartago",
		"cita": "CITA 1",
		"cita_id": 1,
		"paciente": "Luis Solano Mora",
		"paciente_id": 1
	},
	{
		"_id": 3,
		"nombre": "Angie Segura Solano",
		"edad": "18",
		"nacionalidad": "Costa Rica",
		"especialidad": "Neurologa",
		"centromedico": "Heredia",
		"cita": "CITA 2",
		"cita_id": 2,
		"paciente": "Floribeth Mora JIménez",
		"paciente_id": 2
	},
	{
		"_id": 4,
		"nombre": "Ronald Solano López",
		"edad": "50",
		"nacionalidad": "Mexico",
		"especialidad": "Oftalmologia",
		"centromedico": "San Jose",
		"cita": "CITA 2",
		"cita_id": 2,
		"paciente": "Floribeth Mora JIménez",
		"paciente_id": 2
	}
]`

func FindDoctor(id int) *Doctor {
	for _, doctor := range doctores {
		if doctor.Id == id {
			return &doctor
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(doctores)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			doctor := FindDoctor(param)
			if doctor != nil {
				data, _ = json.Marshal(*doctor)
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
	_ = json.Unmarshal([]byte(jsonData), &doctores)
	lambda.Start(handler)
}
