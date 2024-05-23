from concurrent import futures
import grpc
import logging
import recommendation_pb2
import recommendation_pb2_grpc
from groq import Groq

class RecommendationService(recommendation_pb2_grpc.RecommendationServiceServicer):
    def GetRecommendations(self, request, context):
        tweets = request.tweets
        emotions = request.emotions

        # Generar un resumen del análisis
        analysis_summary = self.generate_summary(tweets, emotions)
        recommendations = self.generate_recommendation(analysis_summary)

        return recommendation_pb2.RecommendationResponse(summary=analysis_summary, recommendations=recommendations)

    def generate_summary(self, tweets, emotions):
        # Generar un resumen del análisis
        summary = "Resumen del análisis:\n"
        for tweet, emotion in zip(tweets, emotions):
            summary += f"Tweet: {tweet} - Emoción: {emotion}\n"
        return summary

    def generate_recommendation(self, analysis):
        client = Groq(api_key="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")

        try:
            completion = client.chat.completions.create(
                model="llama3-70b-8192",
                messages=[
                    {"role": "system", "content": "Habla en español latinoamericano"},
                    {"role": "user", "content": f"Hola, aquí están los resultados del análisis: {analysis}. ¿Qué recomendaciones tienes?"}
                ],
                temperature=1,
                max_tokens=1024,
                top_p=1,
                stream=False
            )

            if completion.choices and len(completion.choices) > 0:
                response_content = completion.choices[0].message.content if hasattr(completion.choices[0].message, 'content') else "No hay texto de respuesta disponible"
                return response_content
            else:
                logging.warning("No se recibió una respuesta válida del modelo.")
                return "No se encontraron recomendaciones válidas."

        except Exception as e:
            logging.error(f"Ocurrió un error al generar recomendaciones: {e}")
            return "Error al generar recomendaciones."


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    recommendation_pb2_grpc.add_RecommendationServiceServicer_to_server(RecommendationService(), server)
    server.add_insecure_port('[::]:50053')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    serve()
