package nzbget

import (
	"fmt"
	"strconv"
	"time"
)

// File represents the listFiles RPC endpoint.
//nolint:lll
type File struct {
	ID                int64  `json:"ID"`                // ID of file.
	NZBID             int64  `json:"NZBID"`             // ID of NZB//file.
	NZBFilename       string `json:"NZBFilename"`       // Name of nzb//file. The filename could include fullpath if client sent it by adding the file to queue.
	NZBName           string `json:"NZBName"`           // The name of nzb//file without path and extension. Ready for user//friendly output.
	Subject           string `json:"Subject"`           // Subject of article read from nzb//file.
	Filename          string `json:"Filename"`          // Filename parsed from subject. It could be incorrect since the subject not always correct formated. After the first article for file is read, the correct filename is read from article body.
	FilenameConfirmed bool   `json:"FilenameConfirmed"` // “True” if filename was already read from article’s body. “False” if the name was parsed from subject. For confirmed filenames the destination file on disk will be exactly as specified in field “filename”. For unconfirmed filenames the name could change later.
	DestDir           string `json:"DestDir"`           // Destination directory for output file.
	FileSizeLo        int64  `json:"FileSizeLo"`        // Filesize in bytes, Low 32//bits of 64//bit value.
	FileSizeHi        int64  `json:"FileSizeHi"`        // Filesize in bytes, High 32//bits of 64//bit value.
	RemainingSizeLo   int64  `json:"RemainingSizeLo"`   // Remaining size in bytes, Low 32//bits of 64//bit value.
	RemainingSizeHi   int64  `json:"RemainingSizeHi"`   // Remaining size in bytes, High 32//bits of 64//bit value.
	Paused            bool   `json:"Paused"`            // “True” if file is paused.
	PostTime          int64  `json:"PostTime"`          // Date/time when the file was posted to newsgroup Time is in C/Unix format.
	ActiveDownloads   int64  `json:"ActiveDownloads"`   // Number of active downloads for the file. With this filed can be determined what files is are being currently downloaded.
	Progress          int64  `json:"Progress"`          // v15.0 Download progress, a number in the range 0..1000. Divide it to 10 to get percent//value.
}

// Group represents the listGroups RPC endpoint.
type Group struct {
	NZBID              int64             `json:"NZBID"`
	RemainingSizeLo    int64             `json:"RemainingSizeLo"`
	RemainingSizeHi    int64             `json:"RemainingSizeHi"`
	RemainingSizeMB    int64             `json:"RemainingSizeMB"`
	PausedSizeLo       int64             `json:"PausedSizeLo"`
	PausedSizeHi       int64             `json:"PausedSizeHi"`
	PausedSizeMB       int64             `json:"PausedSizeMB"`
	RemainingFileCount int64             `json:"RemainingFileCount"`
	RemainingParCount  int64             `json:"RemainingParCount"`
	MaxPriority        int64             `json:"MaxPriority"`
	ActiveDownloads    int64             `json:"ActiveDownloads"`
	Status             GroupStatus       `json:"Status"`
	NZBName            string            `json:"NZBName"`
	Kind               string            `json:"Kind"`
	URL                string            `json:"URL"`
	NZBFilename        string            `json:"NZBFilename"`
	DestDir            string            `json:"DestDir"`
	FinalDir           string            `json:"FinalDir"`
	Category           string            `json:"Category"`
	ParStatus          ParStatus         `json:"ParStatus"`
	ExParStatus        ExParStatus       `json:"ExParStatus"`
	UnpackStatus       UnpackStatus      `json:"UnpackStatus"`
	MoveStatus         MoveStatus        `json:"MoveStatus"`
	ScriptStatus       ScriptStatus      `json:"ScriptStatus"`
	DeleteStatus       DeleteStatus      `json:"DeleteStatus"`
	MarkStatus         MarkStatus        `json:"MarkStatus"`
	URLStatus          URLStatus         `json:"UrlStatus"`
	FileSizeLo         int64             `json:"FileSizeLo"`
	FileSizeHi         int64             `json:"FileSizeHi"`
	FileSizeMB         int64             `json:"FileSizeMB"`
	FileCount          int64             `json:"FileCount"`
	MinPostTime        Time              `json:"MinPostTime"`
	MaxPostTime        Time              `json:"MaxPostTime"`
	TotalArticles      int64             `json:"TotalArticles"`
	SuccessArticles    int64             `json:"SuccessArticles"`
	FailedArticles     int64             `json:"FailedArticles"`
	Health             int64             `json:"Health"`
	CriticalHealth     int64             `json:"CriticalHealth"`
	DupeScore          int64             `json:"DupeScore"`
	DupeKey            string            `json:"DupeKey"`
	DupeMode           string            `json:"DupeMode"`
	DownloadedSizeLo   int64             `json:"DownloadedSizeLo"`
	DownloadedSizeHi   int64             `json:"DownloadedSizeHi"`
	DownloadedSizeMB   int64             `json:"DownloadedSizeMB"`
	DownloadTimeSec    int64             `json:"DownloadTimeSec"`
	PostTotalTimeSec   int64             `json:"PostTotalTimeSec"`
	ParTimeSec         int64             `json:"ParTimeSec"`
	RepairTimeSec      int64             `json:"RepairTimeSec"`
	UnpackTimeSec      int64             `json:"UnpackTimeSec"`
	MessageCount       int64             `json:"MessageCount"`
	ExtraParBlocks     int64             `json:"ExtraParBlocks"`
	Parameters         []Parameter       `json:"Parameters"`
	ScriptStatuses     []PerScriptStatus `json:"ScriptStatuses"`
	ServerStats        []ServerStats     `json:"ServerStats"`
	PostInfoText       string            `json:"PostInfoText"`
	PostStageProgress  int64             `json:"PostStageProgress"`
	PostStageTimeSec   int64             `json:"PostStageTimeSec"`
}

