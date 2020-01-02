package git

import (
	"testing"
)

func TestGitClone(t *testing.T) {
	type args struct {
		g *GIT
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GitClone(tt.args.g)
		})
	}
}

func TestCloneYaml(t *testing.T) {
	type args struct {
		g *GIT
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloneYaml(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("CloneYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCloneYamlByTag(t *testing.T) {
	type args struct {
		g   *GIT
		tag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloneYamlByTag(tt.args.g, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("CloneYamlByTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
