package cmd

import "testing"

func TestSayHello(t *testing.T) {

	subtests:=[]struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Harsh"},
			result: "Hello, Harsh!",
		},
		{
			items: []string{"Harsh","John"},
			result: "Hello, Harsh, John!",
		},
	}

	for _,st:=range subtests{
		if s:=Say(st.items); s!=st.result{
			t.Errorf("wanted %s (%v), got %s",st.result,st.items,s)
		}
	}

	// want:="Hello, test!"
	// got:=Say([]string{"test"})

	// if want!=got{
	// 	t.Errorf("wanted %s, got %s",want ,got)
	// }
}
