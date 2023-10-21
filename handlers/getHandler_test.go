package handlers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetHandler", func() {
	When("test is meant to fail", func() {
		It("fails", func() {
			Expect(true).To(BeFalse())
		})
	})
})
