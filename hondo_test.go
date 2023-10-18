package hondo_test

import (
	"math/rand"
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
				Expect(id).ToNot(Equal("1LGIehp1s"))
			})
		})

		When("seed is set to a known value", func() {
			BeforeEach(func() {
				rand.Seed(1) //nolint:staticcheck // ok for simple unit
			})

			It("returns a predictable string", func() {
				Expect(id).To(Equal("1LGIehp1s"))
			})
		})
	})

})
