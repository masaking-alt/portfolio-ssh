package portfolio

import "testing"

func TestDefaultProfileHasDescendingWorks(t *testing.T) {
	profile := DefaultProfile()
	if len(profile.Works) != 10 {
		t.Fatalf("作品数 = %d, want 10", len(profile.Works))
	}

	for i := 1; i < len(profile.Works); i++ {
		if profile.Works[i-1].ID <= profile.Works[i].ID {
			t.Fatalf("作品IDが降順ではありません: %d, %d", profile.Works[i-1].ID, profile.Works[i].ID)
		}
	}
}

func TestDefaultProfileHasContact(t *testing.T) {
	profile := DefaultProfile()
	if profile.Email == "" {
		t.Fatal("メールアドレスが空です")
	}
	if len(profile.Socials) == 0 {
		t.Fatal("SNSリンクが空です")
	}
}
