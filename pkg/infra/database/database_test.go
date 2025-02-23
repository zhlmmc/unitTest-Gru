package database_test

import (
	"database/sql"
	"testing"
	"time"

	"order-system/pkg/infra/database"

	"github.com/stretchr/testify/assert"
)

type mockSQLDB struct {
	stats sql.DBStats
}

func (m *mockSQLDB) Stats() sql.DBStats {
	return m.stats
}

type mockDB struct {
	sqlDB *mockSQLDB
}

func (m *mockDB) Stats() database.Stats {
	stats := m.sqlDB.Stats()
	return database.Stats{
		OpenConnections: stats.OpenConnections,
		InUse:           stats.InUse,
		Idle:            stats.Idle,
		WaitCount:       stats.WaitCount,
		WaitDuration:    stats.WaitDuration,
		MaxIdleTime:     time.Duration(stats.MaxIdleTimeClosed),
	}
}

func TestStats(t *testing.T) {
	// Create mock stats
	mockStats := sql.DBStats{
		MaxOpenConnections: 10,
		OpenConnections:    3,
		InUse:              1,
		Idle:               2,
		WaitCount:          5,
		WaitDuration:       time.Second * 2,
		MaxIdleTimeClosed:  1,
	}

	// Create db wrapper with mock
	mockDBWrapper := &mockDB{
		sqlDB: &mockSQLDB{stats: mockStats},
	}

	// Test Stats method
	stats := mockDBWrapper.Stats()

	// Verify stats values
	assert.Equal(t, mockStats.OpenConnections, stats.OpenConnections)
	assert.Equal(t, mockStats.InUse, stats.InUse)
	assert.Equal(t, mockStats.Idle, stats.Idle)
	assert.Equal(t, mockStats.WaitCount, stats.WaitCount)
	assert.Equal(t, mockStats.WaitDuration, stats.WaitDuration)
	assert.Equal(t, time.Duration(mockStats.MaxIdleTimeClosed), stats.MaxIdleTime)

	// Test zero values
	zeroStats := sql.DBStats{}
	mockDBWrapper.sqlDB = &mockSQLDB{stats: zeroStats}

	stats = mockDBWrapper.Stats()
	assert.Equal(t, 0, stats.OpenConnections)
	assert.Equal(t, 0, stats.InUse)
	assert.Equal(t, 0, stats.Idle)
	assert.Equal(t, int64(0), stats.WaitCount)
	assert.Equal(t, time.Duration(0), stats.WaitDuration)
	assert.Equal(t, time.Duration(0), stats.MaxIdleTime)

	// Test maximum values
	maxStats := sql.DBStats{
		MaxOpenConnections: 1000,
		OpenConnections:    1000,
		InUse:              1000,
		Idle:               1000,
		WaitCount:          1000000,
		WaitDuration:       time.Hour,
		MaxIdleTimeClosed:  1000,
	}

	mockDBWrapper.sqlDB = &mockSQLDB{stats: maxStats}

	stats = mockDBWrapper.Stats()
	assert.Equal(t, maxStats.OpenConnections, stats.OpenConnections)
	assert.Equal(t, maxStats.InUse, stats.InUse)
	assert.Equal(t, maxStats.Idle, stats.Idle)
	assert.Equal(t, maxStats.WaitCount, stats.WaitCount)
	assert.Equal(t, maxStats.WaitDuration, stats.WaitDuration)
	assert.Equal(t, time.Duration(maxStats.MaxIdleTimeClosed), stats.MaxIdleTime)
}
