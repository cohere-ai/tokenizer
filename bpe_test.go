package tokenizer

import (
	"testing"
)

func TestBPESuite(t *testing.T) {
	t.Run("BPE", testBPE)
	t.Run("e2e", testBPEEncode)
}

var loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque sed viverra nunc. Fusce congue luctus ipsum eget tincidunt. Vivamus non consectetur nisi, nec aliquam tortor. Vestibulum scelerisque placerat nisi at hendrerit. Nulla venenatis pharetra auctor. Mauris ac velit pharetra urna pulvinar consequat. Maecenas sodales tempus erat, feugiat maximus purus condimentum sit amet.
Mauris cursus neque at fringilla lobortis. Maecenas eget nisl a felis eleifend pretium. Nulla at massa ligula. Aliquam pulvinar semper cursus. Cras ut tellus vel purus fringilla lacinia quis vel diam. Etiam hendrerit elit vitae mauris mattis egestas. Aliquam ut felis non orci consequat efficitur blandit ut felis. Suspendisse id pharetra mi.
Sed congue varius libero ac euismod. Sed mattis velit lorem, a rutrum est finibus vel. Pellentesque vel elit augue. Nulla et turpis eros. Suspendisse sit amet hendrerit urna. Suspendisse at enim dictum, condimentum leo at, vehicula libero. Mauris a nisl ac leo condimentum imperdiet. Duis tempor, lorem ultrices semper hendrerit, urna lectus convallis lacus, et vulputate ex nibh sed lectus. Nulla placerat ipsum ut vulputate cursus. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Quisque faucibus nibh felis, vitae dignissim est auctor in. Donec laoreet ultrices diam. Maecenas eu viverra est. Aliquam erat volutpat. Aenean pulvinar purus in sem fermentum dapibus.
Proin quis enim sed tellus varius scelerisque eget in dolor. Nullam auctor tellus vitae tortor ullamcorper, ut tempus orci semper. In bibendum ultrices imperdiet. Nullam eleifend dolor ac facilisis cursus. Cras id ligula non sapien elementum laoreet vitae tempor leo. Nulla facilisi. Quisque interdum augue sed dui consequat, et sodales urna accumsan. Duis sit amet congue est, ac ullamcorper urna. Nunc ultricies quis enim eu faucibus. Sed volutpat dapibus nulla, quis rhoncus risus ultrices dignissim. Donec a neque vel nunc ultricies scelerisque id eget orci. Curabitur venenatis ullamcorper lacus ac cursus. Nam sit amet nulla sed mi consectetur pellentesque nec faucibus diam. Proin et magna consequat, suscipit erat sed, efficitur nunc.
Aenean et convallis lacus, non facilisis metus. Sed felis dui, pulvinar eget risus at, semper feugiat dui. Curabitur tempus feugiat ex, nec lobortis turpis aliquet eu. Maecenas imperdiet et sem in scelerisque. Sed aliquam a lacus sed congue. Donec semper nulla porta odio ultrices, nec elementum ipsum aliquam. Aliquam quis diam imperdiet ligula porttitor euismod id ut enim. Sed et est aliquet, laoreet mi quis, molestie lorem. In hac habitasse platea dictumst. Maecenas nec metus porttitor, interdum metus a, ultrices mi. Praesent aliquet ante ipsum, vel eleifend ex venenatis vel. Suspendisse faucibus euismod augue a aliquam.`

func testBPEEncode(t *testing.T) {
	texts := []string{
		"是不",
		"à une opinion répandue wypełn",
		"😁 I'm happy!",
		"a b c d e f g h i j k l 		m n o p q r s t u v w x y z",
		"🐋🐳",
		"起来Qǐlái!！ 不愿Búyuàn做zuò奴隶núlì的de人们rénmen!",
		"Arise, ye who refuse to be slaves!",
	}

	tests := []struct {
		vocabCorpus    string
		expectedTokens [][]int64
	}{
		{
			vocabCorpus: "",
			expectedTokens: [][]int64{
				{
					231, 153, 176, 229, 185, 142,
				},
				{
					196, 161, 33, 118, 111, 102, 33, 112, 113, 106, 111, 106, 112, 111, 33, 115, 196, 170, 113, 98, 111, 101, 118, 102, 33, 120, 122, 113, 102, 198, 131, 111,
				},
				{
					241, 160, 153, 130, 33, 74, 40, 110, 33, 105, 98, 113, 113, 122, 34,
				},
				{
					98, 33, 99, 33, 100, 33, 101, 33, 102, 33, 103, 33, 104, 33, 105, 33, 106, 33, 107, 33, 108, 33, 109, 33, 10, 10, 110, 33, 111, 33, 112, 33, 113, 33, 114, 33, 115, 33, 116, 33, 117, 33, 118, 33, 119, 33, 120, 33, 121, 33, 122, 33, 123,
				},
				{
					241, 160, 145, 140, 241, 160, 145, 180,
				},
				{
					233, 182, 184, 231, 158, 166, 82, 200, 145, 109, 196, 162, 106, 34, 240, 189, 130, 33, 229, 185, 142, 231, 133, 192, 67, 196, 187, 122, 118, 196, 161, 111, 230, 130, 155, 123, 118, 196, 179, 230, 166, 181, 234, 155, 183, 111, 196, 187, 109, 196, 173, 232, 155, 133, 101, 102, 229, 187, 187, 229, 188, 173, 115, 196, 170, 111, 110, 102, 111, 34,
				},
				{
					66, 115, 106, 116, 102, 45, 33, 122, 102, 33, 120, 105, 112, 33, 115, 102, 103, 118, 116, 102, 33, 117, 112, 33, 99, 102, 33, 116, 109, 98, 119, 102, 116, 34,
				},
			},
		},
		{
			vocabCorpus: loremIpsum,
			expectedTokens: [][]int64{
				{
					231, 153, 176, 229, 185, 142,
				},
				{
					196, 161, 33, 364, 102, 33, 112, 113, 394, 106, 272, 417, 196, 170, 113, 439, 101, 118, 102, 33, 120, 122, 113, 102, 198, 131, 111,
				},
				{
					241, 160, 153, 130, 33, 74, 40, 110, 340, 440, 113, 122, 34,
				},
				{
					98, 586, 274, 314, 259, 279, 33, 104, 340, 278, 33, 107, 33, 108, 296, 33, 10, 10, 110, 275, 33, 112, 280, 33, 114, 417, 270, 293, 33, 118, 271, 33, 120, 33, 121, 33, 122, 33, 123,
				},
				{
					241, 160, 145, 140, 241, 160, 145, 180,
				},
				{
					233, 182, 184, 231, 158, 166, 82, 200, 145, 109, 196, 162, 106, 34, 240, 189, 130, 33, 229, 185, 142, 231, 133, 192, 67, 196, 187, 122, 118, 196, 161, 111, 230, 130, 155, 123, 118, 196, 179, 230, 166, 181, 234, 155, 183, 111, 196, 187, 109, 196, 173, 232, 155, 133, 101, 102, 229, 187, 187, 229, 188, 173, 115, 196, 170, 111, 110, 261, 34,
				},
				{
					66, 115, 257, 102, 45, 33, 122, 102, 33, 120, 105, 112, 417, 102, 103, 258, 102, 293, 112, 586, 102, 270, 269, 119, 284, 34,
				},
			},
		},
		{
			vocabCorpus: bigText,
			expectedTokens: [][]int64{
				{
					3594,
				},
				{
					286, 821, 2560, 320, 106, 315, 378, 1214, 571, 979, 290, 1105, 102, 296, 111,
				},
				{
					241, 160, 153, 130, 483, 40, 110, 3801, 34,
				},
				{
					98, 309, 301, 268, 389, 336, 361, 330, 525, 347, 422, 302, 33, 10, 10, 110, 325, 386, 298, 466, 378, 269, 264, 478, 505, 290, 486, 355, 335,
				},
				{
					241, 160, 145, 140, 241, 160, 145, 180,
				},
				{
					2456, 82, 368, 1333, 34, 1166, 33, 1263, 2412, 192, 67, 632, 122, 118, 419, 3529, 123, 1030, 461, 181, 234, 155, 183, 111, 632, 109, 367, 397, 1324, 944, 1860, 1510, 1498, 34,
				},
				{
					66, 115, 4880, 45, 1657, 10190, 892, 103, 2932, 404, 620, 269, 916, 119, 339, 34,
				},
			},
		},
	}

	for _, test := range tests {
		encoderMap, merges, err := BPE(CountString(test.vocabCorpus), 10000, 1)
		if err != nil {
			t.Fatal(err)
		}

		bpeMerges := [][2]string{}
		for _, v := range merges {
			bpeMerges = append(bpeMerges, v.Merge)
		}

		encoder, err := New(encoderMap, bpeMerges)
		if err != nil {
			t.Fatal(err)
		}

		for i, text := range texts {
			tokens, _ := encoder.Encode(text)
			if len(test.expectedTokens[i]) != len(tokens) {
				t.Fatalf("expected %d tokens but only got %d for \"%s\"", len(test.expectedTokens[i]), len(tokens), text)
			}

			for j := range test.expectedTokens[i] {
				if tokens[j] != test.expectedTokens[i][j] {
					t.Fatalf("failed to validate tokens for text %d, expected %v but got %v", i, test.expectedTokens[i], tokens)
				}
			}

			decoded := encoder.Decode(tokens)
			if decoded != text {
				t.Fatalf("expected %v to decode to %s, but got %s", tokens, text, decoded)
			}
		}
	}
}

func testBPE(t *testing.T) {
	tests := []struct {
		frequencies     map[string]int64
		numSymbols      int64
		minFrequency    int64
		expectedEncoder map[string]int64
		expectedMerges  []Merge
	}{
		{
			frequencies: map[string]int64{
				"aaaa": 30,
				"abcd": 30,
				"abab": 30,
				"aaab": 30,
			},
			numSymbols:   30,
			minFrequency: 1,
			expectedEncoder: map[string]int64{
				"Ë": 204, "A": 66, "Ï": 208, "Ĝ": 29, "ĝ": 30, "æ": 231, "=": 62, "aaaa": 259, "d": 101, "ò": 243, "§": 168, "/": 48, "±": 178, "P": 81, "Å": 198, "#": 36, "Į": 141, "z": 123, "È": 201, "g": 104, "abcd": 263, "v": 119, "¡": 162, "p": 113, "M": 78, "<": 61, "ė": 24, "ġ": 128, "¨": 169, "×": 216, "¾": 191, "ê": 235, ";": 60, "ą": 6, "Ċ": 11, "õ": 246, "i": 106, "Þ": 223, "+": 44, "h": 105, "Ó": 212, "l": 109, "Ď": 15, "¼": 189, "ĭ": 140, "í": 238, "ù": 250, "u": 118, "ĸ": 151, "t": 117, "¢": 163, "aa": 257, "Ö": 215, "ë": 236, "D": 69, "©": 170, "?": 64, "Ł": 160, "Ģ": 129, "á": 226, "ab": 258, "abc": 262, "Ă": 3, "S": 84, "Ī": 137, "Ú": 219, "Ĕ": 21, "¯": 176, "ú": 251, "ô": 245, "ă": 4, "Ą": 5, "Ć": 7, "ē": 20, "Č": 13, "o": 112, "ĉ": 10, "î": 239, "ć": 8, "Z": 91, "ç": 232, "k": 108, "½": 190, "ľ": 157, "Ã": 196, "â": 227, "º": 187, "Ü": 221, "ě": 28, "¦": 167, "ø": 249, "Ķ": 149, "Ø": 217, "ĳ": 146, "ö": 247, "q": 114, "$": 37, "8": 57, "¿": 192, "ð": 241, "Ñ": 210, "ì": 237, "å": 230, "·": 184, "K": 76, "³": 180, "²": 179, "ā": 2, "ļ": 155, "ģ": 130, "^": 95, "G": 72, "Ĉ": 9, ")": 42, "ğ": 32, "4": 53, "Ò": 211, "Ä": 197, "É": 202, "_": 96, "Ğ": 31, "3": 52, "»": 188, "Ļ": 154, "Ĵ": 147, "ł": 161, "*": 43, "U": 86, "Ľ": 156, "č": 14, "W": 88, "%": 38, ">": 63, "÷": 248, "ď": 16, "Ù": 218, "û": 252, "e": 102, "Ð": 209, "ī": 138, "Ā": 1, "!": 34, "x": 121, "ü": 253, "Ŀ": 158, "ï": 240, "r": 115, "B": 67, "s": 116, "Ġ": 33, "}": 126, "ý": 254, "c": 100, "ª": 171, "ċ": 12, "£": 164, "2": 51, "ä": 229, "R": 83, "į": 142, "{": 124, "Æ": 199, "&": 39, "Ĭ": 139, "Q": 82, "Đ": 17, "¥": 166, "]": 94, "`": 97, "Ç": 200, "n": 111, "b": 99, "Ĺ": 152, "ķ": 150, "ĩ": 136, "¶": 183, "Ń": 174, "F": 71, "N": 79, "Ê": 203, "Í": 206, "ħ": 134, "Ý": 222, "C": 68, "-": 46, "ñ": 242, "|": 125, ".": 47, "¬": 173, "w": 120, "À": 193, "Ĥ": 131, "İ": 143, "ß": 224, "J": 75, "Î": 207, "aaab": 260, "Ē": 19, "¹": 186, "´": 181, "¸": 185, "H": 73, "Ě": 27, "ĵ": 148, "@": 65, "a": 98, "đ": 18, "0": 49, "«": 172, "Ė": 23, "ı": 144, "Ô": 213, "Ħ": 133, "Ì": 205, "V": 87, "ÿ": 256, "I": 74, "L": 77, "\\": 93, "¤": 165, "ę": 26, "ó": 244, "7": 56, "Y": 90, "Â": 195, "Ĩ": 135, "þ": 255, "O": 80, "ĕ": 22, "f": 103, "Ĳ": 145, ",": 45, "[": 92, "Ę": 25, "abab": 261, "Õ": 214, "ŀ": 159, "ĥ": 132, "é": 234, "'": 40, "\"": 35, "E": 70, "X": 89, "è": 233, "5": 54, "6": 55, "9": 58, "m": 110, "y": 122, "T": 85, "ã": 228, "µ": 182, "j": 107, "Û": 220, "°": 177, "(": 41, "à": 225, "ĺ": 153, ":": 59, "Á": 194, "1": 50, "®": 175, "~": 127,
			},
			expectedMerges: []Merge{
				{
					Merge: [2]string{"a", "a"},
					Count: 150,
				},
				{
					Merge: [2]string{"a", "b"},
					Count: 120,
				},
				{
					Merge: [2]string{"aa", "aa"},
					Count: 30,
				},
				{
					Merge: [2]string{"aa", "ab"},
					Count: 30,
				},
				{
					Merge: [2]string{"ab", "ab"},
					Count: 30,
				},
				{
					Merge: [2]string{"ab", "c"},
					Count: 30,
				},
				{
					Merge: [2]string{"abc", "d"},
					Count: 30,
				},
			},
		},
	}

	for _, test := range tests {
		encoder, merges, err := BPE(test.frequencies, test.numSymbols, test.minFrequency)
		if err != nil {
			t.Fatal(err)
		}

		// uncomment this to get updated test values (double check before updating!)
		// log.Println(encoder)
		// for k, v := range encoder {
		// 	if k == "\"" || k == "\\" {
		// 		k = "\\" + k
		// 	}
		// 	fmt.Printf("\"%s\": %d,", k, v)
		// }

		// for _, m := range merges {
		// 	log.Println(m.Merge, m.Count)
		// }

		if len(encoder) != len(test.expectedEncoder) {
			t.Fatalf("expected %d encoder keys but got %d", len(test.expectedEncoder), len(encoder))
		}

		if len(merges) != len(test.expectedMerges) {
			t.Fatalf("expected %d merges but got %d", len(test.expectedMerges), len(merges))
		}

		for expectedk, expectedv := range test.expectedEncoder {
			v, ok := encoder[expectedk]

			if !ok {
				t.Fatalf("expected frequencies to contain \"%s\"", expectedk)
			}

			if expectedv != v {
				t.Fatalf("expected %s to have index %d but got %d", expectedk, expectedv, v)
			}
		}

		for i := range test.expectedMerges {
			m := merges[i]
			em := test.expectedMerges[i]

			if m.Merge[0] != em.Merge[0] {
				t.Fatalf("expected merge at %d to have first pair %s but got %s", i, m.Merge[0], em.Merge[0])
			}

			if m.Merge[1] != em.Merge[1] {
				t.Fatalf("expected merge at %d to have second pair %s but got %s", i, m.Merge[0], em.Merge[0])
			}

			if m.Count != em.Count {
				t.Fatalf("expected merge at %d to have count %d but got %d", i, m.Count, em.Count)
			}
		}
	}
}

