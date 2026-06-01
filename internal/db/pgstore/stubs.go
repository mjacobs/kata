// Phase 3 stubs: every db.Storage domain method returns ErrNotImplementedPhase3.
// Phase 4 replaces each entity group's stubs with real queries.

package pgstore

import (
	"context"
	"iter"
	"time"

	"go.kenn.io/kata/internal/db"
)

// Compile-time guarantee that *Store satisfies db.Storage. The identity and
// lifecycle methods live on store.go and the Migrate impl lives on migrate.go;
// the remaining ~180 domain methods are stubbed below.
var _ db.Storage = (*Store)(nil)

// ----- projects + aliases -----

// CreateProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateProject(_ context.Context, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// CreateProjectWithUID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateProjectWithUID(_ context.Context, _, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// ProjectByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ProjectByID(_ context.Context, _ int64) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// ProjectByName is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ProjectByName(_ context.Context, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// ProjectByNameIncludingArchived is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ProjectByNameIncludingArchived(_ context.Context, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// ProjectByUID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ProjectByUID(_ context.Context, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// ListProjects is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListProjects(_ context.Context) ([]db.Project, error) {
	return nil, ErrNotImplementedPhase3
}

// ListProjectsIncludingArchived is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListProjectsIncludingArchived(_ context.Context) ([]db.Project, error) {
	return nil, ErrNotImplementedPhase3
}

// RenameProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RenameProject(_ context.Context, _ int64, _ string) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// RemoveProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RemoveProject(_ context.Context, _ db.RemoveProjectParams) (db.Project, *db.Event, error) {
	return db.Project{}, nil, ErrNotImplementedPhase3
}

// RestoreProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RestoreProject(_ context.Context, _ int64, _ string) (db.Project, *db.Event, bool, error) {
	return db.Project{}, nil, false, ErrNotImplementedPhase3
}

// HardDeleteProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) HardDeleteProject(_ context.Context, _ int64) error {
	return ErrNotImplementedPhase3
}

// MergeProjects is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MergeProjects(_ context.Context, _ db.MergeProjectsParams) (db.ProjectMergeResult, error) {
	return db.ProjectMergeResult{}, ErrNotImplementedPhase3
}

// MoveIssueProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MoveIssueProject(_ context.Context, _ db.MoveIssueProjectIn) (db.MoveIssueProjectOut, error) {
	return db.MoveIssueProjectOut{}, ErrNotImplementedPhase3
}

// PatchProjectMetadata is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PatchProjectMetadata(_ context.Context, _ db.PatchProjectMetadataIn) (db.PatchProjectMetadataOut, error) {
	return db.PatchProjectMetadataOut{}, ErrNotImplementedPhase3
}

// BatchProjectStats is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) BatchProjectStats(_ context.Context) (map[int64]db.ProjectStats, error) {
	return nil, ErrNotImplementedPhase3
}

// AliasByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AliasByID(_ context.Context, _ int64) (db.ProjectAlias, error) {
	return db.ProjectAlias{}, ErrNotImplementedPhase3
}

// AliasByIdentity is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AliasByIdentity(_ context.Context, _ string) (db.ProjectAlias, error) {
	return db.ProjectAlias{}, ErrNotImplementedPhase3
}

// AttachAlias is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AttachAlias(_ context.Context, _ int64, _, _, _ string) (db.ProjectAlias, error) {
	return db.ProjectAlias{}, ErrNotImplementedPhase3
}

// ReassignAlias is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ReassignAlias(_ context.Context, _, _ int64, _ string) error {
	return ErrNotImplementedPhase3
}

// DetachProjectAlias is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) DetachProjectAlias(_ context.Context, _ db.DetachAliasParams) (db.ProjectAlias, *db.Event, error) {
	return db.ProjectAlias{}, nil, ErrNotImplementedPhase3
}

// TouchAlias is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) TouchAlias(_ context.Context, _ int64, _ string) error {
	return ErrNotImplementedPhase3
}

// ProjectAliases is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ProjectAliases(_ context.Context, _ int64) ([]db.ProjectAlias, error) {
	return nil, ErrNotImplementedPhase3
}

// LatestAliasForProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LatestAliasForProject(_ context.Context, _ int64) (db.AliasRow, bool, error) {
	return db.AliasRow{}, false, ErrNotImplementedPhase3
}

// ----- issues -----

// CreateIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateIssue(_ context.Context, _ db.CreateIssueParams) (db.Issue, db.Event, error) {
	return db.Issue{}, db.Event{}, ErrNotImplementedPhase3
}

// IssueByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) IssueByID(_ context.Context, _ int64) (db.Issue, error) {
	return db.Issue{}, ErrNotImplementedPhase3
}

// IssueByShortID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) IssueByShortID(_ context.Context, _ int64, _ string, _ db.IncludeDeleted) (db.Issue, error) {
	return db.Issue{}, ErrNotImplementedPhase3
}

// IssueByUID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) IssueByUID(_ context.Context, _ string, _ db.IncludeDeleted) (db.Issue, error) {
	return db.Issue{}, ErrNotImplementedPhase3
}

// IssueUIDPrefixMatch is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) IssueUIDPrefixMatch(_ context.Context, _ string, _ int, _ db.IncludeDeleted) ([]db.Issue, error) {
	return nil, ErrNotImplementedPhase3
}

// ListIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListIssues(_ context.Context, _ db.ListIssuesParams) ([]db.Issue, error) {
	return nil, ErrNotImplementedPhase3
}

// ListAllIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListAllIssues(_ context.Context, _ db.ListAllIssuesParams) ([]db.Issue, error) {
	return nil, ErrNotImplementedPhase3
}

// ReadyIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ReadyIssues(_ context.Context, _ int64, _ int, _ db.ReadyIssuesFilter) ([]db.Issue, error) {
	return nil, ErrNotImplementedPhase3
}

// ReadyIssuesGlobal is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ReadyIssuesGlobal(_ context.Context, _ int) ([]db.ReadyGlobalIssue, error) {
	return nil, ErrNotImplementedPhase3
}

// ChildrenOfIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ChildrenOfIssue(_ context.Context, _, _ int64) ([]db.Issue, error) {
	return nil, ErrNotImplementedPhase3
}

// OpenChildrenOf is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) OpenChildrenOf(_ context.Context, _, _ int64, _ int) ([]db.Issue, int, error) {
	return nil, 0, ErrNotImplementedPhase3
}

// EditIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EditIssue(_ context.Context, _ db.EditIssueParams) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// EditIssueAtomic is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EditIssueAtomic(_ context.Context, _ db.EditIssueAtomicParams) (db.EditIssueAtomicResult, error) {
	return db.EditIssueAtomicResult{}, ErrNotImplementedPhase3
}

// CloseIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CloseIssue(_ context.Context, _ int64, _, _, _ string, _ []db.Evidence) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// CloseIssueWithEvents is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CloseIssueWithEvents(_ context.Context, _ int64, _, _, _ string, _ []db.Evidence) (db.Issue, []db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// ReopenIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ReopenIssue(_ context.Context, _ int64, _ string) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// SoftDeleteIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SoftDeleteIssue(_ context.Context, _ int64, _ string) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// RestoreIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RestoreIssue(_ context.Context, _ int64, _ string) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// PurgeIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PurgeIssue(_ context.Context, _ int64, _ string, _ *string) (db.PurgeLog, error) {
	return db.PurgeLog{}, ErrNotImplementedPhase3
}

// ClaimOwner is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClaimOwner(_ context.Context, _ int64, _ string, _ bool) (db.ClaimResult, error) {
	return db.ClaimResult{}, ErrNotImplementedPhase3
}

// UpdateOwner is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UpdateOwner(_ context.Context, _ int64, _ *string, _ string) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// UpdatePriority is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UpdatePriority(_ context.Context, _ int64, _ *int64, _ string) (db.Issue, *db.Event, bool, error) {
	return db.Issue{}, nil, false, ErrNotImplementedPhase3
}

// PatchIssueMetadata is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PatchIssueMetadata(_ context.Context, _ db.PatchIssueMetadataIn) (db.PatchIssueMetadataOut, error) {
	return db.PatchIssueMetadataOut{}, ErrNotImplementedPhase3
}

// ShortIDsByUIDs is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ShortIDsByUIDs(_ context.Context, _ int64, _ []string) (map[string]string, error) {
	return nil, ErrNotImplementedPhase3
}

// PurgeResetCheck is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PurgeResetCheck(_ context.Context, _, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// ----- comments -----

// CreateComment is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateComment(_ context.Context, _ db.CreateCommentParams) (db.Comment, db.Event, error) {
	return db.Comment{}, db.Event{}, ErrNotImplementedPhase3
}

// CommentBodyByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CommentBodyByID(_ context.Context, _ int64) (string, error) {
	return "", ErrNotImplementedPhase3
}

// CommentsByIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CommentsByIssue(_ context.Context, _ int64) ([]db.Comment, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- labels -----

// AddLabel is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AddLabel(_ context.Context, _ int64, _, _ string) (db.IssueLabel, error) {
	return db.IssueLabel{}, ErrNotImplementedPhase3
}

// AddLabelAndEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AddLabelAndEvent(_ context.Context, _ int64, _ db.LabelEventParams) (db.IssueLabel, db.Event, error) {
	return db.IssueLabel{}, db.Event{}, ErrNotImplementedPhase3
}

// RemoveLabel is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RemoveLabel(_ context.Context, _ int64, _ string) error {
	return ErrNotImplementedPhase3
}

// RemoveLabelAndEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RemoveLabelAndEvent(_ context.Context, _ int64, _ db.LabelEventParams) (db.Event, error) {
	return db.Event{}, ErrNotImplementedPhase3
}

// HasLabel is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) HasLabel(_ context.Context, _ int64, _ string) (bool, error) {
	return false, ErrNotImplementedPhase3
}

// LabelByEndpoints is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LabelByEndpoints(_ context.Context, _ int64, _ string) (db.IssueLabel, error) {
	return db.IssueLabel{}, ErrNotImplementedPhase3
}

// LabelCounts is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LabelCounts(_ context.Context, _ int64) ([]db.LabelCount, error) {
	return nil, ErrNotImplementedPhase3
}

// LabelsByIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LabelsByIssue(_ context.Context, _ int64) ([]db.IssueLabel, error) {
	return nil, ErrNotImplementedPhase3
}

// LabelsByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LabelsByIssues(_ context.Context, _ int64, _ []int64) (map[int64][]string, error) {
	return nil, ErrNotImplementedPhase3
}

// LabelsForIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LabelsForIssue(_ context.Context, _ int64) ([]string, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- links -----

// CreateLink is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateLink(_ context.Context, _ db.CreateLinkParams) (db.Link, error) {
	return db.Link{}, ErrNotImplementedPhase3
}

// CreateLinkAndEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateLinkAndEvent(_ context.Context, _ db.CreateLinkParams, _ db.LinkEventParams) (db.Link, db.Event, error) {
	return db.Link{}, db.Event{}, ErrNotImplementedPhase3
}

// DeleteLinkByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) DeleteLinkByID(_ context.Context, _ int64) error {
	return ErrNotImplementedPhase3
}

// DeleteLinkAndEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) DeleteLinkAndEvent(_ context.Context, _ db.Link, _ db.LinkEventParams) (db.Event, error) {
	return db.Event{}, ErrNotImplementedPhase3
}

// LinkByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LinkByID(_ context.Context, _ int64) (db.Link, error) {
	return db.Link{}, ErrNotImplementedPhase3
}

// LinkByEndpoints is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LinkByEndpoints(_ context.Context, _, _ int64, _ string) (db.Link, error) {
	return db.Link{}, ErrNotImplementedPhase3
}

// LinksByIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LinksByIssue(_ context.Context, _ int64) ([]db.Link, error) {
	return nil, ErrNotImplementedPhase3
}

// ParentOf is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ParentOf(_ context.Context, _ int64) (db.Link, error) {
	return db.Link{}, ErrNotImplementedPhase3
}

// ChildCountsByParents is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ChildCountsByParents(_ context.Context, _ int64, _ []int64) (map[int64]db.ChildCounts, error) {
	return nil, ErrNotImplementedPhase3
}

// ParentNumbersByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ParentNumbersByIssues(_ context.Context, _ int64, _ []int64) (map[int64]int64, error) {
	return nil, ErrNotImplementedPhase3
}

// ParentShortIDsByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ParentShortIDsByIssues(_ context.Context, _ int64, _ []int64) (map[int64]string, error) {
	return nil, ErrNotImplementedPhase3
}

// BlockNumbersByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) BlockNumbersByIssues(_ context.Context, _ int64, _ []int64) (map[int64][]int64, error) {
	return nil, ErrNotImplementedPhase3
}

// BlockedByNumbersByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) BlockedByNumbersByIssues(_ context.Context, _ int64, _ []int64) (map[int64][]int64, error) {
	return nil, ErrNotImplementedPhase3
}

// RelatedNumbersByIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RelatedNumbersByIssues(_ context.Context, _ int64, _ []int64) (map[int64][]int64, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- recurrences -----

// CreateRecurrence is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateRecurrence(_ context.Context, _ db.CreateRecurrenceIn) (db.Recurrence, error) {
	return db.Recurrence{}, ErrNotImplementedPhase3
}

// GetRecurrenceByID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) GetRecurrenceByID(_ context.Context, _ int64) (db.Recurrence, error) {
	return db.Recurrence{}, ErrNotImplementedPhase3
}

// GetRecurrenceByUID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) GetRecurrenceByUID(_ context.Context, _ string) (db.Recurrence, error) {
	return db.Recurrence{}, ErrNotImplementedPhase3
}

// ListRecurrencesByProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListRecurrencesByProject(_ context.Context, _ int64) ([]db.Recurrence, error) {
	return nil, ErrNotImplementedPhase3
}

// PatchRecurrence is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PatchRecurrence(_ context.Context, _ db.PatchRecurrenceIn) (db.PatchRecurrenceOut, error) {
	return db.PatchRecurrenceOut{}, ErrNotImplementedPhase3
}

// SoftDeleteRecurrence is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SoftDeleteRecurrence(_ context.Context, _ int64, _ string) error {
	return ErrNotImplementedPhase3
}

// MaterializeNext is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MaterializeNext(_ context.Context, _ int64, _, _ string) (db.MaterializeNextOut, error) {
	return db.MaterializeNextOut{}, ErrNotImplementedPhase3
}

// ----- events / idempotency / close-throttle -----

// EventsAfter is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EventsAfter(_ context.Context, _ db.EventsAfterParams) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// EventsByUIDs is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EventsByUIDs(_ context.Context, _ int64, _ []string) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// EventsInWindow is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EventsInWindow(_ context.Context, _ db.EventsInWindowParams) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// MaxEventID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MaxEventID(_ context.Context) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// MaxLocalOriginEventID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MaxLocalOriginEventID(_ context.Context, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// MaxFederationBaselineEventID is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MaxFederationBaselineEventID(_ context.Context, _, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// LookupIdempotency is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) LookupIdempotency(_ context.Context, _ int64, _ string, _ time.Time) (*db.IdempotencyMatch, error) {
	return nil, ErrNotImplementedPhase3
}

// InsertCloseThrottledEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) InsertCloseThrottledEvent(_ context.Context, _ int64, _ string, _ db.CloseThrottledPayload) (db.Event, error) {
	return db.Event{}, ErrNotImplementedPhase3
}

// RecentSiblingCloses is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecentSiblingCloses(_ context.Context, _, _, _ int64, _ string, _ time.Time) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// RecentSameMessageClose is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecentSameMessageClose(_ context.Context, _, _, _ int64, _, _ string, _ time.Time) (*db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- search -----

// SearchFTS is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SearchFTS(_ context.Context, _ int64, _ string, _ int, _ bool) ([]db.SearchCandidate, error) {
	return nil, ErrNotImplementedPhase3
}

// SearchFTSAny is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SearchFTSAny(_ context.Context, _ int64, _ string, _ int, _ bool) ([]db.SearchCandidate, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- import support -----

// ImportBatch is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ImportBatch(_ context.Context, _ db.ImportBatchParams) (db.ImportBatchResult, []db.Event, error) {
	return db.ImportBatchResult{}, nil, ErrNotImplementedPhase3
}

// UpsertImportMapping is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UpsertImportMapping(_ context.Context, _ db.ImportMappingParams) (db.ImportMapping, error) {
	return db.ImportMapping{}, ErrNotImplementedPhase3
}

// ImportMappingBySource is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ImportMappingBySource(_ context.Context, _ int64, _, _, _ string) (db.ImportMapping, error) {
	return db.ImportMapping{}, ErrNotImplementedPhase3
}

// ImportMappingsByProjectSource is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ImportMappingsByProjectSource(_ context.Context, _ int64, _ string) ([]db.ImportMapping, error) {
	return nil, ErrNotImplementedPhase3
}

// ImportReplay is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ImportReplay(_ context.Context, _ []db.ImportRecord, _ db.ImportOptions) error {
	return ErrNotImplementedPhase3
}

