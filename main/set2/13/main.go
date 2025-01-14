package main

//ECB cut and paste
import (
	"crypto-pals/crysuite"
	"crypto-pals/util"
	"fmt"
	"regexp"
	"strings"
)

func kvParse(str string) map[string]string {
	kvMap := make(map[string]string)
	strs := strings.Split(str, "&")

	for _, i := range strs {
		kv := strings.Split(i, "=")

		if len(kv) != 2 {
			panic("Unexpected key-value length: ")
		}

		kvMap[kv[0]] = kv[1]
	}

	return kvMap
}

func kvEncode(kvMap map[string]string) string {
	//kv := ""
	//kvS := make([]string, len(kvMap))

	return "email=" + kvMap["email"] + "&uid=" + kvMap["uid"] + "&role=" + kvMap["role"]

	/*
		for k := range kvMap {
			kvS = append(kvS, k)
		}
		sort.Strings(kvS)
		for k, _ := range kvMap {
			kv += (k + "=" + kvMap[k] + "&")
		}

		return strings.TrimRight(kv, "&")
	*/
}

func createProfile(str string) map[string]string {
	re := regexp.MustCompile(`[&=]`)
	str = re.ReplaceAllString(str, "")

	// id := uuid.New()

	return map[string]string{
		"email": str,
		"uid":   "10",
		"role":  "user",
	}
}

func aes(str string, key []byte) []byte {
	prof := createProfile(str)
	kvProf := kvEncode(prof)

	return crysuite.AES_ECB_Encrypt([]byte(kvProf), key)
}

func main() {
	toParse := "foo=bar&baz=qux&zap=zazzle"
	toEncode := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}

	fmt.Println("Inital tests:")
	fmt.Println(kvParse(toParse))
	fmt.Println(kvEncode(toEncode))

	user := "foo@gmail.com&role=admin"
	profile := createProfile(user)
	fmt.Println(profile)

	//kvProfile := kvEncode(profile)
	key := util.GenerateRandomBytes(16)
	aesProfile := aes(user, key)

	fmt.Println("\nEncoder tests:")
	dProfile := crysuite.AES_ECB_Decrypt(aesProfile, key)
	fmt.Println(string(dProfile))
	fmt.Println(kvParse(string(dProfile)))

	fmt.Println("\nBreaking ECB:")
	uProfile := aes("user@gmail.com", key)
	aProfile := aes("user@gmail.comadmin", key)

	encode := append(uProfile[:32], aProfile[16:32]...)
	decoded := string(crysuite.AES_ECB_Decrypt(encode, key))

	fmt.Println(decoded)
	fmt.Println(kvParse(decoded))
}
