package event

import "testing"

func TestPut(t *testing.T) {
    wanted := "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
    got := Put("korn") 
    if got != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}