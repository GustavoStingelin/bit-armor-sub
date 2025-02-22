package internal

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
	"time"
)

var dbPool *pgxpool.Pool

func InitDBPool() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, databaseName)

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		slog.Error("Failed to parse database URL", "err", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		slog.Error("Failed to create connection pool", "err", err)
	}

	// test connection
	var version string
	err = dbPool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		slog.Error("Query failed", "err", err)
	}

	slog.Info("Connected to pg", "version", version)
}

func CloseDBPool() {
	if dbPool != nil {
		dbPool.Close()
	}
}

type ArmoredOutpointDB struct {
	ID                 uuid.UUID `db:"id"`
	CreatedAt          time.Time `db:"created_at"`
	Hash               []byte    `db:"hash"`
	Index              int64     `db:"index"`
	Value              int64     `db:"value"`
	DestinationAddress string    `db:"destination_address"`
}

type PreSignedTxDB struct {
	ID                uuid.UUID `db:"id"`
	CreatedAt         time.Time `db:"created_at"`
	ArmoredOutpointID uuid.UUID `db:"armored_outpoint_id"`
	Fee               int64     `db:"fee"`
	Raw               []byte    `db:"raw"`
}

func GetArmoredOutpointsDB() ([]ArmoredOutpointDB, error) {
	slog.Debug("Getting armored outpoints from DB")
	rows, err := dbPool.Query(context.Background(), "SELECT id, created_at, hash, index, value, destination_address FROM public.armored_outpoint")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var armoredOutpoints []ArmoredOutpointDB
	for rows.Next() {
		var ao ArmoredOutpointDB
		err := rows.Scan(&ao.ID, &ao.CreatedAt, &ao.Hash, &ao.Index, &ao.Value, &ao.DestinationAddress)
		if err != nil {
			return nil, err
		}
		armoredOutpoints = append(armoredOutpoints, ao)
	}
	return armoredOutpoints, nil
}

func GetPreSignedTxsDB() ([]PreSignedTxDB, error) {
	slog.Debug("Getting pre-signed transactions from DB")
	rows, err := dbPool.Query(context.Background(), "SELECT id, created_at, armored_outpoint_id, fee, raw FROM public.pre_signed_tx")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var preSignedTxs []PreSignedTxDB
	for rows.Next() {
		var pst PreSignedTxDB
		err := rows.Scan(&pst.ID, &pst.CreatedAt, &pst.ArmoredOutpointID, &pst.Fee, &pst.Raw)
		if err != nil {
			return nil, err
		}
		preSignedTxs = append(preSignedTxs, pst)
	}
	return preSignedTxs, nil
}

func GetArmoredOutpoints() ([]ArmoredOutpoint, error) {
	slog.Info("Getting armored outpoints")
	armoredOutpoints, err := GetArmoredOutpointsDB()
	if err != nil {
		return nil, err
	}
	preSignedTxs, err := GetPreSignedTxsDB()
	if err != nil {
		return nil, err
	}

	mappedTxs := make(map[uuid.UUID][]PreSignedTx)
	for _, tx := range preSignedTxs {
		mappedTxs[tx.ArmoredOutpointID] = append(mappedTxs[tx.ArmoredOutpointID], PreSignedTx{
			fee: tx.Fee,
			raw: tx.Raw,
		})
	}

	var result []ArmoredOutpoint
	for _, ao := range armoredOutpoints {
		result = append(result, ArmoredOutpoint{
			Hash:               chainhash.Hash(ao.Hash),
			Index:              uint32(ao.Index),
			value:              ao.Value,
			DestinationAddress: ao.DestinationAddress,
			SpendablePreSigned: mappedTxs[ao.ID],
		})
	}

	return result, nil
}
