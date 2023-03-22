package application

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

func IsValidEmail(u *domain.Usuario) bool {
	return validator.Validate(u) == nil
}

func GenerateToken(id string, typeToken string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	duration, err := time.ParseDuration(os.Getenv(typeToken))
	if err != nil {
		log.Fatal("error parsing token duration")
		return "", err
	}
	claims["exp"] = time.Now().Add(duration).Unix()
	if typeToken == "REFRESH_TOKEN_EXP" {
		claims["token_type"] = "refresh"
	} else if typeToken == "TEMP_TOKEN_EXP" {
		claims["token_type"] = "temp"
	}

	tokenString, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		log.Fatal("error parsing token duration")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["token_type"] != "temp" {
			return "", errors.New("invalid token type")
		}
		id := claims["user_id"].(string)
		return id, nil
	} else {
		return "", err
	}
}

func ValidateRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		return GenerateToken(userID, "TEMP_TOKEN_EXP")
	} else {
		return "", err
	}
}

func HashPassword(s string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashed)
}

func ComparePassword(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
}

func Encrypt(stringToEncrypt string) (string, error) {
	key, _ := hex.DecodeString(string([]byte(config.SECRET)))
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func ConvertAndEncrypt(uintToEncrypt uint) (string, error) {
	stringToEncrypt := strconv.Itoa(int(uintToEncrypt))
	return Encrypt(stringToEncrypt)
}

func Decrypt(encryptedString string) (string, error) {
	key, _ := hex.DecodeString(string([]byte(config.SECRET)))
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func DecryptAndConvert(id string) (uint, error) {
	stringID, err := Decrypt(id)
	if err != nil {
		return 0, err
	}

	intID, err := strconv.Atoi(stringID)
	if err != nil {
		return 0, err
	}

	return uint(intID), nil
}

func DecryptAndConvertArray(ids []string) ([]uint, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var result []uint
	for _, id := range ids {
		intID, err := DecryptAndConvert(id)
		if err != nil {

			return nil, err
		}
		result = append(result, uint(intID))
	}
	return result, nil
}

func GetUsuarioFromCtx(ctx context.Context) (*domain.Usuario, error) {
	u, ok := ctx.Value("usuario").(*domain.Usuario)
	if !ok {
		return nil, errors.New("access denied")
	}
	return u, nil
}

func Get[M any](id *string, load func(context.Context, dataloader.Key) dataloader.Thunk) (*M, error) {
	if id == nil {
		return nil, nil
	}

	thunk := load(context.Background(), dataloader.StringKey(*id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*M), nil
}

func GetBatch[F any, D any, M any](keys []string, f *F, getAll func(*F) ([]*D, error), convertOut func(*D) (*M, error)) []*dataloader.Result {
	ids, err := DecryptAndConvertArray(keys)
	if err != nil {
		panic(err)
	}

	fIDs := reflect.ValueOf(f).Elem().FieldByName("IDs")
	idsVal := reflect.ValueOf(&ids)
	fIDs.Set(idsVal)
	oo, err := getAll(f)
	if err != nil {
		panic(err)
	}

	okmap := map[string]*D{}
	for _, o := range oo {
		id := reflect.ValueOf(*o).FieldByName("ID").Interface()
		encrypted, _ := ConvertAndEncrypt(id.(uint))
		okmap[encrypted] = o
	}

	out := make([]*dataloader.Result, len(keys))

	for i, k := range keys {
		o, ok := okmap[k]
		if ok {
			c, err := convertOut(o)
			out[i] = &dataloader.Result{Data: c, Error: err}
		} else {
			err := fmt.Errorf("entity not found: %s", k)
			out[i] = &dataloader.Result{Data: nil, Error: err}
		}
	}

	return out
}

func GetAll[I any, F any, D any, M any, NC any](i *I, p *paging.PageArgs, convertFilter func(*I, *paging.PageArgs) (*F, error), getAll func(*F) ([]*D, error), convertOut func(*D) (*M, error), buildNodeConnection func([]*M, *paging.PageArgs) (*NC, error)) (*NC, error) {
	pp, err := GetAllArray(i, p, convertFilter, getAll, convertOut)
	if err != nil {
		return nil, err
	}

	return buildNodeConnection(pp, p)
}

func GetAllArray[I any, F any, D any, M any](i *I, p *paging.PageArgs, convertFilter func(*I, *paging.PageArgs) (*F, error), getAll func(*F) ([]*D, error), convertOut func(*D) (*M, error)) ([]*M, error) {
	f, err := convertFilter(i, p)
	if err != nil {
		return nil, err
	}

	oo, err := getAll(f)
	if err != nil {
		return nil, err
	}

	var cc []*M
	for _, o := range oo {
		c, err := convertOut(o)
		if err != nil {
			return nil, err
		}

		cc = append(cc, c)
	}

	return cc, nil
}

func Save[I any, D any, M any](ctx context.Context, i *I, convertIn func(context.Context, *I) (*D, error), save func(*D) error, get func(uint) (*D, error), convertOut func(*D) (*M, error)) (*M, error) {
	d, err := convertIn(ctx, i)
	if err != nil {
		return nil, err
	}

	if err := save(d); err != nil {
		return nil, err
	}

	id := reflect.ValueOf(*d).FieldByName("ID").Interface()
	dd, err := get(id.(uint))
	if err != nil {
		return nil, err
	}
	return convertOut(dd)
}

// La nueva funcion Update primero coge la entidad de la DB. Luego, aplica la nueva funcion ConvertUpdate. Luego ejecuta Update.
// Las nuevas funciones ConvertUpdate en los Services, actualizan los valores opcionales recibidos no nulos del input de GraphQL
// a la entidad.
// Anteriormente se creaba una entidad con solamente los valores recibidos no nulos. Pero algunos tipos, como int, tienen valor "nulo" = 0
// Esto, el ORM lo interpreta como nulo, entonces la posibilidad de hacer updates
func Update[I any, D any, M any](i *I, convertUpdate func(*I, *D) (*D, error), update func(*D) error, get func(uint) (*D, error), convertOut func(*D) (*M, error)) (*M, error) {
	id := reflect.ValueOf(*i).FieldByName("ID").Interface()
	uintid, err := DecryptAndConvert(id.(string))
	if err != nil {
		return nil, err
	}

	o, err := get(uintid)
	if err != nil {
		return nil, err
	}

	u, err := convertUpdate(i, o)
	if err != nil {
		return nil, err
	}

	if err := update(u); err != nil {
		return nil, err
	}

	return convertOut(u)
}

func Delete(d func(uint) error, id string) (string, error) {
	uintid, err := DecryptAndConvert(id)
	if err != nil {
		return "", err
	}

	if err := d(uintid); err != nil {
		return "", err
	}

	return id, nil
}
