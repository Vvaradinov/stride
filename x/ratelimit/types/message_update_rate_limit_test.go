package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Stride-Labs/stride/v4/app/apptesting"
	"github.com/Stride-Labs/stride/v4/x/ratelimit/types"
)

func TestMsgUpdateRateLimit(t *testing.T) {
	apptesting.SetupConfig()
	validAddr, invalidAddr := apptesting.GenerateTestAddrs()

	validDenom := "denom"
	validChannelId := "channel-0"
	validMaxPercentSend := uint64(10)
	validMaxPercentRecv := uint64(10)
	validDurationHours := uint64(60)

	tests := []struct {
		name string
		msg  types.MsgUpdateRateLimit
		err  string
	}{
		{
			name: "successful msg",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  validDurationHours,
			},
		},
		{
			name: "invalid creator",
			msg: types.MsgUpdateRateLimit{
				Creator:        invalidAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  validDurationHours,
			},
			err: "invalid creator address",
		},
		{
			name: "invalid denom",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          "",
				ChannelId:      validChannelId,
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  validDurationHours,
			},
			err: "invalid denom",
		},
		{
			name: "invalid channel-id",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      "chan-0",
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  validDurationHours,
			},
			err: "invalid channel-id",
		},
		{
			name: "invalid send percent",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: 101,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  validDurationHours,
			},
			err: "percent must be between 0 and 100",
		},
		{
			name: "invalid receive percent",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: 101,
				DurationHours:  validDurationHours,
			},
			err: "percent must be between 0 and 100",
		},
		{
			name: "invalid send and receive percent",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: 0,
				MaxPercentRecv: 0,
				DurationHours:  validDurationHours,
			},
			err: "either the max send or max receive threshold must be greater than 0",
		},
		{
			name: "invalid duration",
			msg: types.MsgUpdateRateLimit{
				Creator:        validAddr,
				Denom:          validDenom,
				ChannelId:      validChannelId,
				MaxPercentSend: validMaxPercentSend,
				MaxPercentRecv: validMaxPercentRecv,
				DurationHours:  0,
			},
			err: "duration can not be zero",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.err == "" {
				require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
				require.Equal(t, test.msg.Route(), types.RouterKey)
				require.Equal(t, test.msg.Type(), "update_rate_limit")

				signers := test.msg.GetSigners()
				require.Equal(t, len(signers), 1)
				require.Equal(t, signers[0].String(), validAddr)

				require.Equal(t, test.msg.Denom, validDenom, "denom")
				require.Equal(t, test.msg.ChannelId, validChannelId, "channelId")
				require.Equal(t, test.msg.MaxPercentSend, validMaxPercentSend, "maxPercentSend")
				require.Equal(t, test.msg.MaxPercentRecv, validMaxPercentRecv, "maxPercentRecv")
				require.Equal(t, test.msg.DurationHours, validDurationHours, "durationHours")
			} else {
				require.ErrorContains(t, test.msg.ValidateBasic(), test.err, "test: %v", test.name)
			}
		})
	}
}