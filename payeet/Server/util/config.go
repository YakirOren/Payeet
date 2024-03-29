package util

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Config holds the config data
type Config struct {
	Port                          string   `mapstructure:"PORT"`
	AccessTokenKey                string   `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey               string   `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration           string   `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration          string   `mapstructure:"REFRESH_TOKEN_DURATION"`
	ConnectionString              string   `mapstructure:"CONNECTION_STRING"`
	DBName                        string   `mapstructure:"DB_NAME"`
	UserCollection                string   `mapstructure:"USER_COLLECTION"`
	TransactionCollection         string   `mapstructure:"TRANSACTION_COLLECTION"`
	LogsCollection                string   `mapstructure:"LOGS_COLLECTION"`
	BaseDailyBonus                string   `mapstructure:"BASE_DAILY_BONUS"`
	StreakDailyBonus              string   `mapstructure:"STREAK_DAILY_BONUS"`
	MinimumRequiredTransferDays   string   `mapstructure:"KARMA_MINIMUM_REQUIRED_TRANSFER_DAYS"`
	MinimumRequiredTransferAmount string   `mapstructure:"KARMA_MINIMUM_REQUIRED_TRANSFER_AMOUNT"`
	KarmaMultiplierFactor         string   `mapstructure:"KARMA_MULTIPLIER_FACTOR"`
	MinimumRequiredUniqueUsers    string   `mapstructure:"KARMA_MINIMUM_REQUIRED_UNIQUE_USERS"`
	MaximumTransfersToSameUser    string   `mapstructure:"KARMA_MAXIMUM_TRANSFERS_TO_SAME_USER"`
	SystemEmail                   string   `mapstructure:"SYSTEM_EMAIL"`
	SystemEmailPassword           string   `mapstructure:"SYSTEM_EMAIL_PASSWORD"`
	ProfileImages                 []string `mapstructure:"PROFILE_IMAGES"`
	TotalImages                   int
}

// LoadConfig is used to load the config from the config file.
func LoadConfig(path string) (c Config, err error) {
	log.Infof("Loading config...")

	c.Port = os.Getenv("PORT")
	c.AccessTokenKey = os.Getenv("ACCESS_TOKEN_KEY")
	c.RefreshTokenKey = os.Getenv("REFRESH_TOKEN_KEY")
	c.AccessTokenDuration = os.Getenv("ACCESS_TOKEN_DURATION")
	c.RefreshTokenDuration = os.Getenv("REFRESH_TOKEN_DURATION")
	c.ConnectionString = os.Getenv("CONNECTION_STRING")
	c.DBName = os.Getenv("DB_NAME")
	c.UserCollection = os.Getenv("USER_COLLECTION")
	c.TransactionCollection = os.Getenv("TRANSACTION_COLLECTION")
	c.LogsCollection = os.Getenv("LOGS_COLLECTION")
	c.BaseDailyBonus = os.Getenv("BASE_DAILY_BONUS")
	c.StreakDailyBonus = os.Getenv("STREAK_DAILY_BONUS")
	c.MinimumRequiredTransferDays = os.Getenv("KARMA_MINIMUM_REQUIRED_TRANSFER_DAYS")
	c.MinimumRequiredTransferAmount = os.Getenv("KARMA_MINIMUM_REQUIRED_TRANSFER_AMOUNT")
	c.KarmaMultiplierFactor = os.Getenv("KARMA_MULTIPLIER_FACTOR")
	c.MinimumRequiredUniqueUsers = os.Getenv("KARMA_MINIMUM_REQUIRED_UNIQUE_USERS")
	c.MaximumTransfersToSameUser = os.Getenv("KARMA_MAXIMUM_TRANSFERS_TO_SAME_USER")
	c.SystemEmail = os.Getenv("SYSTEM_EMAIL")
	c.SystemEmailPassword = os.Getenv("SYSTEM_EMAIL_PASSWORD")

	c.ProfileImages = strings.Fields(os.Getenv("PROFILE_IMAGES"))
	c.TotalImages = len(c.ProfileImages)
	log.Infof("Done! ✅")
	return

}
