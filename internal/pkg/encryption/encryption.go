package encryption

import "golang.org/x/crypto/bcrypt"

func Encrypt(pwd string, cost int) ([]byte, error) {

	pwdByte := []byte(pwd)
	hashPwd, err := bcrypt.GenerateFromPassword(pwdByte, cost)
	if err != nil {
		return nil, err
	}

	return hashPwd, nil

}

func ComparePwd(hashedPwd, usrProvidedPwd string) error {

	hashedUsrProvidedPwd := []byte(usrProvidedPwd)

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), hashedUsrProvidedPwd); err != nil {
		return err
	}

	return nil
}