// Parameter is part of a few method's outputs.
type Parameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// LogEntry represents the log RPC endpoint.
type LogEntry struct {
	ID   int64   `json:"ID"`
	Time Time    `json:"Time"`
	Kind LogKind `json:"Kind"`
	Text string  `json:"Text"`
}

// History represents the hisory RPC endpoint.
type History struct {
	NZBID              int64             `json:"NZBID"`
	Name               string            `json:"Name"`
	RemainingFileCount int64             `json:"RemainingFileCount"`
	HistoryTime        Time              `json:"HistoryTime"`
	Status             string            `json:"Status"`
	NZBName            string            `json:"NZBName"`
	Kind               string            `json:"Kind"`
	URL                string            `json:"URL"`
	NZBFilename        string            `json:"NZBFilename"`
	DestDir            string            `json:"DestDir"`
	FinalDir           string            `json:"FinalDir"`
	Category           string            `json:"Category"`
	ParStatus          ParStatus         `json:"ParStatus"`
	ExParStatus        ExParStatus       `json:"ExParStatus"`
	UnpackStatus       UnpackStatus      `json:"UnpackStatus"`
	MoveStatus         MoveStatus        `json:"MoveStatus"`
	ScriptStatus       ScriptStatus      `json:"ScriptStatus"`
	DeleteStatus       DeleteStatus      `json:"DeleteStatus"`
	MarkStatus         MarkStatus        `json:"MarkStatus"`
	URLStatus          URLStatus         `json:"UrlStatus"`
	FileSizeLo         int64             `json:"FileSizeLo"`
	FileSizeHi         int64             `json:"FileSizeHi"`
	FileSizeMB         int64             `json:"FileSizeMB"`
	FileCount          int64             `json:"FileCount"`
	MinPostTime        Time              `json:"MinPostTime"`
	MaxPostTime        Time              `json:"MaxPostTime"`
	TotalArticles      int64             `json:"TotalArticles"`
	SuccessArticles    int64             `json:"SuccessArticles"`
	FailedArticles     int64             `json:"FailedArticles"`
	Health             int64             `json:"Health"`
	CriticalHealth     int64             `json:"CriticalHealth"`
	DupeScore          int64             `json:"DupeScore"`
	DupeKey            string            `json:"DupeKey"`
	DupeMode           string            `json:"DupeMode"`
	DownloadedSizeLo   int64             `json:"DownloadedSizeLo"`
	DownloadedSizeHi   int64             `json:"DownloadedSizeHi"`
	DownloadedSizeMB   int64             `json:"DownloadedSizeMB"`
	DownloadTimeSec    int64             `json:"DownloadTimeSec"`
	PostTotalTimeSec   int64             `json:"PostTotalTimeSec"`
	ParTimeSec         int64             `json:"ParTimeSec"`
	RepairTimeSec      int64             `json:"RepairTimeSec"`
	UnpackTimeSec      int64             `json:"UnpackTimeSec"`
	MessageCount       int64             `json:"MessageCount"`
	ExtraParBlocks     int64             `json:"ExtraParBlocks"`
	RetryData          bool              `json:"RetryData"`
	Parameters         []Parameter       `json:"Parameters"`
	ScriptStatuses     []PerScriptStatus `json:"ScriptStatuses"`
	ServerStats        []ServerStats     `json:"ServerStats"`
}

