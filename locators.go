package trac

import "time"

// Represents a locator string for filtering Agent entities.
type AgentLocator struct {
	Authorized bool   `json:"authorized,omitempty"` // Is the agent authorized.
	Build      string `json:"build,omitempty"`      // Build locator.
	Compatible string `json:"compatible,omitempty"` // Compatible build types locator.
	Connected  bool   `json:"connected,omitempty"`  // Is the agent connected.
	Count      int32  `json:"count,omitempty"`      // For paginated calls, how many entities to return per page.
	Enabled    bool   `json:"enabled,omitempty"`    // Is the agent enabled.
	ID         int32  `json:"id,omitempty"`         // Entity ID.
	IP         string `json:"ip,omitempty"`
	Item       string `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name       string `json:"name,omitempty"`
	Parameter  string `json:"parameter,omitempty"`
	Pool       string `json:"pool,omitempty"`  // Agent pool locator.
	Start      int32  `json:"start,omitempty"` // For paginated calls, from which entity to start rendering the page.
}

type AgentPoolLocator struct {
	Agent   string `json:"agent,omitempty"` // Pool's agents locator.
	Id      string `json:"id,omitempty"`
	Item    string `json:"item,omitempty"` // Supply multiple locators and return a union of the results.
	Name    string `json:"name,omitempty"`
	Project string `json:"project,omitempty"` // Pool's associated projects locator.
}

// Represents a locator string for filtering AuditEvent entities.
type AuditLocator struct {
	Action          string `json:"action,omitempty"`          // Use `$help` to get the full list of supported actions.
	AffectedProject string `json:"affectedProject,omitempty"` // Related project locator.
	BuildType       string `json:"buildType,omitempty"`       // Related build type or template locator.
	Count           int32  `json:"count,omitempty"`           // For paginated calls, how many entities to return per page.
	Id              string `json:"id,omitempty"`
	Item            string `json:"item,omitempty"`        // Supply multiple locators and return a union of the results.
	LookupLimit     int32  `json:"lookupLimit,omitempty"` // Limit processing to the latest `<lookupLimit>` entities.
	Start           int32  `json:"start,omitempty"`       // For paginated calls, from which entity to start rendering the page.
	SystemAction    string `json:"systemAction,omitempty"`
	User            string `json:"user,omitempty"` // Locator of user who caused the audit event.
}

func BuildLocatorID(id int32) BuildLocator {
	return BuildLocator{Id: id}
}

// Represents a locator string for filtering Build entities.
type BuildLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Agent locator.
	Agent string `json:"agent,omitempty"`
	// typeId of agent used to execute build.
	AgentTypeId int32 `json:"agentTypeId,omitempty"`
	// State can be any.
	Any                bool   `json:"any,omitempty"`
	ArtifactDependency string `json:"artifactDependency,omitempty"`
	// Branch locator.
	Branch string `json:"branch,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// Is canceled.
	Canceled bool `json:"canceled,omitempty"`
	// Agent locator.
	CompatibleAgent string `json:"compatibleAgent,omitempty"`
	// Is composite.
	Composite bool `json:"composite,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// If true, applies default filter which returns only \"normal\" builds (finished builds which are not canceled, not failed-to-start, not personal, and on default branch (in branched build configurations)).
	DefaultFilter bool `json:"defaultFilter,omitempty"`
	// Is failed to start.
	FailedToStart bool `json:"failedToStart,omitempty"`
	// Requires either date or build dimension.
	FinishDate string `json:"finishDate,omitempty"`
	// Is finished.
	Finished bool `json:"finished,omitempty"`
	// Is hanging.
	Hanging bool `json:"hanging,omitempty"`
	// Is history build.
	History bool `json:"history,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32  `json:"lookupLimit,omitempty"`
	Number      string `json:"number,omitempty"`
	// Is a personal build.
	Personal bool `json:"personal,omitempty"`
	// Is pinned.
	Pinned bool `json:"pinned,omitempty"`
	// Project (direct parent) locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
	// Is queued.
	Queued bool `json:"queued,omitempty"`
	// Requires either date or build dimension.
	QueuedDate string `json:"queuedDate,omitempty"`
	// Is running.
	Running            bool   `json:"running,omitempty"`
	SnapshotDependency string `json:"snapshotDependency,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// Requires either date or build dimension.
	StartDate string `json:"startDate,omitempty"`
	State     string `json:"state,omitempty"`
	Status    string `json:"status,omitempty"`
	// Tag locator.
	Tag string `json:"tag,omitempty"`
	// ID of a build or build promotion.
	TaskId int32 `json:"taskId,omitempty"`
	// User locator.
	User string `json:"user,omitempty"`
}

// Represents a locator string for filtering Build entities.
type BuildQueueLocator struct {
	// Agent locator.
	Agent string `json:"agent,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Is personal.
	Personal bool `json:"personal,omitempty"`
	// Agent pool locator.
	Pool string `json:"pool,omitempty"`
	// Project locator.
	Project string `json:"project,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start  int32  `json:"start,omitempty"`
	TaskId string `json:"taskId,omitempty"`
	// User locator.
	User string `json:"user,omitempty"`
}

// Represents a locator string for filtering BuildType entities.
type BuildTypeLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count      int32  `json:"count,omitempty"`
	Id         string `json:"id,omitempty"`
	InternalId string `json:"internalId,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Name string `json:"name,omitempty"`
	// Is paused.
	Paused bool `json:"paused,omitempty"`
	// Project (direct parent) locator.
	Project string `json:"project,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// Base template locator.
	Template string `json:"template,omitempty"`
	// Is a template.
	TemplateFlag bool   `json:"templateFlag,omitempty"`
	Uuid         string `json:"uuid,omitempty"`
	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`
	// VCS root instance locator.
	VcsRootInstance string `json:"vcsRootInstance,omitempty"`
}

