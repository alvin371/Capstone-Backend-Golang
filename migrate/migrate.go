package migrate

import (
	"capstone/backend/driver"

	m_news "capstone/backend/features/News"
	m_user "capstone/backend/features/User"
)

func AutoMigrate() {
	driver.DB.AutoMigrate(
		&m_news.NewsCore{}, &m_user.User{},
	)
}
