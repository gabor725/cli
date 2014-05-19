package resources

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func cf_i18n_test_fixtures_en_main_en_us_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x00, 0x82, 0x6a, 0x30, 0x09, 0x02, 0x4a, 0x99, 0x29, 0x4a, 0x56,
		0x0a, 0x4a, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x0a, 0xe5, 0xf9, 0x45, 0x39,
		0x29, 0x8a, 0x4a, 0x3a, 0x08, 0xd9, 0x92, 0xa2, 0xc4, 0xbc, 0xe2, 0x9c,
		0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x0c, 0x65, 0x60, 0x55, 0xb5, 0x3a, 0xf8,
		0x0c, 0xac, 0xae, 0xd6, 0xf3, 0x4b, 0xcc, 0x4d, 0xad, 0xad, 0x25, 0x6c,
		0x28, 0x92, 0x52, 0xbc, 0x06, 0xfb, 0xf8, 0xfb, 0xe0, 0x31, 0xcb, 0x27,
		0xb1, 0x34, 0x3d, 0x43, 0x21, 0xbf, 0xb4, 0x44, 0x21, 0x27, 0xbf, 0x34,
		0x05, 0x6a, 0x12, 0x57, 0x2c, 0x17, 0x20, 0x00, 0x00, 0xff, 0xff, 0x14,
		0x0f, 0x08, 0x75, 0xf7, 0x00, 0x00, 0x00,
	},
		"cf/i18n/test_fixtures/en/main/en_US.all.json",
	)
}

func cf_i18n_test_fixtures_fr_main_fr_fr_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x00, 0x82, 0x6a, 0x30, 0x09, 0x02, 0x4a, 0x99, 0x29, 0x4a, 0x56,
		0x0a, 0x4a, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x0a, 0xe5, 0xf9, 0x45, 0x39,
		0x29, 0x8a, 0x4a, 0x3a, 0x08, 0xd9, 0x92, 0xa2, 0xc4, 0xbc, 0xe2, 0x9c,
		0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x90, 0xb2, 0x98, 0x52, 0x03, 0x83, 0x64,
		0x03, 0xa0, 0xca, 0x9c, 0x54, 0x85, 0xdc, 0xfc, 0xbc, 0x94, 0x54, 0x45,
		0x25, 0xb0, 0xda, 0x5a, 0x1d, 0x7c, 0xc6, 0x56, 0x57, 0xeb, 0xf9, 0x25,
		0xe6, 0xa6, 0xd6, 0xd6, 0x12, 0x65, 0x34, 0x92, 0x6a, 0x88, 0xd9, 0x5c,
		0xb1, 0x5c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x5e, 0x83, 0x35,
		0xb6, 0x00, 0x00, 0x00,
	},
		"cf/i18n/test_fixtures/fr/main/fr_FR.all.json",
	)
}