// PerScriptStatus is part of the history endpoint output.
type PerScriptStatus struct {
	Name   string       `json:"Name"`
	Status ScriptStatus `json:"Status"`
}

// ServerStats is part of a few endpoints output.
type ServerStats struct {
	ServerID        int64 `json:"ServerID"`
	SuccessArticles int64 `json:"SuccessArticles"`
	FailedArticles  int64 `json:"FailedArticles"`
}

// Status represents the status RPC endpoint.
type Status struct {
	RemainingSizeLo     int64         `json:"RemainingSizeLo"`
	RemainingSizeHi     int64         `json:"RemainingSizeHi"`
	RemainingSizeMB     int64         `json:"RemainingSizeMB"`
	ForcedSizeLo        int64         `json:"ForcedSizeLo"`
	ForcedSizeHi        int64         `json:"ForcedSizeHi"`
	ForcedSizeMB        int64         `json:"ForcedSizeMB"`
	DownloadedSizeLo    int64         `json:"DownloadedSizeLo"`
	DownloadedSizeHi    int64         `json:"DownloadedSizeHi"`
	DownloadedSizeMB    int64         `json:"DownloadedSizeMB"`
	MonthSizeLo         int64         `json:"MonthSizeLo"`
	MonthSizeHi         int64         `json:"MonthSizeHi"`
	MonthSizeMB         int64         `json:"MonthSizeMB"`
	DaySizeLo           int64         `json:"DaySizeLo"`
	DaySizeHi           int64         `json:"DaySizeHi"`
	DaySizeMB           int64         `json:"DaySizeMB"`
	ArticleCacheLo      int64         `json:"ArticleCacheLo"`
	ArticleCacheHi      int64         `json:"ArticleCacheHi"`
	ArticleCacheMB      int64         `json:"ArticleCacheMB"`
	DownloadRate        int64         `json:"DownloadRate"`
	AverageDownloadRate int64         `json:"AverageDownloadRate"`
	DownloadLimit       int64         `json:"DownloadLimit"`
	ThreadCount         int64         `json:"ThreadCount"`
	PostJobCount        int64         `json:"PostJobCount"`
	URLCount            int64         `json:"UrlCount"`
	UpTimeSec           int64         `json:"UpTimeSec"`
	DownloadTimeSec     int64         `json:"DownloadTimeSec"`
	FreeDiskSpaceLo     int64         `json:"FreeDiskSpaceLo"`
	FreeDiskSpaceHi     int64         `json:"FreeDiskSpaceHi"`
	FreeDiskSpaceMB     int64         `json:"FreeDiskSpaceMB"`
	ServerTime          Time          `json:"ServerTime"`
	ResumeTime          Time          `json:"ResumeTime"`
	QueueScriptCount    int64         `json:"QueueScriptCount"`
	FeedActive          bool          `json:"FeedActive"`
	DownloadPaused      bool          `json:"DownloadPaused"`
	ServerStandBy       bool          `json:"ServerStandBy"`
	PostPaused          bool          `json:"PostPaused"`
	ScanPaused          bool          `json:"ScanPaused"`
	QuotaReached        bool          `json:"QuotaReached"`
	NewsServers         []NewsServers `json:"NewsServers"`
}

// NewsServers is part of the Status endpoint output.
type NewsServers struct {
	ID     int64 `json:"ID"`
	Active bool  `json:"Active"`
}

// Time defines a timestamp encoded as epoch seconds in JSON.
// NZBGet returns all times as seconds since epoch, so we use a custom type to always return a proper go Time.
type Time time.Time

// MarshalJSON is used to convert the timestamp to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil //nolint:gomnd
}

// UnmarshalJSON is used to convert the timestamp from JSON.
func (t *Time) UnmarshalJSON(s []byte) error {
	q, err := strconv.ParseInt(string(s), 10, 64) //nolint:gomnd
	if err != nil {
		return fmt.Errorf("parsing number: %w", err)
	}

	*(*time.Time)(t) = time.Unix(q, 0)

	return nil
}