// Represents a locator string for filtering Branch entities.
type BranchLocator struct {
	// Is feature branch.
	Branched string `json:"branched,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// Is default branch.
	Default_ string `json:"default,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item   string `json:"item,omitempty"`
	Name   string `json:"name,omitempty"`
	Policy string `json:"policy,omitempty"`
}

// Represents a locator string for filtering VcsRootInstance entities.
type VcsRootInstanceLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Project (direct parent) locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// Type of VCS (e.g. jetbrains.git).
	Type_ string `json:"type,omitempty"`
	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`
	// Is used for versioned settings.
	VersionedSettings bool `json:"versionedSettings,omitempty"`
}

// Represents a locator string for filtering CloudImage entities.
type CloudImageLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Agent locator.
	Agent string `json:"agent,omitempty"`
	// Agent pool locator.
	AgentPool string `json:"agentPool,omitempty"`
	Id        string `json:"id,omitempty"`
	// Cloud instance locator.
	Instance string `json:"instance,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Name string `json:"name,omitempty"`
	// Cloud profile locator.
	Profile string `json:"profile,omitempty"`
	// Project locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
}

// Represents a locator string for filtering Tag entities.
type TagLocator struct {
	Name    string `json:"name,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Private string `json:"private,omitempty"`
}

// Represents a locator string for filtering Change entities.
type ChangeLocator struct {
	// Build locator.
	Build string `json:"build,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	Comment   string `json:"comment,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32  `json:"count,omitempty"`
	File  string `json:"file,omitempty"`
	// Entity ID.
	Id              int32  `json:"id,omitempty"`
	InternalVersion string `json:"internalVersion,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Is pending.
	Pending bool `json:"pending,omitempty"`
	// Project locator.
	Project string `json:"project,omitempty"`
	// Commit SHA since which the changes should be returned.
	SinceChange string `json:"sinceChange,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// User locator.
	User string `json:"user,omitempty"`
	// VCS side username.
	Username string `json:"username,omitempty"`
	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`
	// VCS instance locator.
	VcsRootInstance string `json:"vcsRootInstance,omitempty"`
	Version         string `json:"version,omitempty"`
}

// Represents a locator string for filtering VcsRoot entities.
type VcsRootLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Entity ID.
	Id         int32  `json:"id,omitempty"`
	InternalId string `json:"internalId,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32  `json:"lookupLimit,omitempty"`
	Name        string `json:"name,omitempty"`
	// Project (direct parent) locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
	// Type of VCS (e.g. jetbrains.git).
	Type_ string `json:"type,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
}

// Represents a locator string for filtering CloudInstance entities.
type CloudInstanceLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Agent locator.
	Agent string `json:"agent,omitempty"`
	Id    string `json:"id,omitempty"`
	// Cloud image locator.
	Instance string `json:"instance,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item           string `json:"item,omitempty"`
	NetworkAddress string `json:"networkAddress,omitempty"`
	// Cloud profile locator.
	Profile string `json:"profile,omitempty"`
	// Project locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
}

