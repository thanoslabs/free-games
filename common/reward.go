package common

import (
 "strings"
 "time"
)

// RewardType distinguishes games from cosmetic rewards.
type RewardType string
const (
 RewardGame RewardType = "game"
 RewardCape RewardType = "cape"
)

// RewardStatus describes the acquisition window.
type RewardStatus string
const (
 RewardUpcoming RewardStatus = "upcoming"
 RewardActive RewardStatus = "active"
 RewardExpired RewardStatus = "expired"
)

// Reward represents an offer independently of its discovery source.
type Reward struct {
 ID string `json:"id"`
 Title string `json:"title"`
 Description string `json:"description,omitempty"`
 ImageURL string `json:"imageUrl,omitempty"`
 ClaimURL string `json:"claimUrl,omitempty"`
 SourceURL string `json:"sourceUrl,omitempty"`
 Type RewardType `json:"type"`
 Provider string `json:"provider"`
 Platform string `json:"platform,omitempty"`
 RequiresMembership bool `json:"requiresMembership,omitempty"`
 Requirements string `json:"requirements,omitempty"`
 Editions []string `json:"editions,omitempty"`
 StartDate time.Time `json:"startDate,omitempty"`
 EndDate time.Time `json:"endDate,omitempty"`
}

// Status returns the availability state of the reward at the given time.
func (r Reward) Status(now time.Time) RewardStatus {
 if !r.StartDate.IsZero() && now.Before(r.StartDate) { return RewardUpcoming }
 if !r.EndDate.IsZero() && !now.Before(r.EndDate) { return RewardExpired }
 return RewardActive
}

// SupportsMinecraftEdition reports whether a cape applies to Java or Bedrock.
func (r Reward) SupportsMinecraftEdition() bool {
 if r.Type != RewardCape { return false }
 for _, edition := range r.Editions {
  switch strings.ToLower(strings.TrimSpace(edition)) {
  case "java", "bedrock": return true
  }
 }
 return false
}
