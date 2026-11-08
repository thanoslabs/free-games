package common

import (
 "testing"
 "time"
)
func TestRewardStatus(t *testing.T) {
 now:=time.Date(2026,11,7,12,0,0,0,time.UTC)
 cases:=[]struct{name string;r Reward;want RewardStatus}{
 {"upcoming",Reward{StartDate:now.Add(time.Hour)},RewardUpcoming},
 {"active",Reward{StartDate:now.Add(-time.Hour),EndDate:now.Add(time.Hour)},RewardActive},
 {"expired",Reward{EndDate:now},RewardExpired},
 {"undated",Reward{},RewardActive},
 }
 for _,tt:=range cases {t.Run(tt.name,func(t *testing.T){if got:=tt.r.Status(now);got!=tt.want{t.Fatalf("got %s, want %s",got,tt.want)}})}
}
func TestCapeEditions(t *testing.T) {
 if !(Reward{Type:RewardCape,Editions:[]string{"Java"}}).SupportsMinecraftEdition(){t.Fatal("Java cape not supported")}
 if (Reward{Type:RewardGame,Editions:[]string{"Bedrock"}}).SupportsMinecraftEdition(){t.Fatal("game is not a cape")}
 if (Reward{Type:RewardCape,Editions:[]string{"Dungeons"}}).SupportsMinecraftEdition(){t.Fatal("Dungeons-only cape should be excluded")}
}
