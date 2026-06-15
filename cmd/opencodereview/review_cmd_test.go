package main

import (
	"strings"
	"testing"
)

func TestValidateReviewRefsRejectsOptionLikeCommit(t *testing.T) {
	err := validateReviewRefs(t.TempDir(), reviewOptions{commit: "-O./pwn.sh"})
	if err == nil {
		t.Fatal("expected option-like --commit ref to be rejected")
	}
	if !strings.Contains(err.Error(), "--commit") || !strings.Contains(err.Error(), "must not start with '-'") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateReviewRefsRejectsOptionLikeRangeRef(t *testing.T) {
	err := validateReviewRefs(t.TempDir(), reviewOptions{to: "-O./pwn.sh"})
	if err == nil {
		t.Fatal("expected option-like --to ref to be rejected")
	}
	if !strings.Contains(err.Error(), "--to") || !strings.Contains(err.Error(), "must not start with '-'") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseReviewFlagsTranslateZH(t *testing.T) {
	opts, err := parseReviewFlags([]string{"--translate-zh"})
	if err != nil {
		t.Fatalf("parseReviewFlags: %v", err)
	}
	if !opts.translateZH {
		t.Fatal("translateZH should be true when --translate-zh is set")
	}
}

func TestShouldUseChineseOutput(t *testing.T) {
	tests := []struct {
		name    string
		lang    string
		forceZH bool
		want    bool
	}{
		{name: "force true", lang: "English", forceZH: true, want: true},
		{name: "english", lang: "English", forceZH: false, want: false},
		{name: "chinese", lang: "Chinese", forceZH: false, want: true},
		{name: "zh-cn", lang: "zh-CN", forceZH: false, want: true},
		{name: "zh", lang: "zh", forceZH: false, want: true},
		{name: "chinese chars", lang: "中文", forceZH: false, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shouldUseChineseOutput(tt.lang, tt.forceZH)
			if got != tt.want {
				t.Fatalf("shouldUseChineseOutput(%q, %v) = %v, want %v", tt.lang, tt.forceZH, got, tt.want)
			}
		})
	}
}
