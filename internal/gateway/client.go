package gateway

import (
	"github.com/luciano340/FC_EDA/internal/entity"
)

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
