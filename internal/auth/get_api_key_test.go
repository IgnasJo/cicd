package auth

import (
	"net/http"
	"testing"
	"github.com/google/go-cmp/cmp"
)

type Expected struct {
	Val string
	Err error
}

func TestGetAPIKey(t *testing.T) {
    tests := map[string]struct {
        input 		http.Header
        want  	  Expected
    }{
        "normal":   {input: http.Header{"Authorization": []string{"ApiKey hello"}}, want: Expected{Val: "hello", Err: nil}},
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            got, err := GetAPIKey(tc.input)
            diff := cmp.Diff(tc.want, Expected{Val: got, Err: err})
            if diff != "" {
                t.Fatalf(diff)
            }
        })
    }
}