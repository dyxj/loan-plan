package main_test

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {

	//-------- Initialize Test Environment -------
	//-------- End Initialize Test Environment -------

	// Run Test
	code := m.Run()
	//-------- Clean Up Test Environment --------
	//-------- End Clean Up Test Environment --------
	os.Exit(code)
}
