from fastapi import FastAPI
import os
import json
import uuid
import pandas as pd
import plotly.express as px
from app.analyzer import fetch_data
from app.gemini import analyze_customers_with_gemma
from app.visualizer import flatten_customers, safe_load_json

app = FastAPI()
GATEWAY_URL = "http://api-gateway:8080/v1/lead"

REPORTS_DIR = "reports"
os.makedirs(REPORTS_DIR, exist_ok=True)

def plot_clusters_pdf(df: pd.DataFrame, filename_prefix="customer_clusters") -> str:
    if df.empty:
        return ""
    fig = px.scatter(
        df,
        x="visits",
        y="total_sales",
        color="segment",
        hover_data=["name"],
        title="Клиенты по сегментам и продажам"
    )
    file_path = os.path.join(REPORTS_DIR, f"{filename_prefix}_{uuid.uuid4().hex}.pdf")
    fig.write_image(file_path)
    return file_path

def plot_sales_pdf(df: pd.DataFrame, filename_prefix="sales_forecast") -> str:
    if df.empty or "predicted_sales" not in df.columns or "date" not in df.columns:
        return ""
    fig = px.line(df, x="date", y="predicted_sales", title="Прогноз продаж")
    file_path = os.path.join(REPORTS_DIR, f"{filename_prefix}_{uuid.uuid4().hex}.pdf")
    fig.write_image(file_path)
    return file_path

@app.get("/gemma-analysis")
def gemma_analysis():
    df = fetch_data(GATEWAY_URL)
    customers = df.to_dict(orient="records")

    gemma_response = analyze_customers_with_gemma(customers)
    data = safe_load_json(gemma_response)
    if not data:
        return {"error": "Invalid response from model", "raw": gemma_response}

    df_customers = flatten_customers(data.get("segment_customers", {}))

    clusters_pdf_path = plot_clusters_pdf(df_customers)
    sales_pdf_path = plot_sales_pdf(df_customers)

    return {
        "customers_count": len(df_customers),
        "clusters_pdf_path": clusters_pdf_path,
        "sales_pdf_path": sales_pdf_path,
        "raw_model_response": data
    }