package pre

import "testing"

func TestTextParse(t *testing.T) {
	tests := []*Text{
		NewText(2543654365, `e3o3o3ekdkke3%hello%="Hello, blyat!" fekwpdfo`),
		NewText(2543654365, `Say "%hello%"`),
		NewText(2543654365, `%ciao%=%hello%`),
	}

	for _, test := range tests {
		t.Log(test.Parse())
	}
}

func BenchmarkTextParse(b *testing.B) {
	NewText(0, `%blyat%="Suka nahui blyat!"`).Parse()
	text := NewText(0, `3edjfjiejdi49ir4iri4rj9i49r49rj4ijrfifinfjnfjfrnjfrnrfjr4j%blyat%ri4jrjfrufrufrrir4irirj4ifjirfjfr4rj8i4r8ir4ffjrifjifrji4rj`)

	for i := 0; i < b.N; i++ {
		text.Parse()
	}
}
