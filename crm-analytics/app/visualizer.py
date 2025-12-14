import plotly.express as px
import json
import os
import uuid
import pandas as pd
import matplotlib.pyplot as plt
REPORTS_DIR = "reports"
def flatten_customers(segment_dict):
    rows = []
    for segment, customers in segment_dict.items():
        for c in customers:
            c_copy = c.copy()
            c_copy['segment'] = segment
            rows.append(c_copy)
    return pd.DataFrame(rows)

def safe_load_json(text: str):
    if text.startswith("```"):
        text = text.strip("`")
        text = text.replace("json", "", 1).strip()
    try:
        return json.loads(text)
    except json.JSONDecodeError:
        return None
    
def plot_clusters_png(df: pd.DataFrame, filename_prefix="customer_clusters") -> str:
    if df.empty:
        return ""
    df = df.rename(columns={
        "totalSales": "total_sales",
        "lastPurchaseDays": "last_purchase_days"
    })

    if "total_sales" not in df.columns or "visits" not in df.columns:
        return ""

    plt.figure(figsize=(8, 6))

    for segment in df["segment"].unique():
        segment_df = df[df["segment"] == segment]
        plt.scatter(
            segment_df["visits"],
            segment_df["total_sales"],
            label=segment
        )

    plt.xlabel("Visits")
    plt.ylabel("Total Sales")
    plt.title("Clients by visits and sales")
    plt.legend()

    file_path = os.path.join(
        REPORTS_DIR,
        f"{filename_prefix}_{uuid.uuid4().hex}.png"
    )

    plt.savefig(file_path)
    plt.close()

    return file_path


def plot_sales_png(df: pd.DataFrame, filename_prefix="sales_forecast") -> str:
    if df.empty or "predicted_sales" not in df.columns or "date" not in df.columns:
        return ""
    fig = px.line(df, x="date", y="predicted_sales", title="Прогноз продаж")
    file_path = os.path.join(REPORTS_DIR, f"{filename_prefix}_{uuid.uuid4().hex}.png")
    fig.write_image(file_path)
    return file_path