// ----- API tokens / system project -----

// EnsureSystemProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EnsureSystemProject(_ context.Context) error {
	return ErrNotImplementedPhase3
}

// SystemProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SystemProject(_ context.Context) (db.Project, error) {
	return db.Project{}, ErrNotImplementedPhase3
}

// CreateAPIToken is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateAPIToken(_ context.Context, _ db.CreateAPITokenParams) (db.APIToken, db.Event, error) {
	return db.APIToken{}, db.Event{}, ErrNotImplementedPhase3
}

// RevokeAPIToken is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RevokeAPIToken(_ context.Context, _ int64, _ string) (db.APIToken, db.Event, error) {
	return db.APIToken{}, db.Event{}, ErrNotImplementedPhase3
}

// ResolveAPIToken is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ResolveAPIToken(_ context.Context, _ string) (db.APIToken, error) {
	return db.APIToken{}, ErrNotImplementedPhase3
}

// ListAPITokens is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListAPITokens(_ context.Context) ([]db.APIToken, error) {
	return nil, ErrNotImplementedPhase3
}

// ----- claims -----

// AcquireClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AcquireClaim(_ context.Context, _ db.AcquireClaimParams) (db.LeaseResult, error) {
	return db.LeaseResult{}, ErrNotImplementedPhase3
}

// RenewClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RenewClaim(_ context.Context, _ db.RenewClaimParams) (db.LeaseResult, error) {
	return db.LeaseResult{}, ErrNotImplementedPhase3
}

// ReleaseClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ReleaseClaim(_ context.Context, _ db.ReleaseClaimParams) (db.LeaseResult, error) {
	return db.LeaseResult{}, ErrNotImplementedPhase3
}

// ForceReleaseClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ForceReleaseClaim(_ context.Context, _ db.ForceReleaseClaimParams) (db.LeaseResult, error) {
	return db.LeaseResult{}, ErrNotImplementedPhase3
}

// ClaimStatus is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClaimStatus(_ context.Context, _ int64, _ string, _ time.Time) (db.ClaimStatus, error) {
	return db.ClaimStatus{}, ErrNotImplementedPhase3
}

// ClaimStatusReadOnly is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClaimStatusReadOnly(_ context.Context, _ int64, _ string, _ time.Time) (db.ClaimStatus, error) {
	return db.ClaimStatus{}, ErrNotImplementedPhase3
}

// EnqueuePendingClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EnqueuePendingClaim(_ context.Context, _ db.PendingClaimParams) (db.PendingClaimRequest, error) {
	return db.PendingClaimRequest{}, ErrNotImplementedPhase3
}

// ResolvePendingClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ResolvePendingClaim(_ context.Context, _ string, _ db.IssueClaim) error {
	return ErrNotImplementedPhase3
}

// RejectPendingClaim is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RejectPendingClaim(_ context.Context, _, _ string, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// ListPendingClaimRequests is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListPendingClaimRequests(_ context.Context, _ int64, _ int) ([]db.PendingClaimRequest, error) {
	return nil, ErrNotImplementedPhase3
}

// ListPendingClaimRequestsForIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListPendingClaimRequestsForIssue(_ context.Context, _ int64, _ string, _ int) ([]db.PendingClaimRequest, error) {
	return nil, ErrNotImplementedPhase3
}

// CountLiveClaims is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CountLiveClaims(_ context.Context, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// CountPendingClaims is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CountPendingClaims(_ context.Context, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// MarkPendingClaimAttempt is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MarkPendingClaimAttempt(_ context.Context, _, _ string, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// ClaimStatusRefreshError is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClaimStatusRefreshError(_ context.Context, _ int64, _ string) (db.ClaimStatusRefreshError, error) {
	return db.ClaimStatusRefreshError{}, ErrNotImplementedPhase3
}

// MarkClaimStatusRefreshError is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MarkClaimStatusRefreshError(_ context.Context, _ int64, _ string, _ int, _ string, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// ClearClaimStatusRefreshError is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClearClaimStatusRefreshError(_ context.Context, _ int64, _ string) error {
	return ErrNotImplementedPhase3
}

// UpsertClaimCache is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UpsertClaimCache(_ context.Context, _ db.IssueClaim) error {
	return ErrNotImplementedPhase3
}

