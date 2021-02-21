package emoji

import (
	"encoding/json"
	"strings"
	"sync"
)

type Emoji struct {
	Emoji     string   `json:"emoji"`
	SkinTones bool     `json:"skin_tones"`
	Aliases   []string `json:"aliases"`
}

type EmojiTrieNode struct {
	Emoji    *Emoji
	Children map[rune]*EmojiTrieNode
}

var (
	emojis     []*Emoji
	emojisOnce sync.Once

	emojisTrie     *EmojiTrieNode
	emojisTrieOnce sync.Once
)

func Emojis() []*Emoji {
	emojisOnce.Do(func() {
		err := json.Unmarshal(emojisJSON, &emojis)
		if err != nil {
			panic("failed to parse emojis JSON: " + err.Error())
		}
	})
	return emojis
}

func addEmojiToTrie(emoji *Emoji) {
	current := emojisTrie
	for _, r := range emoji.Emoji {
		child, ok := current.Children[r]
		if !ok {
			child = &EmojiTrieNode{
				Children: map[rune]*EmojiTrieNode{},
			}
			current.Children[r] = child
		}
		current = child
	}
	current.Emoji = emoji
}

func EmojisTrie() *EmojiTrieNode {
	emojisTrieOnce.Do(func() {
		emojisTrie = &EmojiTrieNode{
			Children: map[rune]*EmojiTrieNode{},
		}
		for _, emoji := range Emojis() {
			addEmojiToTrie(emoji)
		}
	})
	return emojisTrie
}

func EmojisToAliases(s string) string {
	emojisTrie := EmojisTrie()

	needsFixing := false

	for _, r := range s {
		if _, ok := emojisTrie.Children[r]; ok {
			needsFixing = true
		}
	}

	if !needsFixing {
		return s
	}

	runes := []rune(s)
	runesLen := len(runes)

	var b strings.Builder

	for i := 0; i < runesLen; i++ {
		r := runes[i]

		if _, ok := emojisTrie.Children[r]; ok {
			var lastEmoji *Emoji
			var lastEmojiJ int

			node := emojisTrie

			for j := 0; i+j < runesLen; j++ {
				r := runes[i+j]

				child, ok := node.Children[r]
				if !ok {
					break
				}

				if child.Emoji != nil {
					lastEmoji = child.Emoji
					lastEmojiJ = j
				}

				node = child
			}

			if lastEmoji != nil &&
				len(lastEmoji.Aliases) > 0 {
				b.WriteRune(':')
				b.WriteString(lastEmoji.Aliases[0])
				b.WriteRune(':')
				i += lastEmojiJ
				continue
			}
		}

		b.WriteRune(r)
	}

	return b.String()
}
