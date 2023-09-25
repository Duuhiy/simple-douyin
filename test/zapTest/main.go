package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	zap.L().Info("err err err ...")
}
