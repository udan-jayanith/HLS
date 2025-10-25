//HLS Go module implements HTTP Live Streaming protocol for Go.
//HLS can encode and decode HTTP Live Streams and also provide a tokenizer and a serializer for low level access.
//This HLS module does not serve HTTP live streams and it's users responsibility to serve HTTP live streams.
//But HLS provides helper methods to server HTTP live streams.

package HLS

//Playlist tags specify either global parameters of the Playlist or
//information about the Media Segments or Media Playlists that appear
//after them.
type PlaylistTag = string

const (
	//Basic Tags
	//These tags are allowed in both Media Playlists and Master Playlists.

	//The EXTM3U tag indicates that the file is an Extended M3U [M3U]
	//Playlist file.  It MUST be the first line of every Media Playlist and
	//every Master Playlist.  Its format is:
	//
	//#EXTM3U
	//
	//[4.3.1.1]
	//
	//[4.3.1.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.1
	EXTM3U PlaylistTag = "EXTM3U"
	//The EXT-X-VERSION tag indicates the compatibility version of the
	//Playlist file, its associated media, and its server.
	//
	//The EXT-X-VERSION tag applies to the entire Playlist file.  Its
	//format is:
	//
	//#EXT-X-VERSION:<n>
	//
	//where n is an integer indicating the protocol compatibility version
	//number.
	//
	//[4.3.1.2]
	//
	//[4.3.1.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.2
	EXT_X_VERSION = "EXT-X-VERSION"

	//Media Segment Tags
	//Each Media Segment is specified by a series of Media Segment tags
	//followed by a URI.  Some Media Segment tags apply to just the next
	//segment; others apply to all subsequent segments until another
	//instance of the same tag.
	//
	//A Media Segment tag MUST NOT appear in a Master Playlist.

	//The EXTINF tag specifies the duration of a Media Segment.  It applies
	//only to the next Media Segment.  This tag is REQUIRED for each Media
	//Segment.  Its format is:
	//
	//#EXTINF:<duration>,[<title>]
	//
	//[4.3.2.1]
	//
	//[4.3.2.1]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.1
	EXTINF = "EXTINF"
	//The EXT-X-BYTERANGE tag indicates that a Media Segment is a sub-range
	//of the resource identified by its URI.  It applies only to the next
	//URI line that follows it in the Playlist.  Its format is:
	//
	//#EXT-X-BYTERANGE:<n>[@<o>]
	//
	//[4.3.2.2]
	//
	//[4.3.2.2]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.2
	EXT_X_BYTERANGE = "EXT-X-BYTERANGE"
	//The EXT-X-DISCONTINUITY tag indicates a discontinuity between the
	//Media Segment that follows it and the one that preceded it.
	//
	//Its format is:
	//
	//#EXT-X-DISCONTINUITY
	//
	//[4.3.2.3]
	//
	//[4.3.2.3]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.3
	EXT_X_DISCONTINUITY = "EXT-X-DISCONTINUITY"
	//Media Segments MAY be encrypted.  The EXT-X-KEY tag specifies how to
	//decrypt them.  It applies to every Media Segment and to every Media
	//Initialization Section declared by an EXT-X-MAP tag that appears
	//between it and the next EXT-X-KEY tag in the Playlist file with the
	//same KEYFORMAT attribute (or the end of the Playlist file).  Two or
	//more EXT-X-KEY tags with different KEYFORMAT attributes MAY apply to
	//the same Media Segment if they ultimately produce the same decryption
	//key.  The format is:
	//
	//#EXT-X-KEY:<attribute-list>
	//
	//[4.3.2.4]
	//
	//[4.3.2.4]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.4
	EXT_X_KEY = "EXT-X-KEY"
	//The EXT-X-MAP tag specifies how to obtain the Media Initialization
	//Section (Section 3) required to parse the applicable Media Segments.
	//It applies to every Media Segment that appears after it in the
	//Playlist until the next EXT-X-MAP tag or until the end of the
	//Playlist.
	//
	//Its format is:
	//
	//#EXT-X-MAP:<attribute-list>
	//
	//[4.3.2.5]
	//
	//[4.3.2.5]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.5
	EXT_X_MAP = "EXT-X-MAP"
	//The EXT-X-PROGRAM-DATE-TIME tag associates the first sample of a
	//Media Segment with an absolute date and/or time.  It applies only to
	//the next Media Segment.  Its format is:
	//
	//#EXT-X-PROGRAM-DATE-TIME:<date-time-msec>
	//
	//[4.3.2.6]
	//
	//[4.3.2.6]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.6
	EXT_X_PROGRAM_DATE_TIME = "EXT-X-PROGRAM-DATE-TIME"
	//  The EXT-X-DATERANGE tag associates a Date Range (i.e., a range of
	//time defined by a starting and ending date) with a set of attribute/
	//value pairs.  Its format is:
	//
	//#EXT-X-DATERANGE:<attribute-list>
	//
	//[4.3.2.7]
	//
	//[4.3.2.7]: https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.7
	EXT_X_DATERANGE = "EXT-X-DATERANGE"
	
	//Media Playlist Tags
	//
	//Media Playlist tags describe global parameters of the Media Playlist.
	//There MUST NOT be more than one Media Playlist tag of each type in
	//any Media Playlist.
	//
	// A Media Playlist tag MUST NOT appear in a Master Playlist.
)