// ApplyClaimStatus is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ApplyClaimStatus(_ context.Context, _ int64, _ string, _ db.ClaimStatus) error {
	return ErrNotImplementedPhase3
}

// CheckClaimGate is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CheckClaimGate(_ context.Context, _ db.ClaimGateParams) error {
	return ErrNotImplementedPhase3
}

// ExpireTimedClaims is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExpireTimedClaims(_ context.Context, _ time.Time, _ int) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// ExpireTimedClaimsForProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExpireTimedClaimsForProject(_ context.Context, _ int64, _ time.Time, _ int) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// UnresolvedClaimViolationsForIssue is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UnresolvedClaimViolationsForIssue(_ context.Context, _ int64, _ string, _ int) ([]db.ClaimViolationSummary, int64, error) {
	return nil, 0, ErrNotImplementedPhase3
}

// UnresolvedClaimViolationsForProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UnresolvedClaimViolationsForProject(_ context.Context, _ int64, _ int) ([]db.ClaimViolationSummary, int64, error) {
	return nil, 0, ErrNotImplementedPhase3
}

// ----- federation: bindings + sync status + quarantines -----

// ListFederationBindings is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListFederationBindings(_ context.Context) ([]db.FederationBinding, error) {
	return nil, ErrNotImplementedPhase3
}

// FederationBindingByProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) FederationBindingByProject(_ context.Context, _ int64) (db.FederationBinding, error) {
	return db.FederationBinding{}, ErrNotImplementedPhase3
}

// FederationSyncStatusByProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) FederationSyncStatusByProject(_ context.Context, _ int64) (db.FederationSyncStatus, error) {
	return db.FederationSyncStatus{}, ErrNotImplementedPhase3
}

// RecordFederationSyncPullStarted is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncPullStarted(_ context.Context, _ int64, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// RecordFederationSyncPullSuccess is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncPullSuccess(_ context.Context, _ int64, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// RecordFederationSyncPushStarted is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncPushStarted(_ context.Context, _ int64, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// RecordFederationSyncPushSuccess is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncPushSuccess(_ context.Context, _ int64, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// RecordFederationSyncReset is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncReset(_ context.Context, _ int64, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// RecordFederationSyncError is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationSyncError(_ context.Context, _ int64, _ error, _ time.Time) error {
	return ErrNotImplementedPhase3
}

// ClearFederationSyncError is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ClearFederationSyncError(_ context.Context, _ int64) error {
	return ErrNotImplementedPhase3
}

// RecordFederationQuarantine is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RecordFederationQuarantine(_ context.Context, _ db.RecordFederationQuarantineParams) (db.FederationQuarantine, error) {
	return db.FederationQuarantine{}, ErrNotImplementedPhase3
}

// ActiveFederationQuarantine is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ActiveFederationQuarantine(_ context.Context, _ int64, _ db.FederationQuarantineDirection) (db.FederationQuarantine, error) {
	return db.FederationQuarantine{}, ErrNotImplementedPhase3
}

// ActiveFederationQuarantinesByProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ActiveFederationQuarantinesByProject(_ context.Context, _ int64) ([]db.FederationQuarantine, error) {
	return nil, ErrNotImplementedPhase3
}

// CountActiveFederationEnrollments is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CountActiveFederationEnrollments(_ context.Context, _ int64) (int64, error) {
	return 0, ErrNotImplementedPhase3
}

// SkipFederationQuarantine is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) SkipFederationQuarantine(_ context.Context, _ db.SkipFederationQuarantineParams) (db.FederationQuarantine, error) {
	return db.FederationQuarantine{}, ErrNotImplementedPhase3
}

// UpsertFederationBinding is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) UpsertFederationBinding(_ context.Context, _ db.FederationBinding) (db.FederationBinding, error) {
	return db.FederationBinding{}, ErrNotImplementedPhase3
}

// AdvanceFederationPullCursor is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AdvanceFederationPullCursor(_ context.Context, _, _ int64) error {
	return ErrNotImplementedPhase3
}

// InsertRemoteEvent is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) InsertRemoteEvent(_ context.Context, _ int64, _ db.RemoteEvent) (bool, error) {
	return false, ErrNotImplementedPhase3
}

// EnableProjectFederation is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EnableProjectFederation(_ context.Context, _ int64, _ string) (db.FederationBinding, error) {
	return db.FederationBinding{}, ErrNotImplementedPhase3
}

