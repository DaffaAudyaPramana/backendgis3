package backendgis3

import (
	"fmt"
	"testing"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	datalokasi := GetAllBangunanLineString(mconn, "petapedia")
	fmt.Println(datalokasi)
}