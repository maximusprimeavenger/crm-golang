from fastapi import FastAPI
import os
from app.analyzer import fetch_data
from app.gemini import analyze_customers_with_gemma
from app.visualizer import flatten_customers, plot_clusters_png, safe_load_json,plot_sales_png

app = FastAPI()
GATEWAY_URL = "http://api-gateway:8080/v1/lead"

REPORTS_DIR = "reports"
os.makedirs(REPORTS_DIR, exist_ok=True)
@app.get("/gemma-analysis")
def gemma_analysis():
    df = fetch_data(GATEWAY_URL)
    customers = df.to_dict(orient="records")

    gemma_response = analyze_customers_with_gemma(customers)
    data = safe_load_json(gemma_response)
    if not data:
        return {"error": "Invalid response from model", "raw": gemma_response}

    df_customers = flatten_customers(data.get("segment_customers", {}))

    clusters_png_path = plot_clusters_png(df_customers)
    sales_png_path = plot_sales_png(df_customers)

    return {
        "customers_count": len(df_customers),
        "clusters_png_path": clusters_png_path,
        "sales_png_path": sales_png_path,
        "raw_model_response": data
    }