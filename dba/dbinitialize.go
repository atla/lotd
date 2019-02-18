package dba

// CreateDatabase ... creates initial database value
func (dba *DBAccess) CreateDatabase() {

	/*
		// reset database
		dba.Session.DB("athenadb").C("ruleset").DropCollection()
		app.DBAccess.Session.DB("athenadb").C("ruleset").Create(&mgo.CollectionInfo{})
		app.DBAccess.Session.DB("athenadb").C("session").DropCollection()
		app.DBAccess.Session.DB("athenadb").C("session").Create(&mgo.CollectionInfo{})

		munchkinRules := model.NewRuleset("Munchkin", "Classic")
		munchkinRules.Ranking = "{playerLevel} descending"
		munchkinRules.WinCondition = "{playerLevel} > 9"

		munchkinRules.AddStat("playerLevel", "Player Level", "int", 1)
		munchkinRules.AddStat("itemLevel", "Item Level", "int", 0)
		munchkinRules.AddStat("battleStrength", "Battle Level", "int", 0)

		duration, _ := time.ParseDuration("2h")

		munchkinSession := model.NewSession(time.Now(), duration.Nanoseconds(), munchkinRules.ID.Hex())
		munchkinSession.AddPlayer(model.NewPlayer("atla"))
		munchkinSession.AddPlayer(model.NewPlayer("claudia"))
		munchkinSession.AddPlayer(model.NewPlayer("daniel"))

		// store sample data
		app.DBAccess.WriteRulesetToDB(munchkinRules)
		app.DBAccess.WriteSessionToDB(munchkinSession)*/
}
