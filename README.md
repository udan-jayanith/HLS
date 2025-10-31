# HLS
HLS Go module implements [HTTP Live Streaming](https://datatracker.ietf.org/doc/html/rfc8216) interface for Go. HLS can encode and decode HTTP Live Streams and also provide a tokenizer and a serializer for low level access. This HLS module does not serve HTTP live streams and it's users responsibility to serve HTTP live streams. But HLS provides helper methods to server HTTP live streams.

# Examples

# Libraries that might help
## Media encoding and decoding
[mp4ff](https://github.com/Eyevinn/mp4ff)
Library and tools for working with MP4 files containing video, audio, subtitles, or metadata. The focus is on fragmented files. Includes mp4ff-info, mp4ff-encrypt, mp4ff-decrypt and other tools.

[Go media](https://github.com/yapingcat/gomedia)
Golang library for rtmp, mpeg-ts,mpeg-ps,flv,mp4,ogg,rtsp

[aac-go](https://github.com/gen2brain/aac-go)
Go bindings for vo-aacenc

## Crypto
[aes](https://pkg.go.dev/crypto/aes)
Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.


