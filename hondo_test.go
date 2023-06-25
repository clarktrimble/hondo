package hondo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/clarktrimble/hondo"
)

func TestHondo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hondo Suite")
}

var _ = Describe("Hondo", func() {

	var (
		length int = 9
		id     string
	)

	JustBeforeEach(func() {
		id = Rand(length)
	})

	Describe("generating a random string", func() {
		When("all is well", func() {
			It("returns a string of requested length", func() {
				Expect(id).To(HaveLen(length))
			})
		})
	})

	// Todo: check more, perhaps with fuzz?

})
