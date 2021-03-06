package utils

const (
	SELECT_ALL_USER = `SELECT 
	mu.id_user,
	mu.nik,
	mu.username,
	mu.tgl_lahir,
	mp.id_pekerjaan,
	mp.label_pekerjaan,
	mpi.id_pendidikan,
	mpi.label_pendidikan,
	mu.user_status,
	mu.created_date 
	FROM m_user mu JOIN m_pekerjaan mp ON mu.id_pekerjaan=mp.id_pekerjaan 
	JOIN m_pendidikan mpi ON mu.id_pendidikan=mpi.id_pendidikan WHERE 
	mu.username LIKE ? AND mu.user_status = '1' LIMIT %s ,%s`

	SELECT_ALL_JOB       = `SELECT * FROM m_pekerjaan`
	SELECT_ALL_EDUCATION = `SELECT * FROM m_pendidikan`
	SELECT_USER_BY_ID    = `SELECT 
	mu.id_user,
	mu.nik,
	mu.username,
	mu.tgl_lahir,
	mp.id_pekerjaan,
	mp.label_pekerjaan,
	mpi.id_pendidikan,
	mpi.label_pendidikan,
	mu.user_status,
	mu.created_date 
	FROM m_user mu JOIN m_pekerjaan mp ON mu.id_pekerjaan=mp.id_pekerjaan 
	JOIN m_pendidikan mpi ON mu.id_pendidikan=mpi.id_pendidikan WHERE mu.id_user = ?`
	UPDATE_USER = `UPDATE m_user SET nik=?,username=?,tgl_lahir=?,id_pekerjaan=?,id_pendidikan=? 
	WHERE id_user = ?`
	INSERT_USER = `INSERT INTO m_user(id_user,nik,username,tgl_lahir,id_pekerjaan,id_pendidikan) values (?,?,?,?,?,?)`
	DELETE_USER = `UPDATE m_user SET user_status=0 where id_user=?`

	// AUTH QUERY
	CREATE_ACCOUNT          = `INSERT INTO m_account(account_id,email,password,token) values (?,?,?,?)`
	SELECT_ACCOUNT_BY_EMAIL = `SELECT * FROM m_account where email=?`
)
