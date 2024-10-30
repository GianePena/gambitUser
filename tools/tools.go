package tools

import (
	"fmt"
	"time"
)

// en mysql los datos se guarda con una formato defecha determinada
func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	//%02d indic que son 2 digitos
}
