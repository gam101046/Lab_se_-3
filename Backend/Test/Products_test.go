package unit

import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"example.com/Lab_se_3/Entity"
)

func TestProduct(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Test Products", func(t *testing.T) {

		Products := Entity.Products{
			NameProduct:   "Me-o",
			Description:   "อาหารแมว",
			Price:          12,
			Stock:          12,
		}

		ok,err := govalidator.ValidateStruct(Products)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})



	t.Run("NameProduct is required", func(t *testing.T) {

		Products := Entity.Products{
			NameProduct:   "",
			Description:   "อาหารแมว",
			Price:          12,
			Stock:          12,
		}

		ok,err := govalidator.ValidateStruct(Products)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NameProduct is required"))

	})


	t.Run("Description is required", func(t *testing.T) {

		Products := Entity.Products{
			NameProduct:   "Me-o",
			Description:   "",
			Price:          12,
			Stock:          12,
		}

		ok,err := govalidator.ValidateStruct(Products)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NameProduct is required"))

	})


}