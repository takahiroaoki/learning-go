package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("When no value, AuthAvatar.GetAvatarURL should return ErrNoAvatarURL.")
	}

	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{
		"avatar_url": testUrl,
	}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("When the value exists, AuthAvatar.GetAvatarURL shoud not return any errors")
	} else {
		if url != testUrl {
			t.Error("The URL is wrong.")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "1460318498c1f53bb880ce2e6d9ef64b",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL shoud not return any errors")
	}
	if url != "//www.gravatar.com/avatar/1460318498c1f53bb880ce2e6d9ef64b" {
		t.Errorf("GravatarAvatar.GetAvatarURL returned a wrong value: %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	os.WriteFile(filename, []byte{}, 0777)
	defer func() {
		os.Remove(filename)
	}()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return any errors")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned wrongly %s", url)
	}
}
