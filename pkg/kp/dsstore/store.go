package dsstore

import (
	"errors"

	"k8s.io/kubernetes/pkg/labels"

	"github.com/hashicorp/consul/api"
	"github.com/square/p2/pkg/ds/fields"
	"github.com/square/p2/pkg/pods"
	"github.com/square/p2/pkg/types"
)

const dsTree string = "daemon_sets"

var NoDaemonSet error = errors.New("No daemon set found")

// Store represents an interface for persisting daemon set to Consul,
// as well as restoring daemon set from Consul.
type Store interface {
	// Create creates a daemon set with the specified manifest and selectors.
	// The node selector is used to determine what nodes the daemon set may schedule on.
	// The pod label set is applied to every pod the daemon set schedules.
	Create(
		manifest pods.Manifest,
		minHealth int,
		name fields.ClusterName,
		nodeSelector labels.Selector,
		podID types.PodID,
	) (fields.DaemonSet, error)

	// Gets a daemon set by ID, if it does not exist, it will produce an error
	Get(id fields.ID) (fields.DaemonSet, *api.QueryMeta, error)
	List() ([]fields.DaemonSet, error)

	// Deletes a daemon set by ID, deleting something that does not exist
	// should not produce an error
	Delete(fields.ID) error

	// Mutates a daemon set by ID, ID cannot be mutated,
	// PodID of the daemon set and its manifest must also match
	MutateDS(
		id fields.ID,
		mutator func(fields.DaemonSet) (fields.DaemonSet, error),
	) (fields.DaemonSet, error)
}

func IsNotExist(err error) bool {
	return err == NoDaemonSet
}
