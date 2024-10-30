package bd

import (
	"fmt"
	"github.com/gambitUser/models"
	"github.com/gambitUser/tools"
	_ "github.com/go-sql-driver/mysql"
)

// grabar en base de datos
func SingUp(sig models.SingUp) error {
	fmt.Println("COMIENZA REGISTRO")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close() //cerrar la conexiona la base de datos

	sentencia := "INSERT INTO users (User_Email, User_UIDD, User_DateAdd) VALUES('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia) //exce es una funcion de amazons
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
