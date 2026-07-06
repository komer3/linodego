package linodego

import "context"

type NFSSnapshotFrequency string

const (
	NFSSnapshotFrequencyHourly  NFSSnapshotFrequency = "hourly"
	NFSSnapshotFrequencyDaily   NFSSnapshotFrequency = "daily"
	NFSSnapshotFrequencyWeekly  NFSSnapshotFrequency = "weekly"
	NFSSnapshotFrequencyMonthly NFSSnapshotFrequency = "monthly"
)

type NFSSnapshotScheduleDayOfWeek string

const (
	NFSSnapshotScheduleDayOfWeekMon NFSSnapshotScheduleDayOfWeek = "mon"
	NFSSnapshotScheduleDayOfWeekTue NFSSnapshotScheduleDayOfWeek = "tue"
	NFSSnapshotScheduleDayOfWeekWed NFSSnapshotScheduleDayOfWeek = "wed"
	NFSSnapshotScheduleDayOfWeekThu NFSSnapshotScheduleDayOfWeek = "thu"
	NFSSnapshotScheduleDayOfWeekFri NFSSnapshotScheduleDayOfWeek = "fri"
	NFSSnapshotScheduleDayOfWeekSat NFSSnapshotScheduleDayOfWeek = "sat"
	NFSSnapshotScheduleDayOfWeekSun NFSSnapshotScheduleDayOfWeek = "sun"
)

type NFSSnapshotPolicy struct {
	FilesystemID int                        `json:"filesystem_id"`
	Enabled      bool                       `json:"enabled"`
	NamePrefix   string                     `json:"name_prefix"`
	Schedules    []NFSSnapshotScheduleEntry `json:"schedules"`
}

type NFSSnapshotScheduleEntry struct {
	Frequency  NFSSnapshotFrequency          `json:"frequency"`
	Minute     *int                          `json:"minute,omitzero"`
	Hour       *int                          `json:"hour,omitzero"`
	DayOfWeek  *NFSSnapshotScheduleDayOfWeek `json:"day_of_week,omitzero"`
	DayOfMonth *int                          `json:"day_of_month,omitzero"`
	Retain     int                           `json:"retain"`
}

type NFSSnapshotPolicyUpdateOptions struct {
	Enabled    *bool                       `json:"enabled,omitzero"`
	NamePrefix string                      `json:"name_prefix"`
	Schedules  *[]NFSSnapshotScheduleEntry `json:"schedules,omitzero"`
}

func (c *Client) GetNFSSnapshotPolicy(
	ctx context.Context,
	spaceID string,
	filesystemID string,
) (*NFSSnapshotPolicy, error) {
	return doGETRequest[NFSSnapshotPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshot-policy", spaceID, filesystemID),
	)
}

func (c *Client) UpsertNFSSnapshotPolicy(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSSnapshotPolicyUpdateOptions,
) (*NFSSnapshotPolicy, error) {
	return doPUTRequest[NFSSnapshotPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshot-policy", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSSnapshotPolicy(ctx context.Context, spaceID string, filesystemID string) error {
	return doDELETERequest(
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshot-policy", spaceID, filesystemID),
	)
}
