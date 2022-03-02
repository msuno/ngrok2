package assets

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

var _assets_server_tls_snakeoil_crt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x94\xc9\xaa\x83\x4e\x1e\x85\xf7\x3e\x45\xef\x43\x13\x35\xc6\x61\x59\x96\x73\x52\xce\xf3\xce\xd9\x98\xc4\x31\x5a\xc6\xa7\x6f\xee\x6d\xf8\x73\xa1\xbb\x96\x87\xfa\xe0\xe3\x77\xe0\xfc\xfb\xe7\x89\xb2\xaa\x9b\xff\x82\xb2\xeb\xeb\x8a\x0e\x81\x2f\xff\xa6\x04\xd2\x75\xd9\x97\x20\x04\x9d\xd7\x00\xac\x8b\xa0\xd1\x0d\x60\xb3\xfa\xb9\x7e\xf4\xd2\xf6\x2a\x10\x20\x55\xe8\x4d\xaa\xa7\xe7\x17\xc9\x91\x45\x88\x03\x80\x64\x66\x87\x07\x30\xc4\xc6\x0c\x09\x11\x24\x3e\x78\x9a\x16\x72\x16\x0c\x9d\x44\x0a\x1d\x47\x97\x80\xa1\x45\x9d\x7c\x47\xe0\xa9\x02\x2a\x90\x45\x8c\xa1\x4b\x0d\xbb\xd4\x01\xf4\x03\x89\x60\x40\x62\x68\x0e\x69\xb4\x3f\x09\xe4\x06\x58\xc6\xbf\xa0\x24\x81\x3d\x48\x23\xa7\x71\x48\xb9\xf1\xdf\xe5\x37\xa7\x17\xac\xb5\x85\x89\x3a\x1d\xa3\x03\xec\xc8\x07\xd8\xf4\x01\x19\xfd\x64\x87\x8e\x51\xa7\x5f\x88\x7f\xc2\x4e\x74\xfe\x58\xa8\x32\x36\x24\xff\x8f\x05\x04\x18\xba\xa4\xf3\x8f\xba\x08\x0a\x04\x9e\x65\x46\x20\x87\xc1\x12\xf8\x85\x6e\x12\x08\xfd\x4c\x0d\x97\x54\x92\x63\x24\x06\xbf\x20\xc0\xd8\xca\xff\x6b\xb3\xe6\xb1\x49\xe5\x6f\x7e\x4d\x68\x06\x37\x8d\xfc\xf8\xb9\x0f\xf1\xf7\x40\xc0\x09\x00\x60\x74\x51\xc2\xe0\xe7\xc3\x0d\x0c\xba\x08\x1c\xf8\x99\x58\x4f\xa1\xaf\xd5\xe3\x18\x2f\xd6\xe6\xae\x93\x5f\xc6\xcf\xef\xae\x18\xb2\x71\x1d\x89\xfe\xa3\x27\xf7\x67\x5f\xbe\x19\x4d\xd0\xa8\xdb\x1e\x54\xa5\x87\xe2\xa9\x74\xab\xc3\x71\x39\x8a\x77\x4b\x6b\x83\xf7\xac\x84\xdd\xd2\xce\xad\x28\xb0\xd9\x3b\x2b\x78\xca\xea\x68\xbf\x56\xcf\x45\x4a\x04\x58\x98\x39\xa5\x68\x9e\xef\x39\xb5\xe1\xc9\xac\xd5\xaa\x75\x0d\x18\xbc\x3b\xc9\x38\x1e\xba\x47\xb2\x6a\x11\x2d\xde\x6d\x1c\xa7\xb4\xfb\xb4\x4c\x44\xd3\x6e\xcb\x16\x32\x53\xae\xbb\x78\xa7\x1a\x62\x9f\xa5\x55\x33\x4c\xe4\x17\xe3\xc8\xef\xaf\x75\xb5\x98\x85\x05\x78\x73\x71\x47\x6a\x70\xbd\x47\x67\x77\xe5\x5f\x62\x21\x7e\x7c\x6b\x36\x9f\xde\xc3\x52\x92\x3b\x5f\x31\xaf\xec\xb1\xb0\x7e\x9e\x11\x69\x6a\xc7\x02\xeb\x24\x39\x48\xa8\xb8\x7f\x90\x89\x62\xe2\x6b\x68\x99\x6c\xa3\x25\xd1\xe1\xe9\x0f\x6f\x82\xcf\xb2\xbd\x84\x27\x2b\x69\xc8\x94\x2a\x0e\x8b\x02\xec\x98\xb2\x6f\xcc\xb3\xc3\x9d\x26\x82\x93\x15\x73\x5c\x72\x7b\x52\x91\x93\xbc\x02\xd3\xdb\x4e\x76\x91\x87\x23\x93\xad\xf5\xa6\x20\xd9\xcd\xa2\x0e\xca\x26\xe9\x36\xdb\x89\xde\xbb\xcc\x16\xdb\x29\x68\x45\x17\x34\x48\x04\x40\xed\x88\xdb\x01\x46\x04\x8b\x9f\xca\x4a\xd9\x71\x1a\x24\xb2\x10\x3e\x86\x3f\x95\xb1\x50\x86\xc3\xfa\xff\xaa\x94\x9c\xc4\x20\x6e\x43\xaa\xb7\x5b\x61\x02\x47\xbe\x8b\x0e\x90\x9a\x46\x17\x81\x68\xb1\xe8\xfa\xe1\xe6\x21\x80\x77\xd3\x98\xe7\x87\x63\x4d\xcf\x02\x96\x4a\x92\x77\x1e\x92\x6c\x78\xff\x6a\xbd\x92\xb0\x30\x5d\x08\x1c\x3a\xc5\x62\xda\x5b\x43\xe5\x17\x8c\x74\x59\x25\xdb\x8e\x15\x95\x98\x1e\xf9\xf5\x78\x39\x6d\xff\x32\xd4\xef\x20\x44\x64\x7c\x0d\xcf\x75\xab\xc7\x9b\x7e\x52\x56\x66\xaa\xc3\x4f\xa6\x0b\xbd\x41\x7c\xfb\x53\x20\xef\x63\x49\xbf\x6a\xaa\x01\x61\x63\xd6\xc7\x37\x96\xe4\x9d\x07\x42\xbb\x34\xd0\x0b\x79\x7b\xa8\xa3\x54\xd0\x62\x88\xd0\xb4\xb7\xa0\xca\x99\xc8\x67\xae\x0c\xdc\x87\xdc\x04\x2d\xc1\x6d\xf0\x2c\xdd\x26\x9a\x8c\x0d\x5a\x90\x25\xc5\xe8\x69\xfa\xf0\x3d\x86\xbe\x5c\x8c\x50\xcd\x56\xce\xb7\xad\xf4\x3c\x36\xed\xb1\x79\x88\xaf\x3e\xc2\xab\x87\x53\x32\x8f\xe6\xee\x98\x52\x65\x33\x44\x94\x93\xd7\x7e\xd4\x1f\x96\xc3\xf9\x1a\x37\x7e\x5d\x5b\x5a\x46\x92\x54\xf4\x70\xf7\xdc\x6b\xe5\x2a\x31\xc7\x91\x74\xb5\x5c\xef\x40\xb6\x6f\x6a\x5f\xae\xba\x9a\xa7\xd9\xa0\x38\x14\xfe\xd8\x7d\x43\x20\xee\x65\x05\xbd\xf6\xf0\xfa\x5a\xf9\x56\xc7\xf9\x19\x8f\xb8\x08\x5e\x6c\x9e\x45\x60\xf0\xcb\x6d\x57\xf3\x4f\x14\x4a\xa8\x19\xf3\xac\xee\x48\xdd\x8a\x5d\xa9\x8f\xb8\x82\x9d\xf9\x28\x3b\x29\x02\x01\xa4\xcc\xbe\xbd\x0f\x9d\x94\xda\x5c\xc6\x37\x91\x2c\xca\x14\x9c\xb5\x1c\x1b\xf3\xc8\xb6\x06\x4a\x03\x9e\x56\x4e\x5f\xd1\x5b\xbe\xf8\x44\x9a\x21\xf0\x17\xe6\xf2\xf9\x9c\xdc\xaa\xb9\x6e\x64\x57\x13\xf8\x35\x5c\x43\xbe\x8b\x21\x35\x07\x74\x59\xd0\xb5\x28\x37\x5f\x75\x17\x46\xa3\x34\xf9\x36\xde\x6d\x46\x23\xb7\x40\xb3\x2d\xde\xce\x2c\xaf\xbc\x46\x57\xfe\x0c\xdd\xec\xdd\x69\x4c\x79\xf5\x8e\x3b\xe1\xa5\xf6\xa9\xe5\xbb\x9a\x51\x23\x16\x71\x93\x50\x27\xd3\xbb\x5a\x56\x3b\xfe\x66\x52\xbc\xf2\x46\xec\x1b\xe3\xec\xce\xb1\xc3\x7e\xc9\xc9\xd9\xb9\xb9\xf8\x5a\x93\x8a\xb6\xcf\x0b\xf1\x25\x28\x7c\xe2\xd3\xbb\xe3\x53\xc1\x9b\xf5\x60\x3a\x7f\x0e\x61\x8c\x45\xa3\x81\x24\xf0\x95\xae\x5f\xfc\xb5\x6b\x3c\x61\x50\x05\x95\x1a\xea\x9c\xf4\x8f\x93\x05\x72\xaa\xc2\x0a\x05\x6d\x72\xbc\x37\x87\xa6\x10\x07\x9b\x2d\x34\x43\x4d\xcc\x50\x08\xf1\xb5\xaa\x0a\x05\x9d\xbe\xd2\x32\xa9\xc7\x22\x0c\x6e\x34\x08\xa7\x08\xd4\xba\x52\x76\x37\xaf\x77\x96\x7c\x4a\xb3\x63\x89\x62\x7c\x7b\x70\x8f\xa5\xf5\x79\x9f\xf8\x1d\x77\xd9\x94\xfe\x77\xf0\xff\x13\x00\x00\xff\xff\x33\xc5\xa3\xed\x0d\x06\x00\x00")

