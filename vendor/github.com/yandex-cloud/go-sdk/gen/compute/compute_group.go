// Code generated by sdkgen. DO NOT EDIT.

package compute

import (
	"context"

	"google.golang.org/grpc"
)

// Compute provides access to "compute" component of Yandex.Cloud
type Compute struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewCompute creates instance of Compute
func NewCompute(g func(ctx context.Context) (*grpc.ClientConn, error)) *Compute {
	return &Compute{g}
}

// Disk gets DiskService client
func (c *Compute) Disk() *DiskServiceClient {
	return &DiskServiceClient{getConn: c.getConn}
}

// DiskType gets DiskTypeService client
func (c *Compute) DiskType() *DiskTypeServiceClient {
	return &DiskTypeServiceClient{getConn: c.getConn}
}

// Image gets ImageService client
func (c *Compute) Image() *ImageServiceClient {
	return &ImageServiceClient{getConn: c.getConn}
}

// Instance gets InstanceService client
func (c *Compute) Instance() *InstanceServiceClient {
	return &InstanceServiceClient{getConn: c.getConn}
}

// Snapshot gets SnapshotService client
func (c *Compute) Snapshot() *SnapshotServiceClient {
	return &SnapshotServiceClient{getConn: c.getConn}
}

// Zone gets ZoneService client
func (c *Compute) Zone() *ZoneServiceClient {
	return &ZoneServiceClient{getConn: c.getConn}
}

// PlacementGroup gets PlacementGroupService client
func (c *Compute) PlacementGroup() *PlacementGroupServiceClient {
	return &PlacementGroupServiceClient{getConn: c.getConn}
}

// HostGroup gets HostGroupService client
func (c *Compute) HostGroup() *HostGroupServiceClient {
	return &HostGroupServiceClient{getConn: c.getConn}
}

// HostType gets HostTypeService client
func (c *Compute) HostType() *HostTypeServiceClient {
	return &HostTypeServiceClient{getConn: c.getConn}
}

// DiskPlacementGroup gets DiskPlacementGroupService client
func (c *Compute) DiskPlacementGroup() *DiskPlacementGroupServiceClient {
	return &DiskPlacementGroupServiceClient{getConn: c.getConn}
}

// Filesystem gets FilesystemService client
func (c *Compute) Filesystem() *FilesystemServiceClient {
	return &FilesystemServiceClient{getConn: c.getConn}
}

// SnapshotSchedule gets SnapshotScheduleService client
func (c *Compute) SnapshotSchedule() *SnapshotScheduleServiceClient {
	return &SnapshotScheduleServiceClient{getConn: c.getConn}
}

// GpuCluster gets GpuClusterService client
func (c *Compute) GpuCluster() *GpuClusterServiceClient {
	return &GpuClusterServiceClient{getConn: c.getConn}
}

// ReservedInstancePool gets ReservedInstancePoolService client
func (c *Compute) ReservedInstancePool() *ReservedInstancePoolServiceClient {
	return &ReservedInstancePoolServiceClient{getConn: c.getConn}
}
