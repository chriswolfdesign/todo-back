package todo_back_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTodoBack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TodoBack Suite")
}
