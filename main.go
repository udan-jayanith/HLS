//HLS Go module implements HTTP Live Streaming interface for Go.
//HLS can encode and decode HTTP Live Streams and also provide a tokenizer and a serializer for low level access.
//This HLS module does not serve HTTP live streams and it's users responsibility to serve HTTP live streams.
//But HLS provides helper methods to server HTTP live streams.

package HLS

//Playlist tags specify either global parameters of the Playlist or
//information about the Media Segments or Media Playlists that appear
//after them.
type PlaylistTag = string

//The section below this comment is AI generated and most of that is unchecked.
const (
	// Basic Tags
	// These tags are allowed in both Media Playlists and Master Playlists.

	// The EXTM3U tag indicates that the file is an Extended M3U [M3U]
	// Playlist file.  It MUST be the first line of every Media Playlist and
	// every Master Playlist.  Its format is:
	//	#EXTM3U
	//
	//[4.3.1.1]
	//
	//[4.3.1.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.1
	EXTM3U PlaylistTag = "EXTM3U"

	// The EXT-X-VERSION tag indicates the compatibility version of the
	// Playlist file, its associated media, and its server.
	//	#EXT-X-VERSION:<n>
	//
	//[4.3.1.2]
	//
	//[4.3.1.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.2
	EXT_X_VERSION PlaylistTag = "EXT-X-VERSION"

	// Media Segment Tags
	// Tags that describe or apply to individual media segments. These
	// MUST NOT appear in a Master Playlist.

	// The EXTINF tag specifies the duration of a Media Segment. It applies
	// only to the next Media Segment. Format:
	//	#EXTINF:<duration>,[<title>]
	//
	//[4.3.2.1]
	//
	//[4.3.2.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.1
	EXTINF PlaylistTag = "EXTINF"

	// The EXT-X-BYTERANGE tag indicates that a Media Segment is a sub-range
	// of the resource identified by its URI. Applies only to the next URI.
	//	#EXT-X-BYTERANGE:<n>[@<o>]
	//
	//[4.3.2.2]
	//
	//[4.3.2.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.2
	EXT_X_BYTERANGE PlaylistTag = "EXT-X-BYTERANGE"

	// The EXT-X-DISCONTINUITY tag indicates a discontinuity between the
	// Media Segment that follows it and the one that preceded it.
	//	#EXT-X-DISCONTINUITY
	//
	//[4.3.2.3]
	//
	//[4.3.2.3]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.3
	EXT_X_DISCONTINUITY PlaylistTag = "EXT-X-DISCONTINUITY"

	// Media Segments MAY be encrypted. The EXT-X-KEY tag specifies how to
	// decrypt them. It applies to every Media Segment and to every Media
	// Initialization Section declared by an EXT-X-MAP tag that appears
	// between it and the next EXT-X-KEY tag in the Playlist file with the
	// same KEYFORMAT attribute (or the end of the Playlist file). Its format is:
	//	#EXT-X-KEY:<attribute-list>
	//
	//[4.3.2.4]
	//
	//[4.3.2.4]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.4
	EXT_X_KEY PlaylistTag = "EXT-X-KEY"

	// The EXT-X-MAP tag specifies how to obtain the Media Initialization
	// Section required to parse following Media Segments (fMP4 use).
	//	#EXT-X-MAP:<attribute-list>
	//
	//[4.3.2.5]
	//
	//[4.3.2.5]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.5
	EXT_X_MAP PlaylistTag = "EXT-X-MAP"

	// The EXT-X-PROGRAM-DATE-TIME tag associates the first sample of a
	// Media Segment with an absolute date/time. Applies only to the next
	// Media Segment.
	//	#EXT-X-PROGRAM-DATE-TIME:<date-time-msec>
	//
	//[4.3.2.6]
	//
	//[4.3.2.6]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.6
	EXT_X_PROGRAM_DATE_TIME PlaylistTag = "EXT-X-PROGRAM-DATE-TIME"

	// The EXT-X-DATERANGE tag associates a Date Range with attribute/value
	// pairs (e.g., for timed metadata, ads, blackout periods).
	//	#EXT-X-DATERANGE:<attribute-list>
	//
	//[4.3.2.7]
	//
	//[4.3.2.7]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.7
	EXT_X_DATERANGE PlaylistTag = "EXT-X-DATERANGE"

	// Media Playlist Tags
	// Global parameters that describe a Media Playlist. At most one of each
	// type may appear in a Media Playlist unless the spec explicitly
	// permits multiple instances.

	// The EXT-X-TARGETDURATION tag specifies the maximum Media Segment
	// duration. Format:
	//	#EXT-X-TARGETDURATION:<s>
	//
	//[4.3.3.1]
	//
	//[4.3.3.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.1
	EXT_X_TARGETDURATION PlaylistTag = "EXT-X-TARGETDURATION"

	// The EXT-X-MEDIA-SEQUENCE tag specifies the Media Sequence Number of
	// the first Media Segment that appears in the Playlist.
	//	#EXT-X-MEDIA-SEQUENCE:<number>
	//
	//[4.3.3.2]
	//
	//[4.3.3.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.2
	EXT_X_MEDIA_SEQUENCE PlaylistTag = "EXT-X-MEDIA-SEQUENCE"

	// The EXT-X-DISCONTINUITY-SEQUENCE tag specifies the value of the
	// discontinuity sequence number for the first Media Segment in the
	// Playlist.
	//	#EXT-X-DISCONTINUITY-SEQUENCE:<number>
	//
	//[4.3.3.3]
	//
	//[4.3.3.3]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.3
	EXT_X_DISCONTINUITY_SEQUENCE PlaylistTag = "EXT-X-DISCONTINUITY-SEQUENCE"

	// The EXT-X-ENDLIST tag indicates no more Media Segments will be
	// added to the Media Playlist (on-demand playback end marker).
	//	#EXT-X-ENDLIST
	//
	//[4.3.3.4]
	//
	//[4.3.3.4]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.4
	EXT_X_ENDLIST PlaylistTag = "EXT-X-ENDLIST"

	// The EXT-X-PLAYLIST-TYPE tag indicates whether the Playlist is "VOD"
	// or "EVENT".
	//	#EXT-X-PLAYLIST-TYPE:<type>
	//
	//[4.3.3.5]
	//
	//[4.3.3.5]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.5
	EXT_X_PLAYLIST_TYPE PlaylistTag = "EXT-X-PLAYLIST-TYPE"

	// The EXT-X-I-FRAMES-ONLY tag indicates the Media Playlist contains
	// only I-frame index records (used for trick play / thumbnails).
	//	#EXT-X-I-FRAMES-ONLY
	//
	//[4.3.3.6]
	//
	//[4.3.3.6]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3.6
	EXT_X_I_FRAMES_ONLY PlaylistTag = "EXT-X-I-FRAMES-ONLY"

	// The EXT-X-START tag indicates a preferred point to start playback
	// within the Playlist.
	//	#EXT-X-START:<attribute-list>
	//
	//[4.3.5.2]
	//
	//[4.3.5.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.5.2
	EXT_X_START PlaylistTag = "EXT-X-START"

	// Master / Multivariant Playlist Tags
	// Tags used in Master Playlists to describe Renditions and Variant
	// Streams.

	// The EXT-X-STREAM-INF tag describes a Variant Stream in a Master
	// Playlist. It is followed by a URI.
	//	#EXT-X-STREAM-INF:<attribute-list>
	//
	//[4.3.4.2]
	//
	//[4.3.4.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.2
	EXT_X_STREAM_INF PlaylistTag = "EXT-X-STREAM-INF"

	// The EXT-X-MEDIA tag describes a Rendition (AUDIO, VIDEO, SUBTITLES,
	// CLOSED-CAPTIONS) in a Master Playlist.
	//	#EXT-X-MEDIA:<attribute-list>
	//
	//[4.3.4.1]
	//
	//[4.3.4.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.1
	EXT_X_MEDIA PlaylistTag = "EXT-X-MEDIA"

	// The EXT-X-I-FRAME-STREAM-INF tag describes an I-frame-only Media
	// Playlist used for trick-play. It is followed by a URI.
	//	#EXT-X-I-FRAME-STREAM-INF:<attribute-list>
	//
	//[4.3.4.3]
	//
	//[4.3.4.3]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.3
	EXT_X_I_FRAME_STREAM_INF PlaylistTag = "EXT-X-I-FRAME-STREAM-INF"

	// The EXT-X-INDEPENDENT-SEGMENTS tag indicates that all Media Segments
	// can be decoded without information from other segments.
	//	#EXT-X-INDEPENDENT-SEGMENTS
	//
	//[4.3.5.1]
	//
	//[4.3.5.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.5.1
	EXT_X_INDEPENDENT_SEGMENTS PlaylistTag = "EXT-X-INDEPENDENT-SEGMENTS"

	// The EXT-X-DEFINE tag allows defining variables for substitution in
	// Playlists.
	//	#EXT-X-DEFINE:<attribute-list>
	//
	//[4.3.4.5]
	//
	//[4.3.4.5]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.5
	EXT_X_DEFINE PlaylistTag = "EXT-X-DEFINE"

	// Session-level Tags
	// Tags that apply to a session (outside individual Playlists) and may
	// carry data or keys applicable to multiple Playlists.

	// The EXT-X-SESSION-DATA tag carries session-level data.
	//	#EXT-X-SESSION-DATA:<attribute-list>
	//
	//[4.3.4.4]
	//
	//[4.3.4.4]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.4
	EXT_X_SESSION_DATA PlaylistTag = "EXT-X-SESSION-DATA"

	// The EXT-X-SESSION-KEY tag specifies encryption keys that apply at
	// the session level (outside of a specific Media Playlist).
	//	#EXT-X-SESSION-KEY:<attribute-list>
	//
	//[4.3.4.5]
	//
	//[4.3.4.5]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4.5
	EXT_X_SESSION_KEY PlaylistTag = "EXT-X-SESSION-KEY"
)

