package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestNewAPIPort(t *testing.T) {
	apiPort1 := NewAPIPort(0)
	require.Equal(t, APIPort(0), apiPort1)
	require.Equal(t, false, apiPort1.Valid())
	require.Equal(t, true, apiPort1.IsZero())

	apiPort2 := NewAPIPort(8000)
	require.Equal(t, APIPort(8000), apiPort2)
	require.Equal(t, true, apiPort2.Valid())
	require.Equal(t, false, apiPort2.IsZero())
}

func TestNodeDetails_UpdateDetails(t *testing.T) {
	tests := []struct {
		name    string
		details NodeDetails
		want    NodeDetails
	}{
		{
			"prices_per_gb is nil",
			NodeDetails{PricesPerGB: nil},
			NodeDetails{},
		}, {
			"prices_per_gb is empty",
			NodeDetails{PricesPerGB: TestCoinsEmpty},
			NodeDetails{},
		}, {
			"prices_per_gb is invalid",
			NodeDetails{PricesPerGB: TestCoinsInvalid},
			NodeDetails{},
		}, {
			"prices_per_gb is negative",
			NodeDetails{PricesPerGB: TestCoinsNeg},
			NodeDetails{},
		}, {
			"prices_per_gb is zero",
			NodeDetails{PricesPerGB: TestCoinsZero},
			NodeDetails{},
		}, {
			"prices_per_gb is positive",
			NodeDetails{PricesPerGB: TestCoinsPos},
			NodeDetails{PricesPerGB: TestCoinsPos},
		}, {
			"net_speed is empty",
			NodeDetails{NetSpeed: types.Bandwidth{}},
			NodeDetails{},
		}, {
			"net_speed is negative",
			NodeDetails{NetSpeed: types.NewBandwidth(TestUploadNeg, TestDownloadNeg)},
			NodeDetails{},
		}, {
			"net_speed is zero",
			NodeDetails{NetSpeed: types.NewBandwidth(TestUploadZero, TestDownloadZero)},
			NodeDetails{},
		}, {
			"net_speed is positive",
			NodeDetails{NetSpeed: types.NewBandwidth(TestUploadPos, TestDownloadPos)},
			NodeDetails{NetSpeed: types.NewBandwidth(TestUploadPos, TestDownloadPos)},
		}, {
			"api_port is zero",
			NodeDetails{APIPort: NewAPIPort(0)},
			NodeDetails{},
		}, {
			"api_port is positive",
			NodeDetails{APIPort: NewAPIPort(8000)},
			NodeDetails{APIPort: NewAPIPort(8000)},
		}, {
			"enc_method is empty",
			NodeDetails{EncMethod: ""},
			NodeDetails{},
		}, {
			"enc_method is valid",
			NodeDetails{EncMethod: TestEncMethod},
			NodeDetails{EncMethod: TestEncMethod},
		}, {
			"node_type is empty",
			NodeDetails{NodeType: ""},
			NodeDetails{},
		}, {
			"node_type is valid",
			NodeDetails{NodeType: TestNodeType},
			NodeDetails{NodeType: TestNodeType},
		}, {
			"version is empty",
			NodeDetails{Version: ""},
			NodeDetails{},
		}, {
			"version is valid",
			NodeDetails{Version: TestVersion},
			NodeDetails{Version: TestVersion},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node := NodeDetails{}
			if node.UpdateDetails(tc.details); !reflect.DeepEqual(node, tc.want) {
				t.Errorf("\ngot = %vwant = %v", node, tc.want)
			}
		})
	}
}

func TestNodeDetails_FindPricePerGB(t *testing.T) {
	var node NodeDetails
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = NodeDetails{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = NodeDetails{PricesPerGB: TestCoinsEmpty}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = NodeDetails{PricesPerGB: TestCoinsPos}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinPos)
}

func TestNodeDetails_CalculateBandwidth(t *testing.T) {
	node := NodeDetails{PricesPerGB: TestCoinsPos}

	b, err := node.CalculateBandwidth(TestCoinNil)
	require.Equal(t, err, ErrorInvalidPriceDenom())
	require.Equal(t, b, types.Bandwidth{})

	b, err = node.CalculateBandwidth(TestCoinZero)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthZero)

	b, err = node.CalculateBandwidth(TestCoinNeg)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthNeg)

	b, err = node.CalculateBandwidth(TestCoinPos)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthPos)
}