// Represents a locator string for filtering CloudProfile entities.
type CloudProfileLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	CloudProviderId string `json:"cloudProviderId,omitempty"`
	Id              string `json:"id,omitempty"`
	// Cloud image locator.
	Instance string `json:"instance,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Name string `json:"name,omitempty"`
	// Project locator.
	Project  string `json:"project,omitempty"`
	Property string `json:"property,omitempty"`
}

// Represents a locator string for filtering Investigation entities.
type InvestigationLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	Assignee        string `json:"assignee,omitempty"`
	// Project (direct parent) locator.
	AssignmentProject string `json:"assignmentProject,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Problem locator.
	Problem    string `json:"problem,omitempty"`
	Reporter   string `json:"reporter,omitempty"`
	Resolution string `json:"resolution,omitempty"`
	// yyyyMMddTHHmmss+ZZZZ
	SinceDate time.Time `json:"sinceDate,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32  `json:"start,omitempty"`
	State string `json:"state,omitempty"`
	// Test locator.
	Test  string `json:"test,omitempty"`
	Type_ string `json:"type,omitempty"`
}

// Represents a locator string for filtering ProblemOccurrence entities.
type ProblemOccurrenceLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Is currently failing.
	CurrentlyFailing bool `json:"currentlyFailing,omitempty"`
	// Is currently investigated.
	CurrentlyInvestigated bool `json:"currentlyInvestigated,omitempty"`
	// Is currently muted.
	CurrentlyMuted bool   `json:"currentlyMuted,omitempty"`
	Identity       string `json:"identity,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Has ever been muted.
	Muted   bool   `json:"muted,omitempty"`
	Problem string `json:"problem,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32  `json:"start,omitempty"`
	Type_ string `json:"type,omitempty"`
}

// Represents a locator string for filtering Problem entities.
type ProblemLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Is currently failing.
	CurrentlyFailing bool `json:"currentlyFailing,omitempty"`
	// Is currently investigated.
	CurrentlyInvestigated bool `json:"currentlyInvestigated,omitempty"`
	// Is currently muted.
	CurrentlyMuted bool `json:"currentlyMuted,omitempty"`
	// Entity ID.
	Id       int32  `json:"id,omitempty"`
	Identity string `json:"identity,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32  `json:"start,omitempty"`
	Type_ string `json:"type,omitempty"`
}

// Represents a locator string for filtering Test entities.
type TestLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Is currently failing.
	CurrentlyFailing bool `json:"currentlyFailing,omitempty"`
	// Is currently investigated.
	CurrentlyInvestigated bool `json:"currentlyInvestigated,omitempty"`
	// Is currently muted.
	CurrentlyMuted bool `json:"currentlyMuted,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Build type locator (for finding out if this test is affected by mutes in build type).
	MuteAffected string `json:"muteAffected,omitempty"`
	Name         string `json:"name,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32 `json:"start,omitempty"`
}

