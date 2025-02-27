package memory

// import (
// 	"architecture/ws/services/bus"
// 	"encoding/hex"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type BulkWriteMemoryRequest struct {
// 	Address string `json:"address"`
// 	Value   string `json:"value"`
// }

// var globalDataBus *bus.DataBus
// var globalMemory *Memory

// func GetRequestForWriteOnMemory() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request BulkWriteMemoryRequest
// 		if err := c.BindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		byteAddress, err := strconv.Atoi(request.Address)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
// 			return
// 		}
// 		byteData, err := hex.DecodeString(request.Value)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
// 			return
// 		}

// 		if globalDataBus == nil {
// 			globalDataBus = bus.NewDataBus()
// 		}
// 		if globalMemory == nil {
// 			globalMemory = NewMemory()
// 		}
// 		globalDataBus.Write(byteAddress, byteData)
// 		data := ReadFromBus(byteAddress)
// 		log.Print("data", data)
// 		globalMemory.Store(byteAddress, data)
// 		globalDataBus.Delete(byteAddress)
// 		c.JSON(http.StatusCreated, gin.H{"message": "Data written successfully"})
// 	}
// }

// func ReadFromMemory() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		address := c.Param("address")

// 		addr, err := strconv.Atoi(address)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
// 			return
// 		}

// 		data := globalMemory.Load(addr)
// 		str := string(data)
// 		hexString := hex.EncodeToString([]byte(str))

// 		c.JSON(http.StatusOK, gin.H{"data": hexString})
// 	}
// }

// func ReadFromBus(address int) []byte {

// 	data := globalDataBus.Read(address)

// 	return data
// }
