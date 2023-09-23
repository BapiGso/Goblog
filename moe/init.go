package moe

import "SMOE/moe/database"

// Init TODO slog的日志等级设定
func (s *Smoe) Init() {
	database.InitDB()
}
