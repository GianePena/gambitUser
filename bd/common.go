package bd

import (
	"database/sql"
	"fmt"
	"github.com/gambitUser/models"
	"github.com/gambitUser/secretm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// funciones que se relaiona con la base de datos
var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB //todo lo que tenga que ver con conexion a base de datos se maneja con punteros

func ReadSecret() error {
	var err error
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// conexion a base de datos
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping() //es un segundo control de conexion --> s devuelve error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("conexxion exitosa de base de datos")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "database-1"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	return dsn
}
