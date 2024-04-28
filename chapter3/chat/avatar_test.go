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

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"email": "user@test.com",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL shoud not return any errors")
	}
	if url != "//www.gravatar.com/avatar/1460318498c1f53bb880ce2e6d9ef64b" {
		t.Errorf("GravatarAvatar.GetAvatarURL returned a wrong value: %s", url)
	}
}