func cf_i18n_resources_en_quota_en_us_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x57,
		0xdf, 0x4f, 0xdb, 0x30, 0x10, 0x7e, 0xe7, 0xaf, 0x38, 0x55, 0x42, 0xda,
		0x24, 0x12, 0xc1, 0xb4, 0x27, 0xde, 0x18, 0x14, 0x84, 0x44, 0xcb, 0xe8,
		0x8f, 0x4d, 0x13, 0x43, 0xc8, 0x24, 0xd7, 0xd6, 0x52, 0x62, 0x77, 0xb6,
		0xd3, 0x0e, 0xb1, 0xfc, 0xef, 0xb3, 0x9d, 0x0e, 0x5a, 0x06, 0x17, 0x7b,
		0x20, 0xb4, 0x4a, 0x40, 0x62, 0xee, 0xfb, 0xee, 0x7c, 0xf7, 0xf5, 0x7c,
		0xbe, 0xdc, 0x02, 0xfb, 0xb9, 0xf3, 0xbf, 0xdd, 0xa7, 0xc3, 0xf3, 0xce,
		0x3e, 0x74, 0x8e, 0x70, 0xc2, 0x05, 0x02, 0x03, 0x81, 0x4b, 0x50, 0xa8,
		0x65, 0xa5, 0x32, 0x84, 0x1f, 0x95, 0x34, 0xac, 0xb3, 0xf3, 0x60, 0x6c,
		0x14, 0x13, 0xba, 0x60, 0x86, 0x4b, 0xd1, 0x86, 0xf2, 0xa0, 0x7a, 0xe7,
		0x69, 0x77, 0x87, 0xc7, 0xd7, 0xfd, 0x83, 0x5e, 0x17, 0x32, 0x85, 0xcc,
		0x60, 0xe2, 0x11, 0x70, 0x31, 0x3e, 0x1f, 0x1d, 0xc0, 0x65, 0x52, 0x42,
		0xaf, 0xdb, 0x3b, 0x1f, 0x7c, 0xbb, 0xb2, 0xcf, 0x0a, 0x06, 0xe7, 0xe3,
		0x51, 0x77, 0xe8, 0x9e, 0x35, 0x0c, 0xbb, 0x83, 0x2f, 0xa7, 0x87, 0xdd,
		0xeb, 0xd3, 0xfe, 0x70, 0x74, 0xd0, 0x3f, 0x6c, 0x96, 0x13, 0x56, 0x14,
		0x72, 0x99, 0xcc, 0x19, 0xcf, 0x13, 0x8d, 0x6a, 0xc1, 0x33, 0x4c, 0xe6,
		0x85, 0x8d, 0xf3, 0x8a, 0x88, 0xfc, 0x8d, 0x02, 0x20, 0x93, 0x30, 0xb2,
		0x3e, 0x0b, 0x10, 0x55, 0x79, 0x83, 0x0a, 0xe4, 0x04, 0x56, 0x58, 0xe0,
		0x42, 0x1b, 0x26, 0x32, 0xd4, 0x44, 0xfc, 0xed, 0x58, 0x3a, 0xff, 0x4c,
		0xc0, 0x5c, 0xc9, 0x05, 0xd7, 0x96, 0xef, 0x01, 0xe4, 0x98, 0xdc, 0x36,
		0xee, 0xe9, 0xfc, 0x36, 0xa8, 0x34, 0x46, 0xf1, 0xd0, 0x21, 0xb9, 0x4a,
		0x70, 0x31, 0x6d, 0xe4, 0x03, 0x77, 0x77, 0xe9, 0x85, 0x7b, 0xe8, 0xb3,
		0x12, 0xeb, 0x1a, 0x98, 0x76, 0x2b, 0x63, 0xcb, 0x27, 0xfc, 0x42, 0x9a,
		0xa6, 0x54, 0x58, 0xd1, 0x5c, 0x01, 0x85, 0x62, 0xa5, 0xac, 0x84, 0x71,
		0x5b, 0x2b, 0xb1, 0x94, 0xea, 0x16, 0xde, 0x61, 0x3a, 0x4d, 0x61, 0x6f,
		0xf7, 0xc3, 0xc7, 0xde, 0x0e, 0xec, 0x9d, 0xd8, 0x9f, 0xdd, 0x93, 0xf7,
		0xad, 0x35, 0x0b, 0xa4, 0x89, 0x52, 0x8e, 0x92, 0x95, 0x89, 0x92, 0xcb,
		0x0a, 0x40, 0x3a, 0x39, 0x15, 0x0b, 0x56, 0xd8, 0x1a, 0xae, 0xc2, 0x2c,
		0x78, 0xc9, 0xcd, 0xbe, 0xcb, 0x5c, 0xcf, 0x2f, 0x9c, 0xb9, 0xf7, 0xba,
		0xfe, 0x2e, 0xec, 0x4a, 0x57, 0xa9, 0xba, 0x26, 0xdc, 0x47, 0x53, 0x91,
		0x81, 0xf9, 0x5a, 0x82, 0xef, 0x3d, 0xdc, 0x79, 0xf8, 0xab, 0xbe, 0x85,
		0x2d, 0x7f, 0x7e, 0x0b, 0xf8, 0x93, 0x6b, 0x43, 0x65, 0x25, 0x92, 0x88,
		0x0c, 0xea, 0x08, 0x0b, 0x34, 0xae, 0x15, 0xb6, 0xf7, 0xcc, 0x0d, 0xc3,
		0xa0, 0x36, 0x99, 0x7b, 0xcc, 0xa3, 0x2e, 0x35, 0x09, 0xe9, 0x6f, 0xcf,
		0x21, 0x49, 0xb7, 0xc7, 0xd2, 0xb5, 0x71, 0x0f, 0x75, 0x59, 0x59, 0x72,
		0x33, 0xb3, 0x7a, 0x81, 0x4c, 0x8a, 0x09, 0x57, 0x65, 0xe3, 0xe3, 0x79,
		0xd7, 0x21, 0xe8, 0x16, 0xe1, 0x65, 0x52, 0x29, 0xcc, 0x0c, 0x8c, 0x35,
		0x9b, 0x22, 0xa9, 0xab, 0x4d, 0xcb, 0xf6, 0x0a, 0xbd, 0x52, 0x83, 0x89,
		0xe7, 0x0a, 0x50, 0xf4, 0x23, 0x96, 0x5c, 0xda, 0x46, 0x2a, 0xa4, 0x69,
		0xe4, 0xd7, 0x2a, 0x63, 0x1a, 0x1d, 0x24, 0xb3, 0x35, 0x95, 0x04, 0x48,
		0x6b, 0xdd, 0x9a, 0xa4, 0x1f, 0xce, 0xe4, 0x72, 0x65, 0xcd, 0xc5, 0x44,
		0x12, 0xd4, 0x8f, 0x2d, 0x49, 0xda, 0x13, 0x34, 0xcf, 0x97, 0xc0, 0xc1,
		0x23, 0x6b, 0xfa, 0x6f, 0x7c, 0x64, 0x88, 0x4d, 0x7b, 0x23, 0x7c, 0xae,
		0x0c, 0x48, 0x92, 0x41, 0x5b, 0x73, 0x1f, 0x04, 0x34, 0xf3, 0x61, 0x73,
		0x10, 0x53, 0x34, 0xf7, 0x26, 0x24, 0xd1, 0xe7, 0x98, 0xf1, 0xe0, 0x09,
		0x63, 0x92, 0xfc, 0xcc, 0x4a, 0x15, 0xd8, 0x82, 0xf1, 0x82, 0xdd, 0x14,
		0x08, 0x95, 0xfb, 0x52, 0x37, 0xf5, 0xa0, 0xbc, 0x50, 0xa8, 0x70, 0xdd,
		0x93, 0x63, 0xce, 0xa6, 0x61, 0xb8, 0x2c, 0xb5, 0xd3, 0xcc, 0xb6, 0x8e,
		0x10, 0xde, 0x1a, 0x82, 0x74, 0xe3, 0x14, 0x48, 0x70, 0xfa, 0x7f, 0x93,
		0x04, 0xeb, 0x67, 0x31, 0x41, 0xb4, 0x61, 0x46, 0x12, 0xb6, 0x0e, 0x21,
		0x21, 0x43, 0x47, 0xcc, 0x0c, 0x1c, 0x39, 0xf3, 0x46, 0x0d, 0xb6, 0xd1,
		0xd3, 0xeb, 0x4a, 0x21, 0xd5, 0x3c, 0xa7, 0xef, 0x13, 0x02, 0xfa, 0xdd,
		0xaf, 0xde, 0xf6, 0xa5, 0xb7, 0x0b, 0xf8, 0x05, 0x49, 0x92, 0x73, 0xfd,
		0x92, 0xeb, 0xcf, 0x7f, 0x15, 0x2e, 0x99, 0xe0, 0xbe, 0xbd, 0x61, 0xb6,
		0x68, 0xfe, 0xde, 0xa4, 0xed, 0xea, 0xe3, 0x8e, 0xc5, 0x57, 0xba, 0xfd,
		0xc4, 0x51, 0x91, 0x81, 0x8d, 0x5d, 0x2d, 0x1e, 0x0e, 0xa0, 0xed, 0xa0,
		0xde, 0x41, 0x80, 0xda, 0x9d, 0xd9, 0x29, 0x54, 0x34, 0xf3, 0x81, 0x63,
		0x08, 0xbe, 0xf5, 0xb7, 0x63, 0xe9, 0x13, 0xa4, 0x40, 0xa6, 0x11, 0xb2,
		0x99, 0x94, 0xf6, 0x0f, 0xda, 0xf1, 0xd0, 0xde, 0x49, 0xbc, 0x2a, 0x40,
		0x2a, 0xf8, 0xa3, 0x90, 0x14, 0x3e, 0x49, 0x33, 0x83, 0x49, 0xc1, 0xa6,
		0x76, 0x4f, 0x0a, 0xfd, 0x2c, 0x33, 0x47, 0x65, 0xdb, 0x90, 0xc1, 0x1c,
		0x8c, 0x84, 0x1b, 0x9b, 0x55, 0xa6, 0xb5, 0x7d, 0xe1, 0x02, 0x2c, 0x09,
		0x68, 0x5b, 0x7b, 0x3b, 0x66, 0x96, 0x25, 0x13, 0x79, 0x0a, 0xd4, 0xb1,
		0xf4, 0x46, 0x11, 0x34, 0x69, 0xd8, 0xba, 0xda, 0xfa, 0x1d, 0x00, 0x00,
		0xff, 0xff, 0x3c, 0x8c, 0x92, 0x23, 0x75, 0x11, 0x00, 0x00,
	},
		"cf/i18n/resources/en/quota/en_US.all.json",
	)
}

