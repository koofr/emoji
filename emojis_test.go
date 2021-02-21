package emoji_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/koofr/emoji"
)

var _ = Describe("Emojis", func() {
	Describe("Emojis", func() {
		It("should return all emojis", func() {
			emojis := Emojis()
			Expect(len(emojis)).To(Equal(1805))
		})
	})

	Describe("EmojisToAliases", func() {
		It("should convert emojis to aliases", func() {
			s := EmojisToAliases("test 👩‍❤️‍💋‍👨👩‍❤️‍💋‍👩😃👩‍❤️‍💋‍ test ❤️❤️💚💛")
			Expect(s).To(Equal("test :couplekiss_man_woman::couplekiss_woman_woman::smiley::woman:\u200d:heart:\u200d:kiss:\u200d test :heart::heart::green_heart::yellow_heart:"))

			Expect(EmojisToAliases("test ❤")).To(Equal("test ❤")) // black heart (old school, 3-byte emoji)
			Expect(EmojisToAliases("test ❤️")).To(Equal("test :heart:"))
			Expect(EmojisToAliases("test ⏱️")).To(Equal("test :stopwatch:"))
		})
	})
})
