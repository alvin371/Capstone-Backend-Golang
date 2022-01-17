package migrate

import (
	"capstone/backend/driver"

	m_news "capstone/backend/features/news"
)

func AutoMigrate() {
	driver.DB.AutoMigrate(
		&m_news.NewsCore{},
	)
}