// RefreshProjectFederationBaseline is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RefreshProjectFederationBaseline(_ context.Context, _ int64, _ string) (db.FederationBinding, bool, error) {
	return db.FederationBinding{}, false, ErrNotImplementedPhase3
}

// MaterializeFederatedProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) MaterializeFederatedProject(_ context.Context, _ int64) error {
	return ErrNotImplementedPhase3
}

// ResetFederatedProject is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ResetFederatedProject(_ context.Context, _, _, _ int64) error {
	return ErrNotImplementedPhase3
}

// ----- federation: enrollments -----

// CreateFederationEnrollment is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) CreateFederationEnrollment(_ context.Context, _ db.CreateFederationEnrollmentParams) (db.CreatedFederationEnrollment, error) {
	return db.CreatedFederationEnrollment{}, ErrNotImplementedPhase3
}

// ListFederationEnrollments is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ListFederationEnrollments(_ context.Context) ([]db.FederationEnrollment, error) {
	return nil, ErrNotImplementedPhase3
}

// RevokeFederationEnrollment is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) RevokeFederationEnrollment(_ context.Context, _ int64) error {
	return ErrNotImplementedPhase3
}

// AuthorizeFederationToken is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AuthorizeFederationToken(_ context.Context, _ string, _ int64, _ string) (db.FederationEnrollment, error) {
	return db.FederationEnrollment{}, ErrNotImplementedPhase3
}

// ----- export (JSONL) -----
//
// Export methods return iter.Seq2 (no outer error), so a stubbed iterator
// must yield the sentinel error as the K=zero, V=ErrNotImplementedPhase3
// shape. Callers (jsonl export) drain the iterator; first-yield error
// surfaces as a normal io path failure.

// ExportMeta is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportMeta(_ context.Context) iter.Seq2[db.MetaKV, error] {
	return func(yield func(db.MetaKV, error) bool) {
		yield(db.MetaKV{}, ErrNotImplementedPhase3)
	}
}

// ExportProjects is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportProjects(_ context.Context, _ db.ExportFilter) iter.Seq2[db.ProjectExport, error] {
	return func(yield func(db.ProjectExport, error) bool) {
		yield(db.ProjectExport{}, ErrNotImplementedPhase3)
	}
}

// ExportProjectAliases is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportProjectAliases(_ context.Context, _ db.ExportFilter) iter.Seq2[db.AliasExport, error] {
	return func(yield func(db.AliasExport, error) bool) {
		yield(db.AliasExport{}, ErrNotImplementedPhase3)
	}
}

// ExportRecurrences is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportRecurrences(_ context.Context, _ db.ExportFilter) iter.Seq2[db.RecurrenceExport, error] {
	return func(yield func(db.RecurrenceExport, error) bool) {
		yield(db.RecurrenceExport{}, ErrNotImplementedPhase3)
	}
}

// ExportIssues is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportIssues(_ context.Context, _ db.ExportFilter) iter.Seq2[db.IssueExport, error] {
	return func(yield func(db.IssueExport, error) bool) {
		yield(db.IssueExport{}, ErrNotImplementedPhase3)
	}
}

// ExportComments is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportComments(_ context.Context, _ db.ExportFilter) iter.Seq2[db.CommentExport, error] {
	return func(yield func(db.CommentExport, error) bool) {
		yield(db.CommentExport{}, ErrNotImplementedPhase3)
	}
}

// ExportIssueLabels is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportIssueLabels(_ context.Context, _ db.ExportFilter) iter.Seq2[db.IssueLabelExport, error] {
	return func(yield func(db.IssueLabelExport, error) bool) {
		yield(db.IssueLabelExport{}, ErrNotImplementedPhase3)
	}
}

// ExportLinks is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportLinks(_ context.Context, _ db.ExportFilter) iter.Seq2[db.LinkExport, error] {
	return func(yield func(db.LinkExport, error) bool) {
		yield(db.LinkExport{}, ErrNotImplementedPhase3)
	}
}

// ExportImportMappings is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportImportMappings(_ context.Context, _ db.ExportFilter) iter.Seq2[db.ImportMappingExport, error] {
	return func(yield func(db.ImportMappingExport, error) bool) {
		yield(db.ImportMappingExport{}, ErrNotImplementedPhase3)
	}
}

