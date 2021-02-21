package emoji_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEmoji(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Emoji Suite")
}
