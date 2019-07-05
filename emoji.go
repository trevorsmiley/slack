package slack

import (
	"context"
	"net/url"
)

type emojiResponseFull struct {
	Emoji map[string]string `json:"emoji"`
	SlackResponse
}

// GetEmoji retrieves all the emojis
func (api *Client) GetEmoji() (map[string]string, error) {
	return api.GetEmojiContext(context.Background())
}

// GetEmojiContext retrieves all the emojis with a custom context
func (api *Client) GetEmojiContext(ctx context.Context) (map[string]string, error) {
	values := url.Values{
		"token": {api.token},
	}
	response := &emojiResponseFull{}

	err := api.postMethod(ctx, "emoji.list", values, response)
	if err != nil {
		return nil, err
	}

	if response.Err() != nil {
		return nil, response.Err()
	}

	return response.Emoji, nil
}

// AddEmoji uploads an emoji
func (api *Client) AddEmoji(image, name string) error {
	return api.AddEmojiContext(context.Background(), image, name)
}

// AddEmojiContext uploads an emoji with a custom context
func (api *Client) AddEmojiContext(ctx context.Context, image, name string) error {
	values := url.Values{
		"token": {api.token},
		"mode":  {"data"},
		"name":  {name},
	}
	response := &emojiResponseFull{}

	err := postLocalWithMultipartResponse(ctx, api.httpclient, api.endpoint+"emoji.add", image, "image", values, response, api)

	if err != nil {
		return err
	}

	if response.Err() != nil {
		return response.Err()
	}

	return nil
}