const (
	// Server / Control Tags (LL-HLS & extensions)
	// The EXT-X-SERVER-CONTROL tag specifies server-side parameters that
	// affect client behavior in low-latency scenarios.
	//	#EXT-X-SERVER-CONTROL:<attribute-list>
	//
	//(LL-HLS extension - not in RFC 8216 core)
	EXT_X_SERVER_CONTROL PlaylistTag = "EXT-X-SERVER-CONTROL"

	// Low-Latency / Part-based Tags (from LL-HLS extensions)
	// These may be present in Media Playlists that implement low-latency
	// behavior (see LL-HLS extensions).

	// The EXT-X-PART tag represents a partial media segment (a "part")
	// used for low-latency delivery.
	//	#EXT-X-PART:<attribute-list>
	//
	//(LL-HLS extension - not in RFC 8216 core)
	EXT_X_PART PlaylistTag = "EXT-X-PART"

	// The EXT-X-PRELOAD-HINT tag gives a hint about a resource the server
	// is likely to upload next (used in LL-HLS push workflows).
	//	#EXT-X-PRELOAD-HINT:<attribute-list>
	//
	//(LL-HLS extension - not in RFC 8216 core)
	EXT_X_PRELOAD_HINT PlaylistTag = "EXT-X-PRELOAD-HINT"

	// The EXT-X-SKIP tag allows indicating skipped segments for fast
	// catch-up (LL-HLS related).
	//	#EXT-X-SKIP:<attribute-list>
	//
	//(LL-HLS extension - not in RFC 8216 core)
	EXT_X_SKIP PlaylistTag = "EXT-X-SKIP"

	// The EXT-X-RENDITION-REPORT tag is a reporting mechanism for
	// renditions (status about availability/progress).
	//	#EXT-X-RENDITION-REPORT:<attribute-list>
	//
	//(LL-HLS extension - not in RFC 8216 core)
	EXT_X_RENDITION_REPORT PlaylistTag = "EXT-X-RENDITION-REPORT"

	// Other Tags (commonly seen or in extensions)
	// The EXT-X-RENDITION-REPORT, EXT-X-KEY, and EXT-X-MAP tags are
	// already defined above; this area can be extended with vendor tags.
)
