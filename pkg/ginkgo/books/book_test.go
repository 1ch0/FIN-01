package books_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
		json string
	)

	BeforeEach(func() {
		json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
	})

	JustBeforeEach(func() {
		book, err = NewBookFromJSON(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				json = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`
			})

			It("should return the zero-value for the books", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})
