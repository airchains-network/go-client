package components

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	JunctionRPC string
	JunctionAPI string
	RpcAUTH     string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}

	JunctionRPC = os.Getenv("JUNCTION_RPC")
	JunctionAPI = os.Getenv("JUNCTION_API")
	RpcAUTH = os.Getenv("RPC_AUTH")
}

func RandomString(n int, allowedChars string) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = rune(allowedChars[rand.Intn(len(allowedChars))])
	}
	return string(s)
}

func GenerateAddresses(n int) []string {
	addresses := make([]string, n)
	allowedChars := "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < n; i++ {
		addresses[i] = "air1" + RandomString(38, allowedChars)
	}
	return addresses
}

func GenerateRandomWithFavour(lowerBound, upperBound int, favourableSet [2]int, favourableProbability float64) int {
	if lowerBound > upperBound || favourableProbability < 0 || favourableProbability > 1 {
		log.Println("Invalid parameters")
		return 0
	}

	totalRange := upperBound - lowerBound + 1
	favourableRange := favourableSet[1] - favourableSet[0] + 1

	if favourableRange <= 0 || favourableRange > totalRange {
		log.Println("Invalid favourable set")
		return 0
	}

	if favourableSet[0] < lowerBound || favourableSet[1] > upperBound {
		log.Println("Invalid favourable set")
		return 0
	}

	favourableOutcomes := int(favourableProbability * float64(totalRange))
	if favourableOutcomes < favourableRange {
		favourableOutcomes = favourableRange
	}

	randNum := rand.Intn(totalRange)
	if randNum < favourableOutcomes {
		randNum = randNum%favourableRange + favourableSet[0]
	} else {
		randNum = randNum%favourableOutcomes + lowerBound
		if randNum >= favourableSet[0] && randNum <= favourableSet[1] {
			randNum = favourableSet[1] + 1 + (randNum - favourableSet[0])
		}
	}

	return randNum
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := io.ReadFull(cryptoRand.Reader, bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

func GenerateEthereumAddress() (string, error) {
	address, err := GenerateRandomHex(20)
	if err != nil {
		return "", err
	}
	return "0x" + address, nil
}

func GenerateTransactionHash() (string, error) {
	txHash, err := GenerateRandomHex(32)
	if err != nil {
		return "", err
	}
	return "0x" + txHash, nil
}

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func GenerateRandomNumber() string {
	tempAmount := GenerateRandomWithFavour(1, 1000000, [2]int{100, 100000}, 0.7)
	return fmt.Sprintf("%d", tempAmount)
}

func GenerateRandomSmallNumber() string {
	tempValue := GenerateRandomWithFavour(1, 1000, [2]int{1, 100}, 0.7)
	return fmt.Sprintf("%d", tempValue)
}

func GenerateRandomBech32Address() string {
	address, err := GenerateRandomHex(20)
	if err != nil {
		return ""
	}
	return "air1" + address
}

func GenerateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

func IncrementVersion(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) == 3 {
		patchVersion, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatalf("Error converting patch version to integer: %v", err)
		}
		if patchVersion >= 9 {
			patchVersion = 0
			minorVersion, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("Error converting minor version to integer: %v", err)
			}
			if minorVersion >= 9 {
				minorVersion = 0
				majorVersion, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("Error converting major version to integer: %v", err)
					return ""
				}
				parts[0] = fmt.Sprintf("%d", majorVersion+1)
			} else {
				parts[1] = fmt.Sprintf("%d", minorVersion+1)
			}
		} else {
			patchVersion++
		}
		parts[2] = fmt.Sprintf("%d", patchVersion)
	}

	finalString := strings.Join(parts, ".")

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := cwd + "/data/schemaVersion.txt"

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating schema version file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(finalString)
	if err != nil {
		log.Fatalf("Error writing to schema version file: %v", err)
	}

	return finalString
}
