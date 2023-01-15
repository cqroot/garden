package app

func Bootstrap() error {
	err := InitConfig()
	if err != nil {
		return err
	}
	InitLogger(Config().LogLevel(), false)
	return nil
}
