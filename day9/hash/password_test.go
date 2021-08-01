package hash_test

import (
	"testing"

	"gitee.com/infraboard/go-course/day9/hash"
	"github.com/stretchr/testify/assert"
)

func TestPasswordCheck(t *testing.T) {
	should := assert.New(t)

	pass, err := hash.NewHashedPassword("123456")
	should.NoError(err)
	t.Log(pass.Password)
	should.Error(pass.CheckPassword("sdfsdf"))
	should.NoError(pass.CheckPassword("123456"))

	new, err := hash.NewHashedPassword("5678")
	should.NoError(err)
	pass.Update(new, 2, false)
	t.Log(pass.CheckPassword("123456"))
	t.Log(pass.CheckPassword("5678"))
}
