package controller

import (
	"cloud.google.com/go/vertexai/genai"
	"demo-travel-guide/internal/model"
	"demo-travel-guide/pkg/gcp"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func MakeItinerary(context *gin.Context) {
	var itineraryRequest model.ItineraryRequest
	if err := context.ShouldBindJSON(&itineraryRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data."})
	}

	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	modelName := os.Getenv("GOOGLE_CLOUD_MODEL")
	temperature := 1

	prompt := []genai.Part{
		genai.Text("<profile> Você é um agente de viagem extremamente experiente em montar roteiros incríveis dentro das limitações do seu cliente final. </profile>"),
		genai.Text("<task> Baseado nos parâmetros informados crie um itinerário de viagem compatível com a localização informada, o orçamento disponível no total (informado na moeda local do destino), quantidade de dias e de forma a considerar as restrições de visitações a evitar indicadas. Retorne isso em formato JSON válido contendo um resumo do roteiro elaborado, o custo total da viagem na moeda local, e o roteiro dividido em dias. Cada dia deve ter um resumo do roteiro daquele dia, e o custo total do roteiro do dia na moeda local e também precisa dentro dele mostrar os itens que compoem a rota escolhida. Cada item da rota deve informar a duração recomendada no passeio, o preço na moeda local (ou FREE se for gratuito), o nome / localização, um resumo pequeno do que esperar nesse local e os meios com os quais se podem chegar nesse local (Transporte por aplicativo, Metrô (citar linha que dá acesso), Onibus (citar linha que dá acesso), Caminhada, Bicicleta). { 'summary': '', 'total_cost': '', 'days': [ { 'summary': '', 'total_cost': '', 'route': [ { 'duration': '', 'price': '', 'location': '', 'what_to_expect': '', 'transport': '' } ] } ] } </task>"),
		genai.Text("<instructions> Seja direto e devolva apenas o JSON, ou seja, sem também acrescentar ```json no inicio e ``` no final, além disso gere sem formatação, retornando o JSON inline. Não repita locais e seja extremamente coerente com pontos turísticos no local indicado, respeitando o orçamento total indicado e não inclindo items que o cliente não deseja visitar. Nos campos que incluem valor monetário, acrescentar o símbolo da moeda, se for em um país nos estados unidos por exemplo acrescentar USD na frente. </instructions>"),
		genai.Text("<parameters> Localização:"), genai.Text(itineraryRequest.Location),
		genai.Text("Dias:"), genai.Text(itineraryRequest.Days),
		genai.Text("Orçamento total na moeda local:"), genai.Text(itineraryRequest.BudgetInLocalCurrency),
		genai.Text("Evitar:"), genai.Text(itineraryRequest.Avoid),
		genai.Text("</parameters>"),
	}

	answer, err := gcp.GenerateMultimodalContent(os.Stdout, prompt, projectId, location, modelName, float32(temperature))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error while generating multimodal content."})
	}

	var itinerary model.Itinerary
	err = json.Unmarshal(answer, &itinerary)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error on unmarshalling json."})
	}

	context.JSON(http.StatusOK, itinerary)
}
