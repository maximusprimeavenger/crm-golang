import plotly.express as px
import json
import os
import uuid
import pandas as pd
REPORTS_DIR = "reports"
def flatten_customers(segment_dict):
    rows = []
    for segment, customers in segment_dict.items():
        for c in customers:
            c_copy = c.copy()
            c_copy['segment'] = segment
            rows.append(c_copy)
    return pd.DataFrame(rows)


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
def safe_load_json(text: str):
    if text.startswith("```"):
        text = text.strip("`")
        text = text.replace("json", "", 1).strip()
    try:
        return json.loads(text)
    except json.JSONDecodeError:
        return None
    
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