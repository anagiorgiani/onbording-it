package repository

const (
	SQLGetAll = `
	SELECT id, data,
	FROM history`
)

const SQLCreate = `INSERT INTO history(data) VALUES(?)`
