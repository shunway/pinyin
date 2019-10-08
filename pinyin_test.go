package pinyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	Want := []string{"GuangRongYuMengXiang", "guang rong yu meng xiang",
		"guāng-róng-yǔ-mèng-xiǎng", "guang rong yu meng xiang"}

	str, err := New("光荣与梦想").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	}
	if str != Want[0] {
		t.Fatalf("Want %v, but got %v", Want[0], str)
	}

	str, err = New("光荣与梦想").Split(" ").Mode(WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	}
	if str != Want[1] {
		t.Fatalf("Want %v, but got %v", Want[1], str)
	}

	str, err = New("光荣与梦想").Split("-").Mode(Tone).Convert()
	if err != nil {
		t.Error(err)
	}
	if str != Want[2] {
		t.Fatalf("Want %v, but got %v", Want[2], str)
	}

	str, err = New("光荣与梦想").Convert()
	if err != nil {
		t.Error(err)
	}
	if str != Want[3] {
		t.Fatalf("Want %v, but got %v", Want[3], str)
	}
}