// nolint: stylecheck
var bigText = `
“I have a golden bedroom,” he said softly to himself as he looked round,
and he prepared to go to sleep; but just as he was putting his head under
his wing a large drop of water fell on him.  “What a curious thing!” he
cried; “there is not a single cloud in the sky, the stars are quite clear
and bright, and yet it is raining.  The climate in the north of Europe is
really dreadful.  The Reed used to like the rain, but that was merely her
selfishness.”
Then another drop fell.
“What is the use of a statue if it cannot keep the rain off?” he said; “I
must look for a good chimney-pot,” and he determined to fly away.
But before he had opened his wings, a third drop fell, and he looked up,
and saw—Ah! what did he see?
The eyes of the Happy Prince were filled with tears, and tears were
running down his golden cheeks.  His face was so beautiful in the
moonlight that the little Swallow was filled with pity.
“Who are you?” he said.
“I am the Happy Prince.”
“Why are you weeping then?” asked the Swallow; “you have quite drenched
me.”
“When I was alive and had a human heart,” answered the statue, “I did not
know what tears were, for I lived in the Palace of Sans-Souci, where
sorrow is not allowed to enter.  In the daytime I played with my
companions in the garden, and in the evening I led the dance in the Great
Hall.  Round the garden ran a very lofty wall, but I never cared to ask
what lay beyond it, everything about me was so beautiful.  My courtiers
called me the Happy Prince, and happy indeed I was, if pleasure be
happiness.  So I lived, and so I died.  And now that I am dead they have
set me up here so high that I can see all the ugliness and all the misery
of my city, and though my heart is made of lead yet I cannot chose but
weep.”
“What! is he not solid gold?” said the Swallow to himself.  He was too
polite to make any personal remarks out loud.
“Far away,” continued the statue in a low musical voice, “far away in a
little street there is a poor house.  One of the windows is open, and
through it I can see a woman seated at a table.  Her face is thin and
worn, and she has coarse, red hands, all pricked by the needle, for she
is a seamstress.  She is embroidering passion-flowers on a satin gown for
the loveliest of the Queen’s maids-of-honour to wear at the next
Court-ball.  In a bed in the corner of the room her little boy is lying
ill.  He has a fever, and is asking for oranges.  His mother has nothing
to give him but river water, so he is crying.  Swallow, Swallow, little
Swallow, will you not bring her the ruby out of my sword-hilt?  My feet
are fastened to this pedestal and I cannot move.”
“I am waited for in Egypt,” said the Swallow.  “My friends are flying up
and down the Nile, and talking to the large lotus-flowers.  Soon they
will go to sleep in the tomb of the great King.  The King is there
himself in his painted coffin.  He is wrapped in yellow linen, and
embalmed with spices.  Round his neck is a chain of pale green jade, and
his hands are like withered leaves.”
“Swallow, Swallow, little Swallow,” said the Prince, “will you not stay
with me for one night, and be my messenger?  The boy is so thirsty, and
the mother so sad.”
“I don’t think I like boys,” answered the Swallow.  “Last summer, when I
was staying on the river, there were two rude boys, the miller’s sons,
who were always throwing stones at me.  They never hit me, of course; we
swallows fly far too well for that, and besides, I come of a family
famous for its agility; but still, it was a mark of disrespect.”
But the Happy Prince looked so sad that the little Swallow was sorry.
“It is very cold here,” he said; “but I will stay with you for one night,
and be your messenger.”
“Thank you, little Swallow,” said the Prince.
So the Swallow picked out the great ruby from the Prince’s sword, and
flew away with it in his beak over the roofs of the town.
He passed by the cathedral tower, where the white marble angels were
sculptured.  He passed by the palace and heard the sound of dancing.  A
beautiful girl came out on the balcony with her lover.  “How wonderful
the stars are,” he said to her, “and how wonderful is the power of love!”
“I hope my dress will be ready in time for the State-ball,” she answered;
“I have ordered passion-flowers to be embroidered on it; but the
seamstresses are so lazy.”
He passed over the river, and saw the lanterns hanging to the masts of
the ships.  He passed over the Ghetto, and saw the old Jews bargaining
with each other, and weighing out money in copper scales.  At last he
came to the poor house and looked in.  The boy was tossing feverishly on
his bed, and the mother had fallen asleep, she was so tired.  In he
hopped, and laid the great ruby on the table beside the woman’s thimble.
Then he flew gently round the bed, fanning the boy’s forehead with his
wings.  “How cool I feel,” said the boy, “I must be getting better”; and
he sank into a delicious slumber.
Then the Swallow flew back to the Happy Prince, and told him what he had
done.  “It is curious,” he remarked, “but I feel quite warm now, although
it is so cold.”
“That is because you have done a good action,” said the Prince.  And the
little Swallow began to think, and then he fell asleep.  Thinking always
made him sleepy.
When day broke he flew down to the river and had a bath.  “What a
remarkable phenomenon,” said the Professor of Ornithology as he was
passing over the bridge.  “A swallow in winter!”  And he wrote a long
letter about it to the local newspaper.  Every one quoted it, it was full
of so many words that they could not understand.
“To-night I go to Egypt,” said the Swallow, and he was in high spirits at
the prospect.  He visited all the public monuments, and sat a long time
on top of the church steeple.  Wherever he went the Sparrows chirruped,
and said to each other, “What a distinguished stranger!” so he enjoyed
himself very much.
When the moon rose he flew back to the Happy Prince.  “Have you any
commissions for Egypt?” he cried; “I am just starting.”
“Swallow, Swallow, little Swallow,” said the Prince, “will you not stay
with me one night longer?”
“I am waited for in Egypt,” answered the Swallow.  “To-morrow my friends
will fly up to the Second Cataract.  The river-horse couches there among
the bulrushes, and on a great granite throne sits the God Memnon.  All
night long he watches the stars, and when the morning star shines he
utters one cry of joy, and then he is silent.  At noon the yellow lions
come down to the water’s edge to drink.  They have eyes like green
beryls, and their roar is louder than the roar of the cataract.”
“Swallow, Swallow, little Swallow,” said the Prince, “far away across the
city I see a young man in a garret.  He is leaning over a desk covered
with papers, and in a tumbler by his side there is a bunch of withered
violets.  His hair is brown and crisp, and his lips are red as a
pomegranate, and he has large and dreamy eyes.  He is trying to finish a
play for the Director of the Theatre, but he is too cold to write any
more.  There is no fire in the grate, and hunger has made him faint.”
“I will wait with you one night longer,” said the Swallow, who really had
a good heart.  “Shall I take him another ruby?”
“Alas!  I have no ruby now,” said the Prince; “my eyes are all that I
have left.  They are made of rare sapphires, which were brought out of
India a thousand years ago.  Pluck out one of them and take it to him.
He will sell it to the jeweller, and buy food and firewood, and finish
his play.”
“Dear Prince,” said the Swallow, “I cannot do that”; and he began to
weep.
“Swallow, Swallow, little Swallow,” said the Prince, “do as I command
you.”
So the Swallow plucked out the Prince’s eye, and flew away to the
student’s garret.  It was easy enough to get in, as there was a hole in
the roof.  Through this he darted, and came into the room.  The young man
had his head buried in his hands, so he did not hear the flutter of the
bird’s wings, and when he looked up he found the beautiful sapphire lying
on the withered violets.
“I am beginning to be appreciated,” he cried; “this is from some great
admirer.  Now I can finish my play,” and he looked quite happy.
The next day the Swallow flew down to the harbour.  He sat on the mast of
a large vessel and watched the sailors hauling big chests out of the hold
with ropes.  “Heave a-hoy!” they shouted as each chest came up.  “I am
going to Egypt”! cried the Swallow, but nobody minded, and when the moon
rose he flew back to the Happy Prince.
5000/5000
Character limit: 5000
“我有一间金色的卧室，”他环顾四周轻声对自己说，
他准备去睡觉；但是就像他低下头一样
他的机翼上掉了一大滴水。 “真奇怪！”他
哭了“天上没有一片云，星星很清楚
和明亮，但正在下雨。欧洲北部的气候是
真可怕。芦苇曾经喜欢下雨，但这仅仅是她
自私。”
然后又下降了。
“如果不能挡雨，雕像的用途是什么？”他说; “一世
必须寻找一个好的烟囱锅。”他决心飞走。
但是在他张开翅膀之前，第三滴落下，他抬起头，
看到了-啊！他看到了什么？
快乐王子的眼睛充满了眼泪，眼泪
顺着他的金色的脸颊。他的脸好漂亮
月光下，小燕子充满了怜悯。
“你是谁？”他说。
“我是快乐王子。”
“那你为什么哭呢？”燕子问； “你浑身湿透了
我。”
雕像回答说：“当我还活着并拥有一颗人类的心时，我没有
知道眼泪是什么，因为我住在无忧宫
悲伤是不允许进入的。白天我和我一起玩
陪同下在花园里，晚上我带领大舞蹈
大厅。花园里绕着一堵非常高大的墙，但我从不问
超越它的一切，关于我的一切是如此美丽。我的朝臣
称我为快乐王子，如果享乐，我确实是快乐的
幸福。所以我住了，所以我死了。现在我已经死了
让我在这里坐得很高，我可以看到所有的丑陋和所有的痛苦
虽然我的心是铅制成的，但我别无选择
泣。”
“什么！他不是纯金吗？”燕子对自己说。他也是
有礼貌地大声发表任何个人评论。
雕像以低沉的音乐声继续说道：“远去
小街上有一间贫民窟。其中一个窗口打开，并且
通过它，我可以看到一个女人坐在桌子旁。她的脸很瘦而且
戴了，她的手是粗红的，全部被针刺了
是裁缝。她在缎面礼服上绣西番莲
女王的女仆中最可爱的，下次穿
球场球。她的小男孩躺在房间角落里的床上
生病。他发烧了，要桔子。他妈妈什么都没有
只给他河水，所以他在哭。燕子小燕子
吞下去，你不会把她的红宝石从我的剑柄中拿出来吗？我的脚
被固定在这个基座上，我无法动弹。”
燕子说：“我在埃及等着。” “我的朋友们飞起来了
下到尼罗河，和大朵莲花聊天。很快他们
将在伟大国王的坟墓中入睡。国王在那里
自己穿着涂满漆的棺材。他用黄色亚麻布包裹着，
用香料防腐。脖子上是一串淡绿色的玉，
他的手像枯萎的叶子。”
王子说：“燕子，燕子，小燕子。”
和我住一晚，做我的使者？这个男孩好渴，
母亲好伤心。”
燕子回答：“我不喜欢男孩。” “去年夏天，当我
留在河上，有两个粗鲁的男孩，米勒的儿子，
总是向我扔石头。当然，他们从没打我。我们
燕子飞得太远了，此外，我来自一个家庭
以敏捷着称；但仍然是不尊重的标志。”
但是快乐王子看起来很难过，小燕子对此感到抱歉。
他说：“这里很冷。” “但是我会和你在一起住一晚，
成为你的使者。”
“谢谢你，小燕子。”王子说。
因此燕子从王子的剑中挑出了红宝石，
带着它的喙飞过城镇的屋顶。
他路过大教堂的塔楼，那里是白色大理石天使
雕刻。他经过宫殿，听到了跳舞的声音。一个
美丽的女孩和她的爱人一起在阳台上走了出来。 “多么美妙
他对她说，星星是，爱的力量是多么美妙！
她回答说：“我希望我的着装能及时为国家舞会做好准备。”
“我已下令在上面绣上西番莲；但是
女裁缝是如此懒惰。”
他经过这条河，看见灯笼挂在
船。他越过了贫民窟，看到古老的犹太人讨价还价
彼此之间，并用铜秤称出金钱。最后他
来到那座贫民窟的房子，向他望去。那个男孩在疯狂地折腾着
他的床，母亲睡着了，她好累。在他
跳了起来，把大红宝石放在女人的顶针旁边的桌子上。
然后他轻轻地绕着床飞行，用他的扇形扇着男孩的额头
翅膀。男孩说：“我感觉有多酷，我必须变得更好”；和
他陷入沉睡中。
然后燕子飞回快乐王子
“Wǒ yǒuyī jiàn jīnsè de wòshì,” tā huángù sìzhōu qīngshēng duì zìjǐ shuō,
tā zhǔnbèi qù shuìjiào; dànshì jiù xiàng tā dīxià tou yīyàng
tā de jī yì shàng diàole yī dà dīshuǐ. “Zhēn qíguài!” Tā
kūle “tiānshàng méiyǒuyīpiàn yún, xīngxīng hěn qīngchǔ
hé míngliàng, dàn zhèngzài xià yǔ. Ōuzhōu běibù de qìhòu shì
zhēn kěpà. Lúwěi céngjīng xǐhuān xià yǔ, dàn zhè jǐnjǐn shì tā
zìsī.”
Ránhòu yòu xiàjiàngle.
“Rúguǒ bùnéng dǎng yǔ, diāoxiàng de yòngtú shì shénme?” Tā shuō; “yīshì
bìxū xúnzhǎo yīgè hǎo de yāncōng guō.” Tā juéxīn fēi zǒu.
Dànshì zài tā zhāng kāi chìbǎng zhīqián, dì sān dī luòxià, tā tái qǐtóu,
kàn dàole-a! Tā kàn dàole shénme?
Kuàilè wángzǐ de yǎnjīng chōngmǎnle yǎnlèi, yǎnlèi
shùnzhe tā de jīnsè de liǎnjiá. Tā de liǎn hǎo piàoliang
yuèguāng xià, xiǎo yànzi chōngmǎnle liánmǐn.
“Nǐ shì shéi?” Tā shuō.
“Wǒ shì kuàilè wángzǐ.”
“Nà nǐ wèishéme kū ne?” Yànzi wèn; “nǐ húnshēn shī tòule
wǒ.”
Diāoxiàng huídá shuō:“Dāng wǒ hái huózhe bìng yǒngyǒu yī kē rénlèi de xīn shí, wǒ méiyǒu
zhīdào yǎnlèi shì shénme, yīnwèi wǒ zhù zài wú yōu gōng
bēishāng shì bù yǔnxǔ jìnrù de. Báitiān wǒ hé wǒ yīqǐ wán
péitóng xià zài huāyuán lǐ, wǎnshàng wǒ dàilǐng dà wǔdǎo
dàtīng. Huāyuán lǐ ràozhe yī dǔ fēicháng gāodà de qiáng, dàn wǒ cóng bù wèn
chāoyuè tā de yīqiè, guānyú wǒ de yīqiè shì rúcǐ měilì. Wǒ de cháochén
chēng wǒ wéi kuàilè wángzǐ, rúguǒ xiǎnglè, wǒ quèshí shì kuàilè de
xìngfú. Suǒyǐ wǒ zhùle, suǒyǐ wǒ sǐle. Xiànzài wǒ yǐjīng sǐle
ràng wǒ zài zhèlǐ zuò dé hěn gāo, wǒ kěyǐ kàn dào suǒyǒu de chǒulòu hé suǒyǒu de tòngkǔ
suīrán wǒ de xīn shì qiān zhì chéng de, dàn wǒ bié wú xuǎnzé
qì.”
“Shénme! Tā bùshì chún jīn ma?” Yànzi duì zìjǐ shuō. Tā yěshì
yǒu lǐmào dì dàshēng fābiǎo rènhé gèrén pínglùn.
Diāoxiàng yǐ dīchén de yīnyuè shēng jìxù shuōdao:“Yuǎn qù
xiǎo jiē shàng yǒu yī jiàn pínmínkū. Qízhōng yīgè chuāngkǒu dǎkāi, bìngqiě
tōngguò tā, wǒ kěyǐ kàn dào yīgè nǚrén zuò zài zhuōzi páng. Tā de liǎn hěn shòu érqiě
dàile, tā de shǒu shì cū hóng de, quánbù bèi zhēn cìle
shì cáiféng. Tā zài duàn miàn lǐfú shàng xiù xī fān lián
nǚwáng de nǚpū zhōng zuì kě'ài de, xià cì chuān
qiúchǎng qiú. Tā de xiǎo nánhái tǎng zài fángjiān jiǎoluò lǐ de chuángshàng
shēngbìng. Tā fāshāole, yào júzi. Tā māmā shénme dōu méiyǒu
zhǐ gěi tā héshuǐ, suǒyǐ tā zài kū. Yànzi xiǎo yànzi
tūn xiàqù, nǐ bù huì bǎ tā de hóngbǎoshí cóng wǒ de jiàn bǐng zhōng ná chūlái ma? Wǒ de jiǎo
bèi gùdìng zài zhège jī zuò shàng, wǒ wúfǎ dòngtán.”
Yànzi shuō:“Wǒ zài āijí děngzhe.” “Wǒ de péngyǒumen fēi qǐláile
xià dào níluóhé, hé dà duǒ liánhuā liáotiān. Hěn kuài tāmen
jiàng zài wěidà guówáng de fénmù zhōng rùshuì. Guówáng zài nàlǐ
zìjǐ chuānzhuó tú mǎn qī de guāncai. Tā yòng huángsè yàmá bù bāoguǒzhe,
yòng xiāngliào fángfǔ. Bózi shàng shì yī chuàn dàn lǜsè de yù,
tā de shǒu xiàng kūwěi de yèzi.”
Wángzǐ shuō:“Yànzi, yànzi, xiǎo yànzi.”
Hé wǒ zhù yī wǎn, zuò wǒ de shǐzhě? Zhège nánhái hǎo kě,
mǔqīn hǎo shāngxīn.”
Yàn zǐ huídá:“Wǒ bù xǐhuān nánhái.” “Qùnián xiàtiān, dāng wǒ
liú zài héshàng, yǒu liǎng gè cūlǔ de nánhái, mǐ lēi de érzi,
zǒng shì xiàng wǒ rēng shítou. Dāngrán, tāmen cóng méi dǎ wǒ. Wǒmen
yàn zǐ fēi dé tài yuǎnle, cǐwài, wǒ láizì yīgè jiātíng
yǐ mǐnjiézhe chēng; dàn réngrán shì bù zūnzhòng de biāozhì.”
Dànshì kuàilè wángzǐ kàn qǐlái hěn nánguò, xiǎo yànzi duì cǐ gǎndào bàoqiàn.
Tā shuō:“Zhèlǐ hěn lěng.” “Dànshì wǒ huì hé nǐ zài yīqǐ zhù yī wǎn,
chéngwéi nǐ de shǐzhě.”
“Xièxiè nǐ, xiǎo yànzi.” Wángzǐ shuō.
Yīncǐ yànzi cóng wángzǐ de jiàn zhōng tiāo chūle hóngbǎoshí,
dàizhe tā de huì fēiguò chéngzhèn de wūdǐng.
Tā lù guo dà jiàotáng de tǎlóu, nàlǐ shì báisè dàlǐshí tiānshǐ
diāokè. Tā jīngguò gōngdiàn, tīng dàole tiàowǔ de shēngyīn. Yīgè
měilì de nǚhái hé tā de àirén yīqǐ zài yángtái shàng zǒule chūlái. “Duōme měimiào
tā duì tā shuō, xīngxīng shì, ài de lìliàng shì duōme měimiào!
Tā huídá shuō:“Wǒ xīwàng wǒ de zhuózhuāng néng jíshí wèi guójiā wǔhuì zuò hǎo zhǔnbèi.”
“Wǒ yǐ xiàlìng zài shàngmiàn xiù shàng xī fān lián; dànshì
nǚ cáiféng shì rúcǐ lǎnduò.”
Tā jīngguò zhè tiáo hé, kànjiàn dēnglóng guà zài
chuán. Tā yuèguòle pínmínkū, kàn dào gǔlǎo de yóutàirén tǎojiàhuánjià
bǐcǐ zhī jiān, bìngyòng tóng chèng chēng chū jīnqián. Zuìhòu tā
lái dào nà zuò pínmínkū de fángzi, xiàng tā wàng qù. Nàgè nánhái zài fēngkuáng de zhētengzhe
tā de chuáng, mǔqīn shuìzhele, tā hǎo lèi. Zài tā
tiàole qǐlái, bǎ dà hóngbǎoshí fàng zài nǚrén de dǐngzhēn pángbiān de zhuōzi shàng.
Ránhòu tā qīng qīng de ràozhe chuáng fēixíng, yòng tā de shànxíng shànzhe nánhái de étóu
chìbǎng. Nánhái shuō:“Wǒ gǎnjué yǒu duō kù, wǒ bìxū biàn dé gèng hǎo”; hé
tā xiànrù chénshuì zhōng.
Ránhòu yàn zǐ fēi huí kuàilè wángzǐ
"J'ai une chambre dorée", se dit-il doucement en regardant autour de lui,
et il se prépara à dormir; mais tout comme il mettait sa tête sous
son aile une grosse goutte d'eau est tombée sur lui. "Quelle chose curieuse!" il
pleuré; "Il n'y a pas un seul nuage dans le ciel, les étoiles sont assez claires
et lumineux, et pourtant il pleut. Le climat dans le nord de l'Europe est
vraiment affreux. Le roseau aimait la pluie, mais ce n'était que son
égoïsme."
Puis une autre goutte est tombée.
"À quoi sert une statue si elle ne peut pas empêcher la pluie de tomber?" il a dit; "JE
doit chercher une bonne cheminée, »et il a décidé de s'envoler.
Mais avant d'avoir ouvert ses ailes, une troisième goutte est tombée, et il a levé les yeux,
et j'ai vu… Ah! Qu'est-ce qu'il a vu?
Les yeux du Happy Prince étaient remplis de larmes, et les larmes étaient
coulant sur ses joues dorées. Son visage était si beau dans le
clair de lune que la petite Hirondelle était remplie de pitié.
"Qui êtes vous?" il a dit.
"Je suis le Prince Heureux."
"Pourquoi pleures-tu alors?" demanda l'Hirondelle; "Vous avez tout à fait trempé
moi."
"Quand j'étais vivant et que j'avais un cœur humain", a répondu la statue, "je n'ai pas
sais ce que les larmes étaient, car je vivais au Palais de Sans-Souci, où
le chagrin n'est pas autorisé à entrer. Le jour, je jouais avec mon
compagnons dans le jardin, et le soir j'ai dirigé la danse dans le Grand
Salle. Autour du jardin courait un mur très haut, mais je ne me suis jamais soucié de demander
ce qui se trouvait au-delà, tout en moi était si beau. Mes courtisans
m'a appelé le Prince Heureux, et j'étais vraiment heureux, si le plaisir était
bonheur. J'ai donc vécu et je suis mort. Et maintenant que je suis mort, ils ont
installe-moi ici si haut que je peux voir toute la laideur et toute la misère
de ma ville, et bien que mon cœur soit fait de plomb, je ne peux que choisir
pleurer."
"Quoi! n'est-il pas en or massif? se dit l'Hirondelle. Il était trop
poli de faire des remarques personnelles à haute voix.
"Loin", continua la statue à voix basse, "loin dans un
petite rue il y a une maison pauvre. L'une des fenêtres est ouverte et
à travers elle, je peux voir une femme assise à une table. Son visage est mince et
usée, et elle a des mains grossières et rouges, toutes piquées par l'aiguille, car elle
est couturière. Elle est en train de broder des fleurs de la passion sur une robe en satin pour
la plus belle des demoiselles d'honneur de la Reine à porter à la prochaine
Court-ball. Dans un lit dans le coin de la pièce, son petit garçon est couché
mauvais. Il a de la fièvre et demande des oranges. Sa mère n'a rien
pour lui donner de l'eau de rivière, alors il pleure. Avaler, avaler, peu
Avale, ne veux-tu pas lui faire sortir le rubis de ma garde d'épée? Mes pieds
sont attachés à ce piédestal et je ne peux pas bouger. »
«Je suis attendu en Égypte», a déclaré l'Hirondelle. «Mes amis s'envolent
et le long du Nil, et parler aux grandes fleurs de lotus. Bientôt, ils
ira dormir dans la tombe du grand roi. Le roi est là
lui-même dans son cercueil peint. Il est enveloppé dans du lin jaune, et
embaumé d'épices. Autour de son cou est une chaîne de jade vert pâle, et
ses mains sont comme des feuilles flétries.
"Avale, avale, petite hirondelle", dit le Prince, "ne resteras-tu pas
avec moi pour une nuit, et être mon messager? Le garçon a tellement soif, et
la mère si triste. "
"Je ne pense pas que j'aime les garçons", a répondu l'Hirondelle. «L'été dernier, quand j'ai
restait sur la rivière, il y avait deux garçons grossiers, les fils du meunier,
qui me jetaient toujours des pierres. Ils ne m'ont jamais frappé, bien sûr; nous
les hirondelles volent beaucoup trop bien pour ça, et d'ailleurs je viens d'une famille
célèbre pour son agilité; mais c'était quand même une marque d'irrespect. »
Mais le Prince Heureux avait l'air si triste que la petite Hirondelle était désolée.
«Il fait très froid ici», a-t-il dit; "Mais je resterai avec toi une nuit,
et soyez votre messager. "
"Merci, petit Swallow", dit le Prince.
Alors l'Hirondelle a choisi le grand rubis de l'épée du Prince, et
s'envola avec elle dans son bec sur les toits de la ville.
Il est passé par la tour de la cathédrale, où les anges de marbre blanc étaient
sculpté. Il est passé devant le palais et a entendu le bruit de la danse. UNE
belle fille est sortie sur le balcon avec son amant. "Merveilleux
les étoiles sont, lui dit-il, et comme la puissance de l'amour est merveilleuse! »
"J'espère que ma robe sera prête à temps pour le State-ball", a-t-elle répondu;
«J'ai commandé des fleurs de passion à broder dessus; mais le
les couturières sont tellement paresseuses.
Il est passé sur la rivière et a vu les lanternes accrochées aux mâts de
Les bateaux. Il est passé au-dessus du Ghetto et a vu les vieux juifs négocier
les uns avec les autres, et pesant de l'argent dans des échelles de cuivre. Enfin, il
est venu dans la maison pauvre et a regardé à l'intérieur. Le garçon
son lit, et la mère s'était endormie, elle était tellement fatiguée. En il
sauta et déposa le grand rubis sur la table à côté du dé à coudre de la femme.
Puis il a volé doucement autour du lit, caressant le front du garçon avec son
ailes. «Comme je me sens cool», a déclaré le garçon, «je dois aller mieux»; et
il sombra dans un délicieux sommeil.
Puis l'Hirondelle est retournée au Happy Prince
«У меня золотая спальня», тихо сказал он себе, оглядываясь,
и он приготовился идти спать; но так же, как он кладет голову под
его крыло на него упала большая капля воды. «Какая странная вещь!» он
плакала; «В небе нет ни одного облака, звезды довольно чистые
и ярко, и все же идет дождь. Климат на севере Европы
действительно ужасно Рид любил дождь, но это была только она
эгоизм."
Затем упала еще одна капля.
«Какая польза от статуи, если она не может предотвратить дождь?» он сказал; "Я
надо искать хороший дымоход », и он решил улететь.
Но прежде чем он открыл свои крылья, упала третья капля, и он поднял голову,
и увидел - ах! что он увидел?
Глаза Счастливого Принца наполнились слезами, а слезы были
бежит по его золотым щекам. Его лицо было так красиво в
лунный свет, что маленькая Ласточка была наполнена жалостью.
"Кто ты?" он сказал.
«Я счастливый принц»
«Почему ты тогда плачешь?» спросила ласточка; «Вы довольно залитые
меня."
«Когда я был жив и имел человеческое сердце, - ответил статуя, - я не
знаю, какие были слезы, потому что я жил во дворце Сан-Суси, где
горе не допускается. Днем я играл со своим
товарищи в саду, а вечером я привел танец в Великом
Холл. Вокруг сада проходила очень высокая стена, но я никогда не удосужился спросить
что лежало за этим, все во мне было так прекрасно. Мои придворные
назвал меня счастливым принцем, и я действительно счастлив, если удовольствие будет
счастье. Так я жил, и поэтому я умер. И теперь, когда я умер, у них есть
поставь меня здесь так высоко, чтобы я мог видеть все уродство и все страдания
моего города, и хотя мое сердце сделано из свинца, но я не могу выбрать, но
плакать."
"Какая! разве он не чистое золото? сказала Ласточка про себя. Он был слишком
вежливо, чтобы сделать какие-либо личные замечания вслух.
«Далеко», продолжил статуя тихим музыкальным голосом, «далеко в
Улица есть бедный дом. Одно из окон открыто, и
сквозь него я вижу женщину, сидящую за столом. Ее лицо худое и
но у нее грубые красные руки, все укололись иглой, потому что она
швея Она вышивает цветы из страсти на атласном платье для
самые красивые из подружек невесты королевы носить на следующем
Куры-мяч. В кровати в углу комнаты лежит ее маленький мальчик
больной. У него жар, и он просит апельсинов. Его мать не имеет ничего
дать ему кроме речной воды, поэтому он плачет. Ласточка, Ласточка, маленькая
Ласточка, ты не принесешь ей рубин из моей рукояти меча? Мои ноги
прикреплены к этому постаменту, и я не могу двигаться ».
«Меня ждут в Египте», - сказала ласточка. «Мои друзья взлетают
и вниз по Нилу, и говорить с большими цветами лотоса. Скоро они
ложусь спать в могиле великого короля. Король там
сам в своем расписном гробу. Он завернут в желтое белье, и
бальзамируется со специями. На шее у него цепочка бледно-зеленого нефрита, и
его руки подобны засохшим листьям.
«Ласточка, Ласточка, маленькая Ласточка, - сказал принц, - ты не останешься?
со мной на одну ночь, и будь моим посланником? Мальчик так хочет пить, и
мать такая грустная.
«Я не думаю, что мне нравятся мальчики», - ответила Ласточка. «Прошлым летом, когда я
находился на реке, там были два грубых мальчика, сыновья мельника,
которые всегда бросали в меня камни. Конечно, они никогда не били меня; мы
ласточки летают слишком хорошо для этого, и, кроме того, я из семьи
славится своей ловкостью; но все же это был знак неуважения ».
Но счастливый принц выглядел настолько грустным, что маленькой Ласточке было жаль.
«Здесь очень холодно», - сказал он. «Но я останусь с тобой на одну ночь,
и будь твоим посланником.
«Спасибо, маленькая Ласточка», - сказал принц.
Таким образом, Ласточка выбрала большой рубин из меча принца, и
улетел с ним в клюве над крышами города.
Он прошел мимо соборной башни, где были ангелы из белого мрамора.
лепили. Он прошел мимо дворца и услышал звуки танца.
красивая девушка вышла на балкон со своим любовником. "Как чудесно
звезды, - сказал он ей, - и как прекрасна сила любви! »
«Я надеюсь, что мое платье будет готово к государственному балу», - ответила она.
«Я приказал, чтобы на нем вышивали цветы страсти; но
швеи такие ленивые.
Он прошел через реку и увидел фонари, висящие на мачтах
корабли. Он прошел через гетто и увидел, как старые евреи торгуются
друг с другом, и взвешивание денег в медных весах. Наконец он
пришел в бедный дом и заглянул внутрь. Мальчик лихорадочно швырял
его кровать, а мать уснула, она так устала. В он
прыгал и положил большой рубин на стол рядом с наперстком женщины.
Затем он мягко облетел вокруг кровати, раздувая лоб мальчика
крылья. «Как здорово я себя чувствую, - сказал мальчик, - должно быть, мне становится лучше»; а также
он погрузился в восхитительный сон.
Затем ласточка полетела обратно к счастливому принцу
5000/5000
Character limit: 5000
「私は金色の寝室を持っています」と彼は見回しながら、優しく自分に言いました。
そして彼は眠りにつく準備をした。ちょうど彼が頭を下に置いていたように
彼の翼に大きな水滴が落ちた。 「なんて奇妙なことだ！」彼
叫んだ。 「空には単一の雲はなく、星は非常にはっきりしています
明るく、まだ雨が降っています。ヨーロッパ北部の気候は
本当に恐ろしい。リードはかつて雨が好きでしたが、それは彼女だけでした
わがまま。」
その後、もう一滴落ちました。
「雨を防ぐことができない場合、像の使用は何ですか？」彼は言った; "私
良い煙突の鍋を探す必要がある」と彼は飛び立つことを決心した。
しかし、彼が翼を開く前に、3滴目が下がり、彼は上を見上げました。
そして見た—ああ！彼は何を見ましたか？
幸せな王子の目は涙でいっぱいで、涙は
黄金の頬を駆け下りる。彼の顔はとても美しかった
小さなツバメが哀れに満ちた月明かり。
"あなたは誰？"彼は言った。
「私は幸福の王子です。」
「なんでそんなに泣いているの？」ツバメは尋ねました。 「あなたはかなりびしょぬれになった
私。」
「私が生きていて、人間の心を持っていたとき」と彫像は答えました。
私はサンスーシ宮殿に住んでいたので、涙が何であったかを知っています
悲しみは入りません。昼間は自分で遊んだ
庭の仲間、そして夕方には大王のダンスを主導しました
ホール。庭の周りは非常に高い壁を走りましたが、私は尋ねることを気にしませんでした
それを超えて何があったか、私についてのすべてがとても美しかった。私の廷臣
私を幸せな王子と呼んだ、そしてもし喜びがあったら、私は本当に幸せだった
幸福。だから私は生き、そして死んだ。そして今、私は死んでいるので、彼らは持っています
私をここまで高く設定して、醜さと悲惨さをすべて見ることができます
私の街の、そして私の心はリードから作られていますが、私は選択することはできませんが
泣きました。」
"何！彼は純金ではないのですか？」ツバメは自分に言いました。彼もそうだった
個人的な発言を大声で出すように礼儀正しい。
「遠くに」と低い声で像を続けた、「遠くに
小さな通りには貧しい家があります。ウィンドウの1つが開いています。
それを通して、私はテーブルに座っている女性を見ることができます。彼女の顔は薄くて
着用して、彼女は粗い、赤い手、すべて針で刺されています、彼女のために
仕立て屋です。彼女はサテンのガウンにパッションフラワーを刺繍しています。
女王の次のときに着用する名誉のメイドの中で最も美しい
コートボール。部屋の隅にあるベッドで彼女の小さな男の子が横たわっています
病気。彼は熱があり、オレンジを求めています。彼の母親は何も持っていません
彼に川の水を与えるために、彼は泣いています。ツバメ、ツバメ、少し
ツバメ、彼女のルビーを私の剣の柄から出してくれませんか？私の足
この台座に固定されており、私は移動できません。」
「私はエジプトで待っています」とツバメは言いました。 「私の友達は飛んでいます
ナイル川を下りて、大きな蓮の花と話しています。すぐに
偉大な王の墓で眠りにつくでしょう。王がいます
自分の塗られた棺の中にいる。彼は黄色のリネンに包まれており、
スパイスで防腐処理。彼の首の周りは薄緑色のヒスイの鎖であり、
彼の手は枯れた葉のようなものです。」
「ツバメ、ツバメ、小さなツバメ」と王子は言いました
一晩私と一緒に、そして私のメッセンジャーになりますか？男の子はとても喉が渇いています、そして
母親はとても悲しい。」
「私は男の子が好きだとは思いません」とツバメは答えました。 「去年の夏、私が
川に滞在していて、二人の失礼な男の子、ミラーの息子がいました、
いつも私に石を投げていました。もちろん、彼らは私を襲ったことはありません。我々
ツバメはそのためにあまりにもうまく飛ぶ、そして私は家族のもとに来ます
その俊敏性で有名です。それでも、それは失礼の印でした。」
しかし、幸せな王子はとても悲しそうに見えたので、小さなツバメは残念でした。
「ここはとても寒い」と彼は言った。 「しかし私はあなたと一晩滞在します、
そしてあなたのメッセンジャーになります。」
「ありがとう、小さなツバメ」と王子は言いました。
それでツバメは王子の剣から素晴らしいルビーを選び、
彼のくちばしでそれを町の屋根の上に飛んで行きました。
彼は白い大理石の天使たちがいた大聖堂の塔を通り過ぎました
彫刻された。彼は宮殿を通り過ぎて踊りの音を聞いた。あ
美しい少女が恋人と一緒にバルコニーに出てきました。 "なんて素敵なの
星はある」と彼は彼女に言った、そして「愛の力はどれほど素晴らしいのでしょう！」
「私はドレスが州のボールに間に合うように準備ができていることを願っています」と彼女は答えた。
「私はパッションフラワーに刺繍するように命令しました。しかし
仕立て屋はとても怠惰です。」
彼は川を越えて、灯台がマストにぶら下がっているのを見ました
船。彼はゲットーを過ぎて、古いユダヤ人が交渉しているのを見ました
銅のはかりでお金を計ります。ついに彼
かわいそうな家にやって来て、のぞき込んだ。少年は熱っぽく投げた
彼のベッド、そして母親は眠りに落ちて、彼女はとても疲れていました。彼の中で
飛び跳ねて、女性の指ぬきの横にあるテーブルに大きなルビーを置いた。
それから彼はベッドの周りを優しく飛んで、少年の額に
翼。 「なんてクールな気分だ」と少年は言った。そして
彼はおいしい眠りに沈んだ。
それからツバメはハッピープリンスに戻ってきました
4999/5000
"لدي غرفة نوم ذهبية" ، قال بلطف لنفسه وهو ينظر مستديرًا ،
واستعد للنوم. ولكن كما كان يضع رأسه تحته
على جناحه سقطت قطرة ماء كبيرة عليه. "يا له من شيء غريب!" هو
بكت؛ "لا توجد سحابة واحدة في السماء ، والنجوم واضحة تمامًا
ومشرقة ومع ذلك تمطر. المناخ في شمال أوروبا
مروع حقا. اعتادت ريد أن تحب المطر ، لكن ذلك كان مجردها
الأنانية ".
ثم سقطت قطرة أخرى.
"ما فائدة التمثال إذا كان لا يستطيع منع المطر؟" هو قال؛ "أنا
يجب أن يبحثوا عن وعاء جيد للمدخنة "، وقرر أن يطير بعيدًا.
ولكن قبل أن يفتح جناحيه ، سقطت قطرة ثالثة ، ونظر إلى الأعلى ،
ورأيت - آه! ماذا قال؟
كانت عيون الأمير السعيد مليئة بالدموع والدموع
يركض خديه الذهبي. كان وجهه جميلا جدا في
ضوء القمر أن السنونو الصغير كان مليئًا بالشفقة.
"من أنت؟" هو قال.
"أنا الأمير السعيد."
"لماذا تبكين إذن؟" سأل السنونو. "أنت غارق تمامًا
أنا."
أجاب التمثال: "عندما كنت على قيد الحياة وكان لدي قلب بشري ، لم أفعل
أعرف ما هي الدموع ، لأنني عشت في قصر سانسوسي ، حيث
لا يسمح للحزن بالدخول. في النهار لعبت مع بلدي
الصحابة في الحديقة ، وفي المساء كنت أقود الرقص في العظمة
صالة. حول الحديقة كان يدير جدارًا رفيعًا جدًا ، لكنني لم أكن أهتم أبدًا بالسؤال
ما وراءه ، كان كل شيء عني جميلًا جدًا. حاشاتي
اتصل بي الأمير السعيد ، وسعدت حقًا ، إذا كنت سعيدًا
السعادة. لذلك عشت ، وماتت. والآن أنا ميت لديهم
وضعني هنا عالياً لدرجة أنني أستطيع رؤية كل البشاعة وكل البؤس
لمدينتي ، وعلى الرغم من أن قلبي مصنوع من الرصاص إلا أنني لا أستطيع أن أختار لكن
بكاء ".
"ماذا! أليس هو ذهب خالص؟ " قال السنونو له. كان كذلك
مهذبا لإبداء أي ملاحظات شخصية بصوت عال.
"بعيد" ، تابع التمثال بصوت موسيقي منخفض ، "بعيدًا في
الشارع الصغير هناك منزل فقير. إحدى النوافذ مفتوحة و
من خلالها يمكنني رؤية امرأة جالسة على طاولة. وجهها رقيق و
ترتديه ، ولديها أيد حمراء خشنة ، وخزتها جميعها الإبرة ، لأنها
خياطة. إنها تطرز زهور العاطفة على ثوب من الساتان
أجمل خادمات الملكة لارتدائهن في اليوم التالي
كرة المحكمة. في سرير في زاوية الغرفة يرقد طفلها الصغير
سوف. يعاني من الحمى ، ويطلب البرتقال. والدته ليس لديها شيء
لإعطائه سوى مياه النهر ، لذلك يبكي. ابتلاع ، ابتلاع ، القليل
ابتلاع ، ألن تحضر لها الياقوت من ذيل سيفي؟ قدمي
يتم تثبيتها على هذه القاعدة ولا يمكنني التحرك ".
قال السنونو: "أنتظرني في مصر". "أصدقائي يطيرون
وأسفل النيل ، والتحدث مع أزهار اللوتس الكبيرة. هم قريبا
سينام في قبر الملك العظيم. الملك هناك
نفسه في نعشه المطلي. ملفوفة في الكتان الأصفر ، و
محنط مع البهارات. حول رقبته سلسلة من اليشم الأخضر الشاحب
يديه مثل أوراق ذابلة. "
قال الأمير: "ابتلع ، ابتلع ، ابتلاع صغير" ، "لن تبقى
معي لليلة واحدة ، وأكون رسولي؟ الصبي عطشان جدا
الأم حزينة جدا ".
أجاب السنونو "لا أعتقد أنني أحب الأولاد". "الصيف الماضي ، عندما كنت
كان يقيم على النهر ، كان هناك صبيان وقحين ، أبناء ميلر ،
الذين كانوا يرمونني بالحجارة دائمًا. لم يضربوني أبداً بالطبع. نحن
يبتلع السنون جيدًا جدًا لذلك ، بالإضافة إلى أنني أتيت من عائلة
مشهورة بخفة الحركة ؛ ولكن مع ذلك ، كانت علامة على عدم الاحترام ".
لكن الأمير السعيد بدا حزينًا لدرجة أن السنونو الصغير كان آسفًا.
قال: "الجو بارد جداً هنا". "ولكن سأبقى معك لليلة واحدة ،
وكن رسولك ".
قال الأمير: "شكرا لك أيها السنونو الصغير".
لذا التقط السنونو الياقوت العظيم من سيف الأمير ، و
طار معها في منقاره فوق أسطح المدينة.
مر بجانب برج الكاتدرائية ، حيث كانت الملائكة الرخامية البيضاء
منحوت. مر بجانب القصر وسمع صوت الرقص. أ
ظهرت فتاة جميلة على الشرفة مع عشيقها. "كم هو رائع
قال لها النجوم "وكم هي رائعة قوة الحب!"
أجابت: "آمل أن يكون ثوبي جاهزًا في الوقت المناسب من أجل الكرة الرسمية".
"لقد أمرت بتطريز أزهار العاطفة عليها ؛ لكن ال
الخياطات كسالى جدا. "
مر فوق النهر ، ورأى الفوانيس معلقة على الصواري
السفن. لقد مر فوق الحي اليهودي ، ورأى اليهود المسنين يساومون
مع بعضهم البعض ، ووزن المال في موازين النحاس. أخيرا هو
جاء إلى المنزل الفقير ونظر. الصبي كان يرمي بحرارة
سريره ، وكانت الأم نائمة ، كانت متعبة للغاية. فيه
قافز ، ووضع الياقوت الكبير على الطاولة بجانب كشتبان المرأة.
ثم طار حول السرير برفق ، واثارت جبهته بالفتى
أجنحة. قال الصبي ، "كم أشعر بالبرد ، لا بد لي من أن أتحسن". و
غرق في سبات لذيذ.
ثم طار السنونو إلى الأمير السعيد
4999/5000
«Έχω μια χρυσή κρεβατοκάμαρα», είπε απαλά στον εαυτό του καθώς κοίταξε γύρω του,
και ετοιμάστηκε να πάει για ύπνο. αλλά ακριβώς όπως έβαζε το κεφάλι του κάτω
στην πτέρυγα του πέφτει μια μεγάλη σταγόνα νερό. «Τι περίεργο πράγμα!» αυτός
φώναξε? «Δεν υπάρχει ούτε ένα σύννεφο στον ουρανό, τα αστέρια είναι αρκετά καθαρά
και φωτεινό, και όμως βρέχει. Το κλίμα στα βόρεια της Ευρώπης είναι
πραγματικά φοβερή. Ο Ριντ συμπαθούσε τη βροχή, αλλά αυτή ήταν μόνο η ίδια
ιδιοτέλεια."
Μετά έπεσε μια άλλη σταγόνα.
"Ποια είναι η χρήση ενός αγάλματος αν δεν μπορεί να κρατήσει τη βροχή;" αυτός είπε; "ΕΓΩ
πρέπει να ψάξει για μια καλή καμινάδα »και αποφάσισε να πετάξει μακριά.
Αλλά πριν ανοίξει τα φτερά του, έπεσε μια τρίτη σταγόνα, και κοίταξε ψηλά,
και είδα — Αχ! τι είδε;
Τα μάτια του ευτυχούς πρίγκιπα ήταν γεμάτα δάκρυα και δάκρυα
τρέχει κάτω από τα χρυσά μάγουλά του. Το πρόσωπό του ήταν τόσο όμορφο στο
σεληνόφως ότι το μικρό Χελιδόνι γέμισε με οίκτο.
"Ποιος είσαι?" αυτός είπε.
«Είμαι ο ευτυχισμένος πρίγκιπας».
«Γιατί κλαις τότε;» ρώτησε το Swallow. «Είσαι αρκετά βρεγμένος
μου."
«Όταν ήμουν ζωντανός και είχα ανθρώπινη καρδιά», απάντησε το άγαλμα, «δεν το έκανα
ξέρω τι ήταν τα δάκρυα, γιατί έζησα στο Παλάτι του Sans-Souci, όπου
Δεν επιτρέπεται η είσοδος στη θλίψη. Κατά τη διάρκεια της ημέρας έπαιζα με το δικό μου
σύντροφοι στον κήπο, και το βράδυ οδήγησα τον χορό στο Μέγα
Αίθουσα. Γύρω από τον κήπο έτρεχε ένας πολύ υψηλός τοίχος, αλλά ποτέ δεν μου άρεσε να ρωτήσω
αυτό που ήταν πέρα ​​από αυτό, όλα για μένα ήταν τόσο όμορφα. Οι αυλοί μου
μου τηλεφώνησε ο Χαρούμενος Πρίγκιπας και χαρούμενος πράγματι ήμουν, αν ήταν χαρά
ευτυχία. Έτσι έζησα και πέθανα. Και τώρα που είμαι νεκρός
ετοιμαστείτε εδώ τόσο ψηλά που μπορώ να δω όλη την ασχήμια και όλη τη δυστυχία
της πόλης μου, και παρόλο που η καρδιά μου είναι από μόλυβδο, δεν μπορώ παρά να επιλέξω
κλαίω."
"Τι! δεν είναι στερεός χρυσός; " είπε ο Χελιδόνι στον εαυτό του. Ήταν επίσης
ευγενικό να κάνει δυνατές προσωπικές παρατηρήσεις.
«Μακριά», συνέχισε το άγαλμα με χαμηλή μουσική φωνή, «πολύ μακριά σε ένα
μικρό δρόμο υπάρχει ένα φτωχό σπίτι. Ένα από τα παράθυρα είναι ανοιχτό και
μπορώ να δω μια γυναίκα που κάθεται σε ένα τραπέζι. Το πρόσωπό της είναι λεπτό και
φοριέται και έχει χονδροειδή, κόκκινα χέρια, όλα τρυπημένα από τη βελόνα
είναι μοδίστρα. Κέντημα λουλουδιών πάθους σε σατέν φόρεμα
την ομορφότερη από τις βασίλισσες της τιμής της Βασίλισσας που φορούν την επόμενη
Γήπεδο. Σε ένα κρεβάτι στη γωνία του δωματίου, το μικρό αγόρι της βρίσκεται
Εγώ θα. Έχει πυρετό και ζητά πορτοκάλια. Η μητέρα του δεν έχει τίποτα
για να του δώσει νερό, αλλά να κλαίει. Καταπιείτε, Χελιδόνι, λίγο
Χελιδόνι, θα σας δεν της το ρουμπίνι από το σπαθί μου, λαβή φέρει; Τα πόδια μου
στερεώνονται σε αυτό το βάθρο και δεν μπορώ να κινηθώ. "
«Με περίμενα στην Αίγυπτο», είπε το Swallow. «Οι φίλοι μου ανεβαίνουν
και κάτω από το Νείλο, και μιλώντας με τα μεγάλα άνθη λωτού. Σύντομα αυτοί
θα κοιμηθεί στον τάφο του μεγάλου Βασιλιά. Ο Βασιλιάς είναι εκεί
ο ίδιος στο βαμμένο φέρετρο του. Είναι τυλιγμένο σε κίτρινα λινά και
καλυμμένο με μπαχαρικά. Γύρω από το λαιμό του είναι μια αλυσίδα από ανοιχτό πράσινο νεφρίτη, και
τα χέρια του είναι σαν μαραμένα φύλλα. "
«Swallow, Swallow, Little Swallow», είπε ο πρίγκιπας, «δεν θα μείνεις
μαζί μου για μια νύχτα, και γίνε αγγελιοφόρος μου; Το αγόρι είναι τόσο διψασμένο και
η μητέρα τόσο λυπημένη. "
«Δεν νομίζω ότι μου αρέσουν τα αγόρια», απάντησε το Swallow. «Το περασμένο καλοκαίρι, όταν εγώ
έμενε στον ποταμό, υπήρχαν δύο αγενή αγόρια, οι γιοι του μυλωνά,
που μου έριχναν πάντα πέτρες. Φυσικά δεν με χτύπησαν. εμείς
τα χελιδόνια πετούν πολύ καλά για αυτό, και εκτός αυτού, προέρχομαι από μια οικογένεια
φημίζεται για την ευκινησία του. αλλά ακόμα, ήταν ένα σημάδι ασεβείας. "
Αλλά ο ευτυχισμένος πρίγκιπας φαινόταν τόσο λυπημένος που το μικρό χελιδόνι λυπούταν.
«Είναι πολύ κρύο εδώ», είπε. "Αλλά θα μείνω μαζί σου για μία νύχτα,
και γίνε αγγελιοφόρος σου. "
«Ευχαριστώ, Μικρό Χελιδόνι», είπε ο Πρίγκιπας.
Έτσι το Swallow διάλεξε το μεγάλο ρουμπίνι από το σπαθί του πρίγκιπα, και
πέταξε μαζί του στο ράμφος του πάνω από τις στέγες της πόλης.
Πέρασε από τον πύργο του καθεδρικού ναού, όπου βρίσκονταν οι λευκοί μαρμάρινοι άγγελοι
σκαλιστός. Πέρασε από το παλάτι και άκουσε τον ήχο του χορού. ΕΝΑ
όμορφη κοπέλα βγήκε στο μπαλκόνι με τον εραστή της. "Πόσο θαυμάσιο
τα αστέρια είναι, "της είπε," και πόσο υπέροχη είναι η δύναμη της αγάπης! "
«Ελπίζω ότι το φόρεμά μου θα είναι έτοιμο εγκαίρως για το State-ball», απάντησε.
«Έχω διατάξει να κεντηθούν τα λουλούδια του πάθους. αλλά το
Οι μοδίστρες είναι τόσο τεμπέλης. "
Πέρασε πάνω από τον ποταμό, και είδε τα φανάρια να κρέμονται στους ιστούς του
τα πλοία. Πέρασε πάνω από το Γκέτο, και είδε τους παλιούς Εβραίους να διαπραγματεύονται
το ένα με το άλλο, και ζυγίζοντας χρήματα σε κλίμακες χαλκού. Επιτέλους αυτός
ήρθε στο φτωχό σπίτι και κοίταξε μέσα. Το αγόρι πετούσε πυρετωδώς
το κρεβάτι του, και η μητέρα είχε κοιμηθεί, ήταν τόσο κουρασμένη. Σε αυτόν
πήδηξε και έβαλε το μεγάλο ρουμπίνι στο τραπέζι δίπλα στη δαχτυλήθρα της γυναίκας.
Στη συνέχεια, πέταξε απαλά γύρω από το κρεβάτι, αερίζοντας το μέτωπο του αγοριού με το δικό του
παρασκήνια. «Πόσο δροσερό νιώθω», είπε το αγόρι, «Πρέπει να γίνω καλύτερος». και
βυθίστηκε σε ένα υπέροχο ύπνο.
4999/5000
„Mam złotą sypialnię”, powiedział cicho do siebie, rozglądając się,
i przygotował się do snu; ale tak jak kładł głowę
na jego skrzydle spadła na niego duża kropla wody. „Co za dziwna rzecz!” on
płakał; „Na niebie nie ma ani jednej chmury, gwiazdy są całkiem jasne
i jasne, a jednak pada deszcz. Klimat na północy Europy to
naprawdę straszne. Trzcina lubiła deszcz, ale to była tylko ona
egoizm."
Potem spadła kolejna kropla.
„Jaki jest pożytek z posągu, jeśli nie może on powstrzymać deszczu?” powiedział; "JA
musi poszukać dobrego komina ”i postanowił odlecieć.
Ale zanim otworzył skrzydła, spadła trzecia kropla i spojrzał w górę,
i zobaczyłem - Ach! Co on zobaczył?
Oczy Szczęśliwego Księcia były pełne łez i łez
spływały po jego złotych policzkach. Jego twarz była taka piękna
światło księżyca, że ​​mała Jaskółka była pełna litości.
"Kim jesteś?" powiedział.
„Jestem Szczęśliwym Księciem”.
„Dlaczego więc płaczesz?” zapytał Jaskółka; „Całkiem przemokłeś
mnie."
„Kiedy żyłem i miałem ludzkie serce”, odpowiedział posąg, „nie zrobiłem tego
wiem, czym były łzy, bo mieszkałem w Pałacu Sans-Souci, gdzie
smutek nie może wejść. W ciągu dnia bawiłem się ze mną
towarzysze w ogrodzie, a wieczorem poprowadziłem taniec w Wielkim
Hol. Wokół ogrodu biegł bardzo wyniosły mur, ale nigdy nie miałem ochoty pytać
co leżało poza tym, wszystko we mnie było takie piękne. Moi dworzanie
nazwał mnie Szczęśliwym Księciem, i rzeczywiście byłam szczęśliwa, jeśli przyjemność
szczęście. Więc żyłem i tak umarłem. A teraz, kiedy jestem martwy, mają
postaw mnie tutaj tak wysoko, że mogę zobaczyć całą brzydotę i całą nędzę
mojego miasta i chociaż moje serce jest z ołowiu, nie mogę wybrać, ale
płakać."
"Co! czy on nie jest z litego złota? powiedział do siebie Jaskółka. On też był
uprzejmie wypowiadać na głos wszelkie osobiste uwagi.
„Daleko” kontynuował posąg niskim, muzycznym głosem, „daleko w
przy małej uliczce jest biedny dom. Jedno z okien jest otwarte i
przez to widzę kobietę siedzącą przy stole. Jej twarz jest chuda i
noszona, a ona ma szorstkie, czerwone ręce, wszystkie nakłute igłą, bo ona
jest krawcową. Haftuje kwiaty męczennicy na satynowej sukni
najładniejsza z pokojówek Królowej, którą można nosić w następnym
Piłka sądowa. W łóżku w rogu pokoju leży jej mały chłopiec
chory. Ma gorączkę i prosi o pomarańcze. Jego matka nie ma nic
dać mu oprócz wody rzecznej, więc płacze. Jaskółka, Jaskółka, mała
Jaskółka, czy nie wyciągniesz jej rubinu z mojej rękojeści miecza? Moje stopy
są przymocowane do tego cokołu i nie mogę się ruszyć. ”
„Jestem oczekiwany w Egipcie” - powiedziała Jaskółka. „Moi przyjaciele lecą w górę
i w dół Nilu, i rozmawiając z dużymi kwiatami lotosu. Wkrótce oni
pójdzie spać w grobie wielkiego króla. Król tam jest
w malowanej trumnie. Jest owinięty żółtym lnem i
balsamowane przyprawy. Na szyi ma łańcuch jasnozielonego jadeitu i
jego ręce są jak zwiędłe liście. ”
„Połknij, połknij, mała połknij”, powiedział książę, „nie zostaniesz
ze mną na jedną noc i być moim posłańcem? Chłopiec jest tak spragniony i
matka taka smutna. ”
„Nie sądzę, że lubię chłopców”, odpowiedziała Jaskółka. „Zeszłego lata, kiedy ja
przebywał na rzece, było dwóch niegrzecznych chłopców, synów młynarza,
którzy zawsze rzucali we mnie kamieniami. Oczywiście nigdy mnie nie uderzyli; my
jaskółki latają o wiele za dobrze, a poza tym pochodzę z rodziny
słynie ze swojej zwinności; ale wciąż był to znak braku szacunku. ”
Ale Szczęśliwy Książę wyglądał tak smutno, że mała Jaskółka żałowała.
„Jest tu bardzo zimno” - powiedział; „Ale zostanę z tobą na jedną noc,
i bądź swoim posłańcem ”.
„Dziękuję, mała Jaskółko”, powiedział książę.
Więc Jaskółka wybrała wielki rubin z miecza księcia i ...
odleciał z nim w dziobie nad dachami miasta.
Minął wieżę katedralną, w której znajdowały się białe marmurowe anioły
rzeźbione. Minął pałac i usłyszał dźwięk tańca. ZA
piękna dziewczyna wyszła na balkon ze swoim kochankiem. "Jak cudownie
gwiazdy są - powiedział do niej - i jak cudowna jest moc miłości!
„Mam nadzieję, że moja suknia będzie gotowa na bal państwowy”, odpowiedziała;
„Zamówiłem na nim haftowane kwiaty męczennicy; ale
szwaczki są takie leniwe. ”
Minął rzekę i zobaczył latarnie zawieszone na masztach
związki. Przeszedł przez getto i zobaczył targujących się starych Żydów
ze sobą i odważanie pieniędzy w miedzianych skalach. Nareszcie on
przyszedł do biednego domu i zajrzał do środka. Chłopiec gorączkowo miotał się
jego łóżko, a matka zasnęła, była tak zmęczona. W on
podskoczył i położył wielki rubin na stole obok naparstka kobiety.
Potem delikatnie otoczył łóżko, wachlując czoło chłopca swoim
skrzydełka. „Jak fajnie się czuję”, powiedział chłopiec, „muszę być coraz lepszy”; i
pogrążył się w pysznym śnie.
Potem Jaskółka poleciała z powrotem do Szczęśliwego Księcia
Send feedback
History
Saved
Community
`
