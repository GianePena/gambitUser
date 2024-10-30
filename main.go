package main

import (
	"context" //paquete propioi de go
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gambitUser/awsgo"
	"github.com/gambitUser/bd"
	"github.com/gambitUser/models"
	"os"
)

func main() {
	lambda.Start(EjecutoLambda)
}
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() { // DEVUELVE UN FALSE
		fmt.Println("Error en los parametros, debe enviar secretName")
		err := errors.New("error en los parametros debe enviar secretName")
		return event, err
	}
	var datos models.SingUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("EMAIL= ", datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("SUB= ", datos.UserUUID)
		}
	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("ERROR AL LEER EL SECRET" + err.Error())
		return event, err
	}
	err = bd.SingUp(datos) //pasasr los datos de la var datos
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName") //que se fueje en la variable de entorno la var secretname, si NO la encuentra devuelve un false
	return traeParametro
}
