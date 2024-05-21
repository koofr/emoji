package emoji_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/koofr/emoji"
)

var _ = Describe("Strings", func() {
	Describe("StringToAliases", func() {
		It("should keep normal text", func() {
			Expect(StringToAliases("text")).To(Equal("text"))
		})

		It("should convert 4-byte UTF-8 to aliases", func() {
			Expect(StringToAliases("𩸽")).To(Equal(":u29e3d:"))
		})

		It("should convert emojis and 4-byte UTF-8 to aliases", func() {
			s := StringToAliases("test 👩‍❤️‍💋‍👨👩‍❤️‍💋‍👩😃👩‍❤️‍💋‍ 𩸽 test ❤️❤️💚💛")
			Expect(s).To(Equal("test :couplekiss_man_woman::couplekiss_woman_woman::smiley::woman:\u200d:heart:\u200d:kiss:\u200d :u29e3d: test :heart::heart::green_heart::yellow_heart:"))
		})
	})
})

func BenchmarkStringToAliasesPlain(b *testing.B) {
	StringToAliases("")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringToAliases("test")
	}
}

func BenchmarkStringToAliasesUnicode(b *testing.B) {
	StringToAliases("")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringToAliases("test 👩‍❤️‍💋‍👨👩‍❤️‍💋‍👩😃👩‍❤️‍💋‍ 𩸽 test ❤️❤️💚💛")
	}
}
