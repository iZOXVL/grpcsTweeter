package main

import (
	"context"
	"log"
	"time"

	pb2 "grpc-go-example/analysis"
	pb "grpc-go-example/helloworld"
	pb3 "grpc-go-example/recommendation"

	"google.golang.org/grpc"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func main() {
	log.Println(Green + "Conectando al servidor gRPC..." + Reset)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf(Red+"No se pudo conectar: %v"+Reset, err)
	}
	defer conn.Close()
	c := pb.NewTweetServiceClient(conn)

	// Aumenta el tiempo de espera del contexto a 10 segundos para la solicitud de tweets
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println(Green + "Enviando solicitud al servidor gRPC..." + Reset)
	r, err := c.GetTweets(ctx, &pb.TweetRequest{TweetId: "1793082109743272142"})
	if err != nil {
		log.Fatalf(Red+"No se pudieron obtener los tweets: %v"+Reset, err)
	}

	log.Println(Green + "Tweets recibidos:" + Reset)
	for _, tweetText := range r.GetTweetTexts() {
		log.Printf(Cyan+"Tweet: %s"+Reset, tweetText)
	}

	// Crear una nueva conexión al servidor gRPC de análisis
	log.Println(Green + "Conectando al servidor gRPC de análisis..." + Reset)
	conn2, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf(Red+"No se pudo conectar al servidor de análisis: %v"+Reset, err)
	}
	defer conn2.Close()
	c2 := pb2.NewAnalysisServiceClient(conn2)

	// Preparar los datos para enviar al servidor de análisis
	var tweets []*pb2.Tweet
	for _, tweetText := range r.GetTweetTexts() {
		tweets = append(tweets, &pb2.Tweet{Tweet: tweetText})
	}

	log.Println(Green + "Enviando tweets al servidor gRPC de análisis..." + Reset)

	// Aumenta el tiempo de espera del contexto a 30 segundos para la solicitud de análisis
	analysisCtx, analysisCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer analysisCancel()

	// Enviar los datos al servidor de análisis
	analyzeRequest := &pb2.AnalyzeRequest{Tweets: tweets}
	resp, err := c2.ChatCompletions(analysisCtx, analyzeRequest)
	if err != nil {
		log.Fatalf(Red+"No se pudo obtener el análisis: %v"+Reset, err)
	}

	log.Println(Green + "Resultados del análisis recibidos:" + Reset)
	var emotions []string
	for i, item := range resp.GetResponses() {
		log.Printf(Cyan+"Tweet: %s - Emoción: %s"+Reset, r.GetTweetTexts()[i], item.GetEmotion())
		emotions = append(emotions, item.GetEmotion())
	}

	// Crear una nueva conexión al servidor gRPC de recomendaciones
	log.Println(Green + "Conectando al servidor gRPC de recomendaciones..." + Reset)
	conn3, err := grpc.Dial("localhost:50053", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf(Red+"No se pudo conectar al servidor de recomendaciones: %v"+Reset, err)
	}
	defer conn3.Close()
	c3 := pb3.NewRecommendationServiceClient(conn3)

	// Enviar los datos al servidor de recomendaciones
	recommendationRequest := &pb3.RecommendationRequest{Tweets: r.GetTweetTexts(), Emotions: emotions}
	recommendationCtx, recommendationCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer recommendationCancel()

	log.Println(Green + "Enviando tweets y emociones al servidor gRPC de recomendaciones..." + Reset)
	recommendationResp, err := c3.GetRecommendations(recommendationCtx, recommendationRequest)
	if err != nil {
		log.Fatalf(Red+"No se pudo obtener las recomendaciones: %v"+Reset, err)
	}

	log.Println(Green + "Resultados de las recomendaciones recibidos:" + Reset)
	log.Printf(Cyan+"Resumen: %s"+Reset, recommendationResp.GetSummary())
	log.Printf(Cyan+"Recomendaciones: %s"+Reset, recommendationResp.GetRecommendations())
	log.Println(Green + "Solicitud completada." + Reset)
}