func cf_i18n_resources_fr_quota_fr_fr_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x58,
		0xef, 0x4e, 0xe3, 0x46, 0x10, 0xff, 0x7e, 0x4f, 0x31, 0x42, 0x42, 0xf4,
		0x24, 0x12, 0xdd, 0x55, 0xfd, 0x74, 0xdf, 0x28, 0x18, 0x84, 0x44, 0x92,
		0x23, 0x7f, 0x5a, 0x55, 0x57, 0x14, 0x2d, 0xf1, 0x04, 0xb6, 0xb2, 0x77,
		0xdd, 0xdd, 0x75, 0x38, 0x4a, 0xf3, 0x00, 0x7d, 0x8b, 0x7e, 0xbc, 0x3c,
		0x87, 0x5f, 0xac, 0xb3, 0x6b, 0x27, 0x81, 0x90, 0x8c, 0x9d, 0xde, 0xa9,
		0x6a, 0x24, 0xc0, 0x36, 0x33, 0xbf, 0x99, 0xdd, 0xf9, 0xf9, 0x37, 0xb3,
		0xf9, 0xf4, 0x06, 0xe8, 0xf3, 0x14, 0x7e, 0xfb, 0xcf, 0x81, 0x8c, 0x0f,
		0x3e, 0xc0, 0xc1, 0x19, 0x4e, 0xa5, 0x42, 0x10, 0xa0, 0xf0, 0x01, 0x0c,
		0x5a, 0x9d, 0x9b, 0x09, 0xc2, 0xef, 0xb9, 0x76, 0xe2, 0xe0, 0x78, 0x6d,
		0xec, 0x8c, 0x50, 0x36, 0x11, 0x4e, 0x6a, 0xb5, 0xf2, 0x92, 0x06, 0x72,
		0x05, 0x4a, 0xe7, 0x33, 0x14, 0x95, 0x47, 0x70, 0x98, 0x1f, 0x6f, 0x0f,
		0x75, 0x7a, 0x3e, 0xee, 0x9e, 0x74, 0x22, 0x98, 0x18, 0x14, 0x0e, 0x5b,
		0xc1, 0x03, 0xae, 0x47, 0xbd, 0xe1, 0x09, 0x7c, 0x6a, 0xa5, 0xd0, 0x89,
		0x3a, 0xbd, 0xfe, 0x2f, 0x37, 0x74, 0x6d, 0xa0, 0xdf, 0x1b, 0x0d, 0xa3,
		0x81, 0xbf, 0xb6, 0x30, 0x88, 0xfa, 0x3f, 0x5d, 0x9e, 0x46, 0xe3, 0xcb,
		0xee, 0x60, 0x78, 0xd2, 0x3d, 0x2d, 0x1f, 0xb7, 0x44, 0x92, 0xe8, 0x87,
		0x56, 0x26, 0x64, 0xdc, 0xb2, 0x68, 0x66, 0x72, 0x82, 0xad, 0x2c, 0xa1,
		0x1c, 0x6f, 0x98, 0xac, 0xeb, 0x12, 0x28, 0xfe, 0xea, 0xf4, 0x2e, 0xfb,
		0xd1, 0xeb, 0x14, 0x96, 0xa1, 0xc7, 0x67, 0xd1, 0xb8, 0x4a, 0xa7, 0x2e,
		0x0b, 0x76, 0x27, 0x86, 0x14, 0x38, 0x01, 0x95, 0xa7, 0xb7, 0x68, 0x40,
		0x4f, 0xa1, 0xf2, 0x05, 0xa9, 0xac, 0x13, 0x6a, 0x82, 0x96, 0x59, 0x44,
		0x57, 0xa7, 0xb7, 0x54, 0x27, 0x70, 0x01, 0x23, 0x3e, 0x5a, 0xfa, 0x40,
		0x8c, 0x4b, 0x1c, 0x5b, 0x53, 0x07, 0xa1, 0x20, 0x33, 0x7a, 0x26, 0x2d,
		0x41, 0xae, 0x63, 0xfa, 0x44, 0xfc, 0x4a, 0x56, 0xd9, 0x84, 0x95, 0x30,
		0x99, 0x7c, 0xc4, 0xfc, 0xf3, 0x1a, 0x88, 0x56, 0x12, 0x13, 0xc8, 0x1a,
		0xee, 0x59, 0x3e, 0x90, 0x11, 0xad, 0x20, 0x29, 0x21, 0x29, 0xc8, 0x63,
		0xb1, 0xa8, 0x49, 0xd1, 0x57, 0x48, 0xaa, 0xbb, 0x92, 0x56, 0xf0, 0xf4,
		0xd4, 0xbe, 0xf6, 0x17, 0x5d, 0x91, 0xe2, 0x7c, 0x0e, 0xc2, 0xfa, 0x27,
		0x23, 0x42, 0x57, 0xe1, 0x41, 0xbb, 0xdd, 0xe6, 0xaa, 0x6e, 0x8a, 0x05,
		0xfe, 0xb1, 0x1d, 0xa9, 0x58, 0x50, 0xb2, 0xee, 0x35, 0x5a, 0x83, 0xea,
		0x89, 0x54, 0xe7, 0xe4, 0x4a, 0x9b, 0x96, 0x62, 0xaa, 0xcd, 0x23, 0x7c,
		0x87, 0xed, 0xbb, 0x36, 0xbc, 0x7f, 0xf7, 0xfd, 0x0f, 0x9d, 0x63, 0x78,
		0x7f, 0x41, 0x3f, 0xef, 0x2e, 0xde, 0x32, 0x79, 0x5d, 0xe7, 0x14, 0x5a,
		0xba, 0x62, 0x51, 0x96, 0x32, 0x14, 0x30, 0x2d, 0x16, 0xa9, 0x96, 0x06,
		0x9f, 0x83, 0x69, 0x8f, 0xa6, 0x03, 0x9c, 0x7e, 0xbb, 0x1f, 0xaf, 0x8c,
		0xce, 0x1d, 0x4b, 0xa6, 0x6d, 0x39, 0x54, 0x4e, 0x6c, 0xa0, 0x4b, 0x35,
		0x13, 0x09, 0x51, 0xa5, 0x5a, 0x7a, 0x22, 0x53, 0xe9, 0x3e, 0xf8, 0x6d,
		0xec, 0x84, 0x07, 0x57, 0xfe, 0x7e, 0x3e, 0xff, 0x55, 0xd1, 0x93, 0xc8,
		0x98, 0xf9, 0x9c, 0x49, 0x21, 0xd8, 0xbe, 0x5c, 0xbc, 0x2c, 0xd1, 0x59,
		0x40, 0x36, 0xbd, 0x50, 0x62, 0x28, 0x45, 0xca, 0xc7, 0x79, 0x45, 0xa0,
		0x84, 0xf8, 0x15, 0x3f, 0x02, 0x7e, 0x96, 0xd6, 0x71, 0xfb, 0x73, 0x25,
		0x20, 0x2e, 0x16, 0x2b, 0x98, 0x18, 0xb7, 0xd3, 0x28, 0xe0, 0xa0, 0x37,
		0xfd, 0xad, 0xf8, 0x9b, 0x4f, 0xed, 0x0c, 0x13, 0x74, 0x5e, 0x6f, 0xeb,
		0x24, 0x76, 0x90, 0x67, 0x46, 0xa6, 0x29, 0x86, 0xf7, 0x66, 0x0f, 0x75,
		0x8d, 0x43, 0x80, 0x0d, 0x71, 0x9b, 0x36, 0x91, 0xc5, 0x5d, 0x9e, 0x6c,
		0xd8, 0x73, 0x6d, 0x82, 0xf6, 0x90, 0xab, 0xdf, 0xa1, 0x07, 0xe9, 0xee,
		0x89, 0x3f, 0x30, 0xd1, 0x6a, 0x2a, 0x4d, 0x5a, 0xc6, 0xd8, 0x1d, 0x3a,
		0x78, 0xd3, 0x0a, 0x05, 0xd8, 0x3c, 0xcb, 0x48, 0xd6, 0x82, 0x24, 0x59,
		0xb2, 0x79, 0x89, 0x50, 0x43, 0xc6, 0x89, 0x36, 0x06, 0x27, 0x0e, 0x46,
		0x56, 0xdc, 0x21, 0x13, 0x6e, 0xe4, 0x64, 0x22, 0x6d, 0xb8, 0x85, 0x95,
		0x17, 0x36, 0x28, 0xd8, 0x37, 0x92, 0xa2, 0x41, 0x9e, 0x6e, 0x16, 0x15,
		0x0e, 0xed, 0x52, 0x85, 0x0e, 0x6d, 0xad, 0xf4, 0x5c, 0x6f, 0xcb, 0x21,
		0xd6, 0xa4, 0xb0, 0x4a, 0xbb, 0x92, 0x86, 0xec, 0xdb, 0xbe, 0xc5, 0x5b,
		0x1d, 0x15, 0x8b, 0x92, 0xbe, 0x59, 0x92, 0xd7, 0x75, 0x8e, 0x8a, 0x29,
		0xcf, 0x28, 0xc2, 0x75, 0xaa, 0x5e, 0x67, 0x7c, 0x7a, 0xfe, 0xc2, 0x98,
		0x45, 0x1f, 0xdc, 0xeb, 0x87, 0xca, 0x5a, 0xaa, 0xa9, 0x66, 0x90, 0x3b,
		0x5a, 0x39, 0x43, 0x9a, 0x9e, 0x1c, 0x79, 0xc3, 0x8a, 0x24, 0xab, 0xb7,
		0x93, 0x8f, 0x72, 0x81, 0x6e, 0x77, 0x39, 0x3d, 0xdc, 0x9e, 0x35, 0x1d,
		0x1a, 0x3f, 0x00, 0x99, 0x8d, 0x5c, 0x4a, 0xf4, 0xd0, 0xf6, 0x36, 0x5b,
		0xce, 0x97, 0x7f, 0xd1, 0x72, 0x4a, 0x15, 0xe4, 0x76, 0x04, 0x83, 0x74,
		0xf2, 0x28, 0xfd, 0xba, 0x76, 0xd0, 0x6f, 0x20, 0xfd, 0x83, 0xe5, 0x8c,
		0xc1, 0xd0, 0xbc, 0xd1, 0x18, 0xf2, 0x71, 0xaf, 0x59, 0xc3, 0x0f, 0x0e,
		0xeb, 0x89, 0x22, 0xcc, 0x10, 0x5f, 0x78, 0xfc, 0x2b, 0x22, 0x35, 0x88,
		0x99, 0x90, 0x89, 0xb8, 0xa5, 0xb7, 0x2d, 0xf7, 0xda, 0x50, 0xd6, 0x85,
		0x55, 0x7c, 0xc6, 0xab, 0xf9, 0xbb, 0xc1, 0x45, 0xd8, 0x30, 0x6c, 0x4e,
		0x56, 0xeb, 0x99, 0x59, 0x8a, 0xc4, 0x6e, 0xf4, 0x5d, 0x1e, 0x6c, 0x18,
		0xcf, 0x42, 0x06, 0x33, 0xfc, 0x9b, 0x05, 0x78, 0x3e, 0x07, 0x30, 0x40,
		0x2f, 0xcc, 0x58, 0xc0, 0xda, 0xc9, 0xa5, 0xc9, 0x94, 0xb2, 0xcf, 0x58,
		0xfd, 0xda, 0x96, 0x85, 0xde, 0x6b, 0x50, 0xde, 0x62, 0xdc, 0x88, 0x4a,
		0x79, 0x16, 0xf3, 0x07, 0x25, 0x05, 0xdd, 0xe8, 0xe7, 0x60, 0xfb, 0xb5,
		0xc7, 0x26, 0xf8, 0x13, 0x5a, 0xad, 0x98, 0x7a, 0xe3, 0x57, 0x1c, 0xab,
		0xfe, 0x57, 0xe9, 0xb2, 0x1b, 0xdc, 0xa5, 0xb3, 0x6e, 0x0d, 0xe7, 0x57,
		0x26, 0x75, 0x47, 0x29, 0xdf, 0x78, 0xbf, 0xc5, 0x69, 0x6a, 0x6f, 0x28,
		0x36, 0xb1, 0x91, 0xaf, 0xc5, 0xba, 0xcd, 0x1d, 0x36, 0xd2, 0x0e, 0xc6,
		0xa9, 0x3e, 0x18, 0xcd, 0xb4, 0xaa, 0x9c, 0x40, 0x3c, 0x42, 0xe3, 0x6f,
		0x12, 0xea, 0x7d, 0xf9, 0x26, 0x92, 0xa0, 0xb0, 0x08, 0x93, 0x7b, 0xad,
		0xe9, 0x0f, 0xd2, 0xf4, 0x49, 0xad, 0x38, 0xb0, 0x02, 0x34, 0x1d, 0x48,
		0x2b, 0x86, 0xb4, 0xe1, 0x47, 0xed, 0xee, 0x61, 0x9a, 0x88, 0x3b, 0x5a,
		0x13, 0x9d, 0x30, 0xc2, 0x4e, 0xa3, 0x21, 0x19, 0x72, 0x18, 0xd3, 0xd9,
		0x07, 0x6e, 0x7d, 0x4b, 0xb1, 0x96, 0x6e, 0xa4, 0x02, 0x02, 0xa1, 0x19,
		0x34, 0x25, 0x58, 0x9d, 0xa6, 0x42, 0xc5, 0x6d, 0x60, 0x3b, 0xd3, 0x7f,
		0x93, 0x41, 0xb9, 0x0d, 0x6f, 0x6e, 0xde, 0xfc, 0x13, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xec, 0x86, 0x06, 0xc9, 0x11, 0x00, 0x00,
	},
		"cf/i18n/resources/fr/quota/fr_FR.all.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"cf/i18n/test_fixtures/en/main/en_US.all.json": cf_i18n_test_fixtures_en_main_en_us_all_json,
	"cf/i18n/test_fixtures/fr/main/fr_FR.all.json": cf_i18n_test_fixtures_fr_main_fr_fr_all_json,
	"cf/i18n/resources/en/quota/en_US.all.json":    cf_i18n_resources_en_quota_en_us_all_json,
	"cf/i18n/resources/fr/quota/fr_FR.all.json":    cf_i18n_resources_fr_quota_fr_fr_all_json,
}
