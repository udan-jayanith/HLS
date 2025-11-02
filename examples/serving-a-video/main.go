package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"

	"github.com/udan-jayanith/HLS"
)

// Returns the media-playlist.
func fullVideoMediaPlaylist() (string, error) {
	playlist := HLS.NewPlaylist()

	//SetHeader append tag EXTM3U and EXT_X_VERSION. 7 is the version.
	if err := playlist.SetHeader(7); err != nil {
		return "", err
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXT_X_TARGETDURATION,
			Value:   "14",
		}); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "11.266667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg000.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "13.766667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg001.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "7.166667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg002.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "8.533333",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg003.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "11.800000",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg004.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "7.433333",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg005.ts")); err != nil {
			return "", err
		}
	}

	{
		playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXT_X_ENDLIST,
		})
	}
	//	Close must be called before reading the playlist
	playlist.Close()

	var builder strings.Builder
	rd := bufio.NewReader(&playlist)
	rd.WriteTo(&builder)

	return builder.String(), nil
}

func main() {
	mediaPlaylist, err := fullVideoMediaPlaylist()
	if err != nil {
		log.Fatal(err)
	}

	// This function serves video segments.
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCrossOrigin(w)
		handler := http.StripPrefix("/", http.FileServer(http.Dir("./video-fragments")))
		handler.ServeHTTP(w, r)
	}))

	// This function serves a trailer to a movie in a media-playlist file.
	http.HandleFunc("/full-video.m3u8", func(w http.ResponseWriter, r *http.Request) {
		setCrossOrigin(w)
		w.Write([]byte(mediaPlaylist))
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func setCrossOrigin(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}