// ExportFederationBindings is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportFederationBindings(_ context.Context, _ db.ExportFilter) iter.Seq2[db.FederationBindingExport, error] {
	return func(yield func(db.FederationBindingExport, error) bool) {
		yield(db.FederationBindingExport{}, ErrNotImplementedPhase3)
	}
}

// ExportFederationSyncStatus is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportFederationSyncStatus(_ context.Context, _ db.ExportFilter) iter.Seq2[db.FederationSyncStatusExport, error] {
	return func(yield func(db.FederationSyncStatusExport, error) bool) {
		yield(db.FederationSyncStatusExport{}, ErrNotImplementedPhase3)
	}
}

// ExportFederationQuarantine is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportFederationQuarantine(_ context.Context, _ db.ExportFilter) iter.Seq2[db.FederationQuarantineExport, error] {
	return func(yield func(db.FederationQuarantineExport, error) bool) {
		yield(db.FederationQuarantineExport{}, ErrNotImplementedPhase3)
	}
}

// ExportFederationEnrollments is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportFederationEnrollments(_ context.Context, _ db.ExportFilter) iter.Seq2[db.FederationEnrollmentExport, error] {
	return func(yield func(db.FederationEnrollmentExport, error) bool) {
		yield(db.FederationEnrollmentExport{}, ErrNotImplementedPhase3)
	}
}

// ExportIssueClaims is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportIssueClaims(_ context.Context, _ db.ExportFilter) iter.Seq2[db.IssueClaimExport, error] {
	return func(yield func(db.IssueClaimExport, error) bool) {
		yield(db.IssueClaimExport{}, ErrNotImplementedPhase3)
	}
}

// ExportPendingClaimRequests is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportPendingClaimRequests(_ context.Context, _ db.ExportFilter) iter.Seq2[db.PendingClaimRequestExport, error] {
	return func(yield func(db.PendingClaimRequestExport, error) bool) {
		yield(db.PendingClaimRequestExport{}, ErrNotImplementedPhase3)
	}
}

// ExportEvents is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportEvents(_ context.Context, _ db.ExportFilter) iter.Seq2[db.EventExport, error] {
	return func(yield func(db.EventExport, error) bool) {
		yield(db.EventExport{}, ErrNotImplementedPhase3)
	}
}

// ExportPurgeLog is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportPurgeLog(_ context.Context, _ db.ExportFilter) iter.Seq2[db.PurgeLogExport, error] {
	return func(yield func(db.PurgeLogExport, error) bool) {
		yield(db.PurgeLogExport{}, ErrNotImplementedPhase3)
	}
}

// ExportSequences is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ExportSequences(_ context.Context) iter.Seq2[db.SequenceExport, error] {
	return func(yield func(db.SequenceExport, error) bool) {
		yield(db.SequenceExport{}, ErrNotImplementedPhase3)
	}
}

// ----- federation: push + ingest -----

// PendingFederationPushEvents is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PendingFederationPushEvents(_ context.Context, _ int64, _ string, _ int64, _ int) ([]db.Event, error) {
	return nil, ErrNotImplementedPhase3
}

// PendingFederationPushStats is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) PendingFederationPushStats(_ context.Context, _ int64, _ string, _ int64) (int64, int64, error) {
	return 0, 0, ErrNotImplementedPhase3
}

// AdvanceFederationPushCursor is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AdvanceFederationPushCursor(_ context.Context, _, _ int64) error {
	return ErrNotImplementedPhase3
}

// EnableFederationPush is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) EnableFederationPush(_ context.Context, _ int64, _ int64) (db.FederationBinding, error) {
	return db.FederationBinding{}, ErrNotImplementedPhase3
}

// ResetFederatedProjectIfNoPendingPush is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) ResetFederatedProjectIfNoPendingPush(_ context.Context, _, _, _ int64, _ string, _ int64) error {
	return ErrNotImplementedPhase3
}

// IngestFederationEvents is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) IngestFederationEvents(_ context.Context, _ db.FederationIngestParams) (db.FederationIngestResult, error) {
	return db.FederationIngestResult{}, ErrNotImplementedPhase3
}

// AdoptProjectIntoFederation is a Phase 3 stub; Phase 4 replaces it with a real query.
func (s *Store) AdoptProjectIntoFederation(_ context.Context, _ db.AdoptProjectIntoFederationParams) (db.AdoptProjectIntoFederationResult, error) {
	return db.AdoptProjectIntoFederationResult{}, ErrNotImplementedPhase3
}
