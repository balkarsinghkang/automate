package pgdb

import "github.com/pkg/errors"

type UpgradeDbTest struct {
	Error        bool
	ControlError bool
	NeedStatus   bool
}

func (u UpgradeDbTest) GetDayLatestUpgradeFlagValue() (status bool, err error) {
	if u.Error {
		return false, errors.New("Unable to fetch status from database")
	}
	if u.NeedStatus {
		return true, nil
	}
	return false, nil
}

func (u UpgradeDbTest) GetControlLatestUpgradeFlagValue() (status bool, err error) {
	if u.ControlError {
		return false, errors.New("Unable to fetch status from database")
	}
	if u.NeedStatus {
		return true, nil
	}

	return true, nil
}

func (u UpgradeDbTest) UpdateDayLatestFlagToFalse() error {
	return nil
}

func (u UpgradeDbTest) UpdateControlFlagToFalse() error {
	return nil
}
