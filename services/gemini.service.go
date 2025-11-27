package services

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func GenerateDescription(c *gin.Context, menuName string, menuIngredients []string) (string, error){
	godotenv.Load()
	client, err := genai.NewClient(c, nil);
	if  err != nil{
		return "", err
	}

	teks := "Buatkan deskripsi menu makanan yang menarik untuk"+ menuName+ "dengan bahan-bahan:"
	for _, item := range menuIngredients{
		teks += item
	}
	teks += ". maksimal 2 kalimat"
	result, err := client.Models.GenerateContent(
		c,
		"gemini-2.5-flash",
		genai.Text(teks),
		nil,
	)

	if err != nil{
		return "", err
	}

	return result.Text(), nil
}