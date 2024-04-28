package main

import "testing"

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
