package utils

import (
	"strings"
	"testing"
)

// TestPasswordHashAndVerify 测试 PasswordHash 和 PasswordVerify 函数。
func TestPasswordHashAndVerify(t *testing.T) {
	// 测试密码
	password := "myP@ssword123"
	wrongPassword := "wrongPassword123"

	// 生成密码哈希
	hashedPassword, err := PasswordHash(password)
	if err != nil {
		t.Fatalf("PasswordHash() error = %v", err)
	}

	// 确保哈希后的密码不等于原始密码
	if password == hashedPassword {
		t.Errorf("Hashed password should not be the same as the plain password.")
	}

	// 使用正确的密码进行验证
	if !PasswordVerify(password, hashedPassword) {
		t.Errorf("PasswordVerify() should return true for correct password.")
	}

	// 使用错误的密码进行验证
	if PasswordVerify(wrongPassword, hashedPassword) {
		t.Errorf("PasswordVerify() should return false for incorrect password.")
	}

	// 测试密码哈希是否每次都不同
	hashedPassword2, err := PasswordHash(password)
	if err != nil {
		t.Fatalf("PasswordHash() error = %v", err)
	}
	if strings.Compare(hashedPassword, hashedPassword2) == 0 {
		t.Errorf("PasswordHash() should generate different hashes for the same password input.")
	}
}

// BenchmarkPasswordHash 测试 PasswordHash 函数的性能。
func BenchmarkPasswordHash(b *testing.B) {
	password := "myP@ssword123"
	for i := 0; i < b.N; i++ {
		_, _ = PasswordHash(password)
	}
}

// BenchmarkPasswordVerify 测试 PasswordVerify 函数的性能。
func BenchmarkPasswordVerify(b *testing.B) {
	password := "myP@ssword123"
	hashedPassword, _ := PasswordHash(password)
	b.ResetTimer() // 重置计时器，因为我们不想将 PasswordHash 的执行时间计入基准测试。
	for i := 0; i < b.N; i++ {
		_ = PasswordVerify(password, hashedPassword)
	}
}