// Represents a locator string for filtering TestOccurrence entities.
type TestOccurrenceLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	Branch          string `json:"branch,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Is currently failing.
	CurrentlyFailing bool `json:"currentlyFailing,omitempty"`
	// Is currently investigated.
	CurrentlyInvestigated bool   `json:"currentlyInvestigated,omitempty"`
	CurrentlyMuted        string `json:"currentlyMuted,omitempty"`
	// Entity ID.
	Id int32 `json:"id,omitempty"`
	// Is ignored.
	Ignored         bool `json:"ignored,omitempty"`
	IncludePersonal bool `json:"includePersonal,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Limit processing to the latest `<lookupLimit>` entities.
	LookupLimit int32 `json:"lookupLimit,omitempty"`
	// Is muted.
	Muted      bool   `json:"muted,omitempty"`
	Name       string `json:"name,omitempty"`
	NewFailure string `json:"newFailure,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start  int32  `json:"start,omitempty"`
	Status string `json:"status,omitempty"`
	// Test locator.
	Test string `json:"test,omitempty"`
}

// Represents a locator string for filtering Mute entities.
type MuteLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// yyyyMMddTHHmmss+ZZZZ
	CreationDate time.Time `json:"creationDate,omitempty"`
	Id           int32     `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	// Problem locator.
	Problem string `json:"problem,omitempty"`
	// Project (direct parent) locator.
	Project string `json:"project,omitempty"`
	// User who muted this test.
	Reporter   string `json:"reporter,omitempty"`
	Resolution string `json:"resolution,omitempty"`
	// Test locator.
	Test  string `json:"test,omitempty"`
	Type_ string `json:"type,omitempty"`
	// yyyyMMddTHHmmss+ZZZZ
	UnmuteDate time.Time `json:"unmuteDate,omitempty"`
}

// Represents a locator string for filtering Project entities.
type ProjectLocator struct {
	// Project (direct or indirect parent) locator.
	AffectedProject string `json:"affectedProject,omitempty"`
	// Is archived.
	Archived bool `json:"archived,omitempty"`
	// Build locator.
	Build string `json:"build,omitempty"`
	// Build type locator.
	BuildType string `json:"buildType,omitempty"`
	// For paginated calls, how many entities to return per page.
	Count int32 `json:"count,omitempty"`
	// Default template locator.
	DefaultTemplate string `json:"defaultTemplate,omitempty"`
	Id              string `json:"id,omitempty"`
	InternalId      string `json:"internalId,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Name string `json:"name,omitempty"`
	// Associated agent pool locator.
	Pool string `json:"pool,omitempty"`
	// Project (direct parent) locator.
	Project string `json:"project,omitempty"`
	// Project feature locator.
	ProjectFeature string `json:"projectFeature,omitempty"`
	// For paginated calls, from which entity to start rendering the page.
	Start int32  `json:"start,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
	// VCS root locator.
	VcsRoot string `json:"vcsRoot,omitempty"`
}

// Represents a locator string for filtering User entities.
type UserLocator struct {
	// User group (direct or indirect parent) locator.
	AffectedGroup string `json:"affectedGroup,omitempty"`
	Email         string `json:"email,omitempty"`
	// User group (direct parent) locator.
	Group string `json:"group,omitempty"`
	Id    string `json:"id,omitempty"`
	// Supply multiple locators and return a union of the results.
	Item      string    `json:"item,omitempty"`
	LastLogin time.Time `json:"lastLogin,omitempty"`
	Name      string    `json:"name,omitempty"`
	Property  string    `json:"property,omitempty"`
	// Role locator.
	Role     string `json:"role,omitempty"`
	Username string `json:"username,omitempty"`
}

// Represents a locator string for filtering Group entities.
type UserGroupLocator struct {
	// Supply multiple locators and return a union of the results.
	Item string `json:"item,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}