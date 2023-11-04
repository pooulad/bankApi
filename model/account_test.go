package model_test

import (
	"fmt"
	"testing"

	"github.com/pooulad/bankApi/model"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := model.NewAccount("a", "b", "pasdijosjif")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
