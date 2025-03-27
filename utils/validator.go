package utils

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"unicode"
)

// # ValidationErrors
//
// Menyimpan daftar error validasi.
type ValidationErrors []*FieldError

func (err ValidationErrors) Error() string {
	if len(err) == 0 {
		return "No validation errors"
	}
	var msg string
	for _, e := range err {
		msg += e.Error() + "\n"
	}
	return msg
}

// # Validator
//
// Alat untuk validasi struct user.
type Validator struct {
	Errors ValidationErrors
}

// # Validator helper
//
// Menambahkan error ke hasil validasi.
func (v *Validator) addError(err *FieldError) {
	v.Errors = append(v.Errors, err)
}

// # Validasi Email
//
// Memeriksa apakah email valid menggunakan regex.
func (v *Validator) validateEmail(email string) *FieldError {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return &FieldError{Name: "Email", Message: "Email tidak valid"}
	}
	if len(email) > 254 {
		return &FieldError{Name: "Email", Message: "Email maksimal 254 karakter"}
	}
	return nil
}

// # Validasi Username
//
// Harus minimal 3 karakter dan hanya boleh berisi huruf & angka.
func (v *Validator) validateUsername(username string) *FieldError {
	if len(username) < 3 {
		return &FieldError{Name: "Username", Message: "Username minimal 3 karakter"}
	}
	if len(username) > 32 {
		return &FieldError{Name: "Username", Message: "Username maksimal 32 karakter"}
	}
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return &FieldError{Name: "Username", Message: "Username hanya boleh mengandung huruf dan angka"}
		}
	}
	return nil
}

// # Validasi Password
//
// Harus minimal 8 karakter, mengandung huruf kapital, huruf kecil, dan angka.
func (v *Validator) validatePassword(password string) *FieldError {
	if len(password) < 8 {
		return &FieldError{Name: "Password", Message: "Password minimal 8 karakter"}
	}

	// Buat regex tanpa lookahead
	// re := regexp.MustCompile(`[a-zA-Z\d]+`) // Ambil semua karakter yang valid
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	// Pastikan ada huruf kecil, huruf besar, dan angka
	if !hasLower || !hasUpper || !hasDigit {
		return &FieldError{Name: "Password", Message: "Password harus mengandung huruf besar, huruf kecil, dan angka"}
	}

	return nil
}

// # Validasi Nama
//
// Harus diisi jika `required` bernilai `true`.
func (v *Validator) validateName(name string, required bool) *FieldError {
	if required && name == "" {
		return &FieldError{Name: "Name", Message: "Nama wajib diisi"}
	}
	return nil
}

// # Validate
//
// Memvalidasi struct user (request gRPC).
func (v *Validator) Validate(userStruct any) ValidationErrors {
	t := reflect.TypeOf(userStruct)
	val := reflect.ValueOf(userStruct)

	// Jika input adalah pointer, ambil elemennya
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}

	// Pastikan input adalah struct
	if t.Kind() != reflect.Struct {
		log.Panic("validator input data harus berupa struct atau pointer ke struct")
	}

	// Loop semua field dalam struct
	for i := range t.NumField() {
		field := t.Field(i)
		value := val.Field(i)

		switch field.Name {
		case "Username":
			if err := v.validateUsername(value.String()); err != nil {
				v.addError(err)
			}
		case "Email":
			if err := v.validateEmail(value.String()); err != nil {
				v.addError(err)
			}
		case "Password":
			if err := v.validatePassword(value.String()); err != nil {
				v.addError(err)
			}
		case "FirstName":
			if err := v.validateName(value.String(), true); err != nil {
				v.addError(err)
			}
		case "LastName":
			if err := v.validateName(value.String(), false); err != nil {
				v.addError(err)
			}
		}
	}

	return v.Errors
}

// # NewValidator
//
// Constructor untuk validator.
// Gunakan validator baru untuk setiap validasi.
func NewValidator() *Validator {
	return &Validator{
		Errors: ValidationErrors{},
	}
}

// # FieldError
//
// Mendeskripsikan setiap error pada field.
type FieldError struct {
	Name    string
	Message string
}

// Implementasi error untuk *FieldError.
func (err *FieldError) Error() string {
	return fmt.Sprintf("Error pada field %s: %s", err.Name, err.Message)
}
