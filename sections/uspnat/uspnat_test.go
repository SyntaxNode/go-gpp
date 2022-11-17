package uspnat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type uspnatTestData struct {
	description string
	gppString   string
	expected    USPNAT
}

func TestUSPNAT(t *testing.T) {
	testData := []uspnatTestData{
		{
			description: "Test 1",
			gppString:   "DSJgmkoZJSg",
			/*
				000011 01 00 10 00 10 01 10 00 00 100110100100101000011001 0010 01 01 00 10 1
			*/
			expected: USPNAT{
				CoreSegment: USPNATCoreSegment{
					Version:                             3,
					SharingNotice:                       1,
					SaleOptOutNotice:                    0,
					SharingOptOutNotice:                 2,
					TargetedAdvertisingOptOutNotice:     0,
					SensitiveDataProcessingOptOutNotice: 2,
					SensitiveDataLimitUseNotice:         1,
					SaleOptOut:                          2,
					SharingOptOut:                       0,
					TargetedAdvertisingOptOut:           0,
					SensitiveDataProcessing: []byte{
						2, 1, 2, 2, 1, 0, 2, 2, 0, 1, 2, 1,
					},
					KnownChildSensitiveDataConsents: []byte{
						0, 2,
					},
					PersonalDataConsents:    1,
					MspaCoveredTransaction:  1,
					MspaOptOutOptionMode:    0,
					MspaServiceProviderMode: 2,
				},
				GPCSegment: USPNATGPCSegment{
					Gpc: 1,
				},
				SectionID: 7,
				Value:     "DSJgmkoZJSg",
			},
		},
	}

	for _, test := range testData {
		result, err := NewUSPNAT(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
