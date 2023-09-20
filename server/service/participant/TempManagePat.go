package participant

import "github.com/flipped-aurora/gin-vue-admin/server/model/participant/request"

type TempManagePatService struct{}

func (t *TempManagePatService) CreatePatTemp(req request.CreateTempOfPat) error {
	return nil
}

func (t *TempManagePatService) UpdatePatTemp(req request.UpdateTempOfPat) error {
	return nil
}

func (t *TempManagePatService) DeletePatTemp(req request.DeleteTempOfPat) error {
	return nil
}

func (t *TempManagePatService) GetPatTemp(req request.GetTempOfPat) (error error, data string) {
	return nil, ""
}
