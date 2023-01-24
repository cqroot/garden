package model

func (m Model) AutoMigrate() error {
	var err error

	err = m.database.DB().AutoMigrate(&Task{})
	if err != nil {
		return err
	}

	err = m.database.DB().AutoMigrate(&Project{})
	if err != nil {
		return err
	}

	return nil
}
