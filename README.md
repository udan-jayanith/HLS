# [HLS](https://pkg.go.dev/github.com/udan-jayanith/HLS)
HLS Go module implements [HTTP Live Streaming](https://datatracker.ietf.org/doc/html/rfc8216) interface for Go. HLS can encode and decode HTTP Live Streams and also provide a tokenizer and a serializer for low level access. This HLS module does not serve HTTP live streams and it's users responsibility to serve HTTP live streams. But HLS provides helper methods to server HTTP live streams.

# Examples

# Libraries that might help
## Media encoding and decoding
**[mp4ff](https://github.com/Eyevinn/mp4ff)**

Library and tools for working with MP4 files containing video, audio, subtitles, or metadata. The focus is on fragmented files. Includes mp4ff-info, mp4ff-encrypt, mp4ff-decrypt and other tools.

**Codec supports**

| Type| Codec | Sample Entry | Config Box | Other Boxes |
| ----- | ----| ---- | ---- | ---- |
| Video | AVC/H.264 | avc1, avc3 | avcC | btrt, pasp, colr |
| Video | HEVC/H.265 | hvc1, hev1 | hvcC | btrt, pasp, colr |
| Video | AV1 | av01 | av1C | btrt, pasp, colr |
| Video | AVS3 | avs3 | av3c | btrt, pasp, colr |
| Video | VP8/VP9 | vp08, vp09 | vpcC | btrt, pasp, colr |
| Video | VVC/H.266 | vvc1, vvi1 | vvcC | btrt, pasp, colr |
| Video | Encrypted | encv | sinf | btrt |
| Audio | AAC | mp4a | esds | btrt |
| Audio | AC-3 | ac-3 | dac3 | btrt |
| Audio | E-AC-3 | ec-3 | dec3 | btrt |
| Audio | AC-4 | ac-4 | dac4 | btrt |
| Audio | Opus | Opus | dOps | btrt |
| Audio | MPEG-H 3D Audio | mha1, mha2, mhm1, mhm2 | mhaC | btrt |
| Audio | Encrypted | enca | sinf | btrt |
| Subtitles | WebVTT | wvtt | vttC, vlab | vttc, vtte, vtta, vsid, ctim, iden, sttg, payl, btrt |
| Subtitles | TTML | stpp | - | btrt |
| Subtitles | Generic | evte | - | btrt |


[Go media](https://github.com/yapingcat/gomedia)

Golang library for rtmp, mpeg-ts,mpeg-ps,flv,mp4,ogg,rtsp

[aac-go](https://github.com/gen2brain/aac-go)

Go bindings for vo-aacenc

## Crypto
[aes](https://pkg.go.dev/crypto/aes)

Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.


