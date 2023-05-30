package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metloGin "github.com/metlo-labs/metlo/ingestors/golang/gin"
	"github.com/metlo-labs/metlo/ingestors/golang/metlo"
)

const LargeResponse = `{
	"0": {
			"uuid": "f3244b34-707b-4c21-a6b6-9c9763555902",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"1": {
			"uuid": "9316f16e-839c-4491-be0f-1bd3d69b9887",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"2": {
			"uuid": "f058bdbf-33f5-4f48-a942-b111f04b49ff",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"3": {
			"uuid": "e76ca527-513f-4a6b-933e-aa83291f9f6c",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"4": {
			"uuid": "9099d7f3-7983-4cd6-a495-83b9cc4f171f",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"5": {
			"uuid": "a03568e2-2405-4069-b6c3-cd358b951609",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"6": {
			"uuid": "aa0388f3-317b-4dbd-94e7-80b62eedea73",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"7": {
			"uuid": "1e386381-e35c-4677-a733-e6736e605d74",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"8": {
			"uuid": "ac70ad09-7009-46a1-876a-713382b6b28a",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"9": {
			"uuid": "08bf5732-3afc-4e49-9726-47451b019066",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"10": {
			"uuid": "a3c99493-8c66-4dfa-a454-02afa413054b",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"11": {
			"uuid": "ba3273e4-90bf-4799-b250-ea80bca2b8bd",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"12": {
			"uuid": "134cd960-a1cd-4966-a991-b0d5ccc15878",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"13": {
			"uuid": "2fbb913a-8466-4023-9436-969c707346ab",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"14": {
			"uuid": "8db2d504-36f8-49b1-837f-7c44c46f3c7f",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"15": {
			"uuid": "10459221-2a13-45bb-a747-78886733a174",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"16": {
			"uuid": "d87d40fa-be71-4b0c-a96e-cebf77893ed9",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"17": {
			"uuid": "18d632cf-4b4f-493d-8323-89e791feb6c2",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"18": {
			"uuid": "5f5a6e50-dae2-427e-9ab6-9d6f4b2859f0",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"19": {
			"uuid": "81094bbe-4029-4e35-8fe0-e4f0c6c16ec5",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"20": {
			"uuid": "f117601b-3b05-4a0f-ae36-5a8783d3a874",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"21": {
			"uuid": "c4eb127c-7c96-4409-9f1a-10fbc2b351b3",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"22": {
			"uuid": "039c19cc-9413-42f4-bfbd-f60b4e7bb40e",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"23": {
			"uuid": "4c845358-6785-424a-9b1d-57d4e0d8a4fa",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"24": {
			"uuid": "f52f13f9-9d04-45cd-9ddf-f3a1d135e37c",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"25": {
			"uuid": "eb5f7f6a-bf77-4751-ade2-f95ed8888241",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"26": {
			"uuid": "4efb2bfb-b504-457e-9af4-685dcfa17422",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"27": {
			"uuid": "139a246f-9f61-4c2b-8b01-a13bc0f4061f",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"28": {
			"uuid": "e788d69e-46bd-46b2-be8c-bad123adb140",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"29": {
			"uuid": "670ec298-55a4-4a3e-b2aa-29e0c74d857d",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"30": {
			"uuid": "e98c9012-41f3-445d-8e45-101a060ec6ca",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"31": {
			"uuid": "ed7a52bf-a753-4cf2-b127-6da94b305ffd",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"32": {
			"uuid": "b2792417-a48a-48bd-826a-87c29d9c2eb1",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"33": {
			"uuid": "6bbc2b66-af23-4f54-85a8-eadd11337ef4",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"34": {
			"uuid": "7d3023c3-e75a-4b44-ad49-f9ee58b80b84",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"35": {
			"uuid": "e49e9be4-e63a-470c-bc0b-b3af566d8bcc",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"36": {
			"uuid": "be55ac93-bd38-4f20-8eb8-68aced89608c",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"37": {
			"uuid": "6e7a1e44-4632-4783-b3b0-0545afeb2b54",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"38": {
			"uuid": "d5cf76e4-1eb6-4b57-b137-bc9ef4f6b5bf",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"39": {
			"uuid": "70dc54ea-ca7c-4c00-8c7e-9dfdbcf13ab3",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"40": {
			"uuid": "9933e5c4-49e0-4a0c-881e-4302d6372d66",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"41": {
			"uuid": "91ed4840-3285-4617-a7e0-901684e682c5",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"42": {
			"uuid": "af8f66c2-7eac-458c-aa94-15dd0e9719e5",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"43": {
			"uuid": "e26f0405-ef10-4c44-bded-be6e48426bbf",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"44": {
			"uuid": "6d331cd3-dbe6-48e0-b8c9-f9829ae81e7c",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"45": {
			"uuid": "621e59b5-5cac-4e5e-a853-225a54926526",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"46": {
			"uuid": "ba2d5c77-a1f6-43b9-8580-60d4e23cfce1",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"47": {
			"uuid": "2ab891dd-4ffc-43aa-a480-dacda794ffad",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"48": {
			"uuid": "4bcc6405-287c-4bb8-933b-2e180723a2e6",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"49": {
			"uuid": "0da9706b-123f-4748-b681-24d24fd7f547",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"50": {
			"uuid": "58ca011a-8f05-43c6-802d-ad697fd42ed9",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"51": {
			"uuid": "2e2c05cc-006f-4298-9808-6b1980c6fecd",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"52": {
			"uuid": "a6dd8011-2b5f-493a-b5ed-606d7bbc6881",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"53": {
			"uuid": "6f8a9267-7b37-44f4-adfb-b64bc72ee645",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"54": {
			"uuid": "e919148e-9bfe-4d22-b346-503e70105aa5",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"55": {
			"uuid": "a28d4039-f179-44aa-95a8-f041f1e7ae2b",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"56": {
			"uuid": "3b0c583c-4500-4b12-afa7-21842fe41eac",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"57": {
			"uuid": "8b55942f-03bf-4479-83aa-c30ced6849d1",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"58": {
			"uuid": "9107c87c-b447-4af4-9a1a-1325f6b54ab0",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"59": {
			"uuid": "7bc320d8-bc38-48c6-a0da-be2e59d087ea",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"60": {
			"uuid": "659c84f1-eb95-4b11-9bb1-04f5d05efe9c",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"61": {
			"uuid": "95765f55-221b-4274-a47d-481bcfede3fa",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"62": {
			"uuid": "dce7e9e9-5248-4158-8975-4d096c9f5c12",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"63": {
			"uuid": "10691b02-123f-44fb-a5ba-3ee0b935a5bb",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"64": {
			"uuid": "06efd5cf-a82d-400c-90e5-196c7a65a26d",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"65": {
			"uuid": "96b065ee-ed91-44a0-be5b-e9ae8795a684",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"66": {
			"uuid": "ed71fd4f-d628-4022-9bfa-b2d1c5799fab",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"67": {
			"uuid": "30c21da3-b39b-412d-9fd8-556c40ed41ea",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"68": {
			"uuid": "97021b7f-ea0d-4f83-a00f-2ebf75d10c13",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"69": {
			"uuid": "b670de3f-394c-445b-8135-56aea946e00f",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"70": {
			"uuid": "03bfa29f-a0d5-4b77-b064-9dc9a55d3391",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"71": {
			"uuid": "3d21a854-dde7-4c86-93b4-97110f7b78b8",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"72": {
			"uuid": "fedc7257-d556-4bbd-969e-51390cb1b2ee",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"73": {
			"uuid": "de1d07ae-4650-4910-a286-0ae50f022fa3",
			"name": "Product thing",
			"description": "SOme product thing",
			"price": 23
	},
	"ok": true
}`

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	metlo := metlo.InitMetlo("https://app.metlo.com", os.Getenv("METLO_API_KEY"))

	var res map[string]interface{}
	json.Unmarshal([]byte(LargeResponse), &res)

	instrumentation := metloGin.Init(metlo)
	r.Use(instrumentation.Middleware)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/large-response", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:3031")
}
