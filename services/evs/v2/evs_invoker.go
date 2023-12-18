package v2

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/invoker"
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/services/evs/v2/model"
)

type CinderAcceptVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderAcceptVolumeTransferInvoker) Invoke() (*model.CinderAcceptVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

type CinderCreateVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderCreateVolumeTransferInvoker) Invoke() (*model.CinderCreateVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

type CinderDeleteVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderDeleteVolumeTransferInvoker) Invoke() (*model.CinderDeleteVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

type CinderListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListAvailabilityZonesInvoker) Invoke() (*model.CinderListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

type CinderListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListQuotasInvoker) Invoke() (*model.CinderListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListQuotasResponse), nil
	}
}

type CinderListVolumeTransfersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTransfersInvoker) Invoke() (*model.CinderListVolumeTransfersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTransfersResponse), nil
	}
}

type CinderListVolumeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderListVolumeTypesInvoker) Invoke() (*model.CinderListVolumeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderListVolumeTypesResponse), nil
	}
}

type CinderShowVolumeTransferInvoker struct {
	*invoker.BaseInvoker
}

func (i *CinderShowVolumeTransferInvoker) Invoke() (*model.CinderShowVolumeTransferResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CinderShowVolumeTransferResponse), nil
	}
}

type DeleteVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVolumeInvoker) Invoke() (*model.DeleteVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVolumeResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type UpdateVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVolumeInvoker) Invoke() (*model.UpdateVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVolumeResponse), nil
	}
}

type ListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}
