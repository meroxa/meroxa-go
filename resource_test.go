package meroxa

import (
	"testing"
)

func TestEncodeURLCreds(t *testing.T) {
	tests := []struct {
		in   string
		want string
		err  error
	}{
		{"s3://KAHDKJKSA:askkshe+skje/fhds@us-east-1/bucket", "s3://KAHDKJKSA:askkshe+skje%2Ffhds@us-east-1/bucket", nil},
		{"s3://KAHDKJKSA:secretsecret@us-east-1/bucket", "s3://KAHDKJKSA:secretsecret@us-east-1/bucket", nil},
		{"s3://us-east-1/bucket", "s3://us-east-1/bucket", nil},
		{"s3://:apassword@us-east-1/bucket", "s3://:apassword@us-east-1/bucket", nil},
		{"not a URL", "", ErrMissingScheme},
	}

	for _, tt := range tests {
		got, err := encodeURLCreds(tt.in)
		if err != tt.err {
			t.Errorf("expected %+v, got %+v", tt.err, err)
		}
		if got != tt.want {
			t.Errorf("expected %+v, got %+v", tt.want, got)
		}
	}
}
