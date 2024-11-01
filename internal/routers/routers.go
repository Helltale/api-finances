package routers

import (
	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/logger"
)

func Init(logger *logger.CombinedLogger, config *config.Config) {
	income(logger, config)
	income_expected(logger, config)
	account(logger, config)
	expence(logger, config)
	remain(logger, config)
	goal(logger, config)
	cashback(logger, config)
}
