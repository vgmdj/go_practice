package Word_Ladder

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				beginWord: "qa",
				endWord:   "sq",
				wordList: []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb",
					"sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti",
					"ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho",
					"ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as",
					"fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga",
					"li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm",
					"an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am",
					"ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr",
					"sq", "ye"},
			},
			5,
		},
		{
			"case",
			args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			5,
		},
		{
			"case",
			args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "dog", "dot"},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
