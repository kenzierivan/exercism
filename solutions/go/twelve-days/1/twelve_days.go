package twelve

import "fmt"

var verse = map[int]string {
    1: "a Partridge in a Pear Tree.",
    2: "two Turtle Doves, and a Partridge in a Pear Tree.",
    3: "three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    4: "four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    5: "five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    6: "six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    7: "seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    8: "eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    9: "nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    10: "ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    11: "eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
    12: "twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
}

var day = map[int]string {
    1: "first",
    2: "second",
    3: "third",
    4: "fourth",
    5: "fifth",
    6: "sixth",
    7: "seventh",
    8: "eighth",
    9: "ninth",
    10: "tenth",
    11: "eleventh",
    12: "twelfth",
}
func Verse(i int) string {
    var asd string
    repeat := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", day[i])
    asd += repeat
    asd += verse[i]
	return asd
}

func Song() string {
    var lyrics string
	for i := 1; i <= 12; i++ {
        lyrics += Verse(i)
        if i != 12 {
            lyrics += "\n"
        }
        
    }
    return lyrics
}
