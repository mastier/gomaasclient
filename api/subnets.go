package api

import (
	"github.com/maas/gomaasclient/entity"
)

// Subnets represents the MaaS Subnets endpoint
type Subnets interface {
	Get() ([]entity.Subnet, error)
	Create(params *entity.SubnetParams) (*entity.Subnet, error)
}