func assets_server_tls_snakeoil_crt() ([]byte, error) {
	return bindata_read(
		_assets_server_tls_snakeoil_crt,
		"assets/server/tls/snakeoil.crt",
	)
}

var _assets_server_tls_snakeoil_key = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd5\x37\x12\xab\x4a\x02\x85\xe1\x9c\x55\xdc\x9c\x9a\x02\x21\x6c\xf0\x82\x06\x21\x68\xe1\x11\x3e\xc3\x09\x84\x69\xbc\x11\xab\x9f\x9a\x3b\xe9\x3b\xe9\x49\xbe\xec\xff\xcf\xff\x26\xca\x0a\x34\xff\xb8\x6f\xf0\xc7\x76\x61\x00\x3c\xf9\x8f\x26\xc7\x7f\x1f\xcc\x80\x50\x1e\x01\x14\x01\xd0\x24\xe0\xc8\x60\x4e\xb7\xb6\x2e\xb6\x68\x58\xd8\xe2\x9a\xc9\x6c\x6a\xef\x01\x63\x2c\xee\xdb\x95\xca\x34\xe1\xde\x8a\xc4\xbc\xee\xc9\x26\x12\xae\xf0\x5e\x02\x15\x61\x7e\xf3\x64\x91\x1f\x2d\xa4\x5c\x0a\x1f\x39\xf2\xe6\xe3\x4b\x21\x67\xe0\x60\x4a\x3b\x9f\xb2\x1f\x7b\x64\x9b\xde\x20\xb4\xf7\x73\xbf\x2b\x81\x61\xa7\xf8\x19\xc1\xd7\x6b\xa2\xbc\x63\x6f\xa2\x13\x61\x83\xff\x76\xba\xd7\x70\xfc\x72\xfa\x2b\xaf\x56\x8d\x9e\xb3\xfc\x9d\xd2\xb4\x8f\xb9\xb8\x7c\xae\x6b\x1b\x97\x48\xb4\xd4\x6c\x71\xde\x42\x6c\xa4\x07\x57\xff\x3c\x4f\xbe\x6b\xe9\xc7\x48\xb2\xb9\xd9\x30\xdd\xaa\x0c\x9d\xcc\xa1\x20\x1e\xf3\xf7\xb6\x93\xd9\xfe\x72\x22\x90\xf9\xd7\x74\x25\xf2\xd0\xd4\xa1\x44\xa8\xdb\x2b\x1c\x66\xab\xa5\xa8\x3e\xf4\x6e\x44\xd9\x2a\xca\xa1\x98\x01\x43\xaf\x8a\xe8\x61\xb9\x15\x78\x4d\x39\x88\xd4\x73\x21\xbf\xb0\x9d\xaa\x51\x8d\x8b\xdb\xa7\xe9\xa1\xa9\x14\x91\x71\xad\x8e\x35\x85\xe5\xb8\xd8\xd6\x24\x09\x9d\xdd\x74\x38\x4e\x49\xa3\x19\xb4\xca\x2b\x78\xf8\x33\x81\x35\x77\xe5\x96\x96\xca\x8c\xb8\xd3\x13\xe5\xb0\x1b\x8e\xfa\x51\xc8\xb1\x4e\xac\x4b\x4c\x35\x47\x9c\x76\xd0\xf1\x1d\xf8\x00\x0e\x10\xc1\x00\x45\xa0\xac\x38\x7d\xbf\xcf\xe6\x51\xba\x4c\xab\xba\x18\x9f\x72\xb6\x61\x64\x8a\x3a\xdb\x1f\x1e\x51\xcd\x95\x90\x2d\x48\x26\x65\xe4\x23\xd6\xee\x08\xc1\xbc\x62\x8e\x01\x34\x85\x6f\x99\xd1\x5a\x38\x27\x98\xb6\xfb\x15\xb8\xd0\xbc\x85\xf1\xe7\xad\x4b\x18\xe2\xfc\x7b\xd8\xef\x14\x2e\xd8\x6e\x82\xaa\x31\xea\x9e\xaf\x3b\x0d\x8c\xc3\x04\x08\x65\xe6\xb1\xf7\xa0\x1f\xc3\xa5\x9b\x42\xfe\x10\x1d\x72\xd1\xc8\xe5\xd4\x99\x8a\xea\x83\xc6\x27\x0e\x9a\xc6\x62\x4f\x92\x25\x46\x34\x65\x68\xc5\xbe\x6a\xdf\x53\x30\x4e\x29\x7d\xc7\x03\xf5\xe7\x08\x62\x8b\x2e\xff\xab\xf4\x1a\xdf\x72\xba\xcd\xb8\x5d\x77\x76\xf6\x5d\x1f\xf5\x49\x5b\xf7\x43\xae\xce\xe0\x87\xc1\x4a\x7f\x6a\x9d\x47\x98\xd7\xe1\x10\xd3\xc8\x51\xd0\x26\x4d\x85\x68\x42\xd3\x52\xc8\x9f\x4c\x70\xd5\x87\x15\xa3\xcf\xd5\x78\x66\xc7\xc2\x67\x72\x6c\xe6\x4f\x5e\x55\x0d\x48\x19\xca\x23\xfd\xc4\xb2\xba\x9a\x80\x49\x11\x6c\xc2\x1f\x21\xb3\xbd\xf0\x53\xf8\xc1\xdc\xbd\xc8\xa9\xdb\xaa\x33\xa0\x7d\x9e\x3e\x5d\xb4\x12\x4e\x12\xe4\x3b\xf1\x36\xd3\x20\xd4\x75\x3c\x9a\x5a\xbc\xfd\x9e\xc2\x8b\xc4\x38\x22\x08\x1e\x86\x2f\x55\xb1\x0c\x28\x87\x19\x24\xc0\x70\x97\x57\x0e\xe5\x1c\x5c\x97\x20\xaf\xa0\x5d\xae\x22\x94\x02\x96\xff\xa6\x3a\x45\xf2\x0c\x55\x98\xae\x3f\xaa\x6d\x30\xc7\xac\x32\x8a\x58\x51\xf9\x46\x0d\xb7\x79\xd9\xe5\xa3\x77\x48\x73\xe8\x45\x87\x70\xda\xf2\xc5\x67\x47\x66\x89\x6c\xb0\x0e\x4b\x12\xbf\xeb\x9c\x2d\x61\x68\xb5\xbb\x7f\x76\xdf\xcd\x74\x5f\xc9\x02\xa3\xb8\x30\x22\xec\xcd\x7c\x1c\x5e\x28\x0a\x6d\xf4\xa7\xc5\x2a\x33\x53\x17\xc9\xc6\x67\x0b\xe7\xb1\x55\xc5\x87\x7a\x5b\x87\xc6\x16\xbf\xc6\x86\x4a\xe8\x52\x79\xfa\x4b\xa4\x13\x06\xea\xf2\x57\x7c\x99\x2e\x8f\x8d\x42\x91\x9c\xf7\x07\xe9\x0b\x69\xad\x6b\xd7\xe7\x52\xf5\x36\xdb\xd0\x95\x12\x2b\xaa\x79\x52\x2a\xe1\x72\x9c\xae\xe6\xc2\x95\x2e\x2b\x59\xd0\x2f\x41\x5b\x5a\x7f\xf8\x9d\x5a\x29\x3d\x4c\x51\xc3\x78\xed\xae\x6c\x02\xac\xe1\xe2\xb1\xcb\x79\xad\xfe\xb7\xce\x11\x21\xa2\xcc\x46\xf6\x49\xaa\x13\x67\x6b\xb8\xf3\x40\x0b\x9d\x3b\xb1\xc4\x74\xeb\x6d\x9b\x1f\x56\xb4\xdf\xd0\xe4\xa3\xa5\x3b\x00\x96\xe9\x8f\xa4\x8e\xac\xb5\x92\x6b\xbd\x58\xfd\x2a\x4f\x94\xc7\xb8\x4b\x9b\xa4\x5f\x35\xca\x3c\xe5\x4b\x49\x34\x3f\x19\x7f\xc9\xfd\xd7\xed\xf8\x71\xc2\x6b\x44\x8e\x06\x6d\xa6\xf3\x53\x69\x30\x50\x65\x44\x7b\x69\xb7\x97\x73\x23\xfa\x58\x2a\x24\xc3\xb5\xba\x41\x98\xea\x74\xc8\xf2\x12\xe9\x85\xab\x94\xe4\x49\xbb\x2a\xa1\xfc\x9a\x9e\xe6\xa6\xe8\x46\x31\xed\xb2\x04\x62\xbb\x19\x9d\x75\x61\x5c\xba\x57\x59\x64\x35\xac\xb2\x0c\x10\xb7\x57\x4f\x6f\xcf\x84\x9e\x7e\x61\x76\x4c\x96\xde\xe5\x3f\x9d\x0c\xdb\xd4\xa6\x16\x9d\xf8\xd0\xbd\xa5\x92\x00\x36\xcb\x39\x68\x92\xfc\xb3\x99\x6e\x2e\xb1\xe7\xee\xc4\x72\xd4\x92\xe3\x08\x52\x46\x43\x5c\xcb\xa9\x3f\x86\xfb\x3f\x79\xdd\x83\x84\xa0\xa2\x31\xd5\x84\x31\x42\xd1\x2b\x67\xbb\x0c\x36\x8d\x88\xbf\x4e\xb6\xf0\x27\x8f\x82\xc9\x85\xa1\x98\x2b\xa9\xa6\x7b\xe7\x1a\xb5\x50\x7b\xc7\x54\xa4\x7a\x68\xfa\xcc\x3c\xc4\x80\x18\xf2\xfa\x78\x12\xdd\xa9\x24\xa1\xf8\x4e\xea\x8f\x33\x1e\x9d\x4c\xff\x08\x86\x91\x35\xe8\x0e\x3d\xe9\x22\x01\xd3\xec\xc2\x42\x3b\x2a\x17\x85\x93\x04\xf6\xb9\xce\x1d\xa2\x82\xd4\xa4\xd8\xcd\x6f\xcf\xab\x33\x7a\xbe\x07\xf2\x94\xc7\x2f\x70\xed\x39\xc9\x0a\xfc\x69\x49\xb1\x78\x0a\x23\x3d\xa2\xd6\x50\x55\x8c\xdc\xbb\x39\x6c\x72\xa9\x8a\x45\x76\x39\xee\x71\x30\x38\x60\x54\xbe\x13\x92\xf7\x9b\x75\x88\x78\x2e\x03\xf0\x5c\xf6\x1e\x4f\x26\x32\xb8\xc9\xc4\xcb\xb5\x7e\x89\xc4\x76\x9d\xd3\x36\x0c\xeb\xf3\x58\x75\x3e\xec\xe1\xb6\x04\x70\xd1\xe7\x50\xcf\xc7\x8c\xb7\x24\x63\x6a\xb2\x1f\xed\x7e\x04\x60\x69\x48\x6a\xc8\x47\x9a\x1c\x0f\xc2\x3f\x4f\xb3\x19\xf0\x37\x29\x2c\x33\x2d\x9b\x80\xd9\x61\x97\x06\x18\xab\x6d\xca\xec\x9b\xbd\x2e\xf1\x9e\xca\x04\x96\x2d\xc4\x5d\x9e\xcc\xdf\xa2\xad\x95\xca\x7d\x02\x23\x7c\x42\xd9\xa0\x87\x7b\x1b\x18\x71\x29\x2a\xf0\xc4\x2b\x13\xfc\xf3\x0f\xf6\x37\x2b\xb2\xf9\xf8\xf7\xdc\xfc\x37\x00\x00\xff\xff\x24\xa8\x9a\x2e\x8f\x06\x00\x00")

func assets_server_tls_snakeoil_key() ([]byte, error) {
	return bindata_read(
		_assets_server_tls_snakeoil_key,
		"assets/server/tls/snakeoil.key",
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
	"assets/server/tls/snakeoil.crt": assets_server_tls_snakeoil_crt,
	"assets/server/tls/snakeoil.key": assets_server_tls_snakeoil_key,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"server": &_bintree_t{nil, map[string]*_bintree_t{
			"tls": &_bintree_t{nil, map[string]*_bintree_t{
				"snakeoil.crt": &_bintree_t{assets_server_tls_snakeoil_crt, map[string]*_bintree_t{
				}},
				"snakeoil.key": &_bintree_t{assets_server_tls_snakeoil_key, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}
