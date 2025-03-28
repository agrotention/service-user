package repo

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/agrotention/pb-user"
	"github.com/agrotention/service-user/cfg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoUser struct {
	db *sql.DB
}

func NewRepoUser(db *sql.DB) *RepoUser {
	return &RepoUser{db}
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.Fatalf("open database error: %v", err)
	}
	return db
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) CreateUser(
	ctx context.Context,
	valid *pb.CreateUserRequest,
) (*pb.SuccessResponse, error) {
	newID := uuid.New()
	id, err := newID.MarshalBinary() // Konversi ke binary
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to generate UUID")
	}

	var lastName any = valid.LastName
	if valid.LastName == "" {
		lastName = nil
	}

	q := `
		INSERT INTO users
		(id, username, email, first_name, last_name, password)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = r.db.ExecContext(ctx, q,
		id, valid.Username, valid.Email, valid.FirstName, lastName, valid.Password)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to insert user")
	}

	return &pb.SuccessResponse{ID: newID.String()}, nil
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) GetPrivateDetailUser(
	ctx context.Context,
	valid *pb.DetailUserRequest,
) (*pb.PrivateDetailUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateDetailUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) GetPublicDetailUser(
	ctx context.Context,
	valid *pb.DetailUserRequest,
) (*pb.PublicDetailUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicDetailUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) GetPasswordHash(
	ctx context.Context,
	valid *pb.GetPasswordHashRequest,
) (*pb.GetPasswordHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordHash not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) GetListUser(
	ctx context.Context,
	valid *pb.ListUserRequest,
) (*pb.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) GetListDeletedUser(
	ctx context.Context,
	valid *pb.ListUserRequest,
) (*pb.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListDeletedUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) UpdatePublicDetailUser(
	ctx context.Context,
	valid *pb.UpdatePublicDetailUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublicDetailUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) UpdateCredentialUser(
	ctx context.Context,
	valid *pb.UpdateCredentialUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentialUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) DisableUser(
	ctx context.Context,
	valid *pb.DisableUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) EnableUser(
	ctx context.Context,
	valid *pb.EnableUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableUser not implemented")
}

// Lakukan validasi sebelum memanggil fungsi ini
// Tidak ada proses validasi dalam repo
func (r *RepoUser) DeleteUser(
	ctx context.Context,
	valid *pb.DeleteUserRequest,
) (*pb.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
