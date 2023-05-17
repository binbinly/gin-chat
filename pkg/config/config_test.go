package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	type config struct {
		Name     string
		Addr     string
		Username string
		Password string
	}
	var dbConf config

	c := New("../test/config/")
	if err := c.Load("database", &dbConf); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, dbConf.Name, "mysql")
	assert.Equal(t, dbConf.Addr, "localhost:3306")
	assert.Equal(t, dbConf.Username, "root")
	assert.Equal(t, dbConf.Password, "123456")
}
