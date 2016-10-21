package db

var maxPool int

func init() {
	var err error
	//maxPool, err = beego.AppConfig.Int("DBMaxPool")
	maxPool, err = 10, nil
	if err != nil {
		// todo: panic!!!
		// panic(err)
		println(err)
	}
	// init method to start db
	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = "127.0.0.1"
		err := service.New()
		if err != nil {
			// todo: panic!!!
			// panic(err)
			println(err)
		}
	}
}
