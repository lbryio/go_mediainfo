# mediainfo
Golang binding for [libmediainfo](https://mediaarea.net/en/MediaInfo)

Duration, Bitrate, Codec, Streams and a lot of other meta-information about media files can be extracted through it.

For now supports only media files with one stream. Bindings for MediaInfoList is not provided. It can be easy fixed if anybody ask me.

Works through MediaInfoDLL/MediaInfoDLL.h(dynamic load and so on), so your mediainfo installation should has it.

Supports direct reading files by name and reading data from []byte buffers(without copying it for C calls).

Documentation for libmediainfo is poor and ascetic, can be found [here](https://mediaarea.net/en/MediaInfo/Support/SDK).

Your advices and suggestions are welcome!

## Simple Example
```go
package main

import (
    "fmt"
    mediainfo "github.com/upstream-nsk/go_mediainfo"
    "os"
)

func main() {
    info, err := mediainfo.GetMediaInfo(os.Args[1])
    if err != nil {
        fmt.Printf("open failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("%v\n", info)
}
```

## Example 
```go
package main

import (
    "fmt"
    mediainfo "github.com/upstream-nsk/go_mediainfo"
    "os"
)

func main() {
    mi := mediainfo.NewMediaInfo()
    err := mi.OpenFile(os.Args[1])
    if err != nil {
        fmt.Printf("open failed: %v\n", err)
        os.Exit(1)
    }
    defer mi.Close()
    video := &mediainfo.SimpleMediaInfo{}
    g := &video.General
    g.DurationStr = mi.Get(mediainfo.MediaInfo_Stream_General, "Duration/String3")
    g.Duration = mi.GetInt(mediainfo.MediaInfo_Stream_General, "Duration")
    g.Start = mi.GetInt(mediainfo.MediaInfo_Stream_General, "Start")
    g.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_General, "OverallBitRate")
    g.FrameRate = mi.GetInt(mediainfo.MediaInfo_Stream_General, "FrameRate")
    g.FileSize = mi.GetInt(mediainfo.MediaInfo_Stream_General, "FileSize")

    g.GeneralCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "GeneralCount"))
    g.VideoCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "VideoCount"))
    g.AudioCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "AudioCount"))
    g.TextCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "TextCount"))
    g.OtherCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "OtherCount"))
    g.ImageCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "ImageCount"))
    g.MenuCount = int(mi.GetInt(mediainfo.MediaInfo_Stream_General, "MenuCount"))

    v := &video.Video
    v.CodecID = mi.Get(mediainfo.MediaInfo_Stream_Video, "CodecID")
    v.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "BitRate")
    v.Width = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "Width")
    v.Height = mi.GetInt(mediainfo.MediaInfo_Stream_Video, "Height")
    v.Resolution = fmt.Sprintf("%dx%d", v.Width, v.Height)
    v.DAR = mi.Get(mediainfo.MediaInfo_Stream_Video, "DisplayAspectRatio/String")

    a := &video.Audio
    a.CodecID = mi.Get(mediainfo.MediaInfo_Stream_Audio, "CodecID")
    a.SamplingRate = mi.GetInt(mediainfo.MediaInfo_Stream_Audio, "SamplingRate")
    a.BitRate = mi.GetInt(mediainfo.MediaInfo_Stream_Audio, "BitRate")

    video.SubtitlesCnt = int(mi.GetInt(mediainfo.MediaInfo_Stream_Text, "StreamCount"))
    for i := 0; i < video.SubtitlesCnt; i++ {
        Format := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "Format")
        CodecID := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "CodecID")
        Title := mi.GetIdx(mediainfo.MediaInfo_Stream_Text, i, "Title")
        subtitles := mediainfo.SubtitlesInfo{Format: Format, CodecID: CodecID, Title: Title}
        video.Subtitles = append(video.Subtitles, subtitles)
    }
    fmt.Printf("%v\n", video)
}

```

## Output Format
```
General
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
GeneralCount              : Count of general streams
VideoCount                : Count of video streams
AudioCount                : Count of audio streams
TextCount                 : Count of text streams
ChaptersCount             : Count of chapters streams
ImageCount                : Count of image streams
MenuCount                 : Count of menu streams
Video_Format_List         : Video Codecs separated by /
Video_Format_WithHint_Lis : Video Codecs separated by /
Video_Codec_List          : Deprecated
Video_Language_List       : Video languages separated by /
Audio_Format_List         : Audio Codecs separated by /
Audio_Format_WithHint_Lis : Audio Codecs separated by /
Audio_Codec_List          : Deprecated
Audio_Language_List       : Audio languages separated by /
Text_Format_List          : Text Codecs separated by /
Text_Format_WithHint_List : Text Codecs separated by /
Text_Codec_List           : Deprecated
Text_Language_List        : Text languages separated by /
Chapters_Format_List      : Chapters Codecs separated by /
Chapters_Format_WithHint_ : Chapters Codecs separated by /
Chapters_Codec_List       : Deprecated
Chapters_Language_List    : Chapters languages separated by /
Image_Format_List         : Image Codecs separated by /
Image_Format_WithHint_Lis : Image Codecs separated by /
Image_Codec_List          : Deprecated
Image_Language_List       : Image languages separated by /
Menu_Format_List          : Menu Codecs separated by /
Menu_Format_WithHint_List : Menu Codecs separated by /
Menu_Codec_List           : Deprecated
Menu_Language_List        : Menu languages separated by /
CompleteName              : Complete name (Folder+Name+Extension)
FolderName                : Folder name only
FileName                  : File name only
FileExtension             : File extension only
Format                    : Format used
Format/String             : Deprecated
Format/Info               : Info about Format
Format/Url                : Link
Format/Extensions         : Known extensions of format
Format_Version            : Version of this format
Format_Profile            : Profile of the Format
Format_Settings           : Settings needed for decoder used
CodecID                   : Codec ID (found in some containers)
CodecID/Info              : Info about codec ID
CodecID/Hint              : A hint for this codec ID
CodecID/Url               : A link for more details about this codec ID
CodecID_Description       : Manual description given by the container
Interleaved               : If Audio and video are muxed
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Codec/Extensions          : Deprecated
Codec_Settings            : Deprecated
Codec_Settings_Automatic  : Deprecated
FileSize                  : File size in bytes
FileSize/String           : File size (with measure)
FileSize/String1          : File size (with measure, 1 digit mini)
FileSize/String2          : File size (with measure, 2 digit mini)
FileSize/String3          : File size (with measure, 3 digit mini)
FileSize/String4          : File size (with measure, 4 digit mini)
Duration                  : Play time of the stream
Duration/String           : Play time (formated)
Duration/String1          : Play time in format : HHh MMmn SSs MMMms, XX omited if zero
Duration/String2          : Play time in format : XXx YYy only, YYy omited if zero
Duration/String3          : Play time in format : HH:MM:SS.MMM
OverallBitRate_Mode       : Bit rate mode of all streams (VBR, CBR)
OverallBitRate_Mode/Strin : Bit rate mode of all streams (VBR, CBR)
OverallBitRate            : Bit rate of all streams
OverallBitRate/String     : Bit rate of all streams (with measure)
OverallBitRate_Minimum    : Minimum Bit rate in bps
OverallBitRate_Minimum/St : Minimum Bit rate (with measurement)
OverallBitRate_Nominal    : Nominal Bit rate in bps
OverallBitRate_Nominal/St : Nominal Bit rate (with measurement)
OverallBitRate_Maximum    : Maximum Bit rate in bps
OverallBitRate_Maximum/St : Maximum Bit rate (with measurement)
StreamSize                : Stream size in bytes
StreamSize/String
StreamSize/String1
StreamSize/String2
StreamSize/String3
StreamSize/String4
StreamSize/String5        : With proportion
StreamSize_Proportion     : Stream size divided by file size
HeaderSize
DataSize
FooterSize
Title                     : (Generic)Title of file
Title/More                : (Generic)More info about the title of file
Title/Url                 : (Generic)Url
Domain                    : Level=8. Eg : Starwars, Stargate, U2
Collection                : Level=7. Name of the collection. Eg : Starwars movies, Stargate movie, Stargate SG-1, Stargate Atlantis
Season                    : Level=6. Name of the season. Eg : Strawars first Trilogy, Season 1
Season_Position
Season_Position_Total
Movie                     : Level=5. Name of the movie. Eg : Starwars, a new hope
Movie/More
Movie/Url
Album                     : Level=5. Name of album. Eg : The joshua tree
Album/More
Album/Sort
Comic                     : Level=5. Name of the comic.
Comic/More
Comic/Position_Total
Part                      : Level=4. Name of the part. Eg : CD1, CD2
Part/Position
Part/Position_Total
Track                     : Level=3. Name of the track. Eg : track1, track 2
Track/More
Track/Url
Track/Sort
Track/Position
Track/Position_Total
Chapter                   : Level=3. Name of the chapter.
SubTrack                  : Level=2, Name of the subtrack.
Original/Album            : Original name of album, serie...
Original/Movie            : Original name of the movie
Original/Part             : Name of the part in the original support
Original/Track            : Original name of the track in the original support
Performer                 : (Generic)Performer of the file
Performer/Sort
Performer/Url
Original/Performer        : Original artist(s)/performer(s).
Accompaniment             : Band/orchestra/accompaniment/musician.
Composer                  : Name of the original composer.
Composer/Nationality      : Nationality of the main composer of the item, mostly for classical music.
Arranger                  : The person who arranged the piece. e.g. Ravel.
Lyricist                  : The person who wrote the lyrics for a musical item.
Original/Lyricist         : Original lyricist(s)/text writer(s).
Conductor                 : The artist(s) who performed the work. In classical music this would be the conductor, orchestra, soloists.
Director                  : Name of the director.
AssistantDirector         : Name of the assistant director.
DirectorOfPhotography     : The name of the director of photography, also known as cinematographer.
SoundEngineer             : The name of the sound engineer or sound recordist.
ArtDirector               : The person who oversees the artists and craftspeople who build the sets.
ProductionDesigner        : Artist responsible for designing the Overall visual appearance of a movie.
Choregrapher              : The name of the choregrapher.
CostumeDesigner           : The name of the costume designer.
Actor                     : Real name of actor or actress playing a role in the movie.
Actor_Character           : Name of the character an actor or actress plays in this movie.
WrittenBy                 : The author of the story or script.
ScreenplayBy              : The author of the screenplay or scenario (used for movies and TV shows).
EditedBy                  : Editor name
Producer                  : Name of the producer of the movie.
CoProducer                : The name of a co-producer.
ExecutiveProducer         : The name of an executive producer.
DistributedBy             :
MasteredBy                : The engineer who mastered the content for a physical medium or for digital distribution.
EncodedBy                 : Name of the person or organisation that encoded/ripped the audio file.
RemixedBy                 : Interpreted, remixed, or otherwise modified by.
ProductionStudio          :
ThanksTo                  : A very general tag for everyone else that wants to be listed.
Publisher                 : Name of the organization producing the track (i.e. the 'record label').
Publisher/URL             : Publishers official webpage.
Label
Genre                     : The main genre of the audio or video. e.g. classical, ambient-house, synthpop, sci-fi, drama, etc.
Mood                      : Intended to reflect the mood of the item with a few keywords, e.g. Romantic, Sad, Uplifting, etc.
ContentType               : The type of the item. e.g. Documentary, Feature Film, Cartoon, Music Video, Music, Sound FX, etc.
Subject                   : Describes the topic of the file, such as Aerial view of Seattle..
Description               : A short description of the contents, such as Two birds flying.
Keywords                  : Keywords to the item separated by a comma, used for searching.
Summary                   : A plot outline or a summary of the story.
Synopsys                  : A description of the story line of the item.
Period                    : Describes the period that the piece is from or about. e.g. Renaissance.
LawRating                 : Depending on the country it's the format of the rating of a movie (P, R, X in the USA, an age in other countries or a URI defining a logo).
LawRating_Reason          : Reason for the law rating
ICRA                      : The ICRA rating. (Previously RSACi)
Released_Date             : The time that the item was released.
Original/Released_Date    : The time that the item was originaly released.
Recorded_Date             : The time that the recording began.
Encoded_Date              : The time that the encoding of this item was completed began.
Tagged_Date               : The time that the tags were done for this item.
Written_Date              : The time that the composition of the music/script began.
Mastered_Date             : The time that the item was tranfered to a digitalmedium.
File_Created_Date         : The time that the file was created on the file system
File_Modified_Date        : The time that the file was modified on the file system
Recorded_Location         : Location where track was recorded. (See COMPOSITION_LOCATION for format)
Written_Location          : Location that the item was originaly designed/written. Information should be stored in the following format: country code, state/province, city where the coutry code is the same 2 octets as in Internet domains, or possibly ISO-3166. e.g. US, Texas, Austin or US, , Austin.
Archival_Location
Encoded_Application       : Software. Identifies the name of the software package used to create the file, such as Microsoft WaveEdit.
Encoded_Application/Url   : Software. Identifies the name of the software package used to create the file, such as Microsoft WaveEdit.
Encoded_Library           : Software used to create the file
Encoded_Library/String    : Software used to create the file
Encoded_Library/Name      : Info from the software
Encoded_Library/Version   : Version of software
Encoded_Library/Date      : Release date of software
Encoded_Library_Settings  : Parameters used by the software
Tagged_Application        : Software used to tag this file
BPM                       : Average number of beats per minute
ISRC                      : International Standard Recording Code, excluding the ISRC prefix and including hyphens.
ISBN                      : International Standard Book Number.
BarCode                   : EAN-13 (13-digit European Article Numbering) or UPC-A (12-digit Universal Product Code) bar code identifier.
LCCN                      : Library of Congress Control Number.
CatalogNumber             : A label-specific catalogue number used to identify the release. e.g. TIC 01.
LabelCode                 : A 4-digit or 5-digit number to identify the record label, typically printed as (LC) xxxx or (LC) 0xxxx on CDs medias or covers, with only the number being stored.
Owner                     : Owner of the file
Copyright                 : Copyright attribution.
Copyright/Url             : Copyright/legal information.
Producer_Copyright        : The copyright information as per the productioncopyright holder.
TermsOfUse                : License information, e.g., All Rights Reserved,Any Use Permitted.
Broadcaster
Broadcaster/Owner
Broadcaster/Url
Cover                     : Is there a cover
Cover_Description
Cover_Type
Cover_Mime
Cover_Data                : Cover, in binary format encoded BASE64
Lyrics
Comment                   : Any comment related to the content.
Rating                    : A numeric value defining how much a person likes the song/movie. The number is between 0 and 5 with decimal values possible (e.g. 2.7), 5(.0) being the highest possible rating.

Video
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
MenuID                    : A menu ID for this stream in this file
MenuID/String             : A menu ID for this stream in this file
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
Format_Version            : Version of this format
Format_Profile            : Profile of the Format
Format_Settings           : Settings needed for decoder used, summary
Format_Settings_BVOP      : Settings needed for decoder used, detailled
Format_Settings_BVOP/Stri : Settings needed for decoder used, detailled
Format_Settings_QPel      : Settings needed for decoder used, detailled
Format_Settings_QPel/Stri : Settings needed for decoder used, detailled
Format_Settings_GMC       : Settings needed for decoder used, detailled
Format_Settings_GMC/String
Format_Settings_Matrix    : Settings needed for decoder used, detailled
Format_Settings_Matrix/St : Settings needed for decoder used, detailled
Format_Settings_Matrix_Da : Matrix, in binary format encoded BASE64. Order = intra, non-intra, gray intra, gray non-intra
Format_Settings_CABAC     : Settings needed for decoder used, detailled
Format_Settings_CABAC/Str : Settings needed for decoder used, detailled
Format_Settings_RefFrames : Settings needed for decoder used, detailled
Format_Settings_RefFrames : Settings needed for decoder used, detailled
Format_Settings_Pulldown  : Settings needed for decoder used, detailled
MuxingMode                : How this file is muxed in the container
CodecID                   : Codec ID (found in some containers)
CodecID/Info              : Info about codec ID
CodecID/Hint              : A hint for this codec ID
CodecID/Url               : A link for more details about this codec ID
CodecID_Description       : Manual description given by the container
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Family              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Codec/CC                  : Deprecated
Codec_Profile             : Deprecated
Codec_Description         : Deprecated
Codec_Settings            : Deprecated
Codec_Settings_PacketBitS : Deprecated
Codec_Settings_BVOP       : Deprecated
Codec_Settings_QPel       : Deprecated
Codec_Settings_GMC        : Deprecated
Codec_Settings_GMC/String : Deprecated
Codec_Settings_Matrix     : Deprecated
Codec_Settings_Matrix_Dat : Deprecated
Codec_Settings_CABAC      : Deprecated
Codec_Settings_RefFrames  : Deprecated
Duration                  : Play time of the stream
Duration/String           : Play time (formated)
Duration/String1          : Play time in format : HHh MMmn SSs MMMms, XX omited if zero
Duration/String2          : Play time in format : XXx YYy only, YYy omited if zero
Duration/String3          : Play time in format : HH:MM:SS.MMM
BitRate_Mode              : Bit rate mode (VBR, CBR)
BitRate_Mode/String       : Bit rate mode (VBR, CBR)
BitRate                   : Bit rate in bps
BitRate/String            : Bit rate (with measurement)
BitRate_Minimum           : Minimum Bit rate in bps
BitRate_Minimum/String    : Minimum Bit rate (with measurement)
BitRate_Nominal           : Nominal Bit rate in bps
BitRate_Nominal/String    : Nominal Bit rate (with measurement)
BitRate_Maximum           : Maximum Bit rate in bps
BitRate_Maximum/String    : Maximum Bit rate (with measurement)
Width                     : Width
Width/String
Height                    : Height
Height/String
PixelAspectRatio          : Pixel Aspect ratio
PixelAspectRatio/String   : Pixel Aspect ratio
DisplayAspectRatio        : Display Aspect ratio
DisplayAspectRatio/String : Display Aspect ratio
FrameRate_Mode            : Frame rate mode (CFR, VFR)
FrameRate_Mode/String     : Frame rate mode (CFR, VFR)
FrameRate                 : Frame rate
FrameRate/String          : Frame rate
FrameRate_Minimum         : Frame rate
FrameRate_Minimum/String  : Frame rate
FrameRate_Nominal         : Frame rate
FrameRate_Nominal/String  : Frame rate
FrameRate_Maximum         : Frame rate
FrameRate_Maximum/String  : Frame rate
FrameRate_Original        : Frame rate
FrameRate_Original/String : Frame rate
FrameCount                : Frame count
Standard                  : NTSC or PAL
Resolution                : example for MP3 : 16 bits
Resolution/String         : example for MP3 : 16 bits
Colorimetry
ScanType
ScanType/String
ScanOrder
ScanOrder/String
Interlacement             : Deprecated
Interlacement/String      : Deprecated
Bits-(Pixel*Frame)        : bits/(Pixel*Frame) (like Gordian Knot)
Delay                     : Delay fixed in the stream (relative) IN MS
Delay/String
Delay/String1
Delay/String2
Delay/String3
StreamSize                : Stream size in bytes
StreamSize/String
StreamSize/String1
StreamSize/String2
StreamSize/String3
StreamSize/String4
StreamSize/String5        : With proportion
StreamSize_Proportion     : Stream size divided by file size
Alignment                 : How this stream file is aligned in the container
Alignment/String
Title                     : Name of the track
Encoded_Application       : Software. Identifies the name of the software package used to create the file, such as Microsoft WaveEdit.
Encoded_Application/Url   : Software. Identifies the name of the software package used to create the file, such as Microsoft WaveEdit.
Encoded_Library           : Software used to create the file
Encoded_Library/String    : Software used to create the file
Encoded_Library/Name      : Info from the software
Encoded_Library/Version   : Version of software
Encoded_Library/Date      : Release date of software
Encoded_Library_Settings  : Parameters used by the software
Language                  : Language (2 letters)
Language/String           : Language (full)
Language_More             : More info about Language (director's comment...)
Encoded_Date              : The time that the encoding of this item was completed began.
Tagged_Date               : The time that the tags were done for this item.
Encryption

Audio
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
MenuID                    : A menu ID for this stream in this file
MenuID/String             : A menu ID for this stream in this file
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
Format_Version            : Version of this format
Format_Profile            : Profile of the Format
Format_Settings           : Settings needed for decoder used, summary
Format_Settings_SBR
Format_Settings_SBR/String
Format_Settings_PS
Format_Settings_PS/String
Format_Settings_Floor
Format_Settings_Firm
Format_Settings_Endianness
Format_Settings_Sign
Format_Settings_Law
Format_Settings_ITU
MuxingMode                : How this stream is muxed in the container
CodecID                   : Codec ID (found in some containers)
CodecID/Info              : Info about codec ID
CodecID/Hint              : A hint for this codec ID
CodecID/Url               : A link for more details about this codec ID
CodecID_Description       : Manual description given by the container
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Family              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Codec/CC                  : Deprecated
Codec_Description         : Deprecated
Codec_Profile             : Deprecated
Codec_Settings            : Deprecated
Codec_Settings_Automatic  : Deprecated
Codec_Settings_Floor      : Deprecated
Codec_Settings_Firm       : Deprecated
Codec_Settings_Endianness : Deprecated
Codec_Settings_Sign       : Deprecated
Codec_Settings_Law        : Deprecated
Codec_Settings_ITU        : Deprecated
Duration                  : Play time of the stream
Duration/String           : Play time (formated)
Duration/String1          : Play time in format : HHh MMmn SSs MMMms, XX omited if zero
Duration/String2          : Play time in format : XXx YYy only, YYy omited if zero
Duration/String3          : Play time in format : HH:MM:SS.MMM
BitRate_Mode              : Bit rate mode (VBR, CBR)
BitRate_Mode/String       : Bit rate mode (VBR, CBR)
BitRate                   : Bit rate in bps
BitRate/String            : Bit rate (with measurement)
BitRate_Minimum           : Minimum Bit rate in bps
BitRate_Minimum/String    : Minimum Bit rate (with measurement)
BitRate_Nominal           : Nominal Bit rate in bps
BitRate_Nominal/String    : Nominal Bit rate (with measurement)
BitRate_Maximum           : Maximum Bit rate in bps
BitRate_Maximum/String    : Maximum Bit rate (with measurement)
Channel(s)                : Number of channels
Channel(s)/String         : Number of channels
ChannelPositions          : Position of channels
ChannelPositions/String2  : Position of channels (x/y.z format)
SamplingRate              : Sampling rate
SamplingRate/String       : in KHz
SamplingCount             : Frame count
Resolution                : Resolution in bits (8, 16, 20, 24)
Resolution/String
CompressionRatio          : Current stream size divided by uncompressed stream size
Delay                     : Delay fixed in the stream (relative)
Delay/String
Delay/String1
Delay/String2
Delay/String3
Video_Delay               : Delay fixed in the stream (absolute / video)
Video_Delay/String
Video_Delay/String1
Video_Delay/String2
Video_Delay/String3
Video0_Delay              : Deprecated
Video0_Delay/String       : Deprecated
Video0_Delay/String1      : Deprecated
Video0_Delay/String2      : Deprecated
Video0_Delay/String3      : Deprecated
ReplayGain_Gain           : The gain to apply to reach 89dB SPL on playback
ReplayGain_Peak           : The maximum absolute peak value of the item
StreamSize                : Stream size in bytes
StreamSize/String
StreamSize/String1
StreamSize/String2
StreamSize/String3
StreamSize/String4
StreamSize/String5        : With proportion
StreamSize_Proportion     : Stream size divided by file size
Alignment                 : How this stream file is aligned in the container
Alignment/String
Interleave_VideoFrames    : Between how many video frames the stream is inserted
Interleave_Duration       : Between how much time (ms) the stream is inserted
Interleave_Duration/Strin : Between how much time and video frames the stream is inserted (with measurement)
Interleave_Preload        : How much time is buffered before the first video frame
Interleave_Preload/String : How much time is buffered before the first video frame (with measurement)
Title                     : Name of the track
Encoded_Library           : Software used to create the file
Encoded_Library_Settings  : Parameters used by the software
Language                  : Language (2 letters)
Language/String           : Language (full)
Language_More             : More info about Language (director's comment...)
Encoded_Date              : The time that the encoding of this item was completed began.
Tagged_Date               : The time that the tags were done for this item.
Encryption

Text
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
MenuID                    : A menu ID for this stream in this file
MenuID/String             : A menu ID for this stream in this file
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
CodecID                   : Codec ID (found in some containers)
CodecID/Info              : Info about codec ID
CodecID/Hint              : A hint for this codec ID
CodecID/Url               : A link for more details about this codec ID
CodecID_Description       : Manual description given by the container
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Codec/CC                  : Deprecated
Duration                  : Play time of the stream
Duration/String           : Play time (formated)
Duration/String1          : Play time in format : HHh MMmn SSs MMMms, XX omited if zero
Duration/String2          : Play time in format : XXx YYy only, YYy omited if zero
Duration/String3          : Play time in format : HH:MM:SS.MMM
BitRate_Mode              : Bit rate mode (VBR, CBR)
BitRate_Mode/String       : Bit rate mode (VBR, CBR)
BitRate                   : Bit rate in bps
BitRate/String            : Bit rate (with measurement)
Width                     : Width
Width/String
Height                    : Height
Height/String
FrameCount                : Frame count
Resolution
Resolution/String
Delay                     : Delay fixed in the stream (relative)
Delay/String
Delay/String1
Delay/String2
Delay/String3
Video_Delay               : Delay fixed in the stream (absolute / video)
Video_Delay/String
Video_Delay/String1
Video_Delay/String2
Video_Delay/String3
Video0_Delay              : Deprecated
Video0_Delay/String       : Deprecated
Video0_Delay/String1      : Deprecated
Video0_Delay/String2      : Deprecated
Video0_Delay/String3      : Deprecated
StreamSize                : Stream size in bytes
StreamSize/String
StreamSize/String1
StreamSize/String2
StreamSize/String3
StreamSize/String4
StreamSize/String5        : With proportion
StreamSize_Proportion     : Stream size divided by file size
Title                     : Name of the track
Language                  : Language (2 letters)
Language/String           : Language (full)
Language_More             : More info about Language (director's comment...)
Summary
Encoded_Date              : The time that the encoding of this item was completed began.
Tagged_Date               : The time that the tags were done for this item.
Encryption

Chapters
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Total                     : Total number of chapters
Title                     : Name of the track
Language                  : Language (2 letters)
Language/String           : Language (full)

Image
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
Title                     : Name of the track
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
Format_Profile            : Profile of the Format
CodecID                   : Codec ID (found in some containers)
CodecID/Info              : Info about codec ID
CodecID/Hint              : A hint for this codec ID
CodecID/Url               : A link for more details about this codec ID
CodecID_Description       : Manual description given by the container
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Family              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Width                     : Width
Width/String
Height                    : Height
Height/String
Resolution
Resolution/String
StreamSize                : Stream size in bytes
StreamSize/String
StreamSize/String1
StreamSize/String2
StreamSize/String3
StreamSize/String4
StreamSize/String5        : With proportion
StreamSize_Proportion     : Stream size divided by file size
Language                  : Language (2 letters)
Language/String           : Language (full)
Summary
Encoded_Date              : The time that the encoding of this item was completed began.
Tagged_Date               : The time that the tags were done for this item.
Encryption

Menu
Count                     : Count of objects available in this stream
StreamCount               : Count of streams of that kind available
StreamKind                : Stream type name
StreamKind/String         : Stream type name
StreamKindID              : Number of the stream (base=0)
StreamKindPos             : When multiple streams, number of the stream (base=1)
Inform                    : Last **Inform** call
ID                        : A ID for this stream in this file
ID/String                 : A ID for this stream in this file
UniqueID                  : A unique ID for this stream, should be copied with stream copy
MenuID                    : A menu ID for this stream in this file
MenuID/String             : A menu ID for this stream in this file
Format                    : Format used
Format/Info               : Info about Format
Format/Url                : Link
Codec                     : Deprecated
Codec/String              : Deprecated
Codec/Info                : Deprecated
Codec/Url                 : Deprecated
Title                     : Name of this menu
Language                  : Language (2 letters)
Language/String           : Language (full)
Language_More             : More info about Language (director's comment...)
List                      : List of programs available
List/String
```

Read the [documentation](https://godoc.org/github.com/zhulik/go_mediainfo) for other functions

## Contributing
You know=)
