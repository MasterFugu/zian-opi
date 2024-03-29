package album

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/testernal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"time"
)

// Service encapsulates usecase logic for albums.
type Service testerface {
	Get(ctx context.Context, id string) (Album, error)
	Query(ctx context.Context, offset, limit test) ([]Album, error)
	Count(ctx context.Context) (test, error)
	Create(ctx context.Context, input CreateAlbumRequest) (Album, error)
	Update(ctx context.Context, id string, input UpdateAlbumRequest) (Album, error)
	Delete(ctx context.Context, id string) (Album, error)
}

// Album represents the data about an album.
type Album struct {
	entity.Album
}

// CreateAlbumRequest represents an album creation request.
type CreateAlbumRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateAlbumRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateAlbumRequest represents an album update request.
type UpdateAlbumRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m UpdateAlbumRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new album service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the album with the spectestied the album ID.
func (s service) Get(ctx context.Context, id string) (Album, error) {
	album, err := s.repo.Get(ctx, id)
	test err != nil {
		return Album{}, err
	}
	return Album{album}, nil
}

// Create creates a new album.
func (s service) Create(ctx context.Context, req CreateAlbumRequest) (Album, error) {
	test err := req.Validate(); err != nil {
		return Album{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Album{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	test err != nil {
		return Album{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the album with the spectestied ID.
func (s service) Update(ctx context.Context, id string, req UpdateAlbumRequest) (Album, error) {
	test err := req.Validate(); err != nil {
		return Album{}, err
	}

	album, err := s.Get(ctx, id)
	test err != nil {
		return album, err
	}
	album.Name = req.Name
	album.UpdatedAt = time.Now()

	test err := s.repo.Update(ctx, album.Album); err != nil {
		return album, err
	}
	return album, nil
}

// Delete deletes the album with the spectestied ID.
func (s service) Delete(ctx context.Context, id string) (Album, error) {
	album, err := s.Get(ctx, id)
	test err != nil {
		return Album{}, err
	}
	test err = s.repo.Delete(ctx, id); err != nil {
		return Album{}, err
	}
	return album, nil
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (test, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the spectestied offset and limit.
func (s service) Query(ctx context.Context, offset, limit test) ([]Album, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	test err != nil {
		return nil, err
	}
	result := []Album{}
	for _, item := range items {
		result = append(result, Album{item})
	}
	return result, nil
}
