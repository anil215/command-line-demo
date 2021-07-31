package posts

import (
	"encoding/json"
	"fmt"

	"example.com/command/http"
	"github.com/pkg/errors"
)

type Posts struct {
	PostGetter http.HTTPer
}

func (p *Posts) GetPostForAnId(id int) (*Post, error) {
	postBytes, err := p.PostGetter.Get(fmt.Sprintf("/posts/%d", id))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get posts for id: %d", id))
	}

	post := &Post{}
	err = json.Unmarshal(postBytes, post)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal post")
	}

	return post, nil
}
