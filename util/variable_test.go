package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("variable util tests", func() {
	Describe("ResultPrefix tests", func() {
		It("Should return return for no create and depth = 1", func() {
			Expect(ResultPrefix("arg", 1, false)).To(Equal("return"))
		})

		It("Should return assignment for depth > 1", func() {
			Expect(ResultPrefix("arg", 2, false)).To(Equal("arg ="))
		})

		It("Should return declaration for create and depth = 1", func() {
			Expect(ResultPrefix("arg", 1, true)).To(Equal("arg :="))
		})
	})

	Describe("VariableName tests", func() {
		It("Should return a correct name", func() {
			Expect(VariableName("id_ip")).To(Equal("idIP"))
		})

		It("Should return unchanged variable name", func() {
			Expect(VariableName("abc")).To(Equal("abc"))
		})

		It("Should return variable name with prefix to avoid keywords", func() {
			Expect(VariableName("range")).To(Equal("rangeObject"))
		})
	})

	Describe("IndexVariable tests", func() {
		It("Should return a correct index variable name - i", func() {
			Expect(IndexVariable(1)).To(Equal('i'))
		})

		It("Should return a correct index variable name - j", func() {
			Expect(IndexVariable(2)).To(Equal('j'))
		})
	})

	Describe("Indent tests", func() {
		It("Should return an ident with a correct width", func() {
			Expect(Indent(2)).To(Equal("\t\t"))
		})
	})
})
