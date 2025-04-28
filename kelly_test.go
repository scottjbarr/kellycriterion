package kellycriteron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	type args struct {
		winRate float64
		reward  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr string
	}{
		{
			name: "example",
			args: args{
				winRate: 0.55,
				reward:  2.0,
			},
			want: 0.325,
		},
		{
			name: "error_winRate_zero",
			args: args{
				winRate: 0.0,
				reward:  2.0,
			},
			wantErr: "winRate must be >= 0 and <= 1.0",
		},
		{
			name: "error_reward_zero",
			args: args{
				winRate: 0.55,
				reward:  0.0,
			},
			wantErr: "reward must be > 0",
		},
		{
			name: "error_both_empty",
			args: args{
				winRate: 0.0,
				reward:  0.0,
			},
			wantErr: "winRate must be >= 0 and <= 1.0\nreward must be > 0",
		},
		{
			name: "error_winRate_too_large",
			args: args{
				winRate: 1.01,
				reward:  2.0,
			},
			wantErr: "winRate must be >= 0 and <= 1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.winRate, tt.args.reward)
			if len(tt.wantErr) > 0 {
				require.Error(t, err)
				require.Equal(t, tt.wantErr, err.Error())
				return
			}

			require.NoError(t, err)
			require.InEpsilon(t, tt.want, got, 0.0001)
		})
	}
}
