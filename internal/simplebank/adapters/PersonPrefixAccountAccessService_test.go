package adapters

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/model"
	"testing"
)

func Test_personPrefixAccountAccessService_PersonHasAccessToAccount(t *testing.T) {
	tests := []struct {
		name    string
		person  model.PersonID
		account model.AccountID
		want    bool
		wantErr bool
	}{
		{
			name:    "Should be true when identical",
			person:  "abcd",
			account: "abcd",
			want:    true,
			wantErr: false,
		},

		{
			name:    "Should be false (case sensitive)",
			person:  "abcd",
			account: "ABCD",
			want:    false,
			wantErr: false,
		},

		{
			name:    "Should be true when is prefix",
			person:  "abcd",
			account: "abcd01",
			want:    true,
			wantErr: false,
		},

		{
			name:    "Should be true when is prefix",
			person:  "abcd01",
			account: "abcd",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := personPrefixAccountAccessService{}
			got, err := p.PersonHasAccessToAccount(context.Background(), tt.person, tt.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("PersonHasAccessToAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PersonHasAccessToAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
